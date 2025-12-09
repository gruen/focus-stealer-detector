package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type AppInfo struct {
	Name string
	Path string
}

func getFrontmostApp() AppInfo {
	script := `tell application "System Events"
		set frontApp to first application process whose frontmost is true
		set appName to name of frontApp
		set appPath to POSIX path of (file of frontApp as alias)
		return appName & "|" & appPath
	end tell`
	cmd := exec.Command("osascript", "-e", script)
	out, err := cmd.Output()
	if err != nil {
		return AppInfo{}
	}
	parts := strings.SplitN(strings.TrimSpace(string(out)), "|", 2)
	if len(parts) == 2 {
		return AppInfo{Name: parts[0], Path: parts[1]}
	}
	return AppInfo{}
}

func main() {
	interval := flag.Duration("interval", 50*time.Millisecond, "polling interval (e.g., 100ms, 1s)")
	flag.Parse()

	lastApp := ""

	for {
		app := getFrontmostApp()
		if app.Name != "" && app.Name != lastApp {
			fmt.Printf("%s  %s  %s\n", time.Now().Format("2006-01-02 15:04:05.000"), app.Name, app.Path)
			lastApp = app.Name
		}
		time.Sleep(*interval)
	}
}
