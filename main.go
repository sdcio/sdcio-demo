package main

import (
	"github.com/saschagrunert/demo"
	"github.com/sdcio/sdcio-demo/basic_usage"
	"github.com/sdcio/sdcio-demo/install_sdcio"
	"github.com/sdcio/sdcio-demo/setup_environment"
)

func main() {
	// Create a new demo CLI application
	d := demo.New()

	// Register the demo runs
	d.Add(setup_environment.Demo(), "setup_environment", "setup environment")
	d.Add(install_sdcio.Demo(), "install_sdcio", "installing sdcio")
	d.Add(basic_usage.Demo(), "basic_usage", "basic usage demo")
	d.Add(basic_usage.Destroy(), "basic_usage_destroy", "destroy basic usage")
	d.Add(setup_environment.Destroy(), "destroy_environment", "destroy environment")

	// Run the application, which registers all signal handlers and waits for
	// the app to exit
	d.Run()
}
