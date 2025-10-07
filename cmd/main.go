package main

import "github.com/MJ-NJ/DataCollector/modules/parser"

// This will house the main server which will control the server
// This will get input of the site name, specific fields to check for in the URI,
// datatypes for the fields in a JSON file conf file
// We can go 2 ways,
// 1. search in same name it www.website.com/abc/123/abce235asw35(:id)
// --> So here we give the website like this, so we enumerate letters in abc, nums in 123, strings in id
// 2. Other way is to Just give a website like www.website.com
// --> Have a search depth, try all possible things from most common uri words and try again
// till we keep hitting, like a tree and backtrack when needed.
func loadConfig(path string) (*parser.ScrapeConfig, error) {
	// Locate file
	// Read file
	// Convert to struct
	// Return

	return &parser.ScrapeConfig{}, nil
}
func main() {

	// Take input from CLIENT REQ or CLI
	// This is a JSON File

}
