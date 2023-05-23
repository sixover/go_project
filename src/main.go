package main

import "fmt"

func main() {
	//cmd := exec.Command("ipconfig")
	//cmd.Stdout = os.Stdout
	//err := cmd.Run()
	//if err != nil {
	//	return
	//}
	//fmt.Println(cmd.Start())
	//out, err := exec.Command("ipconfig").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf(string(out))
	a := [5]int{1, 2, 3, 4, 5}
	t := a[:2:3]
	fmt.Println(t[1])
}
