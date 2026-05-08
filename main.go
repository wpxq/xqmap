// github.com/wpxq
package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

type ScanResult struct {
	Port    int
	Service string
	Banner  string
}

func getBanner(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(time.Second * 2))
	fmt.Fprint(conn, "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n")
	buffer := make([]byte, 512)
	n, err := conn.Read(buffer)
	if err != nil {
		return ""
	}
	return string(buffer[:n])
}

func main() {
	cyan := color.New(color.FgCyan).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintfFunc()

	flag.Usage = func() {
		fmt.Printf("\n%s\n", white("[ xqmap ]"))
		fmt.Printf("Usage: xqmap %s\n\n", cyan("<host>"))
		fmt.Println(white("Options:"))
		fmt.Println("  -s      Enable service detection (banners)")
		fmt.Println("Example: xqmap -s 192.168.0.1")
	}
	servicePtr := flag.Bool("s", false, "Enable service detection")

	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		return
	}
	targetHost := flag.Arg(0)

	allPorts := map[int]string{
		// Remote management & System
		21:    "FTP",
		22:    "SSH",
		23:    "Telnet",
		3389:  "RDP",
		8291:  "Winbox",
		2022:  "SSH-Alt",
		5900:  "VNC",
		10000: "Webmin",
		111:   "RPCBind/Portmap",
		// Web Services
		80:   "HTTP",
		443:  "HTTPS",
		8080: "HTTP-Proxy",
		8008: "HTTP-Proxy",
		8443: "HTTPS-Alt",
		8009: "AJP13",
		3000: "React/Dev",
		9000: "Portainer",
		// Mail Services
		25:  "SMTP",
		110: "POP3",
		143: "IMAP",
		465: "SMTPS",
		993: "IMAPS",
		995: "POP3S",
		// Database & Cache
		3306:  "MySQL",
		5432:  "PostgreSQL",
		6379:  "Redis",
		27017: "MongoDB",
		1433:  "MS-SQL",
		9200:  "Elasticsearch",
		11211: "Memcached",
		// Gaming
		25565: "Minecraft",
		27015: "Source",
		30120: "FiveM",
		7777:  "Terraria",
		2302:  "Arma 3 / DayZ",
		34197: "Factorio",
		28015: "Rust",
		19132: "Minecraft Bedrock",
		// Infrastructure & IoT
		53:    "DNS",
		161:   "SNMP",
		445:   "SMB",
		554:   "RTSP (Video Stream - Camera)",
		1883:  "MQTT",
		1900:  "UPnP (SSDP - Device Discovery)",
		1935:  "RTMP (Video Stream)",
		2869:  "UPnP (Windows Event)",
		3702:  "WS-Discovery (ONVIF Camera)",
		5060:  "SIP (VoIP Phone)",
		5061:  "SIP-TLS (VoIP Phone)",
		5353:  "mDNS (Apple/Google Discovery)",
		8000:  "HikVision (CCTV Control)/Chromecast",
		8554:  "RTSP (Alternative Stream)",
		34567: "Xiongmai/XMeye (CCTV)",
		37777: "Dahua (CCTV Control)",
		// Windows & Enterprise Services
		88:   "Kerberos",
		135:  "RPC Endpoint Mapper",
		137:  "NetBIOS",
		138:  "NetBIOS",
		139:  "NetBIOS",
		389:  "LDAP",
		636:  "LDAPS",
		3268: "Global Catalog",
		3269: "Global Catalog",
		5985: "WinRM",
		5986: "WinRM",
		// Apple & Linux
		548:  "AFP",
		2049: "NFS",
		515:  "LDP",
		631:  "IPP/CUPS",
		// Monitoring, Backup & DevTools
		1521:  "Oracle DB",
		33060: "MySQL Shell",
		5672:  "RabbitMQ",
		8081:  "Nexus/HTTP-Alt",
		818:   "GlassFish/Jenkins-Alt",
		9090:  "Prometheus/Cockpit",
		9100:  "Node Exporter/Printer",
		9411:  "Zipkin",
		// VPN & Proxy
		1194:  "OpenVPN",
		1723:  "PPTP VPN",
		3128:  "Squid Proxy",
		8188:  "Privoxy",
		8888:  "Fiddler/Burp",
		9050:  "Tor Proxy",
		9051:  "Tor Control",
		1080:  "SOCKS4/5",
		6443:  "Kubernetes API",
		500:   "IKEv2/IPsec",
		4500:  "IKEv2/IPsec",
		51820: "WireGuard",
		// Other & Mixed
		1857: "Startron",
		1271: "EXCW",
		3001: "Nessus",
		7000: "AFS-Fileserver",
		5000: "Flask/UPnP",
		1064: "JSTEL",
		1244: "ISB-Conference",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var foundPorts []ScanResult

	fmt.Printf("\n%s Scanning: [%s]\n", magenta("[*]"), white(targetHost))
	testPort := 54321
	checkAddr := fmt.Sprintf("%s:%d", targetHost, testPort)
	testConn, testErr := net.DialTimeout("tcp", checkAddr, 1500*time.Millisecond)
	if testErr == nil {
		testConn.Close()
		fmt.Printf("%s %s\n", red("[!]"), yellow("WARNING: Firewall/Honeypot detected | Results may be fake."))
	}
	if *servicePtr {
		fmt.Printf("%s Service detection: %s\n", magenta("[*]"), green("[ENABLED]"))
	}
	fmt.Println("-------------------------------------------")

	for port, name := range allPorts {
		wg.Add(1)
		go func(p int, n string) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", targetHost, p)
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err == nil {
				banner := ""
				isWeb := strings.Contains(strings.ToLower(n), "http")
				if *servicePtr && isWeb {
					banner = getBanner(conn)
				}
				conn.Close()

				mu.Lock()
				foundPorts = append(foundPorts, ScanResult{p, n, banner})
				mu.Unlock()
			}
		}(port, name)
		time.Sleep(50 * time.Millisecond)
	}

	wg.Wait()

	sort.Slice(foundPorts, func(i, j int) bool {
		return foundPorts[i].Port < foundPorts[j].Port
	})

	if len(foundPorts) == 0 {
		fmt.Printf("%s No open ports found\n", red("[!]"))
	} else {
		for _, res := range foundPorts {
			fmt.Printf("%s Port %-5d [ %s ] %s", magenta("[*]"), res.Port, res.Service, green("open"))
			if res.Banner != "" {
				lines := strings.Split(res.Banner, "\n")
				var serverInfo string
				for _, line := range lines {
					if strings.Contains(strings.ToLower(line), "server:") {
						serverInfo = strings.TrimSpace(line)
						break
					}
				}
				if serverInfo != "" {
					fmt.Printf(" -> %s", yellow(serverInfo))
				} else if len(lines) > 0 {
					fmt.Printf(" -> %s", yellow(strings.TrimSpace(lines[0])))
				}
			}
			fmt.Println()
		}
	}
}
