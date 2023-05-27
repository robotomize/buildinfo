// Package buildinfo provides information about the build.
package buildinfo

import (
	"flag"
	"runtime/debug"
)

const defaultValue = "debug" // Default value for build information fields.

// Variables for storing build information.
var (
	BuildTime = defaultValue
	BuildTag  = defaultValue
	BuildSHA  = defaultValue
)

// Constants for debug build information fields.
const (
	debugVcsRev  = "vcs.revision" // VCS revision field in debug build information.
	debugVcsTime = "vcs.time"     // VCS time field in debug build information.
)

// info contains the build information as a buildStruct.
var info = defaultBuildStruct()

func defaultBuildStruct() *buildStruct {
	return &buildStruct{
		time: newAtomicString(defaultValue), tag: newAtomicString(defaultValue),
		sha: newAtomicString(defaultValue),
	}
}

// buildStruct is a struct for storing build information.
type buildStruct struct {
	time *atomicString // Build time field.
	tag  *atomicString // Build tag field.
	sha  *atomicString // Build SHA field.
}

// Time returns the build time.
func Time() string {
	return info.time.load()
}

// Tag returns the build tag.
func Tag() string {
	return info.tag.load()
}

// SHA returns the build sha.
func SHA() string {
	return info.sha.load()
}

// Set allows setting specific build information fields.
func Set(tag *string, time *string, sha *string) {
	if tag != nil {
		info.tag.store(*tag)
	}
	if time != nil {
		info.time.store(*time)
	}
	if sha != nil {
		info.sha.store(*sha)
	}
}

// init initializes build information from flag, linker, and debug.
func init() {
	initFromFlag()  // Initializes build information from command line flags.
	initFromLD()    // Initializes build information from linker flags.
	initFromDebug() // Initializes build information from debug build information.
}

// initFromFlag initializes build information from command line flags.
func initFromFlag() {
	flagTime := flag.String("buildtime", defaultValue, "-buildtime <current date>")
	flagSha := flag.String("buildsha", defaultValue, "-buildsha <commit sha>")
	flagTag := flag.String("buildtag", defaultValue, "-buildtag <commit tag>")
	flag.Parse()

	if flagTime != nil {
		info.time.store(*flagTime)
	}

	if flagSha != nil {
		info.sha.store(*flagSha)
	}

	if flagTag != nil {
		info.tag.store(*flagTag)
	}
}

// initFromLD initializes build information from linker flags.
func initFromLD() {
	if isDefaultAttribute(info.tag.load()) && !isDefaultAttribute(BuildTag) {
		info.tag.store(BuildTag)
	}
	if isDefaultAttribute(info.time.load()) && !isDefaultAttribute(BuildTime) {
		info.time.store(BuildTime)
	}
	if isDefaultAttribute(info.sha.load()) && !isDefaultAttribute(BuildSHA) {
		info.sha.store(BuildSHA)
	}
}

// initFromDebug initializes build information from debug build information.
func initFromDebug() {
	if inf, ok := debug.ReadBuildInfo(); ok {
		if isDefaultAttribute(info.tag.load()) && inf.Main.Version != "(devel)" {
			info.tag.store(inf.Main.Version)
		}

		for _, s := range inf.Settings {
			if isDefaultAttribute(info.sha.load()) && s.Key == debugVcsRev {
				info.sha.store(s.Value)
			}

			if isDefaultAttribute(info.time.load()) && s.Key == debugVcsTime {
				info.time.store(s.Value)
			}
		}
	}
}

// isDefaultAttribute returns true if the attribute is the default value.
func isDefaultAttribute(attr string) bool {
	return attr == defaultValue
}
