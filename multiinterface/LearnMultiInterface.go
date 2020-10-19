package main

func main() {
	//var a readableWriteable
	//a = ReadableImpl{"haha"}
	//b := ReadableImpl{"haha"}
	//fmt.Println(reflect.DeepEqual(a, b))
	//a.read("hello")
	//fmt.Println(a.write("go"))
}

type readable interface {
	read(content string)
}

type writeable interface {
	write(content string) string
}

type readableWriteable interface {
	readable
	writeable
}
