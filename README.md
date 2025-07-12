# Stress Tester CLI

A minimal Go CLI tool for stress testing HTTP endpoints with configurable concurrency.

## Features

- Configurable number of requests
- Concurrent request execution with goroutines
- Status code grouping and summary
- Performance metrics (requests per second)

## Usage

### Using Docker (Recommended)

```bash
# Pull and run from Docker Hub
docker run --rm yourusername/stress-tester --url https://httpbin.org/get --requests 50 --concurrency 10

# Or build locally
docker build -t stress-tester .
docker run --rm stress-tester --url https://httpbin.org/get --requests 50 --concurrency 10
```

### Local Development

```bash
# Run directly with Go
go run main.go --url https://httpbin.org/get --requests 50 --concurrency 10

# Build and run binary
go build -o stress-tester
./stress-tester --url https://httpbin.org/get --requests 50 --concurrency 10
```

## Flags

- `--url`: (required) URL endpoint to stress test
- `--requests`: Number of requests to make (default: 10)
- `--concurrency`: Number of concurrent goroutines (default: 1)

## Example Output

```
Starting stress test on https://httpbin.org/get with 50 requests using 10 concurrent workers

--- Results ---
Total requests: 50
Concurrency: 10
Total time: 2.345s
Requests per second: 21.32

Status Code Summary:
  200 OK: 48 requests
  ERROR: 2 requests
``` 