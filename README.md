# awscreds

![test](https://github.com/jedipunkz/awscreds/actions/workflows/ci.yml/badge.svg)
## What is awscreds?

awscreds is a CLI tool to setup aws credentials for terminal shell with MFA device authentication.

## Requirement

- go 1.22 x or later
- awscli

## Installation

```shell
go install github.com/jedipunkz/awscreds@latest
```

## Usage

### Options

| Option | Explanation             | Required |
|--------|-------------------------|--------|
| -m     | MFA Number              | YES    |
| -p     | aws config profile name | YES    |
| -r     | aws region name         | YES    |
| -s     | shell (fish, zsh, bash, sh) | No |

### Fish

```shell
awscreds set -p <aws-config-profile-name> -r <region-name> -m <mfa-number> | source
```

### Zsh, Bash

```shell
eval `awscreds set -p <aws-config-profile-name> -r <region-name> -m <mfa-number>`
```

## License

[Apache License 2.0](https://github.com/jedipunkz/awscreds/blob/main/LICENSE)

## Author

[jedipunkz](https://twitter.com/jedipunkz)
