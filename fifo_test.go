package fifo

import (
	"testing"
)

func TestFIFO(t *testing.T) {
	f := New[int]()

	f.Add(1)
	f.Add(2)
	f.Add(3)
	if size := f.Size(); size != 3 {
		t.Errorf("Tamanho incorreto da FIFO. Esperado: 3, Obtido: %d", size)
	}

	if ok, _ := f.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO neste momento")
	}
	if ok, _ := f.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO neste momento")
	}
	if ok, _ := f.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO neste momento")
	}

	if size := f.Size(); size != 0 {
		t.Errorf("Tamanho incorreto da FIFO após a remoção. Esperado: 0, Obtido: %d", size)
	}
	if !f.IsEmpty() {
		t.Error("A FIFO deveria estar vazia neste momento")
	}

	// Teste de adição e remoção de elementos de diferentes tipos
	f2 := New[string]()
	f2.Add("foo")
	f2.Add("bar")
	f2.Add("baz")
	if size := f2.Size(); size != 3 {
		t.Errorf("Tamanho incorreto da FIFO de strings. Esperado: 3, Obtido: %d", size)
	}

	if ok, _ := f2.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO de strings neste momento")
	}
	if ok, _ := f2.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO de strings neste momento")
	}
	if ok, _ := f2.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO de strings neste momento")
	}
	if size := f2.Size(); size != 0 {
		t.Errorf("Tamanho incorreto da FIFO de strings após a remoção. Esperado: 0, Obtido: %d", size)
	}
	if !f2.IsEmpty() {
		t.Error("A FIFO de strings deveria estar vazia neste momento")
	}
}

func TestFIFO_NextWhenEmpty(t *testing.T) {
	f := New[int]()

	// Teste quando a FIFO está vazia
	if ok, item := f.Next(); ok || item != 0 {
		t.Errorf("Próximo elemento incorreto. Esperado: false, 0. Obtido: %v, %v", ok, item)
	}

	// Teste quando a FIFO está vazia com tipo de referência
	f2 := New[*string]() // Cria uma nova FIFO de ponteiros para strings
	if ok, item := f2.Next(); ok || item != nil {
		t.Errorf("Próximo elemento incorreto. Esperado: false, nil. Obtido: %v, %v", ok, item)
	}
}

type Person struct {
	Name string
	Age  int
}

type Printable interface {
	Print() string
}

func (p Person) Print() string {
	return p.Name
}

func TestStructFIFO(t *testing.T) {
	f := New[Person]()

	f.Add(Person{Name: "Alice", Age: 25})
	f.Add(Person{Name: "Bob", Age: 30})
	if size := f.Size(); size != 2 {
		t.Errorf("Tamanho incorreto da FIFO. Esperado: 2, Obtido: %d", size)
	}

	if ok, _ := f.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO neste momento")
	}
	if ok, _ := f.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO neste momento")
	}

	if size := f.Size(); size != 0 {
		t.Errorf("Tamanho incorreto da FIFO após a remoção. Esperado: 0, Obtido: %d", size)
	}
	if !f.IsEmpty() {
		t.Error("A FIFO deveria estar vazia neste momento")
	}

	f2 := New[Printable]()
	person1 := Person{Name: "Alice", Age: 25}
	person2 := Person{Name: "Bob", Age: 30}
	f2.Add(person1)
	f2.Add(person2)
	if size := f2.Size(); size != 2 {
		t.Errorf("Tamanho incorreto da FIFO de Printable. Esperado: 2, Obtido: %d", size)
	}

	if ok, _ := f2.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO de Printable neste momento")
	}
	if ok, _ := f2.Next(); !ok {
		t.Error("Deveria haver elementos na FIFO de Printable neste momento")
	}
	if size := f2.Size(); size != 0 {
		t.Errorf("Tamanho incorreto da FIFO de Printable após a remoção. Esperado: 0, Obtido: %d", size)
	}
	if !f2.IsEmpty() {
		t.Error("A FIFO de Printable deveria estar vazia neste momento")
	}
}
