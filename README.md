# AWS CLI profile manager

## Introduction

AWS CLI profile manager helps with multiple AWS profiles. Providing --profile with each AWS command can be tedious and dangerous. This tool helps with storing multiple AWS profiles and then setting the desired profile as default.

## Installation steps

1. Download source code.
2. Build project ```go build -o apm.exe```.
3. Add executable to path environment variable.
4. Run program with available commands.

## Available commands

Show available commands.
```
apm --help
```

Show current version.
```
apm --v
```

List all profiles.
```
apm ls
```

Show default profile.
```
apm show default
```

Set default profile.
```
apm default <profile>
```

Add profile (aws_session_token must be provided when using temporary credentials).
```
apm add <profile> <access_key_id> <secret_access_key> [aws_session_token]
```

Remove profile.
```
apm rmv <profile>
```