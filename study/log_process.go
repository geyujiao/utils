package study

import (
	"bufio"
	"github.com/vgmdj/utils/logger"
	"io"
	"os"
	"strings"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type Writer interface {
	Write(wc chan string)
}

type ReadFile struct {
	Path string //文件目录
}

func (r *ReadFile) Read(rc chan []byte) {
	// 读取
	// 打开文件
	f, err := os.Open(r.Path)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	// 从文件末尾开始逐行读取文件内容
	f.Seek(0, 2)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadBytes('\n')
		if err == io.EOF {
			//logger.Info("read to tail... ")
			time.Sleep(50 * time.Microsecond)
			continue
		} else if err != nil {
			logger.Error("ReadBytes err", err.Error())
			return

		}
		logger.Info("read byte", string(line[:len(line)-1]))
		rc <- line[:len(line)-1] //去掉换行符

	}
}

type WriteInfluxDB struct {
	InfluxDB string
}

func (w *WriteInfluxDB) Write(wc chan string) {
	//  写入
	for v := range wc {
		logger.Info("write data", v)
	}

}

type LogProcess struct {
	Rc    chan []byte
	Wc    chan string
	Read  Reader
	Write Writer
}

func (l *LogProcess) Process() {
	// 解析
	for v := range l.Rc{
		// TODO
		// 用正则表达式解析数据
		// 将解析后的数据 写入到 chan 中
		l.Wc <- strings.ToUpper(string(v))
	}

}
