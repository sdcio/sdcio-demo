package install_sdcio

import lib "github.com/saschagrunert/demo"

// example is the single demo run for this application
func Demo() *lib.Run {

	r := lib.NewRun(
		"Setup SDCIO",
	)

	r.Step(lib.S(
		"Install SDCIO",
	), lib.S(
		"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/colocated.yaml",
	))

	return r
}
