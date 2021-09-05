package gofileoperation

import (
	"io/ioutil"
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestTextOperation(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		t.Log("Load Failed")
	}
	//file.Read()
	content, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		t.Log("Read Failed")
	}
	t.Log(len(content))
	contentString := string(content)
	t.Log(len(contentString))
	contentString = strings.Replace(contentString,"\n"," ",-1)
	words := strings.Split(contentString," ")
	t.Log(len(words))
	for _,v:=range words{
		if len(v)!=0{
			t.Log(v,len(v))
		}
	}
	file.Close()
}

func TestWriteFile(t *testing.T)  {
	file,err := os.OpenFile("test_write.txt",syscall.O_RDWR,0666)
	if err!=nil{
		t.Log("Open failed")
	}
	wn,errw := file.WriteString("Hello,World")
	if errw != nil{
		t.Log("Write failed")
	}
	t.Logf("Write %d bytes",wn)
	file.Close()
}