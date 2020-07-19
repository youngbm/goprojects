package main

import "fmt"

func myPrint(i *int){  // ������븱�����ӡ�����ĵ��ǵ���ǰ�ĸ���
    fmt.Println(*i)
}


type People struct {
    name string
}

type Teacher struct {
    People  // ���ﲻ�� p People �̳���People
}

func (p People) ShowA(){
      fmt.Println("p ShowA")
      p.ShowB()
}

func (p People) ShowB(){
      fmt.Println("p ShowB")
}

func (t Teacher) ShowB(){
      fmt.Println("t ShowB")
}


func main(){
    var i = 5
    defer myPrint(&i) // �������ȥ�Ƿ�ָ��or��ַ�����ǵ�ǰ�ĸ�������ӡ����5
    i =  10
    defer myPrint(&i) // �������ȥ�Ƿ�ָ��or��ַ�����ӡ����10
    i = i + 10
    myPrint(&i)

    // Next
    var t Teacher
    t.ShowA()
}