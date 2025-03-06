build:
	cd ./src && go build -o ../dist/check_domain *.go

install:
	cp -v ./dist/check_domain /usr/local/nagios/libexec/
