# How does Jarvis works!

### System Architecture Overview

#### Hexagonal Architecture

The GPT Jarvis CLI system uses hexagonal architecture to promote separation of concerns, enhancing flexibility, maintainability, and scalability.

#### Components

1. **Adapters**:

      - **Role**: Bridge between the application and the external environment.
      - **Types**:
           - **Driven Adapters**: Interact with systems or services like PostgreSQL, OpenAI API, and file systems.
           - **Driving Adapters**: Handle requests from clients, including CLI or potential web interfaces.

2. **Models**:

      - **Role**: Represent domain logic and data structures.
      - **Types**:
           - **Data Models**: Structure of data like configuration settings and commands.
           - **Interactors**: Business logic layer processing user commands.

3. **Ports**:

      - **Role**: Define interfaces for adapter interaction with the application.
      - **Types**:
           - **Input Ports**: Interfaces for driving adapters.
           - **Output Ports**: Interfaces for driven adapters.

4. **Services**:
      - **Role**: Implement business logic and workflows.
      - **Types**:
           - **Configuration Service**: Manages system configurations.
           - **Interactive Services**: Manages interactive user sessions.

### Command History Management

The system stores command history on disk at `~/.jarvis/history/{uuid}`.

### Future Enhancements

1. **Retrieving Past History**: Implementing a feature to access and use past command history.
2. **Jarvis Developer**: Fine-tuning the GPT model with user codebases for better context-aware solutions. Enables file manipulation and command execution.
3. **Customizable Chat Names**: Users will be able to choose and reuse custom names for the chat interface.
4. **Binary installation**: A easy way to install Jarvis independently from Golang compiler
5. **GPT Chat Garbage Collector**: Optimizes chat history by filtering out irrelevant data, retaining useful information using the GPT API.

### OS Integration

The `runners` package includes `osRunner` for executing scripts. It uses the `exec.Command` method for running scripts and handling output, errors, and exit statuses.

### Dev Mode

The system employs a `MODE` environment variable to activate development logs for dynamic configuration adjustments and error logging.

Certainly! Here's the updated section of the `technical.md` document for the GPT Jarvis CLI system, focusing on the Communication Protocols for OS and PostgreSQL interactions, including the exact JSON models from the initiate GPT commands:

### Communication Protocols: OS and PostgreSQL Schemas

Certainly! Here's a cleaner and more readable version of the OS and PostgreSQL schemas for the `technical.md` document:

### Communication Protocols: OS and PostgreSQL Schemas

#### OS Schema

The OS interaction uses a JSON-based communication protocol. Below are the structured formats:

**OS Reply JSON Structure:**

```json
{
  "ReplyToUser": null | string,
  "WaitForUserPrompt": boolean,
  "Script": {
    "Runtime": string,
    "Script": string
  } | null
}
```

- `ReplyToUser`: Communicates with the user, outputting as markdown. Should be concise.
- `WaitForUserPrompt`: Set to true to request a user response.
- `Script`: Contains script to be executed, including the runtime environment and script text.

**OS Prompt JSON Structure:**

```json
{
  "ClientPrompt": null | string,
  "UserPrompt": null | string,
  "Screen": {
    "Width": number,
    "Height": number
  } | null,
  "LastScriptResult": {
    "Stdout": string,
    "Stderr": string,
    "StatusCode": int
  } | null
}
```

- `ClientPrompt`: Message from the client program to GPT.
- `UserPrompt`: Direct user prompt to GPT.
- `Screen`: Information about user's screen dimensions.
- `LastScriptResult`: Details of the last executed script, including output, errors, and status code.

#### PostgreSQL Schema

The PostgreSQL interaction also follows a JSON-based communication format:

**PostgreSQL Reply JSON Structure:**

```json
{
  "ReplyToUser": null | string,
  "WaitForUserPrompt": boolean,
  "QueryRequest": {
    "Query": string
  } | null
}
```

- `ReplyToUser`: Communicates with the user in markdown format.
- `WaitForUserPrompt`: True if the system should wait for user input.
- `QueryRequest`: Contains SQL query to be executed.

**PostgreSQL Prompt JSON Structure:**

```json
{
  "ClientPrompt": null | string,
  "UserPrompt": null | string,
  "Screen": {
    "Width": number,
    "Height": number
  } | null,
  "LastQueryResult": {
    "QueryResponses": [
      {
        "ColumnValues": [any],
        "Columns": [string],
        "ColumnType": [
          {
            "Name": string,
            "Length": int | null,
            "Nullable": bool | null,
            "DatabaseTypeName": string
          }
        ],
        "Err": error
      }
    ]
  } | null
}
```

