package advanced_on_basic

import lib "github.com/saschagrunert/demo"

// example is the single demo run for this application
func Demo() *lib.Run {

	r := lib.NewRun(
		"Setup the advanced usage scenario",
	)

	r.Step(
		lib.S("Verify interface config on device"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface ethernet-1/*"),
	)

	r.Step(
		lib.S("Apply ethernet-1/5 config with low precedence"),
		lib.S("curl -fsSL https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml ; ",
			"kubectl apply -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml",
		),
	)

	r.Step(
		lib.S("Verify interface config remains unchanged on device"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface ethernet-1/*"),
	)

	r.Step(
		lib.S("Apply ethernet-1/5 config higher precedence"),
		lib.S("curl -fsSL https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config02_advanced.yaml ; ",
			"kubectl apply -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config02_advanced.yaml",
		),
	)

	r.Step(
		lib.S("Verify higher precedence config is applied"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface ethernet-1/*"),
	)

	r.Step(
		lib.S("Delete high precedence intent"),
		lib.S("curl -fsSL https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config02_advanced.yaml ; ",
			"kubectl delete -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config02_advanced.yaml",
		),
	)

	r.Step(
		lib.S("Verify initial config is again applied on device"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface ethernet-1/*"),
	)

	return r
}

func Destroy() *lib.Run {

	r := lib.NewRun("Destroy advanced_usage", "remove all the cr's")
	r.Step(
		lib.S(
			"Destroy advanced_usage",
		),
		lib.S(
			"kubectl delete -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml ; ",
		),
	)
	return r
}
