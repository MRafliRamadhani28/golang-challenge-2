package soal

import "fmt"

func calculateBMI(weight, height *float32) float32 {
	return *weight / (*height * *height)
}

func SoalKedua() {
	fmt.Println("\nSoal Kedua : ")

	// Data Uji 1
	var markWeight float32 = 78.0
	var markHeight float32 = 1.69
	var johnWeight float32 = 92.0
	var johnHeight float32 = 1.95

	// calculate BMI for Mark and John
	markBMI := calculateBMI(&markWeight, &markHeight)
	johnBMI := calculateBMI(&johnWeight, &johnHeight)

	markHigherBMI := markBMI > johnBMI

	// show results BMI for Mark and John and comparation who has higher
	fmt.Println("Data Uji 1 : ")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Does Mark have a higher BMI than John? %t\n\n", markHigherBMI)

	// Data uji 2
	markWeight = 95.0
	markHeight = 1.88
	johnWeight = 85.0
	johnHeight = 1.76

	// calculate BMI for Mark and John
	markBMI = calculateBMI(&markWeight, &markHeight)
	johnBMI = calculateBMI(&johnWeight, &johnHeight)

	markHigherBMI = markBMI > johnBMI

	// show results BMI for Mark and John and comparation who has higher
	fmt.Println("Data Uji 2 : ")
	fmt.Printf("BMI Mark: %.2f\n", markBMI)
	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("Does Mark have a higher BMI than John? %t\n", markHigherBMI)
}
