package main

import (
	"github.com/urfave/cli"
	"fmt"
	"net"
	"os"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "DNS CLI"
	app.Usage = "Lets you make DNS Queries"

	app.Flags = []cli.Flag {
		&cli.StringFlag {
			Name: "host",
			Value: "doc.ic.ac.uk",
		},
	}

	app.Commands = []*cli.Command {
		{
			Name: "ns",
			Usage: "look up the authoratative server for the host's domain",
			Flags: app.Flags,
			Action: lookupNS,
		},

		{
			Name: "a",
			Usage: "look up the IP addresses corresponding to the host names",
			Flags: app.Flags,
			Action: lookupIP,
		},

		{
			Name: "cname",
			Usage: "look up the canonical or primary name of the host",
			Flags: app.Flags,
			Action: lookupCNAME,
		},

		{
			Name: "mx",
			Usage: "look up the mail exchange server for the host's domain",
			Flags: app.Flags,
			Action: lookupMX,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func lookupNS(c *cli.Context) error {
	ns, err := net.LookupNS(c.String("host"))
	if err != nil {
		return err
	}
	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i].Host)
	}
	return nil
}

func lookupIP(c *cli.Context) error {
	ip, err := net.LookupIP(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(ip); i++ {
		fmt.Println(ip[i])
	}
	return nil
}

func lookupCNAME(c *cli.Context) error {
	cname, err := net.LookupCNAME(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cname)
	return nil
}

func lookupMX(c *cli.Context) error {
	mx, err := net.LookupMX(c.String("host"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for i := 0; i < len(mx); i++ {
		fmt.Println(mx[i].Host, mx[i].Pref)
	}
	return nil
}