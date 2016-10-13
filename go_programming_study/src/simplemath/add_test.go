package simplemath

import "testing"


/**
在任意目录下运行单元测试如下
go test simplemath
 */


func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed . Got %的， expect 3. ", r)
	}
}
