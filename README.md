# GPT Jarvis CLI

The Jarvis CLI is a versatile command-line tool powered by `ChatGPT` that brings an interactive assistant to your terminal. It can help you with a wide range of tasks, from managing your operating system to working with PostgreSQL databases.

<!-- TODO: it designed to do complex tasks -->

## Installation

You can install the Jarvis CLI using the following command:

```bash
go install github.com/mreza0100/jarvis
```

## Getting Started

Before you can start using Jarvis, you'll need an OpenAI API key. You can obtain one from the OpenAI platform. Once you have your API key, you can use Jarvis for various tasks.

You can obtain the API key from [here](https://platform.openai.com/account/api-keys)

## Commands

### OS assistante:

```
jarvis [interactive/i] os ./.jarvisrc.json # will create config file if not exist
```
![Example-1](https://github.com/mreza0100/jarvis/assets/59161329/890c6884-e087-4d8e-ae7a-c64e7706cb7a)
![Example-2](https://github.com/mreza0100/jarvis/assets/59161329/10aac14a-f889-4357-a60b-1b6c84f4ec04)



### Postgres assistante:

```![Uploading Screenshot from 2023-10-09 01-32-09.png…]()

jarvis [interactive/i] [postgres/pgs] ./.jarvisrc.json # will create config file if not exist
```

## Config File (.jarvisrc.json)

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
				"value": "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"
			},
			"model": "gpt-4",
			"temperature": 1
		}
	}
}
```
