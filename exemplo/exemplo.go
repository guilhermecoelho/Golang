package exemplo

import (
	"errors"
	"fmt"
	model "go/Models"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) andou() (string, error) {

	if p.Nome != "Guilherme" {
		return "", errors.New("nome tem que ser guilherme")
	}
	return p.Nome + " andou", nil
}

func (p *Pessoa) setNome(nome string) {
	p.Nome = nome
	fmt.Println(p.Nome)
}

func main() {

	// guilherme := Pessoa{
	// 	Nome:  "Guilherme",
	// 	Idade: 35,
	// }
	// res, err := guilherme.andou()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(res)

	// guilherme.setNome(("coelho"))
	// fmt.Println(guilherme.Nome)

	product1 := model.NewProduct()
	product1.Name = "Carrinho"

	product2 := model.NewProduct()
	product2.Name = "Boneca"

	products := model.Products{}
	products.Add(product1)
	products.Add(product2)

	s := make([]model.Product, 3)

	s[0] = *product1
	s[1] = *product2

	fmt.Println(s)
	fmt.Println(s[0].Name)
	fmt.Println(s[1])
	l := s[0:2]
	fmt.Println("sl1:", l)

	//fmt.Println(products.Product)

	// play := list.New()
	// play.PushBack(product1)
	// play.PushBack(product2)

	// A_elem := play.Back()
	// fmt.Println(A_elem.Value)
	// fmt.Println(A_elem.Value.(*model.Product).Name)

	// number := play.Len()
	// fmt.Println(number)
}
