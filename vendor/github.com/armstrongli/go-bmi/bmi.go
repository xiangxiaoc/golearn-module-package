package my_bmi

import "fmt"

func BMI(weightKG, heightM float64) (bmi float64, err error) {
	// 重写 BMI 函数
	if weightKG <= 0 {
		err = fmt.Errorf("体重不能为负数或为 0")
		return
	}
	if heightM <= 0 {
		err = fmt.Errorf("身高不能为负数或为 0")
		return
	}
	bmi = weightKG / (heightM * heightM)
	return
}

// 扩展 gobmi 包的能力

func CalFatRate(sex string, bmi float64, age int) (fatRate float64, err error) {
	var sexWeight int
	switch sex {
	case "男":
		sexWeight = 1
	case "女":
		sexWeight = 0
	default:
		err = fmt.Errorf("性别只能输入'男'或'女'")
		return
	}
	if bmi == 0 {
		err = fmt.Errorf("BMI 计算错误，体重是不是输入 0 了？")
		return
	}
	if age <= 0 {
		err = fmt.Errorf("年龄输入有误，不能为负数")
	} else if age > 150 {
		err = fmt.Errorf("年龄不能超过 150 啊")
	}

	fatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexWeight)) / 100
	return
}

func GetSugesstion(fatRate float64) (suggestion string, err error) {
	if fatRate >= 0 && fatRate < 0.1 {
		suggestion = "偏瘦，多补充营养"
		return
	} else if fatRate >= 0.1 && fatRate < 0.16 {
		suggestion = "标准，太棒了，继续保持"
		return
	} else if fatRate >= 0.16 && fatRate < 0.21 {
		suggestion = "偏重，需要加强锻炼"
		return
	} else if fatRate >= 0.21 && fatRate < 0.26 {
		suggestion = "肥胖，开始计划锻炼，得减重了"
		return
	} else if fatRate >= 0.26 {
		suggestion = "严重肥胖，警告！必须制定计划，恢复健康"
		return
	} else {
		err = fmt.Errorf("体脂率为负数，好像输入的数据不对呢 %v", fatRate)
		return
	}
}
