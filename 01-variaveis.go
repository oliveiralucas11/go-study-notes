package main // Indica que é o pacote principal da aplicacao e seu codigo deve comecar por ele.
import "fmt" // Pacote de formatacao que contem as classes de imprimir e formatar o texto
import "reflect" // Habilita a funcao typeof para descobrir quais sao os tipos da variaveis

func main() { // Função da aplicação principal
	
	//Variaveis
	//int, string, float32, float64
	// Quando nao atribuimos valor a uma variavel, o go entende que seu valor é 0 ou um caracter vazio
	// O go nao deixa declarar uma variavel e nao usa-la, logo ele nao compila até a variavel ser usada

	var nome string = "Lucas" // Variavel declarada de forma explicita e tipada
	var versao float32 = 1.1 // Float possui de 32bits ou de 64bits para numeros maiores
	var idade int = 32 // Variavel do tipo inteiro
	var idade1 int // Quando nao atribuimos valor a variavel o go entende que o valor é 0

	var versao1 = 3.3 // Variavel sem tipar para descobrir seu valor com o import reflect

	fmt.Println("Ola, sr.", nome, "sua idade é ", idade) // , usada para concatenar
	fmt.Println("Este programa está na versao", versao, "com sua idade", idade1)

	fmt.Println("O tipo da varaivel versao1 é", reflect.TypeOf(versao1)) // reflect.typeof usada para descobrir qual a tipagem de uma variavel

	// Outra forma de declarar variavel sem precisar atribuir o valor var
	// Esse é o tipo de variavel mais usada

	nome1 := "Lucas"
	versao2 := 2.2
	idade2 := 20

	fmt.Println("Meu nome é ", nome1)
	fmt.Println("Minha idade é: ", idade2)
	fmt.Println("Estou atualizado na versao: ", versao2)

}
