package main // Pacote principal da aplicação

import "fmt" // Pacote externo para print

func main() { // Função da aplicação principal

	// Tipos da variaveis:
	// O go não deixa declarar uma variavel e não atribuir uso a ela, ele não compila
	// Tipo String
	// Quando não for atribuido valor da variavel ele assumira o valor vazio
	var nome string = "Lucas"

	//Tipo Float64 ou Float32 bits para numeros maiores ou menores
	// Quando não for atribuido valor da variavel ele assumira o valor 0
	var versao float32 = 1.1
	var versao2 float64 = 1.232323113131

	// Tipo Inteiro
	// Quando não for atribuido valor da variavel ele assumira o valor 0
	var idade int = 18 

	fmt.Println("Olá, sr. ", nome)
	fmt.Println("Você tem", idade, "anos")
	fmt.Println("Esse programa está na versão", versao)
	fmt.Println("A versão extendida desse programa está na", versao2)
}
