# Jarvis CLI Powered by ChatGPT

The Jarvis CLI is a versatile command-line tool powered by ChatGPT that brings an interactive assistant to your terminal. It can help you with a wide range of tasks, from managing your operating system to working with PostgreSQL databases.

## Installation

You can install the Jarvis CLI using the following command:

```bash
go install github.com/mreza0100/jarvis
```

## Getting Started

Before you can start using Jarvis, you'll need an OpenAI API key. You can obtain one from the OpenAI platform. Once you have your API key, you can use Jarvis for various tasks.

### Interactive OS Assistance

Run the following command to start an interactive session with Jarvis for general OS assistance:

```bash
OPEN_API_KEY="<Your-Token>"
jarvis interactive os
```

### Interactive PostgreSQL Assistant

To interact with Jarvis for PostgreSQL tasks, use the following command. If a configuration file doesn't exist, Jarvis will create a template for you to fill out:

```bash
OPEN_API_KEY="<Your-Token>"
jarvis interactive postgres ./.jarvis.json
```

## Usage Examples

Here are some examples of tasks you can perform with Jarvis:

### Operating System Assistance

- Manage your files and directories.
- Execute system commands and scripts.
- Retrieve information about your system.

### PostgreSQL Interaction

- Execute SQL queries on your PostgreSQL database.
- Retrieve and manage database records.
- Perform database maintenance tasks.

---

**Note:** Please make sure to replace `<Your-Token>` with your actual OpenAI API key when using Jarvis.

For any issues, questions, or suggestions, feel free to [create an issue](https://github.com/mreza0100/jarvis/issues) on GitHub.
