package main
import "fmt"

// no go é sempre bom separar o seu dominio (regras de negocio do resto do codigo) no caso seria o print para podermos testar

func Ola() string { // Criamos uma nova função usando func, mas dessa vez adicionamos outra palavra reservada string na sua definição. Isso significa que essa função terá como retorno uma string (cadeia de caracteres).
	return "Olá, mundo!"
}

func Hello () {
	fmt.Println(Ola())
}