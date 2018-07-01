PLUGINS := Apis/Apis.go Intents/Intents.go \
		IsMalicious/IsMalicious.go Md5/Md5.go PackageName/PackageName.go \
		PackageVersion/PackageVersion.go Permissions/Permissions.go \
		Sha1/Sha1.go Strings/Strings.go Date/Date.go FileSize/FileSize.go \
		Sha256/Sha256.go
PREFIX=go build -buildmode=plugin

.PHONY: all clean

all: $(PLUGINS)
	@for plugin in $(PLUGINS); do echo "Building Plugin $$plugin"; $(PREFIX) $$plugin; done

clean:
	@rm -f *.so
