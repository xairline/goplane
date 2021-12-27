clean:
	rm -r dist

mac:
	GOOS=darwin \
    GOARCH=amd64 \
    CGO_ENABLED=1 \
    CGO_CFLAGS="-I/Users/dzou/Downloads/SDK/CHeaders -DAPL=1 -DIBM=0 -DLIN=0 -DXPLM_DEPRECATED=1" \
    CGO_LDFLAGS="-F/System/Library/Frameworks/ -F/Users/dzou/Downloads/SDK/Libraries/Mac -framework XPLM" \
    go build -buildmode c-shared -o dist/mac.xpl TestPlugin.go
win:
	CGO_CFLAGS="-DIBM=1 -static" \
	CGO_LDFLAGS="-L/Users/dzou/Downloads/SDK/Libraries/Win -lXPLM_64 -static-libgcc -static-libstdc++ -Wl,--exclude-libs,ALL" \
	GOOS=windows \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CC=x86_64-w64-mingw32-gcc \
	CXX=x86_64-w64-mingw32-g++ \
	go build --buildmode c-shared -o dist/win.xpl TestPlugin.go
lin:
	GOOS=linux \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CGO_CFLAGS="-I/Users/dzou/Downloads/SDK/CHeaders -DLIN=1" \
	go build -buildmode c-shared -o dist/lin.xpl TestPlugin.go