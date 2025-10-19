
# CONTRIBUTING.md

## Contributing to Cache Cleaner

First off, thank you for considering contributing to Cache Cleaner! It's people like you that make it such a great tool.

### Code of Conduct

This project adheres to the Contributor Covenant [code of conduct](https://www.contributor-covenant.org/). By participating, you are expected to uphold this code.

### How Can I Contribute?

#### Reporting Bugs

Before creating bug reports, please check the issue list as you might find out that you don't need to create one. When you are creating a bug report, please include as many details as possible:

* **Use a clear and descriptive title**
* **Describe the exact steps which reproduce the problem**
* **Provide specific examples to demonstrate the steps**
* **Describe the behavior you observed after following the steps**
* **Explain which behavior you expected to see instead**
* **Include screenshots if possible**
* **Include your environment details** (OS, Go version, etc.)

#### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

* **Use a clear and descriptive title**
* **Provide a step-by-step description of the suggested enhancement**
* **Provide specific examples to demonstrate the steps**
* **Describe the current behavior and what you want instead**
* **Explain why this enhancement would be useful**

#### Pull Requests

* Fill in the required template
* Follow the Go styleguides
* Include appropriate test cases
* End all files with a newline
* Avoid platform-specific code when possible

### Styleguides

#### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

#### Go Styleguide

* Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
* Use `go fmt` for formatting
* Use meaningful variable names
* Write comments for exported functions
* Keep functions focused and testable

#### Documentation Styleguide

* Use Markdown
* Reference function names in backticks: \`Scan()\`
* Add code examples where relevant

### Testing

* Write tests for new features
* Ensure all tests pass: `go test ./...`
* Maintain or improve code coverage
* Test on multiple platforms when possible

### Development Setup

1. Fork and clone the repository
2. Create a feature branch: `git checkout -b my-new-feature`
3. Make your changes
4. Add/update tests as needed
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -am 'Add new feature'`
7. Push to the branch: `git push origin my-new-feature`
8. Submit a Pull Request

### Additional Notes

* This project uses semantic versioning
* All contributions are subject to the MIT License
