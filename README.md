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

***Zero-API public search engine harvester***
***Shodan · ZoomEye · ViewDNS → raw IPs. No keys. No credits. No limits.***
***Built for bug bounty, red team, and anyone who needs thousands of IPs right now.***

## Install (one command)
```
git clone https://github.com/cristophercervantes/IPHarvester.git && cd IPHarvester && sudo ./build.sh
```

### Or with Go (1.23+):
```
go install github.com/cristophercervantes/IPHarvester@latest
```
Binary lands in /usr/local/bin/ipharvester

## Commands
```
Command,What it does
reap,Shodan harvester (auto-bypasses 1000-result limit)
zm,ZoomEye multi-page ripper
history,ViewDNS historical IP lookup
```
