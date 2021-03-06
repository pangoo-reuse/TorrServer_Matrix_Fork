package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/pkg/browser"

	"server"
	"server/log"
	"server/settings"
	"server/torr"
	"server/version"
	"server/web/api/utils"
)

type args struct {
	Port        string `arg:"-p" help:"web server port"`
	Path        string `arg:"-d" help:"database dir path"`
	LogPath     string `arg:"-l" help:"server log file path"`
	WebLogPath  string `arg:"-w" help:"web access log file path"`
	RDB         bool   `arg:"-r" help:"start in read-only DB mode"`
	HttpAuth    bool   `arg:"-a" help:"enable http auth on all requests"`
	DontKill    bool   `arg:"-k" help:"don't kill server on signal"`
	UI          bool   `arg:"-u" help:"open torrserver page in browser"`
	TorrentsDir string `arg:"-t" help:"autoload torrents from dir"`
}

func (args) Version() string {
	return "TorrServer " + version.Version
}

var params args

func main() {
	arg.MustParse(&params)

	if params.Path == "" {
		params.Path, _ = os.Getwd()
	}

	if params.Port == "" {
		params.Port = "8090"
	}

	settings.Path = params.Path
	settings.HttpAuth = params.HttpAuth
	log.Init(params.LogPath, params.WebLogPath)
	fmt.Println("=========== START ===========")
	fmt.Println("TorrServer", version.Version+",", runtime.Version())
	if params.HttpAuth {
		log.TLogln("Use HTTP Auth file", settings.Path+"/accs.db")
	}

	dnsResolve()
	Preconfig(params.DontKill)

	if params.UI {
		go func() {
			time.Sleep(time.Second)
			browser.OpenURL("http://127.0.0.1:" + params.Port)
		}()
	}

	if params.TorrentsDir != "" {
		go watchTDir(params.TorrentsDir)
	}

	server.Start(params.Port, params.RDB)
	log.TLogln(server.WaitServer())
	log.Close()
	time.Sleep(time.Second * 3)
	os.Exit(0)
}

func dnsResolve() {
	hosts := [6]string{"1.1.1.1", "1.0.0.1", "208.67.222.222", "208.67.220.220", "8.8.8.8", "8.8.4.4"}
	ret := 0
	for _, ip := range hosts {
		ret = toolResolve("www.google.com", ip)
		switch {
		case ret == 2:
			fmt.Println("DNS resolver OK\n")
		case ret == 1:
			fmt.Println("New DNS resolver OK\n")
		case ret == 0:
			fmt.Println("New DNS resolver failed\n")
		}
		if ret == 2 || ret == 1 {
			break
		}
	}
}

func toolResolve(host string, serverDNS string) int {
	addrs, err := net.LookupHost(host)
	addr_dns := fmt.Sprintf("%s:53", serverDNS)
	a := 0
	if len(addrs) == 0 {
		fmt.Println("Check dns", addrs, err)
		fn := func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", addr_dns)
		}
		net.DefaultResolver = &net.Resolver{
			Dial: fn,
		}
		addrs, err = net.LookupHost(host)
		fmt.Println("Check new dns", addrs, err)
		if err == nil || len(addrs) > 0 {
			a = 1
		} else {
			a = 0
		}
	} else {
		a = 2
	}
	return a
}

func watchTDir(dir string) {
	time.Sleep(5 * time.Second)
	path, err := filepath.Abs(dir)
	if err != nil {
		path = dir
	}
	for {
		files, err := ioutil.ReadDir(path)
		if err == nil {
			for _, file := range files {
				filename := filepath.Join(path, file.Name())
				if strings.ToLower(filepath.Ext(file.Name())) == ".torrent" {
					sp, err := utils.ParseLink("file://" + filename)
					if err == nil {
						tor, err := torr.AddTorrent(sp, "", "", "")
						if err == nil {
							if tor.GotInfo() {
								if tor.Title == "" {
									tor.Title = tor.Name()
								}
								torr.SaveTorrentToDB(tor)
								tor.Drop()
								os.Remove(filename)
								time.Sleep(time.Second)
							}
						}
					}
				}
			}
		}
		time.Sleep(time.Second)
	}
}
