package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asilvr/mystery"
	"github.com/urfave/cli/v2"
)
var version string

func main() {
	// flagLanguage stores the value passed in by the user for the language flag
	flagLanguage := ""
	// flagDifficulty stores the value passed in by the user for the difficulty flag
	flagDifficulty := ""
	// flagImports stores the value passed in by the user for the imports flag
	flagImports := ""

	app := &cli.App{
		Name:        "mystery",
		Usage:       "generate mystery ingredients for your next project",
		Description: "Inspired by the Food Network show Chopped, mystery provides you with mystery ingredients to help kickstart your new or existing project.",
		Version: 	  version,
		Authors: []*cli.Author{
			{
				Name:  "Alex Silver",
				Email: "alexaaronsilver@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "language", Usage: "supported values: random, go, kotlin, swift", DefaultText: "random", Destination: &flagLanguage},
			&cli.StringFlag{Name: "difficulty", Usage: "supported values: random, beginner, intermediate, advanced", DefaultText: "random", Destination: &flagDifficulty},
			&cli.StringFlag{Name: "imports", Usage: "supported values: all, native, external", DefaultText: "all", Destination: &flagImports},
		},
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generate mystery ingredients in a sample code snippet",
				Action: func(c *cli.Context) error {
					switch mystery.LangType(flagLanguage) {
					case mystery.NoLang:
						// this means the flag wasn't passed in, so we can pick a random one... sticking to Go for now
						dataset, _ := mystery.LoadDataset(mystery.GoLang)
						ingredients := dataset.Generate("", flagImports)
						mystery.PrintGenerate("Go", ingredients)
					case mystery.GoLang:
						dataset, _ := mystery.LoadDataset(mystery.GoLang)
						ingredients := dataset.Generate("", flagImports)
						mystery.PrintGenerate("Go", ingredients)
					case mystery.KotlinLang:
						dataset, _ := mystery.LoadDataset(mystery.KotlinLang)
						ingredients := dataset.Generate("", flagImports)
						mystery.PrintGenerate("Kotlin", ingredients)
					case mystery.SwiftLang:
						dataset, _ := mystery.LoadDataset(mystery.SwiftLang)
						ingredients := dataset.Generate("", flagImports)
						mystery.PrintGenerate("Swift", ingredients)
					default:
						fmt.Println(fmt.Sprintf("Error: specified language %s is not supported", flagLanguage))
					}
					return nil
				},
			},
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name: "version",
		Aliases: []string{"v, V"},
		Usage: "show version",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
