package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	para := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer suscipit orci id lorem consectetur, in mollis justo ullamcorper. Cras pellentesque magna et ipsum facilisis molestie. Nullam tempor tortor est, a consequat libero blandit vel. Pellentesque quis posuere risus, et iaculis ipsum. Curabitur vel faucibus metus, ac feugiat urna. Mauris semper turpis vel nulla cursus, ut dapibus erat auctor. Donec feugiat aliquet pulvinar. Aliquam nunc odio, vehicula eget sollicitudin ac, euismod ac justo. Quisque hendrerit non sem vitae aliquam. Aenean eu vulputate ante."

	filename, err := os.Create("./file.txt")
	checkErr(err)

	fileInfo, err := filename.Stat()
	checkErr(err)

	io.WriteString(filename, para)
	checkErr(err)

	fmt.Println("Length of file is: ", fileInfo.Size())

}

func checkErr(err error) {
	if err != nil {
		panic("unimplemented")
	}
}
