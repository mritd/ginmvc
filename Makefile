BUILD_VERSION   := $(version)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64" \
        -output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
    	-ldflags   "-X 'main.Version=${BUILD_VERSION}' \
                    -X 'main.BuildTime=${BUILD_TIME}' \
                    -X 'main.CommitID=${COMMIT_SHA1}'"

docker: all
	docker build -t mritd/ginmvc:$(version) .

clean:
	rm -rf dist

install:
	go install

.PHONY : all release docker clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
