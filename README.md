# xqmap

![xqmap](https://github.com/wpxq/xqmap/blob/main/xqmap.png)

---
A concurrent network port scanner written in Go, designed for DevOps and CyberSecurity enthusiasts. It features high-speed TCP scanning using goroutines, service detection, and colorized output for better readability.

---

## Features
- **Fast Scanning:** Utilizes Go's goroutines to scan multiple ports simultaneously.
- **Service Detection:** Optional banner grabbing (`-s` flag) to identify service versions (e.g., HTTP Server headers).
- **Smart Firewall Detection:** Identifies if a target is behind a firewall or honeypot by performing out-of-band validation on non-standard ports to filter false-positive results.
- **Deep Port Map:** Includes a comphrenesive list of ports covering:
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