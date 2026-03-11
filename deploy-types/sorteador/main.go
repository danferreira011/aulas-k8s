package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	numEscolhido, _ := strconv.Atoi(os.Getenv("NUM_ESCOLHIDO"))
	numMaximo, _ := strconv.Atoi(os.Getenv("NUM_MAXIMO"))

	delayStr := os.Getenv("DELAY_SORTEIO")
	if delayStr == "" {
		fmt.Println("A variável DELAY_SORTEIO não está definida. Usando delay padrão de 5 segundos.")
		delayStr = "5"
	}

	delay, err := strconv.Atoi(delayStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sorteando...")
	time.Sleep(time.Duration(delay) * time.Second)

	numeroSorteado := rand.Intn(numMaximo + 1)
	if numeroSorteado == numEscolhido {
		fmt.Printf("O valor sorteado foi %d e você escolheu %d. Parabéns !!!\n", numeroSorteado, numEscolhido)
	} else {
		fmt.Printf("O valor sorteado foi %d e você escolheu %d. Tenta de novo...\n", numeroSorteado, numEscolhido)
		os.Exit(1)
	}
}
