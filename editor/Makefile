# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
#
# SPDX-License-Identifier: MIT

test:
	go test ./... -cover -coverprofile=coverage.txt

coverage: test
	go tool cover -html=coverage.txt

build:
	go build -o editor

lint:
	golint ./...

clean:
	rm -rf main coverage.txt
