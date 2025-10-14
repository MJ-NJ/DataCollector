// This will house the main server which will control the server
// This will get input of the site name, specific fields to check for in the URI,
// datatypes for the fields in a JSON file conf file
package main

import (
	"fmt"
	"os"

	"github.com/MJ-NJ/DataCollector/config"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "scraper",
		Usage: "Configurable data scraping engine",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "path",
				Usage:    "provide the file path of the input here",
				Required: false,
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "use this when need to print the site",
			},
			&cli.BoolFlag{
				Name:  "easter-egg",
				Usage: "this is an easter egg",
			},
		},
		Action: func(ctx *cli.Context) error {

			if ctx.Bool("easter-egg") {
				fmt.Println("This is a easter egg")
				return nil
			}

			configPath := ctx.String("path")
			var conf *config.ScrapeConfig
			var err error

			if configPath == "" {
				conf, err = config.LoadConfig("config.json") // This is the default val if string not present
				if err != nil {
					return err
				}
			} else {
				conf, err = config.LoadConfig(configPath) // Now this will take any path needed
				if err != nil {
					return err
				}
			}

			if ctx.Bool("dry-run") {
				fmt.Printf("[Dry Run] Would scrape the site: %s", conf.URL)
			}

			fmt.Printf("\nstarted scraping")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("error: %v", err)
	}
}
