//go:build mage
// +build mage

package main

import (
	"get.porter.sh/porter/mage/mixins"
	"github.com/carolynvs/magex/shx"

	// Import common targets that all mixins should expose to the user
	// mage:import
	porter "get.porter.sh/magefiles/porter"
	_ "get.porter.sh/porter/mage"
)

const (
	mixinName     = "terraform"
	mixinPackage  = "get.porter.sh/mixin/terraform"
	mixinBin      = "bin/mixins/" + mixinName
	porterVersion = "v1.0.0-beta.2"
)

var (
	magefile = mixins.NewMagefile(mixinPackage, mixinName, mixinBin)
	must     = shx.CommandBuilder{StopOnError: true}
)

// Build the mixin
func Build() {
	magefile.Build()
}

// Cross-compile the mixin before a release
func XBuildAll() {
	magefile.XBuildAll()
}

// Run unit tests
func TestUnit() {
	magefile.TestUnit()
}

func Test() {
	magefile.Test()
	TestIntegration()
}

// Publish the mixin to github
func Publish() {
	magefile.Publish()
}

// Install the mixin
func Install() {
	magefile.Install()
}

// Remove generated build files
func Clean() {
	magefile.Clean()
}

// Install porter locally
func EnsureLocalPorter() {
	porter.UseBinForPorterHome()
	porter.EnsurePorterAt(porterVersion)
}

func TestIntegration() {
	EnsureLocalPorter()
	must.Command("./scripts/test/test-cli.sh").RunV()
}
