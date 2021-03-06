package model

type Venda struct {
	numero  int
	data    Data
	cliente Cliente
	itens   []ItemVenda
}

func (v *Venda) NovaVenda(numero int, data Data, cliente Cliente, itens []ItemVenda) {
	v.SetNumero(numero)
	v.SetData(data)
	v.SetCliente(cliente)
	v.SetItemVenda(itens)
}

func (v *Venda) GetNumero() int {
	return v.numero
}

func (v *Venda) SetNumero(numero int) {
	v.numero = numero
}

func (v *Venda) GetData() Data {
	return v.data
}

func (v *Venda) SetData(data Data) {
	v.data = data
}

func (v *Venda) GetCliente() Cliente {
	return v.cliente
}

func (v *Venda) SetCliente(cliente Cliente) {
	v.cliente = cliente
}

func (v *Venda) GetItemVenda() []ItemVenda {
	return v.itens
}

func (v *Venda) SetItemVenda(itens []ItemVenda) {
	v.itens = itens
}

func (v *Venda) Total() float32 {
	var soma float32 = 0
	for i := range v.itens {
		soma += v.itens[i].Total()
	}
	return soma
}
