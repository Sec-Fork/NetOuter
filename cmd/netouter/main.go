package main

import (
	"NetOuter/pkg/checkdns"
	"NetOuter/pkg/checkhttp"
	"NetOuter/pkg/checkicmp"
	"NetOuter/pkg/checkntp"
	"NetOuter/pkg/checksnmp"
	"NetOuter/pkg/checktcp"
	"NetOuter/pkg/checktftp"
	"flag"
	"os"
)

var version = "0.1.0"

var (
	tcpFullCheckPtr *bool
	snmpCheckPtr    *bool
	tftpCheckPtr    *bool
	customip        *string
)

func main() {

	tcpFullCheckPtr = flag.Bool("tcp", false, "TCP 1-65535 full check use allports.exposed slow")
	snmpCheckPtr = flag.Bool("snmp", false, "snmp custom ip check")
	tftpCheckPtr = flag.Bool("tftp", false, "tftp custom ip check")
	customip = flag.String("ip", "1.1.1.1", "custom ip for snmp or tftp")

	flag.Parse()

	if *snmpCheckPtr {
		checksnmp.Checksnmp(*customip)
		os.Exit(0)
	}
	if *tftpCheckPtr {
		checktftp.Checktftp(*customip)
		os.Exit(0)
	}

	if *tcpFullCheckPtr {
		checktcp.CheckALLtcp()
	} else {

		checkntp.Checkntp()
		checksnmp.Checksnmp("116.162.120.19")
		checktftp.Checktftp("183.62.177.78")
		checkdns.CheckDirectDNS()
		checkdns.CheckLocalDNS()
		checkicmp.Checkicmp()
		checkhttp.Checkhttp()
		checktcp.Checktcp("45.79.204.144", "22")
		checktcp.Checktcp("220.181.38.148", "80")
		checktcp.Checktcp("220.181.38.148", "443")
	}

}
