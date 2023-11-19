# Commented README.md for GPT Jarvis CLI

The GPT Jarvis CLI is a powerful command-line interface powered by ChatGPT, designed for sophisticated tasks in your terminal.

## Installation

To install the Jarvis CLI, use the following command:

```bash
go install github.com/mreza0100/jarvis
```

## Getting Started

### Step 1: Acquiring an OpenAI API Token

First, obtain an OpenAI API key from the OpenAI platform. This key is crucial for accessing Jarvis's functionalities.

### Step 2: Configuring Jarvis

#### Config File Setup

To set up the configuration file (.jarvisrc.json), run the `config bootstrap` command. Use the `--trunk` flag to overwrite an existing configuration:

```bash
jarvis config bootstrap [--trunk]
```

### Edit configuration

`nano ~/.jarvis/.jarvisrc.json`

This command creates a default configuration file in the `.jarvis` directory under your home path. Modify the `.jarvisrc.json` file to include your specific settings, such as the API token and connection details.

A typical configuration file looks like this:

````json
{
    "postgres": {
        "config": {
            "token": {
                "env_name": "OPENAI_API_KEY"
            },
            "model": "gpt-4",
            "temperature": 1
        },
        "connection": {
            "host": "localhost",
            "port": 5432,
            "username": "mamad",
            "password": "mamadspass",
            "database": "mamad_db"
        }
    },
    "os": {
        "config": {
            "token": {
                "env_name": "OPENAI_API_KEY"
            },
            "model": "gpt-4",
            "temperature": 1
        }
    }
}

### Step 3: Running Jarvis

#### OS Assistant

To use the OS assistant, execute:

```bash
jarvis interactive os
````

#### PostgreSQL Assistant

For the PostgreSQL assistant, run:

```bash
jarvis interactive postgres
```

## Command Overview

- `interactive`: Launches Jarvis in interactive mode with OS or PostgreSQL assistants.
- `config bootstrap`: Initializes and sets up the configuration file. Use `--trunk` to overwrite existing configurations.

Ensure your API token and connection details are correctly set in `.jarvisrc.json` before using Jarvis.
