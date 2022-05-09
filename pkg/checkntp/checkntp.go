package checkntp

import (
	"fmt"

	"github.com/beevik/ntp"
)

func Checkntp() {
	_, err := ntp.Time("52.231.114.183")
    if err != nil {
        fmt.Println("[-] NTP protocol(UDP 123) blocked")
        return
    }
    fmt.Println("[*] UDP 123 can access the internet")
}
