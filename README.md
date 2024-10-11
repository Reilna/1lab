Написать программу, которая выводит текущее время и дату.

Создать переменные различных типов (int, float64, string, bool) и вывести их на экран.

Использовать краткую форму объявления переменных для создания и вывода переменных.

Написать программу для выполнения арифметических операций с двумя целыми числами и выводом результатов.

Реализовать функцию для вычисления суммы и разности двух чисел с плавающей запятой.

Написать программу, которая вычисляет среднее значение трех чисел.

  ```package main

  import (
  	"fmt"
  	"time"
  )
  
  func main() {
  
  	//task_3
  	dateNow := time.Now()
  	myInteger := int64(1111)
  	myFloat := float64(11.11)
  	myString := string("'1'")
  	myBool := bool(myInteger >= 1000)
  
  	fmt.Print(
  
  		"\ntask_1\n", "Date/time today is: ", dateNow.Format("2006-01-02 15:04:05"), "\n\n",
  		"task_2\n", "int64: ", myInteger, "\n",
  		"float64: ", myFloat, "\n",
  		"string: ", myString, "\n",
  		"bool: ", myBool, "\n")
  
  	IntegerOperation(myInteger, myInteger)
  	FloatOperation(myFloat, myFloat)
  	AverageValue(myInteger, myInteger, myFloat)
  
  }
  
  func IntegerOperation(a int64, b int64) {
  
  	fmt.Print(
  		"\ntask_4\n", "summ myInteger: ", a+b, "\n",
  		"mutipl myInteger: ", a*b, "\n")
  }
  
  func FloatOperation(a float64, b float64) {
  
  	fmt.Print(
  		"\ntask_5\n", "summ myFloat: ", a+b, "\n",
  		"substr myFloat: ", a-b, "\n")
  }
  
  func AverageValue(a int64, b int64, c float64) {
  
  	average := (float64(a) + float64(b) + c) / 3
  
  	fmt.Print("\ntask_6\n", "Average myInt, myInt, myFl: ", average)
  
  }
```
![Image alt](https://github.com/Reilna/1lab/blob/main/firstLab/image.png)
запускаемый файл - main.go
или же в терминале go run "путь до файла с расширением"
