# Build info

A package for getting application build metadata information from multiple sources

# Usage

```shell
go get -u github.com/robotomize/buildinfo
```

## Example

```go

logger, _ := zap.NewProduction()
sugar := logger.Sugar()

lg := sugar.With(
zap.String("build_tag", buildinfo.Tag()), zap.String("build_time", buildinfo.Time()),
zap.String("build_sha", buildinfo.SHA()),
)

```

### Use flags

```shell
./exampleapp -buildtag=$(git describe --tags --abbrev=0) -buildtime=$(date -u '+%Y-%m-%d-%H:%M') -buildsha=$(git rev-parse HEAD)
```

### Use buildinfo ldflags
```shell
go build -ldflags "-X github.com/robotomize/buildinfo.BuildTag=v0.4.0 \
-X github.com/robotomize/buildinfo.BuildTime=2022-05-27 \
-X github.com/robotomize/buildinfo.BuildSHA=e4601a766ce364b65427cbcfd3f0cbfe233725af"
```

```go
fmt.Println(buildinfo.Tag(),buildinfo.Time(),buildinfo.SHA())
```

## Use ldflags and Set method

```shell
go build -ldflags "-X main.BuildTag=v0.4.0 -X main.BuildTime=2022-05-27 -X main.BuildSHA=e4601a766ce364b65427cbcfd3f0cbfe233725af"
```

```go

var (
    	BuildTag string
	BuildTime string
	BuildSHA string
)

buildinfo.Set(&BuildTag, &BuildTime, &BuildSHA)
fmt.Println(buildinfo.Tag(),buildinfo.Time(),buildinfo.SHA())
```

## Order
* Set method
* build flags
* ldflags
* -buildvcs=true