- `ClientPrompt`: Message from the client program.
- `UserPrompt`: Direct user prompt.
- `Screen`: Screen size details.
- `LastQueryResult`: Results of the last executed query, including response details and any errors.

These JSON structures define the communication format for the GPT Jarvis CLI system when interacting with the OpenAI API for OS and PostgreSQL functionalities, ensuring clear and structured interactions.

### Challenges of Using OpenAI API

Challenges include managing API rate limits, balancing cost and usage, ensuring data privacy and security, dependency on internet connectivity, maintaining response quality, handling complex queries, and integration complexity.

This revised document reflects the latest developments and plans for the GPT Jarvis CLI system, providing a comprehensive overview of its architecture, features, and challenges.

## Diagram

```
┌──────────────────┐   ┌──────────────────────┐       ┌──────────────────┐     ┌───────────────────┐
│  .jarvisrc.json  │   │File Descriptors (FDs)│       │    OpenAI API    │     │ PostgresQL Server │
│                  │   │Stdin                 │       │                  │     │                   │
│                  │   │Stdout                │       │                  │     │                   │
│                  ◄─┐ │Stderr                ◄─┐     │                  ◄─┐   │                   ◄──────────┐
└──────────────────┘ │ └──────────────────────┘ │     └──────────────────┘ │   └───────────────────┘          │
                     │                          │                          │                                  │
                     └────┐                     │                          └────────────────────┐             │
                          │                     │                                               │             │
                          │                     └────────────────────────────────────────────┐  │             │
                          │                                                                  │  │             │
                          │                                                                  │  │             │
                          ├──────────────────────┐                  ┌──────────────────────┐ │  │             │
                          │    Config Provider   │                  │   Terminal Adapter   ├─┘  │             │
                          │                      │                  │                      │    │             │
                          │                      │                  │                      │    │             │
                          │                      ◄─────────┬────────►Renders Markdown      │    │             │
                          │Config Control        │         │        │Read                  │    │             │
                          │GPT Config            │         │        │Write                 │    │             │
                          │History Config        │         │        │Log controll          │    │             │
                          └──────────────────────┘         │        └──────────────────────┘    │             │
                                                           │                                    │             │
                          ┌──────────────────────┐         │        ┌──────────────────────┬────┘             │
                          │   History Adapter    │         │        │    Chat Adapter      │                  │
                          │                      │         │        │                      │                  │
                          │ Dump history on disk │         │        │  Interact with API   │                  │
                          │                      ◄─────────┼────────►                      │                  │
                          │                      │         │        │                      │                  │
                          │                      │         │        │                      │                  │
                          │                      │         │        │                      │                  │
                          └──────────────────────┘         │        └──────────────────────┘                  │
                                                         ┌─┴─┐                                                │
                                                         │   │                                                │
        ┌───────────────────────┐                        │   │                        ┌───────────────────────┤
        │    OS Script Runner   │                        │   │                        │ Postgres Query Runner │
        │                       │                        │   │                        │                       │
        │        Adapter        │                        │   │                        │       Adapter         │
        │                       │                        │   │                        │                       │
        │                       ◄──┐                     │   │                     ┌──►                       │
        │                       │  │                     │   │                     │  │                       │
        │                       │  │                     │   │                     │  │                       │
        │                       │  │                     │   │                     │  │                       │
        └───────────────────────┘  ├─────────────────────┤   ├─────────────────────┤  └───────────────────────┘
                                   │         OS          │   │      Postgres       │
                                   │                     │   │                     │
                                   │   Service Domain    │   │    Service Domain   │
                                   │                     │   │                     │
                                   │                     │   │                     │
                                   │                     │   │                     │
                                   │                     │   │                     │
                                   │                     │   │                     │
                                   └──────────▲──────────┘   └─────────▲───────────┘
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                              │                        │
                                    ┌─────────┴───────┐       ┌────────┴────────┐
                                    │                 │       │                 │
                                    │       OS        │       │    Postgres     │
                                    │   CLI Handler   │       │   CLI Handler   │
                                    │                 │       │                 │
                                    │                 │       │                 │
                                    └─────────────────┘       └─────────────────┘
```
