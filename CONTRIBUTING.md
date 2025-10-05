# Contributing to DataCollector

Thank you for considering contributing to DataCollector.

## Hacktoberfest

This project is participating in **Hacktoberfest 2025**. We welcome quality contributions from developers of all skill levels.

### Hacktoberfest Guidelines

- Make meaningful contributions, not spam
- Follow the project's coding standards
- Test your changes before submitting
- Be respectful and constructive in discussions

## Code of Conduct

By participating in this project, you agree to:
- Be respectful and inclusive
- Welcome newcomers and beginners
- Give and receive constructive feedback gracefully
- Focus on what's best for the community

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the problem
- **Expected behavior** vs. **actual behavior**
- **Go version** and **OS** you're using
- **Code samples** or test cases if applicable

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the suggested enhancement
- **Explain why this enhancement would be useful** to most users
- **List any similar features** in other tools, if applicable

### Your First Code Contribution

Unsure where to begin? Look for issues labeled:

- `good first issue` - Good for newcomers
- `help wanted` - Extra attention needed
- `hacktoberfest` - Hacktoberfest-friendly issues

### Pull Requests

Follow these steps for contributing code:

1. **Fork the repository** and create your branch from `main`
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR-USERNAME/DataCollector.git
   cd DataCollector
   ```

3. **Create a new branch** for your feature:
   ```bash
   git checkout -b feature/amazing-feature
   ```

4. **Make your changes** and commit:
   ```bash
   git add .
   git commit -m "Add some amazing feature"
   ```

5. **Push to your fork**:
   ```bash
   git push origin feature/amazing-feature
   ```

6. **Open a Pull Request** against the `main` branch

## Development Setup

### Prerequisites

- Go 1.23.5 or higher
- Git

### Setting Up Your Development Environment

```bash
# Clone your fork
git clone https://github.com/YOUR-USERNAME/DataCollector.git
cd DataCollector

# Add upstream remote
git remote add upstream https://github.com/MJ-NJ/DataCollector.git

# Install dependencies
go mod download

# Run tests (when available)
go test ./...

# Build the project
go build -o datacollector ./cmd
```

### Keeping Your Fork Updated

```bash
git fetch upstream
git checkout main
git merge upstream/main
git push origin main
```

## Coding Standards

### Go Style Guide

- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Run `go vet` to catch common mistakes
- Write meaningful variable and function names
- Add comments for exported functions and types

### Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally

**Good examples:**
```
Add URI pattern parser
Fix crash when processing empty URLs
Update README with usage examples
```

### Code Organization

```
DataCollector/
├── cmd/           # Main applications
├── pkg/           # Public libraries
├── internal/      # Private application code
├── test/          # Additional test data
└── docs/          # Documentation files
```

## Testing

- Write unit tests for new features
- Ensure all tests pass before submitting PR
- Aim for meaningful test coverage
- Add integration tests for complex features

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/parser
```

## Documentation

- Update README.md if you change functionality
- Add comments to exported functions and types
- Update examples if you add new features
- Consider adding a doc.go file for packages

## Issue and PR Labels

We use labels to organize issues and PRs:

- `bug` - Something isn't working
- `enhancement` - New feature or request
- `documentation` - Documentation improvements
- `good first issue` - Good for newcomers
- `help wanted` - Extra attention needed
- `hacktoberfest` - Hacktoberfest eligible
- `priority: high` - High priority items
- `priority: low` - Nice to have features

## Areas for Contribution

We especially welcome contributions in:

1. **Core Functionality**
   - HTTP client implementation
   - URI pattern parsing
   - File type detection
   - Enumeration algorithms

2. **Testing**
   - Unit tests
   - Integration tests
   - Test utilities

3. **Documentation**
   - Code comments
   - Usage examples
   - Tutorials
   - API documentation

4. **Features**
   - Concurrent processing
   - Rate limiting
   - Progress tracking
   - Export formats

5. **DevOps**
   - CI/CD pipelines
   - Docker support
   - Build automation

## Questions?

Don't hesitate to ask questions. You can:

- Open an issue with the `question` label
- Comment on existing issues or PRs
- Reach out to maintainers

## Thank You

Your contributions to open source make projects like this possible. Thank you for taking the time to contribute.
