# Label Printer
An API server to create a tag image with Code128 or QR codes for a text and print it through NIIMBOT D110 label printer

## Requirements
- Python `>=3.10`
- Golang
- Linux (macOS doesn't support bluetooth sockets, not tested on Windows)

## Getting started
- In `niimprint` directory, run `pip3 install -r requirements.txt`
- Configure image generation and calls to `niimprint` in `print.go`
- Run golang server: `go run .`
- POST request on `http://localhost:8769/print` with payload in following JSON format
	```json
	{
		"text" : "MYLABEL",
		"code_text" : "https://www.example.com/MYLABEL",
		"code_type" : "qr"
	}
	```

### cURL Example

```shell
curl -d '{"text":"MYLABEL", "qr_text":"https://www.example.com/MYLABEL"}' -H "Content-Type: application/json" -X POST http://localhost:8769/print
```
