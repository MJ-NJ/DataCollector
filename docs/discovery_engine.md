# Discovery Engine Module
Enumerate URIs based on the mode to discover routes or Tree based searching


## Mode 1 - URI Enumeration
1. You define a URI template

2. System generates permutations (e.g., /abc/001/xyz, /abc/002/abc)

3. Scrapes if response is valid


## Mode 2 - Tree based
1. Start from www.website.com

2. Explore recursively to depth N

3. Backtrack and hit routes

4. Further TODO: Uses heuristics (e.g., common keywords, link density) to prioritize paths