mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))


clean:
	rm -r dist || true
mac:
	GOOS=darwin \
	GOARCH=amd64 \
	CGO_ENABLED=1 \
	CGO_CFLAGS="-DAPL=1 -DIBM=0 -DLIN=0" \
	CGO_LDFLAGS="-F/System/Library/Frameworks/ -F${CURDIR}/Libraries/Mac -framework XPLM" \
	go build -buildmode c-shared -o ~/X-Plane\ 12/Resources/plugins/dist//mac.xpl TestPlugin.go
win:
	CGO_CFLAGS="-DIBM=1 -static" \
	CGO_LDFLAGS="-L${CURDIR}/Libraries/Win -lXPLM_64 -static-libgcc -static-libstdc++ -Wl,--exclude-libs,ALL" \
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
	CC=/usr/local/bin/x86_64-linux-musl-cc \
	CGO_CFLAGS="-DLIN=1" \
	CGO_LDFLAGS="-shared -rdynamic -nodefaultlibs -undefined_warning" \
	go build -buildmode c-shared -o dist/lin.xpl TestPlugin.go

all: mac win lin