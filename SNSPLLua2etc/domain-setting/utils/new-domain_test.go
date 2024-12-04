package utils

import (
	"testing"
)

func TestNewDomain(t *testing.T) {
	// srcHost := "portal.chinayanghe.com"
	// destHost := "v5-gz-b.douyinvod.com"
	// NewDomain(srcHost, destHost)
	// ModifyFileContent(srcHost, destHost, "v5-gz-b.douyinvod.com.conf")

	srcHost := "v96.douyinvod.com" //"portal.chinayanghe.com"
	destHostList := []string{
		// "v5-e-mc.douyinvod.com",
		// "v5-f-mc.douyinvod.com",
		// "v5-g-mc.douyinvod.com",
		// "v5-h-mc.douyinvod.com",
	}
	NewDomainList(srcHost, destHostList)
	// SplitDomainList(srcHost, destHostList)

}
