package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	proxyFile := flag.String("proxy", "proxy.env", "path to proxy config file")
	rotate := flag.Int("rotate", 1, "number of proxies per tunnel port")
	listenStart := flag.Int("listen-start", 20000, "starting port")
	flag.Parse()

	fmt.Fprintf(os.Stdout, "proxytunnel starting\nproxy=%s\nrotate=%d\nlisten_start=%d\n", *proxyFile, *rotate, *listenStart)
}
