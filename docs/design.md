# Basic Design
----

## Server functionality
1. Conf to setup which data to scrape and find datapoints publically available
2. Background worker to scrape the data
3. Setup 3 types of databases which store the data
4. Prefer KV stores for faster as compared to SQL dbs
5. Start with single worker and then expand to concurrently doing this, using a concurrent ACID db

## Modules

### Config Parser
Read the JSON/YAML config which will be of this type and contain site name, mode, uri for template, depth of search and rate limit

```json
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

```

### Scraper Engine/ Server
Uses Go routines to explore URIs, parse HTML/PDF/Excel

### Discovery Engine
Enumerate URIs based on the mode to discover routes or Tree based searching

### Storage Layer
Hybrid SQL + NoSQL based for Efficient storage as well as efficient lookups With Redis for caching

### Worker Pool
Initially 1, then move on to a set of N worker pool concurrently aggregating data

### Scheduler
Cron based/ type scheduler for the processes

### Metrics and Logging 
Prometheus + structured logs for observability


// We can go 2 ways,
// 1. search in same name it www.website.com/abc/123/abce235asw35(:id)
// --> So here we give the website like this, so we enumerate letters in abc, nums in 123, strings in id
// 2. Other way is to Just give a website like www.website.com
// --> Have a search depth, try all possible things from most common uri words and try again
// till we keep hitting, like a tree and backtrack when needed.
