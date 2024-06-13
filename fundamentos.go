package main // Indica que é o pacote principal da aplicacao e seu codigo deve comecar por ele.
import "fmt" // Pacote de formatacao que contem as classes de imprimir e formatar o texto

func main() { // Função da aplicação principal



	//Variaveis
	//int, string, float32, float64
	// Quando nao atribuimos valor a uma variavel, o go entende que seu valor é 0 ou um caracter vazio
	// O go nao deixa declarar uma variavel e nao usa-la, logo ele nao compila até a variavel ser usada

	var nome string = "Lucas" // Variavel declarada de forma explicita e tipada
	var versao float32 = 1.1 // Float possui de 32bits ou de 64bits para numeros maiores
	var idade int = 32 // Variavel do tipo inteiro
	fmt.Println("Ola, sr.", nome, "sua idade é ", idade) // , usada para concatenar
	fmt.Println("Este programa está na versao", versao, "com sua idade", idade1)

	var idade1 int // Quando nao atribuimos valor a variavel o go entende que o valor é 0

}
