package main // Pacote principal da aplicação

import "fmt" // Pacote externo para print
import "reflect" // Pacote para descobrir o tipo das variaveis

func main() { // Função da aplicação principal

	// Tipos da variaveis:
	// O go não deixa declarar uma variavel e não atribuir uso a ela, ele não compila
	// Tipo String
	// Quando não for atribuido valor da variavel ele assumira o valor vazio
	var nome string = "Lucas"
	nome2 := "Emilly" // Também pode ser declarada dessa forma

	//Tipo Float64 ou Float32 bits para numeros maiores ou menores
	// Quando não for atribuido valor da variavel ele assumira o valor 0
	var versao float32 = 1.1
	var versao2 float64 = 1.232323113131
	versao3 := 1.223

	// Tipo Inteiro
	// Quando não for atribuido valor da variavel ele assumira o valor 0
	var idade int = 18 
	idade2 := 20

	fmt.Println("Olá, sr. ", nome, "sua idade é", idade)
	fmt.Println("Olá, sr. ", nome2, "sua idade é", idade2)

	
	fmt.Println("Esse programa está na versão", versao)
	fmt.Println("A versão extendida desse programa está na", versao2)
	fmt.Println("A versão extendida desse programa está na", versao3)


	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))
	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(idade))
}
