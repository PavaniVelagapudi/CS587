package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}


	data    := MyType { 5,3,2,1 }
	efmt.Println("values before operation",data.x1,data.y1,data.x2,data.y2)

	efmt.Println("values of Difference between X Co-Ordinates",(data.x1-data.x2))

	efmt.Println("values of Difference between Y co-ordinates",(data.y1-data.y2))
	data.Write(fd)
	data.WriteVar(path +"foobar")

}
