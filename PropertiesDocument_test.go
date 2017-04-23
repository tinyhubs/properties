package properties

import (
    "testing"
    "strings"
    "bytes"
    "fmt"
    "os"
)

func Test_Load(t *testing.T) {
    s := `
    a=aa
    b=bbb
    c ccc = cccc
    dd
    ee: r-rt rr
    `
    file, err := os.Open("test1.properties")
    if nil != err {
        return
    }
    //p, err := Load(strings.NewReader(s))
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
    
    v = p.String("ee")
    if "r-rt rr" != v {
        t.Error("Get string failed")
        return
    }
}


func Test_LoadFromFile(t *testing.T) {
    file, err := os.Open("test1.properties")
    if nil != err {
        return
    }
    
    doc, err := Load(file)
    if nil != err {
        t.Error("加载失败")
        return
    }
    
    fmt.Println(doc.String("key"))
}

func Test_New(t *testing.T) {
    doc := New()
    doc.Set("a", "aaa")
    doc.Comment("a", "This is a comment for a")
    
    buf := bytes.NewBufferString("")
    Save(doc, buf)
    
    if "#This is a comment for a\na=aaa\n" != buf.String() {
        fmt.Println("Dump failed:[" + buf.String() + "]")
        t.Error("Dump failed")
        return
    }
}

func Test_Save(t *testing.T) {
    doc := New()
    doc.Set("a", "aaa")
    doc.Comment("a", "This is a comment for a")
    
    buf := bytes.NewBufferString("")
    Save(doc, buf)
    
    if "#This is a comment for a\na=aaa\n" != buf.String() {
        t.Error("Dump failed")
        return
    }
}


func Test_Comment(t *testing.T) {
    doc := New()
    doc.Set("a", "aaa")
    doc.Comment("a", "This is a \ncomment \nfor a")
    
    buf := bytes.NewBufferString("")
    Save(doc, buf)
    
    if "#This is a \n#comment \n#for a\na=aaa\n" != buf.String() {
        t.Error("Dump failed")
        return
    }
}

