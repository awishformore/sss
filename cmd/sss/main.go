package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
			Usage: "split secret into several shares",
			Flags: []cli.Flag{
				cli.UintFlag{Name: "n"},
				cli.UintFlag{Name: "k"},
				cli.StringFlag{Name: "secret"},
			},
			Action: func(ctx *cli.Context) error {
				n := ctx.Uint("n")
				k := ctx.Uint("k")
				secret := []byte(ctx.String("secret"))
				if len(secret) == 0 {
					return cli.NewExitError("secret empty", 1)
				}
				shares, err := sss.Split(byte(n), byte(k), secret)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				for index, share := range shares {
					fmt.Printf("%vx%v\n", index, hex.EncodeToString(share))
				}
				return nil
			},
		},
		{
			Name:  "combine",
			Usage: "combine secret from several shares",
			Flags: []cli.Flag{
				cli.StringSliceFlag{Name: "shares"},
			},
			Action: func(ctx *cli.Context) error {
				values := ctx.StringSlice("shares")
				shares := make(map[byte][]byte, len(values))
				for _, value := range values {
					parts := strings.Split(value, "x")
					if len(parts) != 2 {
						return cli.NewExitError("invalid share", 1)
					}
					index, err := strconv.ParseUint(parts[0], 8, 10)
					if err != nil {
						return cli.NewExitError(err.Error(), 1)
					}
					share, err := hex.DecodeString(parts[1])
					if err != nil {
						return cli.NewExitError(err.Error(), 1)
					}
					shares[byte(index)] = share
				}
				secret := sss.Combine(shares)
				fmt.Println(string(secret))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
