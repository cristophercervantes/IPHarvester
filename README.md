# IPHarvester v1.0
```
    ___ ____  __  __    __   __           __
   (_  )__  \/ / / /_  / /  / /_  _______/ /_____  _____
  / / / / / / /_  __ \/ /  / __ \/ __   / __/ __ \/ ___/
 / / / / / / / / /_/ / /  / / / / /_/  / /_/ /_/ / /
/_/_/ / /_/ /_/ .___/_/  /_/ /_/\__,_/\__/\____/_/
             /_/
    IPHarvester V1.0
```

**Zero-API public search engine harvester**  
**Shodan · ZoomEye · ViewDNS → raw IPs. No keys. No credits. No limits.**  
**Built for bug bounty, red team, and anyone who needs thousands of IPs right now.**

## Install (one command)
```
git clone https://github.com/cristophercervantes/IPHarvester.git && cd IPHarvester && sudo ./build.sh
```

### Or with Go (1.23+):
```
go install github.com/cristophercervantes/IPHarvester@latest
```
Binary lands in /usr/local/bin/ipharvester

### Commands

| Command    | Description                                              |
|------------|----------------------------------------------------------|
| `reap`     | Shodan harvester (auto-bypasses 1000-result limit)       |
| `zm`       | ZoomEye multi-page ripper                                |
| `history`  | ViewDNS historical IP lookup                             |

### Usage Examples
```
# GitHub real origin IPs
echo 'ssl.cert.subject.cn:"github.com"' | ipharvester reap -t 80

# Cloudflare origins (expect 10k+)
echo 'org:"Cloudflare"' | ipharvester reap -f ip -t 120

# Russian SSH servers
echo 'port:22 country:RU' | ipharvester reap -t 50

# ZoomEye — Chinese web servers
echo 'app:"nginx" country:"cn"' | ipharvester zm -p 20 -t 100

# Old IPs for a domain
echo "microsoft.com" | ipharvester history

# Silent mode (perfect for pipelines)
cat queries.txt | ipharvester -s reap | sort -u > targets.txt
```

### Flags
```
-v, --version      Print version (v1.0)
-s, --silent       Suppress banner
```

### Build from source
```
git clone https://github.com/cristophercervantes/IPHarvester.git
cd IPHarvester
sudo ./build.sh
```

### Disclaimer
**This tool only queries publicly available data.**
**Use responsibly. Use proxies/VPS when hammering hard.**

### License
**MIT © 2025 cristophercervantes**
**IPHarvester v1.0 — Harvest or be harvested.**


