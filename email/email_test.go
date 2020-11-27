package email

import (
	"fmt"
	"testing"
)

func TestEmailClient_SendMessage(t *testing.T) {
	username := "750313228@qq.com"
	host := "smtp.qq.com"
	password := "zibjzgvjkpqqbfhe"
	port := 465

	subject := "邀请"
	content := `尊敬的张先生先生/女士：<br>
"北京脑科学与类脑研究中心"诚挚地邀请您担任<l style="color:red">2021年北京脑中心第*批博士研究生申请</l>材料<l style="color:red">的初评/面试</l>专家，并为申报人进行材料评审评分。<br>
评审流程：1、请您在线查看申报人申报材料；2、请您在线为申报人评分，确认无误后提交最终评审结果。完成以上两个步骤，您的工作流程已结束。<br>
注意事项：评分截止日期为2020-02-12。<br>
我们诚挚地期待能够通过您的专业意见，筛选出符合条件的优秀候选人。如蒙同意，我们深表荣幸。后续工作中，还望您按时按要求完成评分为盼。由衷感谢您的支持！`
	contentType := "text/html"
	attach := ""
	from := "2813474070@qq.com" //发件人
	to := []string{"gyj17718326036@163.com"} //收件人
	cc := []string{"gyj17718326036@163.com"} //抄送人

	message := NewEmailMessage(from, subject, contentType, content, attach, to, cc)
	email := NewEmailClient(host, username, password, port, message)
	_, err := email.SendMessage()
	if err != nil {
		fmt.Println("失败", err)
	}
}
