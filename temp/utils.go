package main

import "fmt"

var domain_mainos_map = map[string]map[string]string{
	"v5-e-mc.douyinvod.com": {"mainos": "https://mainos.bytedance.com", "secondos": "https://secondos.bytedance.com"},
	"v5-f-mc.douyinvod.com": {"mainos": "https://mainos.bytedance.com", "secondos": "https://secondos.bytedance.com"},
}

func main() {
	for domain, servermap := range domain_mainos_map {
		for mainos, secondos := range servermap {
			fmt.Println(domain + "  " + mainos + "   " + secondos)
		}
	}
	var req_host = "v5-e-mc.douyinvod.com"
	if servermap, ok := domain_mainos_map[req_host]; ok {
		fmt.Println(servermap["mainos"])
		fmt.Println(servermap["secondos"])
	} else {
		fmt.Println("not found")
	}

	switch req_host {
	case "v5-e-mc.douyinvod.com":
		fmt.Println("v5-e-mc.douyinvod.com")
	default:
		fmt.Println("default")
	}

}
