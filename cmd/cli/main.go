package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Network-cli(nc)"
	app.Usage = "Easily handle all network queries"
	// Flags
	gFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Lookup for host nameservers",
			Flags: gFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := range ns {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Lookup for host IP Addresses",
			Flags: gFlags,
			Action: func(c *cli.Context) error {
				addresses, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := range addresses {
					fmt.Println(addresses[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Lookup for host cname",
			Flags: gFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Lookup for host mx records",
			Flags: gFlags,
			Action: func(c *cli.Context) error {
				records, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for i := range records {
					fmt.Println(records[i].Host, records[i].Pref)
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
