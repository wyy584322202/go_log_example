package main

import "log"
import "os"

func logInit() {
	file, err := os.Create("2016-12-22-test.log")
	if err != nil {
		log.Fatal("open log file error")
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("debug::")
}
func main() {
	logInit()
	log.Println("test1")
	log.Println("test2")
	log.Println("test3")
	log.Println("test4")

}
