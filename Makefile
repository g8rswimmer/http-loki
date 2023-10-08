example-srvr:
	go run cmd/server/*.go -port=8000 -mock_dir="mock-dir"

example-basic:
	echo "Running Loki Mock Basic Example Port 8000"
	go run cmd/server/*.go -port=8000 -mock_dir="_examples/basic/mock-files"	

example-req-validation:
	echo "Running Loki Mock Request Validation Example Port 8000"
	go run cmd/server/*.go -port=8000 -mock_dir="_examples/request-validation/mock-files"	