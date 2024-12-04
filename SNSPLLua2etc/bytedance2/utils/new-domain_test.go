package utils

import "testing"

// func TestCopyFile(t *testing.T) {
// 	srcHost := "v96.douyinvod.com" //"portal.chinayanghe.com"
// 	destHost := "v5-e-mc.douyinvod.com"
// 	oldfile := fmt.Sprintf("./forbidden/toutiao/forbidden_url_%s.conf", srcHost)
// 	newFile := fmt.Sprintf("./forbidden/toutiao/forbidden_url_%s.conf", destHost)
// 	CopyFile(oldfile, newFile)
// }

func TestNewDomain(t *testing.T) {
	// srcHost := "portal.chinayanghe.com"
	// destHost := "v5-gz-b.douyinvod.com"
	// NewDomain(srcHost, destHost)
	// ModifyFileContent(srcHost, destHost, "v5-gz-b.douyinvod.com.conf")

	srcHost := "v96.douyinvod.com" //"portal.chinayanghe.com"
	destHostList := []string{
		"v5-mc-reading-video-a.qznovelvod.com",
		"v5-mc-reading-video-b.qznovelvod.com",
		"v5-mc-be-pack.pglstatp-toutiao.com",
		"v5-mc-reading-video.fqnovelvod.com",
		// "v5-e-mc.douyinvod.com",
		// "v5-f-mc.douyinvod.com",
		// "v5-g-mc.douyinvod.com",
		// "v5-h-mc.douyinvod.com",
		// "v5-mc-wha.douyinvod.com",
		// "v5-mc-whb.douyinvod.com",
		// "v5-mc-cold.douyinvod.com",
		// "v5-mc-lite-wha.douyinvod.com",
		// "v5-mc-hsv2-a.douyinvod.com",
		// "v5-mc-lite.douyinvod.com",
		// "v5-mc-lite-a.douyinvod.com",
		// "v5-mc-lite-b.douyinvod.com",
		// "v5-mc-lite-c.douyinvod.com",
	}
	NewDomainList(srcHost, destHostList)
	// SplitDomainList(srcHost, destHostList)

}
