# 课后作业01

- 使用github上的lib：github.com/armstrongli/go-bmi 完成体脂计算器
- 本地添加module的replace，并在本地项目扩展 github.com/armstrongli/go-bmi 以支持BMP、FatRate的计算
- 使用vendor保证代码的完整性与可运行

## 程序执行

项目目录下，./fat_rate_calculator.go 包含 main 函数入口，即需要在项目根目录执行：

```shell
go run .
```

# 课后作业02

- 为体脂计算器编写单元测试并完善体脂计算器的验证逻辑
    - BMI计算
        1. 录入正常身高、体重，确保计算结果符合预期
        2. 录入0或负数身高，返回错误
        3. 录入0或负数体重，返回错误
    - 体脂率计算
        1. 录入正常BMI、年龄、性别，确保计算结果符合预期
        2. 录入非法BMI、年龄、性别（0、负数、超过150的年龄、非男女的性别输入），返回错误
        3. 录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期 