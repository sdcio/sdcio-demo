package basic_usage

import lib "github.com/saschagrunert/demo"

// example is the single demo run for this application
func Demo() *lib.Run {

	r := lib.NewRun(
		"Setup the advanced usage scenario",
	)

	r.Step(
		lib.S("Verify interface config on device"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface ethernet/*"),
	)

	r.Step(
		lib.S("Apply ethernet config"),
		lib.S("curl -fsSL https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml ;",
			"kubectl apply -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml",
		),
	)

	r.Step(
		lib.S("Verify config CR status"),
		lib.S("kubectl describe configs.config.sdcio.dev dev1-myserviceinterface_disable"),
	)

	r.Step(
		lib.S("Verify interface config is applied on device"),
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
			"kubectl delete -f https://raw.githubusercontent.com/sdcio/sdcio-demo/main/advanced_on_basic/config01_advanced.yaml ;",
		),
	)
	return r
}
