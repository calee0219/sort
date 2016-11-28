package sort_test

import (
  "github.com/go-Algorithm/sort"
  "fmt"
)

type student struct{
  id    int
  name  string
}

func cmp(i,j interface{})bool{
  ii:=i.(student)
  jj:=j.(student)
  return ii.id < jj.id
}

func Example(){
  stu:=[]student{{5,"sam"},{3,"lily"},{7,"jacky"},{1,"willy"},{2,"steve"}}
  fmt.Println("original: \t",stu)

  out:=sort.MergeSortPara(stu,cmp)
  fmt.Println("[]interface{}: \t",out)

  stu2:=make([]student,len(out))
  for a:=0;a<len(out);a++{
    stu2[a]=out[a].(student)
  }
  fmt.Println("[]student: \t",stu2)
  // Output:
  // original:      [{5 sam} {3 lily} {7 jacky} {1 willy} {2 steve}]
  // []interface{}: [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]
  // []student:     [{1 willy} {2 steve} {3 lily} {5 sam} {7 jacky}]
}