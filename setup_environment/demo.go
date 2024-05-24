package setup_environment

import lib "github.com/saschagrunert/demo"

// example is the single demo run for this application
func Demo() *lib.Run {

	r := lib.NewRun(
		"Setup Environment",
	)

	// setup kind cluster
	r.Step(
		lib.S("create k8s kind cluster"),
		lib.S("kind create cluster"),
	)

	// setup iptables rules
	r.Step(
		lib.S("Allow the kind cluster to communicate with the later created containerlab topology"),
		lib.S("sudo iptables -I DOCKER-USER -o br-$(docker network inspect -f '{{ printf \"%.12s\" .ID }}' kind) -j ACCEPT"),
	)

	// Deploy Containerlab Topo
	r.Step(
		lib.S("Deploy Containerlab topology"),
		lib.S("sudo containerlab deploy -t https://docs.sdcio.dev/artifacts/basic-usage/basic-usage.clab.yaml --reconfigure"),
	)

	// Install CertManager
	r.Step(
		lib.S("Install CertManager"),
		lib.S("kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.yaml"),
	)

	// Wait for CertManager
	r.Step(
		lib.S(
			"If the SDCIO resources, see below are being applied to fast, the webhook of the cert-manager is not already there.",
			"Hence we need to wait for the resource be become Available",
		),
		lib.S("kubectl wait -n cert-manager --for=condition=Available=True --timeout=300s deployments.apps cert-manager-webhook"),
	)

	return r
}

func Destroy() *lib.Run {

	r := lib.NewRun(
		"Setup Environment",
	)

	// setup iptables rules
	r.Step(
		lib.S("drop the iptables rule"),
		lib.S("sudo iptables -D DOCKER-USER -o br-$(docker network inspect -f '{{ printf \"%.12s\" .ID }}' kind) -j ACCEPT"),
	)

	// delete kind cluster
	r.Step(
		lib.S("delete k8s kind cluster"),
		lib.S("kind delete cluster"),
	)

	// Destroy Containerlab Topo
	r.Step(
		lib.S("Destroy Containerlab topology"),
		lib.S(
			"curl https://docs.sdcio.dev/artifacts/basic-usage/basic-usage.clab.yaml ; ",
			"sudo containerlab destroy -t https://docs.sdcio.dev/artifacts/basic-usage/basic-usage.clab.yaml --cleanup",
		),
	)

	return r
}
