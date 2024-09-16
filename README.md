# RAM Usage API

This API provides information about the current RAM usage of the system. It offers a simple HTTP endpoint that returns RAM usage statistics in various units.

## Table of Contents

1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Usage](#usage)
4. [API Endpoints](#api-endpoints)
5. [Authentication](#authentication)
6. [Units of Measurement](#units-of-measurement)
7. [Example Responses](#example-responses)
8. [Troubleshooting](#troubleshooting)
9. [Contributing](#contributing)
10. [License](#license)

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/astrooom/ram-usage-api.git
   cd ram-usage-api
   ```

2. Build the application:
   ```
   go build -o ram-usage-api ./cmd/server
   ```

## Configuration

1. Create a `.env` file in the root directory with the following content:
   ```
   PORT=8645
   API_TOKEN=your_secret_token_here
   ```

2. Replace `your_secret_token_here` with a strong, unique token of your choice. No token set means the API is open for everyone!

## Usage

Run the server:
```
./ram-usage-api
```

The server will start on the port specified in your `.env` file (default is 8645).

## API Endpoints

### GET /

Returns the current RAM usage statistics.

Query Parameters:
- `unit`: Specifies the unit of measurement for RAM values. Options are:
  - `b` (bytes, default)
  - `kb` (kilobytes)
  - `mb` (megabytes)
  - `gb` (gigabytes)

## Authentication

All requests must include an `Authorization` header with the API token specified in your `.env` file.

Example:
```
Authorization: your_secret_token_here
```

## Units of Measurement

You can specify the unit of measurement for RAM values using the `unit` query parameter:

1. Bytes (default):
   ```
   curl -H "Authorization: your_secret_token_here" http://localhost:8645/
   ```

2. Kilobytes:
   ```
   curl -H "Authorization: your_secret_token_here" "http://localhost:8645/?unit=kb"
   ```

3. Megabytes:
   ```
   curl -H "Authorization: your_secret_token_here" "http://localhost:8645/?unit=mb"
   ```

4. Gigabytes:
   ```
   curl -H "Authorization: your_secret_token_here" "http://localhost:8645/?unit=gb"
   ```

## Example Responses

```json
{
  "total": 16777216000,
  "available": 8388608000,
  "used": 8388608000,
  "free": 8388608000,
  "percent": 50.0,
  "rounded_percent": 50
}
```

## Troubleshooting

- If you receive a 401 Unauthorized error, check that you're using the correct API token in the Authorization header.
- If you receive a 404 Not Found error, ensure you're using the correct endpoint (/).
- If the server won't start, make sure the specified port is not in use by another application.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.