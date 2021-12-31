package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"github.com/xiangxiaoc/golearn-module-package/type_in"
	"os"
)

func main() {
	name, sex, age, height, weight := type_in.TypeIN()
	bmi, err := gobmi.BMI(weight, height)
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	fatRate, err2 := gobmi.CalFatRate(sex, bmi, age)
	if err2 != nil {
		fmt.Println("ERROR: ", err2)
		os.Exit(1)
	}

	fmt.Println("======= 输出结果 =======")
	fmt.Printf("姓名：%s BMI：%.2f 体脂率：%.2f", name, bmi, fatRate)

}
