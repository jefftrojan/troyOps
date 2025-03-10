# Contributing to TroyOps

Thank you for your interest in contributing to TroyOps! This document provides guidelines and instructions for contributing to this project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
  - [Development Environment Setup](#development-environment-setup)
  - [Project Structure](#project-structure)
- [How to Contribute](#how-to-contribute)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Enhancements](#suggesting-enhancements)
  - [Pull Requests](#pull-requests)
- [Development Workflow](#development-workflow)
  - [Branching Strategy](#branching-strategy)
  - [Commit Messages](#commit-messages)
  - [Testing](#testing)
- [Style Guidelines](#style-guidelines)
  - [Go Code Style](#go-code-style)
  - [Documentation Style](#documentation-style)
- [Community](#community)

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [jeff@example.com](mailto:jeff@example.com).

## Getting Started

### Development Environment Setup

1. **Prerequisites**:
   - Go 1.24 or higher
   - kubectl
   - Flux CLI
   - Helm
   - Kustomize
   - Docker (for building container images)

2. **Clone the repository**:
   ```bash
   git clone https://github.com/jefftrojan/troyops.git
   cd troyops
   ```

3. **Install dependencies**:
   ```bash
   go mod download
   ```

4. **Build the project**:
   ```bash
   ./troyops.sh
   ```

### Project Structure

- `cmd/`: Contains the main application entry point
- `ci/`: CI/CD integration code
- `flux/`: Flux CD integration code
- `kustomize/`: Kustomize deployment code
- `policies/`: Policy enforcement code
- `secrets/`: Secret management code
- `charts/`: Helm charts for TroyOps
- `.github/`: GitHub-specific files (issue templates, workflows)

## How to Contribute

### Reporting Bugs

1. Check if the bug has already been reported by searching the [Issues](https://github.com/jefftrojan/troyops/issues).
2. If the bug hasn't been reported, [create a new issue](https://github.com/jefftrojan/troyops/issues/new?template=bug_report.md) using the bug report template.
3. Provide as much detail as possible, including steps to reproduce, expected behavior, and your environment.

### Suggesting Enhancements

1. Check if the enhancement has already been suggested by searching the [Issues](https://github.com/jefftrojan/troyops/issues).
2. If the enhancement hasn't been suggested, [create a new issue](https://github.com/jefftrojan/troyops/issues/new?template=feature_request.md) using the feature request template.
3. Clearly describe the enhancement, its benefits, and potential implementation approach.

### Pull Requests

1. Fork the repository.
2. Create a new branch from `main` for your changes.
3. Make your changes, following the [style guidelines](#style-guidelines).
4. Add or update tests as necessary.
5. Update documentation as necessary.
6. Submit a pull request to the `main` branch.
7. Ensure the PR description clearly describes the changes and references any related issues.

## Development Workflow

### Branching Strategy

- `main`: The main development branch. All PRs should target this branch.
- `feature/*`: For new features.
- `bugfix/*`: For bug fixes.
- `docs/*`: For documentation changes.
- `release/*`: For release preparation.

### Commit Messages

We follow the [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages:

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation changes
- `style`: Changes that do not affect the meaning of the code (formatting, etc.)
- `refactor`: Code changes that neither fix a bug nor add a feature
- `test`: Adding or updating tests
- `chore`: Changes to the build process or auxiliary tools

Examples:
- `feat(flux): Add multi-cluster support for Flux CD integration`
- `fix(kustomize): Resolve timeout issue with large manifests`
- `docs: Update secret management documentation`

### Testing

- Write unit tests for all new code.
- Ensure all tests pass before submitting a PR.
- Run tests with:
  ```bash
  go test ./...
  ```

## Style Guidelines

### Go Code Style

- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments).
- Use `gofmt` to format your code.
- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines.
- Document all exported functions, types, and constants.

### Documentation Style

- Use Markdown for documentation.
- Keep documentation up-to-date with code changes.
- Use clear, concise language.
- Include examples where appropriate.

## Community

- coming soon[Todo]

Thank you for contributing to TroyOps! 