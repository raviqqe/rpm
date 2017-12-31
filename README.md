# rpm

RPM (Requests Per Minute) reporter in Go.

## Installation

```sh
go get github.com/raviqqe/rpm
```

## Usage

```
rpm, the RPM reporter

Usage:
	rpm [-n <number>] [-c <number>] [-h] <url>

Options:
	-n <number>  Number of total requests. [default: 10000]
	-c <number>  Number of concurrent requests. [default: 1000]
	-h           Show this help.
```

## License

[MIT](LICENSE)
