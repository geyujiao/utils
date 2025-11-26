package utils

import (
	"testing"
)

func TestNewDomain(t *testing.T) {
	// srcHost := "portal.chinayanghe.com"
	// destHost := "v5-gz-b.douyinvod.com"
	// NewDomain(srcHost, destHost)
	// ModifyFileContent(srcHost, destHost, "v5-gz-b.douyinvod.com.conf")

	srcHost := "findertt.video.qq.com.60.cdnhwcqir15.com"
	destHostList := []string{
		"finderks.video.qq.com.03.cdnhwcqir15.com",
		"finderqv.video.qq.com.97.cdnhwcqir15.com",
		"finderbk.video.qq.com.30.cdnhwcqir15.com",
		"finderdma.video.qq.com.15.cdnhwcqir15.com",
		"finderd1p.video.qq.com.07.cdnhwcqir15.com",
		"finderdsp.video.qq.com.73.cdnhwcqir15.com",
		"finderhy.video.qq.com.20.cdnhwcqir15.com",
		"finderv1.video.qq.com.11.cdnhwcqir15.com",
	}
	NewDomainList(srcHost, destHostList)
	// SplitDomainList(srcHost, destHostList)

}
