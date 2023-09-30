package tools

import "fmt"

type Double[Ta interface{}, Tb interface{}] struct {
	First  Ta
	Second Tb
}

func NewDouble[Ta interface{}, Tb interface{}](first Ta, second Tb) Double[Ta, Tb] {
	return Double[Ta, Tb]{first, second}
}

func (d Double[Ta, Tb]) String() string {
	return fmt.Sprintf("(%v, %v)", d.First, d.Second)
}
