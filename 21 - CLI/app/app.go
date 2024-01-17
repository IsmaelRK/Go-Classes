package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Cli Application"
	app.Usage = "Search for ips and server names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for ips",
			Flags:  flags,
			Action: searchIps,
		},

		{
			Name:   "servers",
			Usage:  "Search for servers name",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {

	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {

	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
