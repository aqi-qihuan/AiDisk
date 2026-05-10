package util

import (
	"fmt"
	"net"
	"os"
)

type CheckResult struct {
	Name    string
	OK      bool
	Latency string
	Skip    bool
}

func PrintStartupBanner(port string, elapsedMs int64, checks []CheckResult) {
	hostname, _ := os.Hostname()
	localIP := getLocalIP()

	addr := fmt.Sprintf("http://%s:%s", localIP, port)

	fmt.Println()
	fmt.Println("┌──────────────────────────────────────────────────────┐")
	fmt.Println("│                 AqiCloud-AgentPan-Go 启动完成            │")
	fmt.Println("├───────────────┬──────────────────────────────────────┤")
	fmt.Printf("│  hostname     │  %-38s│\n", hostname)
	fmt.Printf("│  profiles     │  %-38s│\n", profile())
	fmt.Printf("│  port         │  %-38s│\n", port)
	fmt.Printf("│  address      │  %-38s│\n", addr)
	fmt.Printf("│  elapsed      │  %6dms%32s│\n", elapsedMs, "")
	fmt.Println("├───────────────┴──────────────────────────────────────┤")
	fmt.Printf("│  API Base     │  %-38s│\n", addr+"/api")
	fmt.Printf("│  Swagger      │  %-38s│\n", addr+"/swagger/index.html")
	fmt.Println("└──────────────────────────────────────────────────────┘")

	fmt.Println()
	fmt.Println("┌─ 外部服务诊断 ──────────────────────────────────────┐")
	for _, c := range checks {
		if c.Skip {
			fmt.Printf("│  %-10s:  ⏭  SKIPPED%32s│\n", c.Name, "")
			continue
		}
		icon := "✅ OK"
		if !c.OK {
			icon = "❌ FAIL"
		}
		fmt.Printf("│  %-10s:  %s  (%s)%19s│\n", c.Name, icon, c.Latency, "")
	}
	fmt.Println("└─────────────────────────────────────────────────────┘")
	fmt.Println()
}

func getLocalIP() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ip := ipnet.IP.To4().String()
			// Skip link-local (169.254.x.x)
			if !ipnet.IP.IsLinkLocalUnicast() {
				return ip
			}
		}
	}
	return "127.0.0.1"
}

func profile() string {
	if os.Getenv("DEBUG") == "true" {
		return "debug"
	}
	return "release"
}
