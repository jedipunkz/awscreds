# awscreds

## What is awscreds?

awscreds is CLI tool to setup aws credentials with MFA device.

## Requirement

- go 1.17.x or earlier

## Installation

```shell
go install github.com/jedipunkz/awscreds@latest
```

## Usage

```shell
awscreds set -p <aws-config-profile-name> -r <region-name> -m <mfa-number> | source
```

## License

MIT

## Author

[jedipunkz](https://twitter.com/jedipunkz)