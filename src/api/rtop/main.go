/*

rtop - the remote system monitoring utility

Copyright (c) 2015 RapidLoop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package rtop

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
)

// func main() {
// 	GetHostStats("117.78.19.76", 15429)
// }

func GetHostStats(HostIP string, Port int) (ret int, stats Stats) {
	// key
	usr, err := user.Current()
	if err != nil {
		log.Print(err)
		return -1, Stats{}
	}

	keyPath := filepath.Join(usr.HomeDir, ".ssh", "id_rsa")
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		keyPath = ""
		return -1, Stats{}
	}

	addr := fmt.Sprintf("%s:%d", HostIP, Port)
	client := sshConnect(usr.Username, addr, keyPath)

	stats = Stats{}
	getAllStats(client, &stats)
	DisplayStats(stats)
	return 0, stats
}

const (
	escClear       = "\033[H\033[2J"
	escRed         = "\033[31m"
	escReset       = "\033[0m"
	escBrightWhite = "\033[37;1m"
)

func DisplayStats(stats Stats) {
	fmt.Printf(
		`%s%s%s%s up %s%s%s

Load:
    %s%s %s %s%s

Processes:
    %s%s%s running of %s%s%s total

Memory:
    free    = %s%s%s
    used    = %s%s%s
    buffers = %s%s%s
    cached  = %s%s%s
    swap    = %s%s%s free of %s%s%s

`,
		escClear,
		escBrightWhite, stats.Hostname, escReset,
		escBrightWhite, fmtUptime(&stats), escReset,
		escBrightWhite, stats.Load1, stats.Load5, stats.Load10, escReset,
		escBrightWhite, stats.RunningProcs, escReset,
		escBrightWhite, stats.TotalProcs, escReset,
		escBrightWhite, fmtBytes(stats.MemFree), escReset,
		escBrightWhite, fmtBytes(stats.MemTotal-stats.MemFree-stats.MemBuffers-stats.MemCached), escReset,
		escBrightWhite, fmtBytes(stats.MemBuffers), escReset,
		escBrightWhite, fmtBytes(stats.MemCached), escReset,
		escBrightWhite, fmtBytes(stats.SwapFree), escReset,
		escBrightWhite, fmtBytes(stats.SwapTotal), escReset,
	)
	if len(stats.FSInfos) > 0 {
		fmt.Println("Filesystems:")
		for _, fs := range stats.FSInfos {
			fmt.Printf("    %s%8s%s: %s%s%s free of %s%s%s\n",
				escBrightWhite, fs.MountPoint, escReset,
				escBrightWhite, fmtBytes(fs.Free), escReset,
				escBrightWhite, fmtBytes(fs.Used+fs.Free), escReset,
			)
		}
		fmt.Println()
	}
	if len(stats.NetIntf) > 0 {
		fmt.Println("Network Interfaces:")
		keys := make([]string, 0, len(stats.NetIntf))
		for intf := range stats.NetIntf {
			keys = append(keys, intf)
		}
		sort.Strings(keys)
		for _, intf := range keys {
			info := stats.NetIntf[intf]
			fmt.Printf("    %s%s%s - %s%s%s, %s%s%s\n",
				escBrightWhite, intf, escReset,
				escBrightWhite, info.IPv4, escReset,
				escBrightWhite, info.IPv6, escReset,
			)
			fmt.Printf("      rx = %s%s%s, tx = %s%s%s\n",
				escBrightWhite, fmtBytes(info.Rx), escReset,
				escBrightWhite, fmtBytes(info.Tx), escReset,
			)
			fmt.Println()
		}
		fmt.Println()
	}
}

func showStats(client *ssh.Client) {
	stats := Stats{}
	getAllStats(client, &stats)
	DisplayStats(stats)
}

func StringStats(stats Stats) (ret string) {

	ret += fmt.Sprintf(
		`%s%s%s%s up %s%s%s

Load:
    %s%s %s %s%s

Processes:
    %s%s%s running of %s%s%s total

Memory:
    free    = %s%s%s
    used    = %s%s%s
    buffers = %s%s%s
    cached  = %s%s%s
    swap    = %s%s%s free of %s%s%s

`,
		"",
		"", stats.Hostname, "",
		"", fmtUptime(&stats), "",
		"", stats.Load1, stats.Load5, stats.Load10, "",
		"", stats.RunningProcs, "",
		"", stats.TotalProcs, "",
		"", fmtBytes(stats.MemFree), "",
		"", fmtBytes(stats.MemTotal-stats.MemFree-stats.MemBuffers-stats.MemCached), "",
		"", fmtBytes(stats.MemBuffers), "",
		"", fmtBytes(stats.MemCached), "",
		"", fmtBytes(stats.SwapFree), "",
		"", fmtBytes(stats.SwapTotal), "",
	)

	if len(stats.FSInfos) > 0 {
		fmt.Println("Filesystems:")
		for _, fs := range stats.FSInfos {
			ret += fmt.Sprintf("    %s%8s%s: %s%s%s free of %s%s%s\n",
				"", fs.MountPoint, "",
				"", fmtBytes(fs.Free), "",
				"", fmtBytes(fs.Used+fs.Free), "",
			)
		}
		ret += "\n"
	}

	if len(stats.NetIntf) > 0 {
		fmt.Println("Network Interfaces:")
		keys := make([]string, 0, len(stats.NetIntf))
		for intf := range stats.NetIntf {
			keys = append(keys, intf)
		}
		sort.Strings(keys)
		for _, intf := range keys {
			info := stats.NetIntf[intf]
			ret += fmt.Sprintf("    %s%s%s - %s%s%s, %s%s%s\n",
				"", intf, "",
				"", info.IPv4, "",
				"", info.IPv6, "",
			)
			ret += fmt.Sprintf("      rx = %s%s%s, tx = %s%s%s\n",
				"", fmtBytes(info.Rx), "",
				"", fmtBytes(info.Tx), "",
			)
			ret += "\n"
		}
		ret += "\n"
	}

	return
}
