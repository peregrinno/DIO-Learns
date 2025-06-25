package main

import (
	"bytes"
	"container/list"
	"crypto/sha1"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"sort"
	"strings"
)

func learnOne() {
	str := " Papipaquigrafo forte"
	fmt.Println(strings.Contains(str, "paaa"))     // Verifica se a string contém a substring
	fmt.Println(strings.Count(str, "p"))           // Conta quantas vezes a substring aparece na string
	fmt.Println(strings.Replace(str, "p", "b", 1)) // Substitui a substring na string
	fmt.Println(strings.ToUpper(str))              // Converte a string para maiúsculo
	fmt.Println(strings.ToLower(str))              // Converte a string para minúsculo
	fmt.Println(strings.TrimSpace(str))            // Remove espaços em branco do início e do fim da string
}

func learnTwo() {
	var buffer bytes.Buffer
	buffer.WriteString("Hello Bytes")
	buffer.WriteString("World ")
	fmt.Println(buffer.String())

	buffer.Reset()
	fmt.Println(buffer.String())

	io.WriteString(&buffer, "Hello io")
	io.WriteString(&buffer, "World")
	fmt.Println(buffer.String())
}

func learnThree() {
	file, err := os.Create("learns.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.WriteString(file, "Hello io")
	io.WriteString(file, "World")
	file.Sync()
}

func learnFour() {
	file, err := os.Open("learns.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	content, err := os.ReadFile("learns.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func learnFive() {
	l := list.New()
	l.PushBack("Hello")
	l.PushBack("World")
	l.InsertAfter("!!", l.Back())
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)
}

type Dados struct {
	Nome  string
	Idade int
}
type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool {
	return ps[i].Nome < ps[j].Nome
}

func (ps ParaNome) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func learnSix() {
	criancas := []Dados{
		{"Pedro", 5},
		{"Maria", 7},
		{"João", 9},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(ParaNome(criancas))
}

func learnSeven() {
	//Encriptação
	senha := "minhasenha"
	h := sha1.New()
	h.Write([]byte(senha))
	bs := h.Sum([]byte{})
	fmt.Println(senha)
	fmt.Println(bs)

	//renomeia - arquivo
	nomeArquivo := "learns.txt"
	hash_nomeArquivo := fnv.New64a()
	os.Rename(nomeArquivo, "learns_"+string(hash_nomeArquivo.Sum(nil))+".txt")
}

func main() {
	//learnOne()
	//learnTwo()
	//learnThree()
	//learnFour()
	//learnFive()
	//learnSix()
	learnSeven()
}
