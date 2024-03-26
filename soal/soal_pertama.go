package soal

import "fmt"

func calculateAvg(number *[]float32) float32 {
	var result float32
	for _, value := range *number {
		result += value
	}
	return result / float32(len(*number))
}

func SoalPertama() {
	fmt.Println("Soal Pertama : ")

	scoreKoala := []float32{88, 91, 110}
	scoreLumbaLumba := []float32{96, 108, 89}

	avgScoreKoala := calculateAvg(&scoreKoala)
	avgScoreLumbaLumba := calculateAvg(&scoreLumbaLumba)

	// show average
	fmt.Printf("Average score team koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Average score team lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("The Results : ")

	// compare score
	if avgScoreKoala == avgScoreLumbaLumba {
		fmt.Println("Draw")
	} else if avgScoreKoala > avgScoreLumbaLumba {
		fmt.Println("Koala Team winner")
	} else {
		fmt.Println("Lumba-lumba Team winner")
	}

	fmt.Println("\nBonus 1 : ")

	scoreKoala = []float32{109, 95, 123}
	scoreLumbaLumba = []float32{97, 112, 101}

	avgScoreKoala = calculateAvg(&scoreKoala)
	avgScoreLumbaLumba = calculateAvg(&scoreLumbaLumba)

	// show average with minimal score 100
	fmt.Printf("Average score team koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Average score team lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("The Results : ")

	// compare score with minimal score 100
	if avgScoreKoala == avgScoreLumbaLumba {
		fmt.Println("Draw")
	} else if (avgScoreKoala > avgScoreLumbaLumba) && (avgScoreKoala >= 100) {
		fmt.Println("Koala Team winner")
	} else if (avgScoreKoala < avgScoreLumbaLumba) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba Team winner")
	} else {
		fmt.Println("There are no winners")
	}

	fmt.Println("\nBonus 2 : ")

	scoreKoala = []float32{109, 95, 106}
	scoreLumbaLumba = []float32{97, 112, 101}

	avgScoreKoala = calculateAvg(&scoreKoala)
	avgScoreLumbaLumba = calculateAvg(&scoreLumbaLumba)

	// show average with minimal score 100
	fmt.Printf("Average score team koala: %.2f\n", avgScoreKoala)
	fmt.Printf("Average score team lumba-lumba: %.2f\n", avgScoreLumbaLumba)

	fmt.Print("The Results : ")

	// compare score with minimal score 100
	if (avgScoreKoala == avgScoreLumbaLumba) && (avgScoreKoala >= 100) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Draw")
	} else if (avgScoreKoala > avgScoreLumbaLumba) && (avgScoreKoala >= 100) {
		fmt.Println("Koala Team winner")
	} else if (avgScoreKoala < avgScoreLumbaLumba) && (avgScoreLumbaLumba >= 100) {
		fmt.Println("Lumba-lumba Team winner")
	} else {
		fmt.Println("There are no winners")
	}
}
