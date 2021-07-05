package study

import (
	"testing"
	"time"
)

func TestLogProcess_Process(t *testing.T) {
	reader := &ReadFile{
		Path: "./log.txt",
	}
	writer := &WriteInfluxDB{
		InfluxDB: "",
	}

	log := &LogProcess{
		Rc: make(chan string, 1),
		Wc: make(chan string, 1),
		Read: reader,
		Write: writer,
	}
	go log.Read.Read(log.Rc)
	go log.Process()
	go log.Write.Write(log.Rc)

	time.Sleep(5 *time.Second)
}