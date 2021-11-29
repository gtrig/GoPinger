# GoPinger
A fast Rest Api to perform ping in a single IP or an IP range.

## Examples
http://localhost:8020/ping/192.168.1.1
http://localhost:8020/pingrange/192.168.1.1/192.168.1.254

## Configuration
The server can easily be configured from the .env file
### Posible variables
* GOPINGER_API_PORT=8020
* GOPINGER_COUNT=5
* GOPINGER_TIMEOUT_MS=5000
* GOPINGER_INTERVAL_MS=500