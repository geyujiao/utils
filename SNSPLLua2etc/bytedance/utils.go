package bytedance

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func NewRemap() {

	domainList := []string{
		"v5-gz-a.douyinvod.com",
		"v5-gz-b.douyinvod.com",
		"v5-gz-c.douyinvod.com",
		"v5-gz-colda.douyinvod.com",
		"v5-gz-coldb.douyinvod.com",
		"v5-gz-coldc.douyinvod.com",
		"v5-gz-coldd.douyinvod.com",
		"v5-gz-d.douyinvod.com",
		"v5-gzb-a.douyinvod.com",
		"v5-gzb-b.douyinvod.com",
		"v5-gzb-c.douyinvod.com",
		"v96-bd-a.douyinvod.com",
		"v96-bd-h.douyinvod.com",
		"v96-be-pack-bd.pglstatp-toutiao.com",
		"v96-dy.ixigua.com",
		"v96-gray.toutiaovod.com",
		"v96-gz-schedule-un.ixigua.com",
		"v96-gzb-schedule-un.ixigua.com",
		"v96-hc.douyinvod.com",
		"v96-hcc.douyinvod.com",
		"v96-hcd.douyinvod.com",
		"v96-jn-a.douyinvod.com",
		"v96-jn.douyinvod.com",
		"v96-l.bdxiguavod.com",
		"v96-lite-bd-a.douyinvod.com",
		"v96-lite-wha.douyinvod.com",
		"v96-litea.douyinvod.com",
		"v96-ml.douyinvod.com",
		"v96-nj-a.douyinvod.com",
		"v96-nj.douyinvod.com",
		"v96-novelapp-bd.fqnovelvod.com",
		"v96-novelfm-bd.novelfmvod.com",
		"v96-qos-control.douyinvod.com",
		"v96-qos-daily.douyinvod.com",
		"v96-qos-hourly.douyinvod.com",
		"v96-quic.douyinvod.com",
		"v96-reading-video-bd-a.qznovelvod.com",
		"v96-schedule-un.ixigua.com",
		"v96-test.douyinvod.com",
		"v96-web.douyinvod.com",
		"v96-wh-a.douyinvod.com",
		"v96-wh.douyinvod.com",
		"v96-wha-bd-a.douyinvod.com",
		"v96-wha.douyinvod.com",
		"v96.bdxiguavod.com",
		"v96.douyinvod.com",
		"v96.idouyinvod.com",
		"v96.itoutiaovod.com",
		"v96.toutiaovod.com",
		"v99-cold.douyinvod.com",
		"v99-coldx.douyinvod.com",
		"v99-default.douyinvod.com",
		"v99-hca.douyinvod.com",
		"v99-hcb.douyinvod.com",
		"v99.douyinvod.com",
	}

	for _, info := range domainList {
		// 修改文件内容
		filename := fmt.Sprintf("./20230706-字节新架构/%s/CCSL3/ats/remap.config", info)
		err := WriteFile(filename, info)
		if err != nil {
			fmt.Println("filename:"+filename+", Error:", err)
			os.Exit(1)
		}
		fmt.Println("filename:" + filename)

		// break
	}

	// time.Sleep(10 * time.Second)
	// fmt.Println("Replacement completed successfully.")
}

func NewDomainNginx() {

	domainList := []string{
		"v5-gz-a.douyinvod.com",
		"v5-gz-b.douyinvod.com",
		"v5-gz-c.douyinvod.com",
		"v5-gz-colda.douyinvod.com",
		"v5-gz-coldb.douyinvod.com",
		"v5-gz-coldc.douyinvod.com",
		"v5-gz-coldd.douyinvod.com",
		"v5-gz-d.douyinvod.com",
		"v5-gzb-a.douyinvod.com",
		"v5-gzb-b.douyinvod.com",
		"v5-gzb-c.douyinvod.com",
		"v96-bd-a.douyinvod.com",
		"v96-bd-h.douyinvod.com",
		"v96-be-pack-bd.pglstatp-toutiao.com",
		"v96-dy.ixigua.com",
		"v96-gray.toutiaovod.com",
		"v96-gz-schedule-un.ixigua.com",
		"v96-gzb-schedule-un.ixigua.com",
		"v96-hc.douyinvod.com",
		"v96-hcc.douyinvod.com",
		"v96-hcd.douyinvod.com",
		"v96-jn-a.douyinvod.com",
		"v96-jn.douyinvod.com",
		"v96-l.bdxiguavod.com",
		"v96-lite-bd-a.douyinvod.com",
		"v96-lite-wha.douyinvod.com",
		"v96-litea.douyinvod.com",
		"v96-ml.douyinvod.com",
		"v96-nj-a.douyinvod.com",
		"v96-nj.douyinvod.com",
		"v96-novelapp-bd.fqnovelvod.com",
		"v96-novelfm-bd.novelfmvod.com",
		"v96-qos-control.douyinvod.com",
		"v96-qos-daily.douyinvod.com",
		"v96-qos-hourly.douyinvod.com",
		"v96-quic.douyinvod.com",
		"v96-reading-video-bd-a.qznovelvod.com",
		"v96-schedule-un.ixigua.com",
		"v96-test.douyinvod.com",
		"v96-web.douyinvod.com",
		"v96-wh-a.douyinvod.com",
		"v96-wh.douyinvod.com",
		"v96-wha-bd-a.douyinvod.com",
		"v96-wha.douyinvod.com",
		"v96.bdxiguavod.com",
		"v96.douyinvod.com",
		"v96.idouyinvod.com",
		"v96.itoutiaovod.com",
		"v96.toutiaovod.com",
		"v99-cold.douyinvod.com",
		"v99-coldx.douyinvod.com",
		"v99-default.douyinvod.com",
		"v99-hca.douyinvod.com",
		"v99-hcb.douyinvod.com",
		"v99.douyinvod.com",
	}

	for _, info := range domainList {
		// 修改文件内容
		filename := fmt.Sprintf("./20230706-字节新架构/%s/CCSL3/nginx/%s3-l2.conf", info, info)
		originalList := ReadOriginalData(filename)
		err := WriteNginxFile(filename, originalList)
		if err != nil {
			fmt.Println("filename:"+filename+", Error:", err)
			os.Exit(1)
		}
		fmt.Println("filename:" + filename)

		// break
	}

	// time.Sleep(10 * time.Second)
	// fmt.Println("Replacement completed successfully.")
}

