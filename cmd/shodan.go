package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var (
	targetFacet    string
	threads        int
	batchDelay     int
	showSource     bool
	dumpAllTheShit bool
)

var reapCmd = &cobra.Command{
	Use:     "reap",
	Aliases: []string{"shodan", "rip", "harvest"},
	Short:   "Rip every fucking IP Shodan knows about",
	Long: `Brute-force harvests IPs from Shodan using facet splitting when results > 1000.
Feed it raw queries via stdin — it will gut Shodan alive.

Examples:
  echo 'ssl.cert.subject.cn:"github.com"' | ipharvester reap -f ip
  echo 'org:"Cloudflare"' | ipharvester reap -f ip -t 15
  cat targets.txt | ipharvester reap --all
`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		// Compile regex once, not every loop
		reTotal := regexp.MustCompile(`Total:\s*([\d,]+)`)
		reValues := regexp.MustCompile(`<strong>([^<]+)</strong>`)

		for scanner.Scan() {
			query := strings.TrimSpace(scanner.Text())
			if query == "" {
				continue
			}

			urlQuery := strings.ReplaceAll(query, " ", "+")
			total := getTotalCount(urlQuery, reTotal)

			if total == 0 {
				continue
			}

			if total <= 1000 || dumpAllTheShit {
				grabFacetValues(urlQuery, "", targetFacet, query, reValues)
			} else {
				cities := getSplitFacets(urlQuery, "city", reValues)
				if len(cities) == 0 {
					fmt.Printf("[!] No city split for %s — dumping raw anyway\n", query)
					grabFacetValues(urlQuery, "", targetFacet, query, reValues)
					continue
				}

				var wg sync.WaitGroup
				sem := make(chan struct{}, threads)

				for _, city := range cities {
					wg.Add(1)
					sem <- struct{}{}

					go func(c string) {
						defer wg.Done()
						defer func() { <-sem }()
						filtered := fmt.Sprintf("%s+city:\"%s\"", urlQuery, strings.ReplaceAll(c, " ", "+"))
						grabFacetValues(filtered, "", targetFacet, query, reValues)
						if batchDelay > 0 {
							time.Sleep(time.Duration(batchDelay) * time.Millisecond)
						}
					}(city)
				}
				wg.Wait()
			}
		}
	},
}

func getTotalCount(q string, reTotal *regexp.Regexp) int {
	out, _ := exec.Command("curl", "-s", fmt.Sprintf("https://www.shodan.io/search/facet?query=%s&facet=ip", q)).Output()
	match := reTotal.FindSubmatch(out)
	if len(match) < 2 {
		return 0
	}
	clean := strings.ReplaceAll(string(match[1]), ",", "")
	n, _ := strconv.Atoi(clean)
	return n
}

func getSplitFacets(q, facet string, reValues *regexp.Regexp) []string {
	out, _ := exec.Command("curl", "-s", fmt.Sprintf("https://www.shodan.io/search/facet?query=%s&facet=%s", q, facet)).Output()
	matches := reValues.FindAllStringSubmatch(string(out), -1)
	list := make([]string, 0, len(matches))
	for _, m := range matches {
		if len(m) > 1 {
			list = append(list, strings.TrimSpace(m[1]))
		}
	}
	return list
}

func grabFacetValues(q, extra, facet, src string, reValues *regexp.Regexp) {
	url := fmt.Sprintf("https://www.shodan.io/search/facet?query=%s%s&facet=%s", q, extra, facet)
	out, _ := exec.Command("curl", "-s", url).Output()

	matches := reValues.FindAllStringSubmatch(string(out), -1)
	for _, m := range matches {
		if len(m) > 1 {
			val := strings.TrimSpace(m[1])
			if showSource {
				fmt.Printf("%s::%s\n", src, val)
			} else {
				fmt.Println(val)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(reapCmd)

	reapCmd.Flags().StringVarP(&targetFacet, "facet", "f", "ip", "What to harvest (ip, domain, org, etc.)")
	reapCmd.Flags().IntVarP(&threads, "threads", "t", 10, "How many cities to rape at once")
	reapCmd.Flags().IntVarP(&batchDelay, "delay", "d", 0, "Delay between city batches (ms)")
	reapCmd.Flags().BoolVar(&showSource, "src", false, "Tag output with original query")
	reapCmd.Flags().BoolVar(&dumpAllTheShit, "all", false, "Force dump even under 1000 (skip city split)")
}
