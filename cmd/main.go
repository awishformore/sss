package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codahale/sss"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "sss"
	app.Usage = "Shamir Secret Sharing"
	app.Commands = []cli.Command{
		{
			Name:  "split",
			Usage: "split secret into several parts",
			Flags: []cli.Flag{
				cli.UintFlag{Name: "n"},
				cli.UintFlag{Name: "k"},
			},
			Action: func(ctx *cli.Context) error {
				n := ctx.Uint("n")
				if n > 255 {
					return cli.NewExitError("n > 255", 1)
				}
				k := ctx.Uint("k")
				if k > 255 {
					return cli.NewExitError("k > 255", 1)
				}
				if k > n {
					return cli.NewExitError("k > n", 1)
				}
				var secret []byte
				shares, err := sss.Split(byte(n), byte(k), secret)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				for i, share := range shares {
					fmt.Printf("%v:%v\n", i, string(share))
				}
				return nil
			},
		},
		{
			Name:  "combine",
			Usage: "combine secret from several parts",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(1)
}
