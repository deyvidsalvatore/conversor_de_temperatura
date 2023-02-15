package main

import (
	"fmt"
)

/*
Um professor do ensino médio solicitou
aos seus alunos que desenvolvessem um programa
para converter o valor do ponto de ebulição da
água de Kelvin para graus Celsius
*/
func main() {
	conversaoKelvin_Celsius(302)
}

func conversaoKelvin_Celsius(temperaturaK float64) {
	fmt.Printf("A temperatura em Kevin para Celsius é: %.2f", (temperaturaK-273.15))
}
