package main

// This will house the main server which will control the server
// This will get input of the site name, specific fields to check for in the URI,
// datatypes for the fields in a JSON file conf file
// We can go 2 ways,
// 1. search in same name it www.website.com/abc/123/abce235asw35(:id)
// --> So here we give the website like this, so we enumerate letters in abc, nums in 123, strings in id
// 2. Other way is to Just give a website like www.website.com
// --> Have a search depth, try all possible things from most common uri words and try again
// till we keep hitting, like a tree and backtrack when needed.



import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MJ-NJ/DataCollector/internal/config"
	"github.com/MJ-NJ/DataCollector/pkg/collector"
)

func main() {
	// Command line flags
	url := flag.String("url", "", "Target URL to crawl")
	depth := flag.Int("depth", 2, "Maximum crawl depth")
	configFile := flag.String("config", "", "Path to configuration file")
	filetypes := flag.String("filetypes", "", "Comma-separated list of file types to download (e.g., pdf,jpg,png)")
	
	flag.Parse()

	var cfg *config.Config
	var err error

	// Load config from file or use flags
	if *configFile != "" {
		cfg, err = config.Load(*configFile)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	} else {
		// Create config from flags
		if *url == "" {
			fmt.Println("Usage: datacollector -url <URL> [-depth <number>] [-filetypes <types>]")
			fmt.Println("   or: datacollector -config <config.json>")
			os.Exit(1)
		}

		cfg = &config.Config{
			Target:    *url,
			Depth:     *depth,
			FileTypes: parseFileTypes(*filetypes),
		}
	}

	// Create collector and start
	c, err := collector.New(cfg.Target, cfg.FileTypes, cfg.Depth)
	if err != nil {
		log.Fatalf("Failed to create collector: %v", err)
	}

	fmt.Println("DataCollector starting...")
	if err := c.Collect(); err != nil {
		log.Fatalf("Collection failed: %v", err)
	}

	fmt.Println("Collection completed successfully")
}

func parseFileTypes(input string) []string {
	if input == "" {
		return []string{}
	}
	
	var types []string
	current := ""
	for _, ch := range input {
		if ch == ',' {
			if current != "" {
				types = append(types, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		types = append(types, current)
	}
	
	return types
}

