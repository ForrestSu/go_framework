package go_framework

import (
	. "./src"
	"fmt"
	"testing"
	"time"
)

func TestZuvo(t *testing.T){
	t.Log("begin..")

	frame := NewFrameWork(100)
	c1 := NewDemoService("c1", "work1")
	c2 := NewDemoService("c2", "work2")
	_ = frame.RegisterService("c1", c1)
	_ = frame.RegisterService("c2", c2)
	if err := frame.Start(); err != nil {
		fmt.Printf("start error %v\n", err)
	}
	fmt.Println(frame.Start())
	time.Sleep(time.Second * 1)
	err := frame.Stop()
	_ = frame.Destroy()
	t.Log("OK!", err)
}