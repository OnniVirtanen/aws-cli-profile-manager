# AWS CLI profile manager

## Introduction

AWS CLI profile manager helps with multiple AWS profiles. Providing --profile with each AWS command can be tedious and dangerous. This tool helps with storing multiple AWS profiles and then setting the desired profile as default.

## Installation steps

1. Download source code.
2. Build project ```go build -o aws-profile.exe```.
3. Add executable to path environment variable.
4. Run program with available commands.

## Available commands

Set default profile.
```
aws-profile <profile>
```