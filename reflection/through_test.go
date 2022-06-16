package reflection

import (
	"reflect"
	"testing"
)

// A reflexão em computação é a habilidade de um programa examinar sua própria estrutura
// particularmente através de tipos; é uma forma de metaprogramação. Também é uma ótima fonte de confusão.

func TestThrought(t *testing.T) {

	cases := []struct {
		Name        string
		Entry       interface{}
		ExpectCalls []string
	}{
		{
			"struct with string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string

			Through(test.Entry, func(entry string) {
				result = append(result, entry)
			})

			if !reflect.DeepEqual(result, test.ExpectCalls) {
				t.Errorf("result %s, expect %s", result, test.ExpectCalls)
			}
		})
	}

}
