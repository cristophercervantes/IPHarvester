package cmd

import (
	"bufio"
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
	userAgent string
	timeout   int
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"viewdns", "dnsip", "oldips"},
	Short:   "Rip every historical IP a domain ever had",
	Long: `Pulls full IP history from viewdns.info — no subdomains, just the raw root domain.
Feed it clean domains via stdin and watch the corpses pile up.

Examples:
  echo "github.com" | ipharvester history
  cat roots.txt | ipharvester history > old_github_ips.txt
  echo "tesla.com" | ipharvester history -t 15
`,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}

		scanner := bufio.NewScanner(os.Stdin)
		ipRe := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)

		for scanner.Scan() {
			domain := strings.TrimSpace(scanner.Text())
			if domain == "" || strings.Contains(domain, ".") == false {
				continue
			}

			url := "https://viewdns.info/iphistory/?domain=" + domain
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", userAgent)
			req.Header.Set("Referer", "https://google.com") // extra stealth

			resp, err := client.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[!] %s → request failed: %v\n", domain, err)
				continue
			}

			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			ips := ipRe.FindAll(body, -1)
			if len(ips) == 0 {
				// fmt.Fprintf(os.Stderr, "[!] %s → no history found\n", domain)
				continue
			}

			for _, ip := range ips {
				fmt.Printf("%s\n", ip)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)

	historyCmd.Flags().StringVarP(&userAgent, "ua", "u", 
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/131.0 Safari/537.36",
		"Custom User-Agent")
	historyCmd.Flags().IntVarP(&timeout, "timeout", "t", 20, "Request timeout in seconds")
}
