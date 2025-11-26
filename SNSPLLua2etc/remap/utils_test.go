package remap

import "testing"
func TestRemap(t *testing.T) {
	GetMap()
}

// -map http://otttv.bj.chinamobile.com http://otttv.bj.chinamobile.com @plugin=cachepolicy.so @plugin=resolv.so @pparam=127.0.0.1:65000 @plugin=tslua.so @pparam=/etc/trafficserver/lua/chinamobile_reg_setcachekey-l2.lua @pparam=otttv.bj.chinamobile.com @plugin=cache_range_requests.so @plugin=no_cache_status.so @pparam=403,404,416,503,302,502 @plugin=conf_remap.so @pparam=proxy.config.http.connect_attempts_timeout=6 @pparam=proxy.config.http.transaction_no_activity_timeout_out=10
// 1111---map http://wapx.cmvideo.cn http://cache.wapx.cmvideo.cn @plugin=resolv.so @pparam=cache.wapx.cmvideo.cn:8080 plugin=cachepolicy.so @pparam=match-suffix-status=jpg|png|jpeg|gif|webp|sif:2XX|304:90d,mp4:2XX|304:1d,*:4XX|5XX:0 @plugin=cachekey.so @pparam=--remove-all-params=true @plugin=background_fetch.so @plugin=conf_remap.so @pparam=proxy.config.http.cache.required_headers=2
// 1111---map http://ywotttv.bj.chinamobile.com http://ywotttv.bj.chinamobile.com @plugin=cachepolicy.so @plugin=resolv.so @pparam=127.0.0.1:65000 @plugin=tslua.so @pparam=/etc/trafficserver/lua/chinamobile_reg_setcachekey-l2.lua @pparam=ywotttv.bj.chinamobile.com @plugin=cache_range_requests.so @plugin=no_cache_status.so @pparam=403,404,416,503,302,502 @plugin=conf_remap.so @pparam=proxy.config.http.connect_attempts_timeout=6 @pparam=proxy.config.http.transaction_no_activity_timeout_out=10