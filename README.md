# xqmap

![xqmap](https://github.com/wpxq/xqmap/blob/main/xqmap.png)

---

A high-performance, concurrent network port scanner written in Go. It utilizes a controlled worker pool pattern for efficient resource management, allowing for high-speed TCP scanning without overwhelming the host system. Features include banner grabbing for service detection, firewall/honeypot alerts, and a clean, colorized CLI interface.

---

## Features
- **High-Performance Scanning:** Implements a controlled worker pool pattern using Go's goroutines and channels, ensuring maximum speed without exhausting system resources (preventing "too many open files" errors).
- **Service Detection:** Optional banner grabbing (`-s` flag) to identify service versions for **HTTP**, **SSH**, **FTP**, **SMTP** and more.
- **Firewall Detection:** Identifies if a target is behind a firewall or honeypot by performing out-of-band validation on non-standard ports to filter false-positive results.
- **Deep Port Map:** Includes a comprehensive list of ports covering:
    - Infrastructure (SSH, RDP, WinBox)
    - Databases (MySQL, PostgreSQL, Redis, MongoDB)
    - CCTV/IoT (HikVision, Dahua, RTSP, MQTT)
    - Gaming (Minecraft, FiveM, Rust, Arma)
- **Colorized Output:** Clean and organized terminal UI using the `fatih/color` library.

## Installation
- Go (1.18 or higher)
### Quick Install (Linux/MacOS)
if you are using the provided installation script:
1. Clone the repo:
```bash
git clone https://github.com/wpxq/xqmap
cd xqmap
```
2. Run the installer:
```bash
chmod +x xqmap_setup.sh
./xqmap_setup.sh
```
## Usage
### Examples
#### Basic Scan:
```bash
xqmap 192.168.1.1
```
#### Scan with service detection
```bash
xqmap -s scanme.org
```
## Disclaimer
This tool is for educational and ethical testing purposes only. The author is not responsible for any misuse or damage caused by this program. Always obtain permission before scanning any network.