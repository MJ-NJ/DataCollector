package collector

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Collector handles web crawling and data collection
type Collector struct {
	BaseURL    string
	FileTypes  []string
	MaxDepth   int
	Client     *http.Client
	visited    map[string]bool
}

// New creates a new Collector instance
func New(baseURL string, fileTypes []string, maxDepth int) (*Collector, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("baseURL cannot be empty")
	}

	// Validate URL format
	_, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	return &Collector{
		BaseURL:   baseURL,
		FileTypes: fileTypes,
		MaxDepth:  maxDepth,
		Client:    &http.Client{},
		visited:   make(map[string]bool),
	}, nil
}

// Collect starts the data collection process
func (c *Collector) Collect() error {
	fmt.Printf("Starting collection from: %s\n", c.BaseURL)
	fmt.Printf("Looking for file types: %v\n", c.FileTypes)
	fmt.Printf("Max depth: %d\n", c.MaxDepth)
	
	return c.crawl(c.BaseURL, 0)
}

// crawl recursively processes URLs
func (c *Collector) crawl(targetURL string, depth int) error {
	// Stop if max depth exceeded
	if depth > c.MaxDepth {
		return nil
	}

	// Skip if already visited
	if c.visited[targetURL] {
		return nil
	}

	// Mark as visited
	c.visited[targetURL] = true

	fmt.Printf("Crawling [depth=%d]: %s\n", depth, targetURL)

	// Fetch the URL
	resp, err := c.Client.Get(targetURL)
	if err != nil {
		return fmt.Errorf("failed to fetch URL %s: %w", targetURL, err)
	}
	defer resp.Body.Close()

	// Check if this file should be downloaded
	if c.shouldDownload(targetURL) {
		fmt.Printf("Found matching file: %s\n", targetURL)
		// TODO: Implement download logic
	}

	// TODO: Parse HTML and extract links for further crawling
	// This will be implemented in future updates

	return nil
}

// shouldDownload checks if URL matches desired file types
func (c *Collector) shouldDownload(targetURL string) bool {
	if len(c.FileTypes) == 0 {
		return false
	}

	targetURL = strings.ToLower(targetURL)
	for _, fileType := range c.FileTypes {
		if strings.HasSuffix(targetURL, "."+strings.ToLower(fileType)) {
			return true
		}
	}
	return false
}

// Download saves a file from the given URL to a writer
func (c *Collector) Download(targetURL string, writer io.Writer) error {
	resp, err := c.Client.Get(targetURL)
	if err != nil {
		return fmt.Errorf("failed to download %s: %w", targetURL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write data: %w", err)
	}

	return nil
}
