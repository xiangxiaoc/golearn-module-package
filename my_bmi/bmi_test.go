package my_bmi

import (
	"testing"
)
import "github.com/shopspring/decimal"

var _HeightM, _WeightKG float64

// BMI 函数相关测试

func TestCalBMI(t *testing.T) {
	_HeightM, _WeightKG = 1.7, 70
	expectedBMI := 24.22
	calResult, _ := BMI(_WeightKG, _HeightM)
	roundedResult, _ := decimal.NewFromFloat(calResult).Round(2).Float64()
	if roundedResult != expectedBMI {
		t.Errorf("BMI计算与预期不符，计算结果为：%v", calResult)
	}
}

func TestTypeHeightError(t *testing.T) {
	_HeightM, _WeightKG = 0, 70
	_, err := BMI(_WeightKG, _HeightM)
	if err.Error() != "身高不能为负数或为 0" {
		t.Errorf("没有判断出身高输入错误，输入身高为：%v", _HeightM)
	}
}

func TestTypeWeightError(t *testing.T) {
	_HeightM, _WeightKG = 1.7, -10
	_, err := BMI(_WeightKG, _HeightM)
	if err.Error() != "体重不能为负数或为 0" {
		t.Errorf("没有判断出体重输入错误，输入体重为：%v", _WeightKG)
	}
}

// 体脂率相关测试

var (
	_bmi float64
	_age int
	_sex string
)

func TestCalFatRate(t *testing.T) {
	_bmi, _age, _sex = 24.22, 28, "男"
	expectedFatRate := 0.19304
	calResult, _ := CalFatRate(_sex, _bmi, _age)
	if calResult != expectedFatRate {
		t.Errorf("体脂率计算与预期不符，计算结果为: %v", calResult)
	}
}

func TestTypeErrorForCalFatRate(t *testing.T) {
	{
		_bmi, _age, _sex = 0, 28, "女"
		_, err := CalFatRate(_sex, _bmi, _age)
		if err.Error() != "BMI 计算错误，体重是不是输入 0 了？" {
			t.Errorf("函数没有抓到BMI的计算错误")
		}
	}
	{
		_bmi, _age, _sex = 24.22, -10, "男"
		_, err := CalFatRate(_sex, _bmi, _age)
		if err.Error() != "年龄输入有误，不能为负数" {
			t.Errorf("函数没有抓到年龄为负数的输入错误")
		}
	}
	{
		_bmi, _age, _sex = 24.22, 160, "男"
		_, err := CalFatRate(_sex, _bmi, _age)
		if err.Error() != "年龄不能超过 150 啊" {
			t.Errorf("函数没有抓到年龄超过150的输入限制")
		}
	}
	{
		_bmi, _age, _sex = 24.22, 18, "非男女"
		_, err := CalFatRate(_sex, _bmi, _age)
		if err.Error() != "性别只能输入'男'或'女'" {
			t.Errorf("函数没有抓到性别输入的错误")
		}
	}
}

func TestSuggestion(t *testing.T) {
	_sex, _age, _HeightM, _WeightKG = "男", 28, 1.7, 70
	expectedSuggestion := "偏重，需要加强锻炼"
	actualBMI, _ := BMI(_WeightKG, _HeightM)
	actualFatRate, _ := CalFatRate(_sex, actualBMI, _age)
	actualSuggestion, _ := GetSugesstion(actualFatRate)
	if actualSuggestion != expectedSuggestion {
		t.Errorf("与预期的建议不符")
	}
}
