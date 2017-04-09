package properties

import (
    "testing"
    "strings"
    "fmt"
    "time"
)

func Test_String(t *testing.T) {
    s := `
    a=aa
    b=bbb
    c ccc = cccc
    dd
    `
    
    p, err := Load(strings.NewReader(s))
    if nil != err {
        t.Error("加载失败")
        return
    }
    
    v := ""
    
    v = p.String("a")
    if "aa" != v {
        t.Error("Get string failed")
        return
    }
    
    v = p.String("b")
    if "bbb" != v {
        t.Error("Get string failed")
        return
    }
    
    v = p.String("Z")
    if "" != v {
        t.Error("Get string failed")
        return
    }
    
    v = p.String("c ccc")
    if "cccc" != v {
        t.Error("Get string failed")
        return
    }
    
    v = p.String("dd")
    if "" != v {
        t.Error("Get string failed")
        return
    }
}
