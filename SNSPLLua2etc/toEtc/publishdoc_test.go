package snspllua2etc

import (
	"testing"
)

func TestWritePublishExcel(t *testing.T) {
	WritePublishExcel("requirement.txt", "path.txt", "publish.xlsx")
	WritePublishExcel2("requirement.txt", "path.txt", "publishR.xlsx")

	// requirementMap, domainMap := GetDomainDateMap("requirement.txt", "path.txt")
	// bytes, _ := json.Marshal(requirementMap)
	// println(string(bytes))
	// println("111111")
	// bytes, _ = json.Marshal(domainMap)
	// println(string(bytes))

}
