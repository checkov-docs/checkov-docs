[![codecov](https://codecov.io/gh/checkov-docs/checkov-docs/branch/main/graph/badge.svg?token=X8ESNTI58A)](https://codecov.io/gh/checkov-docs/checkov-docs)

# checkov-docs

Generate documentation from [`checkov`](https://github.com/bridgecrewio/checkov) results.

## Features

- Generate Markdown table of `checkov` *skipped* results.
- Supported formats: `json`.

## Installation

```console
brew install checkov-docs/tap/checkov-docs
```

## Usage

To generate markdown table of `checkov` skipped rules, run the following command:

```console
checkov-docs -i path/to/input/file -o path/to/output/file
```

## Compatibility

This project follows the [Go support policy](https://go.dev/doc/devel/release#policy). Only two latest major releases of Go are supported by the project.

Currently, that means **Go 1.19** or later must be used when developing or testing code.

## Credits

This project is inspired by [`terraform-docs`](https://github.com/terraform-docs/terraform-docs/tree/master).

## License

[MIT License](./LICENSE)
