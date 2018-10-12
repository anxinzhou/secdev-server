package main

import (
	"contract"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"sort"
)

// 10000,1,["0x2131","0x21321","0x111"]
type Arguments struct {
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "genkey",
			Usage: "generate keystore and store in -f",
			Action: func(c *cli.Context) error {
				ct := new(contract.Contract)
				_, err := ct.GenerateKeyStore(c.String("file"), c.String("password"))
				if err != nil {
					fmt.Print(err.Error())
				}
				return err

			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Value: "../etc/keystore",
					Usage: "store generated keystore in `FILE`",
				},
				cli.StringFlag{
					Name:  "password,p",
					Usage: "`password` for keystore file",
				},
			},
		},
		{
			Name:  "deploy",
			Usage: "deploy contract",
			Action: func(c *cli.Context) error {

				if c.String("chain") == "public" {
					pbc := new(contract.Pbc)
					pbc.Connect(c.String("port"))
					pbc.LoadKey(c.String("file"), c.String("password"))
				} else if c.String("chain") == "private" {
				} else {
					panic("wrong chain type, public or private")
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file,f",
					Value: "../etc/keystore",
					Usage: "load  generated keystore in `FILE`",
				},
				cli.StringFlag{
					Name:  "args,a",
					Usage: "`args` for deployment, json format",
				},
				cli.StringFlag{
					Name:  "port,P",
					Usage: "`port` to connect to chain",
				},
				cli.StringFlag{
					Name:  "chain,c",
					Usage: "specify `chain`, public or private",
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
