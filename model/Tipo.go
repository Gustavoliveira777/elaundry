package model

type Tipo struct {
	tipoId              int
	nome                string
	precoHora           float32
	precoFixoCapacidade float32
}

func TipoFactory(tipoId int, nome string, precoHora float32, precoFixoCapacidade float32) Tipo {
	return Tipo{tipoId: tipoId, nome: nome, precoHora: precoHora, precoFixoCapacidade: precoFixoCapacidade}
}

func (t *Tipo) SetTipoId(id int) {
	t.tipoId = id
}
func (t *Tipo) SetNome(nome string) {
	t.nome = nome
}
func (t *Tipo) SetPrecoHora(precoHora float32) {
	t.precoHora = precoHora
}
func (t *Tipo) SetPrecoFixoCapacidade(precoFixoCapacidade float32) {
	t.precoFixoCapacidade = precoFixoCapacidade
}

func (t Tipo) GetTipoId() int {
	return t.tipoId
}

func (t Tipo) GetNome() string {
	return t.nome
}

func (t Tipo) GetPrecoHora() float32 {
	return t.precoHora
}

func (t Tipo) GetPrecoFixoCapacidade() float32 {
	return t.precoHora
}
