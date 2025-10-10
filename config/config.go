// Parse JSON/YAML configs, validate fields, and prepare job definitions.

package config

/*
{
  "site": "www.website.com",
  "mode": "enumerate", // or "tree"
  "uri_template": "/abc/{num}/{id}",
  "fields": {
    "num": "int",
    "id": "string"
  },
  "search_depth": 3,
  "rate_limit": "5req/s"
}
*/

type ScrapeConfig struct {
	URL          string            `json:"url" yaml:"url" validate:"required,url"`
	Mode         string            `json:"mode" yaml:"mode" validate:"required,oneof=enumerate tree"`
	URI_Template string            `json:"uri_template" yaml:"uri_template"`
	Fields       map[string]string `json:"fields" yaml:"fields"`             // This will hold like template ka fields mapped to type
	Search_depth int               `json:"search_depth" yaml:"search_depth"` // For connected hyperlinks
	Rate_Limit   int               `json:"rate_limit" yaml:"rate_limit"`     // This will keep the number of calls to a server below this count in it is the number of req/sec
}
