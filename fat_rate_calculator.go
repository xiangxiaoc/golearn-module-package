package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"github.com/xiangxiaoc/golearn-module-package/type_in"
)

func main() {
	name, sex, age, height, weight := type_in.TypeIN()
	bmi, err := gobmi.BMI(weight, height)
	if err != nil {
		return
	}

	fatRate := gobmi.CalFatRate(sex, bmi, age)

	fmt.Println("======= 输出结果 =======")
	fmt.Printf("姓名：%s BMI：%.2f 体脂率：%.2f", name, bmi, fatRate)

}