func WriteFile(filename, strPre string) (err error) {

	// 创建/打开文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		return err
	}
	defer file.Close() // 确保文件在函数结束时关闭

	content := "map http://" + strPre + " http://v-cm.pstatp.com @plugin=cachekey.so @pparam=--static-prefix=bytedancevediohycdn @pparam=--remove-all-params=true @pparam=--include-headers=Range @plugin=header_rewrite.so @pparam=/etc/trafficserver/plugins/headerrewrite/bytedance_vedio-l2.hr @plugin=cache_range_requests.so @plugin=conf_remap.so @pparam=proxy.config.http.insert_request_via_str=0"
	lines := make([]string, 0)
	content_str := string(content)
	// fmt.Println("content_str=" + content_str + "\n")
	// lines = append(lines, content_str)
	// lines = append(lines, strings.Replace(content_str, " http://v-cm.pstatp.com", ":80 http://v-cm.pstatp.com", -1))
	// lines = append(lines, strings.Replace(content_str, " http://v-cm.pstatp.com", ":443 http://v-cm.pstatp.com", -1))
	lines = append(lines, strings.Replace(content_str, " http://v-cm.pstatp.com", ":8080 http://v-cm.pstatp.com", -1))
	// lines = append(lines, strings.Replace(content_str, " http://v-cm.pstatp.com", ":8443 http://v-cm.pstatp.com", -1))

	// bytes, _ := json.Marshal(lines)
	// fmt.Println("new lines=" + string(bytes) + "\n")

	// 创建一个*Writer，用于按行写入
	writer := bufio.NewWriter(file)

	// 按行写入数据
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	// 确保所有数据都被刷新到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return err
	}

	// fmt.Println("文件写入成功，内容如下：")
	// // 打印文件内容
	// newContent, _ := os.ReadFile(filename)
	// fmt.Print(string(newContent))

	return err

}

func WriteNginxFile(filename string, contentList []string) (err error) {

	// 创建/打开文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		return err
	}
	defer file.Close() // 确保文件在函数结束时关闭

	// bytes, _ := json.Marshal(lines)
	// fmt.Println("new lines=" + string(bytes) + "\n")

	// 创建一个*Writer，用于按行写入
	writer := bufio.NewWriter(file)

	// 按行写入数据
	// for _, line := range contentList {
	// 	_, err := writer.WriteString(line + "\n")
	// 	if err != nil {
	// 		fmt.Println("Error writing to file:", err)
	// 		return err
	// 	}
	// }
	// writer.WriteString("\n")

	for k, line := range contentList {
		// if k < 5 {
		// 	continue
		// }
		if k > 95 {
			break
		}
		if k == 6 {
			line = strings.ReplaceAll(line, "*:80;", "*:8080;")
		}
		if k == 7 {
			line = strings.ReplaceAll(line, "[::]:80;", "[::]:8080;")
		}
		if k == 8 {
			line = strings.ReplaceAll(line, "*:443 ssl;", "*:8443 ssl;")
		}
		if k == 9 {
			line = strings.ReplaceAll(line, "[::]:443 ssl;", "[::]:8443 ssl;")
		}
		if k == 75 {
			// line = strings.ReplaceAll(line, "Host $host;", "Host $host:$server_port;")
			line = strings.ReplaceAll(line, "Host $host;", "Host $host:8080;")
		}
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return err
		}
	}

	// 确保所有数据都被刷新到文件中
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return err
	}

	// fmt.Println("文件写入成功，内容如下：")
	// // 打印文件内容
	// newContent, _ := os.ReadFile(filename)
	// fmt.Print(string(newContent))

	return err

}

// 1. 按行读文件，得到list
func ReadOriginalData(fileName string) (originalList []string) {
	// 读取一个文件的内容
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 处理结束后关闭文件
	defer file.Close()

	// 使用bufio读取
	r := bufio.NewReader(file)

	for {
		// 分行读取文件  ReadLine返回单个行，不包括行尾字节(\n  或 \r\n)
		data, _, err := r.ReadLine()
		// 读取到末尾退出
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("read err", err.Error())
			break
		}
		// if string(data) == "" {
		// 	continue
		// }
		originalList = append(originalList, string(data))
	}
	return
}
