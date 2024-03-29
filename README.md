# TechProfiling
## Installation
```Bash
sudo mkdir -p /usr/local/tools/techprofiling && cd /usr/local/tools/techprofiling
sudo chown -R <user>:<user> ../techprofiling
go mod init techprofiling
go install -v github.com/projectdiscovery/wappalyzergo/cmd/update-fingerprints@latest
go get github.com/projectdiscovery/wappalyzergo
wget https://raw.githubusercontent.com/mjmoreee/techprofiling/main/techprofiling.go
go build techprofiling.go && sudo cp techprofiling /usr/local/go/bin
```

## Usage
```Bash
cat live-subdomain.txt | techprofiling
```
