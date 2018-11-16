// sudo apt search libgtk-3-dev libwebkit2gtk-4.0-dev

package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Wasap - Web Whatsapp",
		"https://web.whatsapp.com", 1024, 768, true)
}
