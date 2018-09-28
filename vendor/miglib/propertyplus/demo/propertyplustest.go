package main

import "miglib/propertyplus"
import "fmt"
import "time"

func main() {
	// must provide propertyname and servername
	pp := propertyplus.CreatePropertyPlus("targoTest", "normal.DockerTest.TgwTest")

	vkey := []string{"normal.DockerTest.vinaTest.DSZ3", "normal.DockerTest.vinaTest", "test", "test"}
	vvalue := []float32{1, 2, 3}

	for {
		pp.Report(vkey, vvalue)
		fmt.Println("finish report")
		time.Sleep(10 * time.Second)
	}

}
