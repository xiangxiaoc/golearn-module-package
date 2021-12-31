package type_in

import "fmt"

func TypeIN() (name string, sex string, age int, height float64, weight float64) {

	fmt.Println("======= 录入信息 =======")
	fmt.Print("姓名：")
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	fmt.Print("性别（男/女）：")
	_, err2 := fmt.Scanln(&sex)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("年龄：")
	_, err3 := fmt.Scanln(&age)
	if err3 != nil {
		panic(err3)
	}

	fmt.Print("身高（米）：")
	_, err4 := fmt.Scanln(&height)
	if err4 != nil {
		panic(err4)
	}

	fmt.Print("体重（公斤）：")
	_, err5 := fmt.Scanln(&weight)
	if err5 != nil {
		panic(err5)
	}

	return
}
