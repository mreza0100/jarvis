# GPT Jarvis CLI

The GPT Jarvis CLI is a versatile command-line tool powered by ChatGPT, designed for complex tasks in your terminal.


## Installation

You can easily install the Jarvis CLI using the following command:

```bash
go install github.com/mreza0100/jarvis
```

# Getting Started

## Step 1: Obtaining Your API Token
Before diving into Jarvis's capabilities, you'll need to obtain an OpenAI API key. Simply head over to the OpenAI platform to acquire your key. This key is essential for unlocking the full potential of Jarvis.

## Step 2: Creating a Configuration File
To use Jarvis, you need to create a configuration file named .jarvisrc.json. Here's an example of what it should look like:

```
{
	"postgres": {
		"config": {
			"token": {
				"env_name": "OPEN_API_KEY"
			},
			"model": "gpt-3.5-turbo-16k",
			"temperature": 1
		},
		"connection": {
			"host": "localhost",
			"port": 5432,
			"username": "username",
			"password": "password",
			"database": "databasename"
		}
	},
	"os": {
		"config": {
			"token": {
				"value": "sk-xxxx"
			},
			"model": "gpt-4",
			"temperature": 1
		}
	}
}
```
In this file, you can customize settings for both the OS and PostgreSQL assistants, including the API token you obtained in Step 1.


## Step 3: Running Jarvis Commands
Now that you have your API token and configuration file set up, you're ready to start using Jarvis. Here are a few example commands:

### OS Assistant:
To engage the OS assistant, execute the following command:

```
jarvis interactive os ./.jarvisrc.json # This command creates a config file if it doesn't exist
```


```
1 Welcome!
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

I am Jarvis, your go-to system administrator. How can I assist you today?


==> get all docker containers and then perfrom a health-check on each of them
Tokens Used: 5472


Sure, I can help you with that.

To get all Docker containers, you can use the following command: docker ps -a

Tokens Used: 7698


Here are the Docker containers:

┃ CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS                     PORTS                   NAMES
┃ 43b8c8e0b8fe   mongo           "docker-entrypoint.s…"   8 minutes ago   Up 7 minutes               27017/tcp               mongo-container
┃ 242a02f728a7   ubuntu          "/bin/bash"              8 minutes ago   Exited (0) 8 minutes ago                           ubuntu-container
┃ 50ec36fb5aed   postgres        "docker-entrypoint.s…"   8 minutes ago   Exited (1) 8 minutes ago                           pg-container
┃ 6e68edb0a1bc   postgres:15.4   "docker-entrypoint.s…"   3 weeks ago     Up 3 weeks (healthy)       0.0.0.0:5432->5432/tcp  pg-jarvis-test

Tokens Used: 8563


To perform a health check on each Docker container, we will need to inspect the containers and check their status. Here's a script to perform the health check:

┃ #!/bin/bash
┃ 
┃ containers=$(docker ps -a --format "{{.ID}}")
┃ 
┃ for container in $containers
┃ 
┃ do
┃ docker inspect --format='{{.State.Health.Status}}' $container | grep -q "healthy"
┃ 
┃ if [ $? -eq 0 ]; then
┃     echo "Container $container is healthy."
┃ else
┃     echo "Container $container is unhealthy."
┃ fi
┃ done

Tokens Used: 11246

Here is the result of the health check:

┃ Container 43b8c8e0b8fe is unhealthy.
┃ Container 242a02f728a7 is unhealthy.
┃ Container 50ec36fb5aed is unhealthy.
┃ Container 6e68edb0a1bc is healthy.
```

### PostgreSQL Assistant:
To utilize the PostgreSQL assistant, execute the following command:
```
jarvis interactive postgres ./.jarvisrc.json # This command creates a config file if it doesn't exist
```
![](https://github-production-user-asset-6210df.s3.amazonaws.com/59161329/273481379-34f646c9-9b8c-4d8c-a769-0b12d1ec0340.png)

Unlock the potential of your terminal with the GPT Jarvis CLI!
