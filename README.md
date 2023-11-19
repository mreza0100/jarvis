# Commented README.md for GPT Jarvis CLI

The GPT Jarvis CLI is a powerful command-line interface powered by ChatGPT, designed for sophisticated tasks in your terminal.

## Installation

To install the Jarvis CLI, use the following command:

```bash
go install github.com/mreza0100/jarvis@latest
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

```json
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
				"value": "sk-xxx"
			},
			"model": "gpt-4",
			"temperature": 1
		}
	}
}
```

### Step 3: Running Jarvis

#### OS Assistant

To use the OS assistant, execute:

```bash
jarvis interactive os
```

#### PostgreSQL Assistant

For the PostgreSQL assistant, run:

```bash
jarvis interactive postgres
```

## Command Overview

- `interactive`: Launches Jarvis in interactive mode with OS or PostgreSQL assistants.
- `config bootstrap`: Initializes and sets up the configuration file. Use `--trunk` to overwrite existing configurations.

Ensure your API token and connection details are correctly set in `.jarvisrc.json` before using Jarvis.

# Examples

## OS

```text
1 Welcome
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

1.1 What can I do for you today, Tony?



==> Do not ask for permissions, find all the docker containers and show them to me, then perform a healthcheck on each of them, and restart the unhealthy ones
Tokens Used: 5966


0.1 Performing Docker Containers Check

I will perform the following steps:
1. Get a list of all Docker containers.
2. Perform a health check on each container.
3. Restart any containers that are found to be unhealthy.

┃ # Script to execute
┃ docker ps -a


Tokens Used: 8652


0.1 Docker Containers List

Here is the list of all Docker containers:
1. Container ID: 43b8c8e0b8fe, Image: mongo, Status: Exited (255) 3 weeks ago, Name: mongo-container
2. Container ID: 242a02f728a7, Image: ubuntu, Status: Exited (0) 4 weeks ago, Name: ubuntu-container
3. Container ID: 50ec36fb5aed, Image: postgres, Status: Exited (1) 4 weeks ago, Name: pg-container
4. Container ID: 6e68edb0a1bc, Image: postgres:15.4, Status: Exited (0) About a minute ago, Name: pg-jarvis-test

Now, I will perform a health check on each container.

┃ # Script to execute
┃ docker inspect --format '{{.State.Health.Status}}' 43b8c8e0b8fe
┃ docker inspect --format '{{.State.Health.Status}}' 242a02f728a7
┃ docker inspect --format '{{.State.Health.Status}}' 50ec36fb5aed
┃ docker inspect --format '{{.State.Health.Status}}' 6e68edb0a1bc


Tokens Used: 10239


0.1 Docker Containers Health Status

Here is the health status of each Docker container:
1. Container ID: 43b8c8e0b8fe, Health Status: No health check configured
2. Container ID: 242a02f728a7, Health Status: No health check configured
3. Container ID: 50ec36fb5aed, Health Status: No health check configured
4. Container ID: 6e68edb0a1bc, Health Status: Unhealthy

The container with ID 6e68edb0a1bc is unhealthy. I will restart this container.

┃ # Script to execute
┃ docker restart 6e68edb0a1bc


Tokens Used: 10735


0.1 Docker Container Restarted

The Docker container with ID 6e68edb0a1bc has been successfully restarted.

Please let me know if there's anything else you need.
```

## Postgres

```text
1 Hello Tony
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

1.1 What can I assist you with today?



==> Do not ask for permissions. Fetch all the existing tables and there structure and show me, then perform a performance check on each table based on 'created_at' column, then optimize the query and run the performance check again to see the difference.
Tokens Used: 5760


0.0.1 Step 1: Fetching List Of Existing Tables And Their Structure


Tokens Used: 6635


Currently there are 2 tables in the database: users and posts. Next, we will fetch their structure.

Tokens Used: 8527


Here are the structures for the tables:

Table: Users
• ID (Integer)
• Created_at (Timestamp without time zone)
• Username (Text)
• Email (Text)

Table: Posts
• ID (Integer)
• User_id (Integer)
• Created_at (Timestamp without time zone)
• Title (Text)
• Content (Text)

Next, we will conduct a performance check for these tables based on the created_at column.

Tokens Used: 10606


The performance check results are as follows:

**Users Table: **
• Planning Time: 0.869 ms
• Execution Time: 0.127 ms

**Posts Table: **
• Planning Time: 0.150 ms
• Execution Time: 0.012 ms

To optimize, we will create an index on the created_at column for both tables.

Tokens Used: 11326


The indexing on created_at for both the users and posts tables completed successfully. Now let's run the performance check again to see the difference.

Tokens Used: 12947


After optimization, the performance check results are as follows:

**Users Table: **
• Planning Time: 0.775 ms
• Execution Time: 0.046 ms

**Posts Table: **
• Planning Time: 0.437 ms
• Execution Time: 0.009 ms

As we can see, creating an index on the created_at column for both tables has made our query more efficient.
```
