# DataCollector

A web data collector written in Go that enumerates websites from URIs and downloads specific file types.

## Overview

DataCollector intelligently crawls websites and collects data based on configurable patterns and file types. It supports two main modes of operation:

1. **Pattern-based enumeration**: Specify URI patterns with variable segments
2. **Intelligent crawling**: Automated discovery with configurable depth

## Features (Planned)

- **URI Pattern Enumeration**: Define patterns like `website.com/abc/123/id` and automatically enumerate combinations
- **Tree-based Crawling**: Depth-first search with backtracking
- **File Type Filtering**: Download specific file types (PDF, images, documents, etc.)
- **JSON Configuration**: Easy configuration through JSON files
- **Concurrent Processing**: Parallel processing of multiple URIs
- **Progress Tracking**: Real-time progress monitoring

## Installation

### Prerequisites

- Go 1.23.5 or higher
- Git

### Clone the Repository

```bash
git clone https://github.com/MJ-NJ/DataCollector.git
cd DataCollector
```

### Install Dependencies

```bash
go mod download
```

### Build

```bash
go build -o datacollector ./cmd
```

## Usage

### Basic Usage

```bash
./datacollector --url https://example.com --depth 2 --filetypes pdf,jpg,png
```

### Pattern-based Enumeration

```bash
# Enumerate letters in 'abc', numbers in '123', and strings in 'id'
./datacollector --pattern "https://example.com/abc/123/id"
```

### Configuration File

Create a `config.json`:

```json
{
  "target": "https://example.com",
  "depth": 3,
  "filetypes": ["pdf", "docx", "jpg"],
  "patterns": {
    "abc": "alpha",
    "123": "numeric",
    "id": "alphanumeric"
  }
}
```

Then run:

```bash
./datacollector --config config.json
```

## Roadmap

### Phase 1: Core Functionality
- [ ] Implement basic HTTP client
- [ ] Add URI pattern parser
- [ ] Create enumeration engine
- [ ] Implement file type detection

### Phase 2: Advanced Features
- [ ] Add concurrent processing
- [ ] Implement smart crawling with backtracking
- [ ] Add rate limiting and politeness delays
- [ ] Create configuration file parser

### Phase 3: Enhancement
- [ ] Add logging and error handling
- [ ] Implement progress bars and statistics
- [ ] Add export formats (CSV, JSON)
- [ ] Create web UI dashboard

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to get started.

This project is participating in **Hacktoberfest 2025**.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Authors

- **MJ-NJ** - *Initial work* - [MJ-NJ](https://github.com/MJ-NJ)

## Acknowledgments

Thanks to all contributors who help improve this project.

## Contact

For questions or suggestions, please open an issue on GitHub.
