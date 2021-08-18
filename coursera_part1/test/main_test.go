package main

import (
	"bufio"
	"strings"
	"bytes"
	"testing"
)

var testOK=`1
2
3
4
5`

var testOkResult=`1
2
3
4
5
`

func TestOk(t *testing.T){
	in:=bufio.NewReader(strings.NewReader(testOK))
	out:=new(bytes.Buffer)
	err:=uniq(in,out)
	if err!=nil{
		t.Errorf("test for OK failed - error")
	}
	result:=out.String()
	if out.String()!=testOkResult{
		t.Errorf("test for OK failed - results not match\n %v %v",result,testOkResult)
	}
}

var testFail=`1
2
1`

func TestForError(t *testing.T){
	in:=bufio.NewReader(strings.NewReader(testFail))
	out:=new(bytes.Buffer)
	err:=uniq(in,out)
	if err==nil{
		t.Errorf("test for OK failed - error: %v",err)
	}
}