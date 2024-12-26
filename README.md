# AWS CLI profile manager

## Introduction

AWS CLI profile manager helps with multiple AWS profiles. Providing --profile with each AWS command can be tedious and dangerous. This tool helps with storing multiple AWS profiles and then setting the desired profile as default.

## Installation steps

1. Download source code.
2. Build project ```go build acpm.exe```.
3. Move create directory $HOME/acpm and move acpm.exe there.
4. Add executable to path environment variable.
5. Run program with available commands.

## Available commands

List all profiles.
```
acpm ls
```

Show default profile.
```
acpm show default
```

Set default profile.
```
acpm df <profile>
```

Add profile.
```
acpm ap <profile> <access_key_id> <secret_access_key>
```

Remove profile.
```
acpm rp <profile>
```