package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// PatternType defines the type of pattern segment
type PatternType string

const (
	Alpha        PatternType = "alpha"        // Letters only
	Numeric      PatternType = "numeric"      // Numbers only
	Alphanumeric PatternType = "alphanumeric" // Letters and numbers
)

// Pattern defines a URI pattern for enumeration
type Pattern struct {
	Template string
	Segments []Segment
}

// Segment defines a variable segment in the URI pattern
type Segment struct {
	Name  string
	Type  PatternType
	Start int
	End   int
}

// ParsePattern parses a URI pattern string
// Example: "https://example.com/{abc:alpha}/{123:numeric}/{id:alphanumeric}"
func ParsePattern(pattern string) (*Pattern, error) {
	if pattern == "" {
		return nil, fmt.Errorf("pattern cannot be empty")
	}

	p := &Pattern{
		Template: pattern,
		Segments: []Segment{},
	}

	// Match {name:type} patterns
	re := regexp.MustCompile(`\{([^:]+):([^}]+)\}`)
	matches := re.FindAllStringSubmatchIndex(pattern, -1)

	for _, match := range matches {
		if len(match) < 6 {
			continue
		}

		name := pattern[match[2]:match[3]]
		typeStr := pattern[match[4]:match[5]]

		segmentType, err := parsePatternType(typeStr)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern type '%s' for segment '%s': %w", typeStr, name, err)
		}

		p.Segments = append(p.Segments, Segment{
			Name:  name,
			Type:  segmentType,
			Start: match[0],
			End:   match[1],
		})
	}

	return p, nil
}

// parsePatternType converts string to PatternType
func parsePatternType(typeStr string) (PatternType, error) {
	switch strings.ToLower(typeStr) {
	case "alpha":
		return Alpha, nil
	case "numeric":
		return Numeric, nil
	case "alphanumeric":
		return Alphanumeric, nil
	default:
		return "", fmt.Errorf("unknown pattern type: %s", typeStr)
	}
}

// Generate creates possible URIs from the pattern
// Note: This is a simplified version to avoid generating excessive URLs
func (p *Pattern) Generate(maxPerSegment int) []string {
	if len(p.Segments) == 0 {
		return []string{p.Template}
	}

	// TODO: Implement enumeration logic
	return []string{p.Template}
}

// ValidateURL checks URL validity
func ValidateURL(urlStr string) error {
	if urlStr == "" {
		return fmt.Errorf("URL cannot be empty")
	}

	if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
		return fmt.Errorf("URL must start with http:// or https://")
	}

	return nil
}
