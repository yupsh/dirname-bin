package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/dirname"
)

const (
	flagZero = "zero"
)

func main() {
	app := &cli.App{
		Name:  "dirname",
		Usage: "strip last component from file name",
		UsageText: `dirname [OPTION] NAME...

   Output each NAME with its last non-slash component and trailing slashes
   removed; if NAME contains no /'s, output '.' (meaning the current directory).`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    flagZero,
				Aliases: []string{"z"},
				Usage:   "end each output line with NUL, not newline",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "dirname: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add all arguments
	for i := 0; i < c.NArg(); i++ {
		params = append(params, c.Args().Get(i))
	}

	// Add flags based on CLI options
	if c.Bool(flagZero) {
		params = append(params, Zero)
	}

	// Create and execute the dirname command
	cmd := Dirname(params...)
	return gloo.Run(cmd)
}
