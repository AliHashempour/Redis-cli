# go-redis

This project is a simple in-memory key-value database system written in [golang](https://go.dev/) with the same features
of
[redis](https://github.com/redis/redis) .

## Features

### Set a Value

Sets a value for a given key in the current database.

- **Usage**: `set <key> <value>`
- **Example**:

```shell
> set mykey "Hello World"
 OK
```

### Get a Value

Retrieves the value for a given key from the current database.

- **Usage**: `get <key>`
- **Example**:

```shell
> get mykey
"Hello World"
```

### Delete a Key

Deletes a key and its associated value from the current database.

- **Usage**: `del <key>`
- **Example**:

```bash
> del mykey
OK
```

### List Keys by Regex

Lists all keys that match a given regex pattern in the current database.

- **Usage**: `<regex_pattern>`
- **Example**:

```bash
> keys my*
["mykey"]
```

### Switch Database

Switches the current working database to another database by name,
creating it if it does not exist.

- **Usage**: `use <db_name>`
- **Example**:

```bash
> use new_db
Switched to database: new_db
```

### List All Databases

Lists all databases currently managed by the system.

- **Usage**: `list`
- **Example**:

```bash
> list
["default", "users_db"]
```

### Dump Database

Dumps the current state of the specified database to a file in JSON format.

- **Usage**: `dump <db_name> <file_path>`
- **Example**:

```bash
> dump default ./default_dump.json
Database dumped successfully.
```

### Load Database

Loads the state of a database from a specified file.

- **Usage**: `load <file_path> <db_name>`
- **Example**:

```bash
> load ./default_dump.json default
Database default loaded successfully.
```

### Exit

Exits the application.

- **Usage**: `exit`
- **Example**:

```bash
> exit
Exiting...
```

## Installation and Running

### Prerequisites:

Ensure you have Go installed on your system. You can download it from Go's [official site](https://go.dev/).

### Run

for installing dependencies run:

```bash
$ go mod tidy
```

finally for running the project run:

```bash
$ go build -o go-redis ./cmd/app

./go-redis
```
