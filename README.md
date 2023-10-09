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
![](https://github-production-user-asset-6210df.s3.amazonaws.com/59161329/273481589-1a33296d-76ef-4509-8eed-bf8f0337a1b7.png)

### PostgreSQL Assistant:
To utilize the PostgreSQL assistant, execute the following command:
```
jarvis interactive postgres ./.jarvisrc.json # This command creates a config file if it doesn't exist
```
![](https://github-production-user-asset-6210df.s3.amazonaws.com/59161329/273481379-34f646c9-9b8c-4d8c-a769-0b12d1ec0340.png)

Unlock the potential of your terminal with the GPT Jarvis CLI!
