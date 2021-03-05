check_install:
	which swagger || O111MODULE==off go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	# GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models
	swagger generate spec -o ./swagger.yml --scan-models