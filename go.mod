module github.com/xiangxiaoc/golearn-module-package

go 1.17

require github.com/armstrongli/go-bmi v0.0.1

replace (
	github.com/armstrongli/go-bmi v0.0.1 => ./my_bmi
)