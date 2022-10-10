package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func main() {
	//normalCopy()

	//deepCopy_1()

	//deepCopy_2()
}

type Book struct {
}

func deepCopy_2(src []Book) (*[]Book, error) {
	var dst = new([]Book)
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, dst)
	return dst, err
}

func deepCopy_1(dst, src interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func normalCopy() {
	//切片拷贝，dst的结构类型必须与src相同，不然失败
	//a := []int{1}
	//b := make([]int, 1)
	//copy(b, a)
	//fmt.Printf("%v", b)

	//copy是浅拷贝,里面的类型也是引用
	//a := map[string]string{
	//	"name": "yuliyang",
	//}
	//arr := []map[string]string{a}
	//
	//arrCopy := make([]map[string]string, 1)
	//copy(arrCopy, arr)
	//a["name"] = "changed"
	//fmt.Printf("%v", arrCopy)
}
