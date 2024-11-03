package main

import (
	"fmt"
)

type Variaveis []rune

func (p Variaveis) toString() string {
	s := ""
	for _, v := range p {
		s += fmt.Sprintf("%c", v)
	}
	return s
}

type Abstracao interface {
	Tipo() int
	Comparar(rune) bool
	ToString() string
	Interior() Lambda
}

const (
	NUMERO = iota
	VARIAVEL
	FUNCAO
)

type Numero struct {
	Num int
}

func (n Numero) Tipo() int {
	return NUMERO
}

func (n Numero) Comparar(_ rune) bool {
	return false
}

func (n Numero) ToString() string {
	return fmt.Sprintf("%d", n.Num)
}

func (n Numero) Interior() Lambda {
	return Lambda{}
}

type Variavel struct {
	Var rune
}

func (v Variavel) Tipo() int {
	return VARIAVEL
}

func (v Variavel) Comparar(r rune) bool {
	return r == v.Var
}

func (v Variavel) ToString() string {
	return fmt.Sprintf("%c", v.Var)
}

func (n Variavel) Interior() Lambda {
	return Lambda{}
}

type Lambda struct {
	Vars Variaveis
	Func []Abstracao
}

func (l Lambda) Printar() {
	fmt.Println(l.ToString())
}

func (l Lambda) Tipo() int {
	return FUNCAO
}

func (l Lambda) Comparar(_ rune) bool {
	return false
}

func (l Lambda) ToString() string {
	s := ""
	if len(l.Vars) >= 1 {
		s += fmt.Sprintf("λ%v.", l.Vars.toString())
	}
	for _, v := range l.Func {
		s += v.ToString()
	}
	return s
}

func (l *Lambda) Beta(b Abstracao) {
	fmt.Printf("reducao beta com arg: %v\n", b.ToString())
	if len(l.Vars) == 0 {
		return
	}
	vari := l.Vars[0]
	l.Vars = l.Vars[1:]
	for k, v := range l.Func {
		if v.Tipo() == VARIAVEL && v.Comparar(vari) {
			l.Func[k] = b
		}
	}
	l.RetirarCamada()
}

func (l *Lambda) RetirarCamada() {
	if len(l.Vars) == 0 && len(l.Func) == 1 && l.Func[0].Tipo() == FUNCAO {
		interior := l.Func[0].Interior()
		l.Vars = interior.Vars
		l.Func = interior.Func
	}
}

func (l Lambda) Interior() Lambda {
	return l
}

func main() {
	f1 := Lambda{Vars: Variaveis{'a'}, Func: []Abstracao{Variavel{'a'}}}
	f2 := Lambda{Vars: Variaveis{'b'}, Func: []Abstracao{Variavel{'b'}}}
	f1.Printar()
	f1.Beta(f2)
	f1.Printar()
	f1.Beta(Numero{1})
	f1.Printar()
}
