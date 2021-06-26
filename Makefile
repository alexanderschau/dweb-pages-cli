build:
	env GOOS=darwin GOARCH=amd64 go build -o dist/${package}_${version}_Darwin_x86_64

	env GOOS=linux GOARCH=arm go build -o dist/${package}_${version}_Linux_armv6
	env GOOS=linux GOARCH=386 go build -o dist/${package}_${version}_Linux_i386
	env GOOS=linux GOARCH=amd64 go build -o dist/${package}_${version}_Linux_x86_64

	env GOOS=windows GOARCH=386 go build -o dist/${package}_${version}_Windows_i386.exe
	env GOOS=windows GOARCH=amd64 go build -o dist/${package}_${version}_Windows_x86_64.exe
