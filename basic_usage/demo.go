package basic_usage

import lib "github.com/saschagrunert/demo"

// example is the single demo run for this application
func Demo() *lib.Run {

	r := lib.NewRun(
		"Setup the basic usage scenario",
	)

	r.Step(
		lib.S("Nokia SR Linux Yang Schema"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/schema-nokia-srl-23.10.1.yaml | yq ;",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/schema-nokia-srl-23.10.1.yaml",
		),
	)

	r.Step(
		lib.S("Connection Profile"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/target-conn-profile-gnmi.yaml | yq ;",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/target-conn-profile-gnmi.yaml",
		),
	)

	r.Step(
		lib.S("Sync Profile"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/target-sync-profile-gnmi.yaml | yq ;",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/target-sync-profile-gnmi.yaml",
		),
	)

	r.Step(
		lib.S("SRL Secret"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/secret-srl.yaml | yq ; ",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/secret-srl.yaml",
		),
	)

	r.Step(
		lib.S("Discovery Rule"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/discovery_address.yaml | yq ; ",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/discovery_address.yaml",
		),
	)

	r.Step(
		lib.S("Verify SDCIO objects"),
		lib.S("kubectl get sdc"),
	)

	r.Step(
		lib.S("Verify SDCIO objects"),
		lib.S("kubectl get sdc"),
	)

	r.Step(
		lib.S("Retrieve running config"),
		lib.S("kubectl get runningconfigs.config.sdcio.dev dev1 -o yaml | yq"),
	)

	r.Step(
		lib.S("Verify interface system0 is not present"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface system0"),
	)

	r.Step(
		lib.S("Apply Config"),
		lib.S("curl https://docs.sdcio.dev/artifacts/basic-usage/config.yaml ;",
			"kubectl apply -f https://docs.sdcio.dev/artifacts/basic-usage/config.yaml",
		),
	)

	r.Step(
		lib.S("Verify config CR status"),
		lib.S("kubectl describe configs.config.sdcio.dev test"),
	)

	r.Step(
		lib.S("Verify config is applied on device"),
		lib.S("docker exec clab-basic-usage-dev1 sr_cli -- info from running interface system0"),
	)

	return r
}

func Destroy() *lib.Run {

	r := lib.NewRun("Destroy basic_usage", "remove all the cr's")
	r.Step(
		lib.S(
			"Destroy basic_usage",
		),
		lib.S(
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/config.yaml ;",
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/discovery_address.yaml ;",
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/secret-srl.yaml ;",
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/target-sync-profile-gnmi.yaml ; ",
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/target-conn-profile-gnmi.yaml ; ",
			"kubectl delete -f https://docs.sdcio.dev/artifacts/basic-usage/schema-nokia-srl-23.10.1.yaml",
		),
	)
	return r
}
