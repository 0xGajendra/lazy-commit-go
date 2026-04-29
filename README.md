# lazy-commit-go - AI-Powered Git Commit & Push Automation CLI

A modern CLI tool that automates the entire git workflow: staging files, generating commit messages with AI, committing, and pushing to remote - all through an interactive interface.

## What is lazy-commit-go?

lazy-commit-go is a command-line tool that simplifies git workflows by automating the add, commit, and push process with intelligent commit message suggestions powered by Groq's AI. Perfect for developers who want faster, consistent commits without the mental overhead of writing commit messages.

## Use Cases

- Quickly stage and commit changes without typing git commands
- Generate professional, conventional commit messages automatically
- Automate git add + commit + push in one command
- Learn git workflow through interactive CLI
- Speed up daily development routine

## Key Features

- Interactive file selection - choose which files to stage
- AI-powered commit messages - Groq API generates 3 smart suggestions
- Edit before commit - modify selected message freely
- Auto-push option - push to remote after commit
- Git repository detection - auto-initialize if not a repo
- API key storage - save Groq key for future use

## Installation

### Option 1: Download Binary (No Go Required)

1. Visit the [Releases Page](https://github.com/0xGajendra/lazy-commit-go/releases)
2. Download the compressed file for your operating system:
   - Windows: `lazy-commit-go_Windows_x86_64.zip`
   - macOS: `lazy-commit-go_Darwin_x86_64.tar.gz` (Intel) or `_arm64.tar.gz` (Apple Silicon)
   - Linux: `lazy-commit-go_Linux_x86_64.tar.gz`
3. Extract the file
4. Run the executable

### Option 2: Go Install

```bash
go install github.com/0xGajendra/lazy-commit-go@latest
```

Add to PATH if needed: `export PATH=$PATH:$HOME/go/bin`

## Requirements

- Git installed and configured
- Groq API key (free at [groq.com](https://groq.com))
- Terminal with color support

## Quick Start

```bash
# Run from any git repository
lazy-commit-go

# Follow the interactive prompts:
# 1. Select files to stage
# 2. Choose AI-generated commit message
# 3. Edit message if needed
# 4. Confirm commit
# 5. Optional: push to remote
```

## How It Works

1. **Repository Check** - Detects if current directory is a git repo, offers to initialize if not
2. **File Selection** - Shows all changed files, lets you select which to stage
3. **Diff Generation** - Creates git diff of staged changes
4. **AI Generation** - Sends diff to Groq API, receives 3 commit message suggestions
5. **Selection** - Pick your preferred message, edit if needed
6. **Commit** - Executes `git commit`
7. **Push** - Optional push to remote branch

## Tech Stack

- Go 1.22+
- Groq API (LLM for commit message generation)
- promptui (interactive CLI prompts)
- survey (multi-select file selection)

## API Key

On first run, you'll be prompted to enter your Groq API key. It's stored in `~/.lzc_config` for future use. Get a free key at [groq.com](https://groq.com).

## Alternatives

Looking for similar tools? Search for:
- git commit assistant
- AI git commit generator
- git workflow automation CLI
- auto git commit tool
- conventional commit generator

## License

MIT License

## Contributing

Contributions welcome! Open issues and pull requests on [GitHub](https://github.com/0xGajendra/lazy-commit-go).

## Related Tools

- commitizen - conventional commit CLI
- gitmoji - emoji in commit messages
- cz-cli - commitizen CLI
- conventional-changelog - commit message tools