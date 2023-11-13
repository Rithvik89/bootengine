beb:
	go build -o devboot .

localtest: 
	go build -o devboot . && sudo mv ./devboot /usr/local/bin