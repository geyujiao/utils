package utils

import (
	"testing"
)

func TestNewDomain(t *testing.T) {
	// srcHost := "portal.chinayanghe.com"
	// destHost := "v5-gz-b.douyinvod.com"
	// NewDomain(srcHost, destHost)
	// ModifyFileContent(srcHost, destHost, "v5-gz-b.douyinvod.com.conf")

	srcHost := "portal.chinayanghe.com"
	destHostList := []string{
		"pic.cm.ahiptv.ahntv.cn",
		// "v5-f-mc.douyinvod.com",
		// "v5-g-mc.douyinvod.com",
		// "v5-h-mc.douyinvod.com",
	}
	NewDomainList(srcHost, destHostList)
	// SplitDomainList(srcHost, destHostList)

}
