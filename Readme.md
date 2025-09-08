# 🕵️ Go Web Scraper — Fetch Website Titles  

A simple **GoLang** project that fetches and displays the `<title>` tags of multiple websites concurrently using **goroutines** and **channels**. The results are served via an **HTTP server**.

---

## 🚀 Features  
- 🌐 Fetches titles from multiple websites  
- ⚡ Uses **goroutines** for concurrent processing  
- 📡 Serves data over an HTTP endpoint  
- 🛡️ Handles errors gracefully  
- 🏷️ Parses HTML using `golang.org/x/net/html`  

---




## 🛠️ Installation  

### 1. Clone the repository  
```bash
git clone https://github.com/yourusername/go-web-scraper.git
cd go-web-scraper

go mod init main.go
