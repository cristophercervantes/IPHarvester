package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	zmThreads int
	zmPages   int
	zmDelay   int
	zmUA      string
)

type zmMatch struct {
	IP []string `json:"ip"`
}

type zmResp struct {
	Matches []zmMatch `json:"matches"`
}

var zmCmd = &cobra.Command{
	Use:     "zm",
	Aliases: []string{"zoomeye", "zoom", "zeye"},
	Short:   "Dump every IP ZoomEye has ever seen",
	Long:    `... (same as before)`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{Timeout: 25 * time.Second}
		ipRe := regexp.MustCompile(`(\b(?:\d{1,3}\.){3}\d{1,3}\b|[a-fA-F0-9:]+:[a-fA-F0-9:]+)`)

		scanner := bufio.NewScanner(os.Stdin)
		sem := make(chan struct{}, zmThreads)

		for scanner.Scan() {
			query := strings.TrimSpace(scanner.Text())
			if query == "" {
				continue
			}

			for page := 1; page <= zmPages; page++ {
				sem <- struct{}{}
				go func(q string, p int) {
					defer func() { <-sem }()
					url := fmt.Sprintf("https://www.zoomeye.hk/api/search?q=%s&page=%d&t=v4+v6+web", urlEncode(q), p)

					req, _ := http.NewRequest("GET", url, nil)
					req.Header.Set("User-Agent", zmUA)
					req.Header.Set("Accept", "application/json")
					req.Header.Set("Origin", "https://www.zoomeye.hk")

					resp, err := client.Do(req)
					if err != nil || resp.StatusCode != 200 {
						if resp != nil {
							resp.Body.Close()
						}
						return
					}

					body, _ := io.ReadAll(resp.Body)
					resp.Body.Close()

					var data zmResp
					if json.Unmarshal(body, &data) != nil {
						// Fallback to regex if JSON is fucked
						matches := ipRe.FindAll(body, -1)
						for _, ip := range matches {
							fmt.Println(string(ip))
						}
						return
					}

					for _, m := range data.Matches {
						for _, ip := range m.IP {
							fmt.Println(ip)
						}
					}

					if zmDelay > 0 {
						time.Sleep(time.Duration(zmDelay) * time.Millisecond)
					}
				}(query, page)
			}
		}

		// Drain semaphore
		for i := 0; i < cap(sem); i++ {
			sem <- struct{}{}
		}
	},
}

func urlEncode(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", "%20"), "\"", "%22")
}

func init() {
	rootCmd.AddCommand(zmCmd)
	zmCmd.Flags().IntVarP(&zmThreads, "threads", "t", 20, "Concurrent requests")
	zmCmd.Flags().IntVarP(&zmPages, "pages", "p", 3, "Pages to rip")
	zmCmd.Flags().IntVarP(&zmDelay, "delay", "d", 0, "Delay between requests (ms)")
	zmCmd.Flags().StringVar(&zmUA, "ua", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36", "User-Agent")
}
