# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
#
# SPDX-License-Identifier: MIT

test:
	go test ./... -cover -coverprofile=coverage.txt

coverage: test
	go tool cover -html=coverage.txt

build:
	go build -o main

lint:
	golint ./...

zip: build
	zip ../terraform/function.zip main

clean:
	rm -rf main ../terraform/function.zip coverage.txt
