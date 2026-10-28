// Команда hexlet-path-size печатает размер файла или директории.
package main

import (
	"context"
	"fmt"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "recursive", Aliases: []string{"r"}, Usage: "recursive size of directories"},
			&cli.BoolFlag{Name: "human", Aliases: []string{"H"}, Usage: "human-readable sizes (auto-select unit)"},
			&cli.BoolFlag{Name: "all", Aliases: []string{"a"}, Usage: "include hidden files and directories"},
		},
		Action: run,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run печатает размер в формате "<размер>\t<путь>".
func run(_ context.Context, cmd *cli.Command) error {
	if cmd.NArg() != 1 {
		return fmt.Errorf("expected exactly one path, got %d", cmd.NArg())
	}

	path := cmd.Args().First()

	size, err := code.GetPathSize(path, cmd.Bool("recursive"), cmd.Bool("human"), cmd.Bool("all"))
	if err != nil {
		return fmt.Errorf("get path size: %w", err)
	}

	fmt.Printf("%s\t%s\n", size, path)

	return nil
}
