# Gopher Watchdog
[![GoDoc](https://godoc.org/github.com/google/trillian?status.svg)](https://godoc.org/github.com/google/trillian)

A lightweight, concurrent HTTP monitoring tool written in Go. <br>
Gopher Watchdog continuously monitors the uptime and HTTP status of multiple target URLs. It uses Go's concurrency model (Goroutines and WaitGroups) to perform health checks efficiently without blocking, and employs `sync.Mutex` for thread-safe state management.

## Disclaimer & Warnings

Please use this tool responsibly. Since Gopher Watchdog can generate concurrent HTTP requests at high speeds, improper configuration can lead to unintended consequences:

* **Rate Limiting & IP Bans:** Setting the `interval_seconds` too low (e.g., `< 5 seconds`) against external servers might trigger Web Application Firewalls (WAF) like Cloudflare. This can result in your IP address being temporarily or permanently banned (HTTP 429 or 403).
* **Accidental DoS:** Do not use this tool with an aggressive polling rate against small, unoptimized, or third-party servers you do not own. It may cause performance degradation or system crashes on the target server.
* **Local OS Limits:** If you configure hundreds or thousands of targets, your local machine might hit the OS File Descriptor limits (e.g., `ulimit -n`). Ensure your system is configured to handle high numbers of concurrent network sockets if you plan to monitor at a massive scale.

**Note:** This tool is intended for educational purposes and monitoring your own infrastructure. The developer is not responsible for any misuse.

## Features
* **Concurrent Monitoring:** Pings multiple URLs simultaneously using Goroutines.
* **Custom HTTP Methods:** Supports `GET`, `POST`, `PUT`, `DELETE`, etc., via configuration.
* **Dynamic Configuration:** Reads targets and settings from a `config.json` file. No hardcoding required.
* **Interval & Limit Control:** Run continuously (Infinite Loop) or set a specific number of execution rounds with custom delays.
* **Thread-Safe:** Prevents Data Race conditions using Mutex locks when writing status results.

## Prerequisites
* Go 1.25.6 or higher

## Getting the Source Code
Using `git clone` allows you to work in whatever directory you want. You will
still need to set GOPATH in order to build some apps (recommended to put this in
a cache dir). E.g.:

```
$ cd ${WORKDIR}
$ git clone https://github.com/Thitipong-PP/gopher-watchdog.git
$ cd gopher-watchdog
```

Create a config.json file in the root directory (see Configuration section below).

And try to run
```
$ go run main.go
```

## Configuration
The tool requires a config.json file to run.<br>
* interval_seconds: Delay between each monitoring cycle (in seconds).<br>
* limit: Number of times to run the check. Set to -1 for an infinite loop.<br>
* targets: Array of target objects containing the url and HTTP method.<br>

Example ``` config.json ```:
```
{
  "interval_seconds": 5,
  "limit": -1,
  "targets": [
    {
      "url": "[https://google.com](https://google.com)",
      "method": "GET"
    },
    {
      "url": "[https://api.github.com](https://api.github.com)",
      "method": "GET"
    },
    {
      "url": "[https://this-web-does-not-exist.com](https://this-web-does-not-exist.com)",
      "method": "POST"
    }
  ]
}
```

## How It Works Under The Hood
- Reads config.json.<br>
- Loops through targets based on the limit condition.<br>
- Spawns a Goroutine for each target to send an HTTP request concurrently.<br>
- Safely writes the HTTP Status Code (or 0 if unreachable) to a shared Map using sync.Mutex.<br>
- Prints the results to the terminal and waits for interval_seconds before the next cycle.

## License
- <a href="https://github.com/Thitipong-PP/gopher-watchdog/blob/main/LICENSE">MIT License</a>