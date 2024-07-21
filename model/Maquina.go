package model

type Maquina struct {
	maquinaId    int
	marca        string
	tempoCiclo   int
	capacidade   int
	instrucaoUso string
	tipoMaquina  Tipo
}

func MaquinaFactory(maquinaId int, marca string, tempoCiclo int, capacidade int, instrucaoUso string, tipoMaquina Tipo) Maquina {
	return Maquina{maquinaId: maquinaId, marca: marca, tempoCiclo: tempoCiclo, capacidade: capacidade, instrucaoUso: instrucaoUso, tipoMaquina: tipoMaquina}
}

func (m *Maquina) SetMaquinaId(id int) {
	m.maquinaId = id
}

func (m *Maquina) SetMarca(marca string) {
	m.marca = marca
}
func (m *Maquina) SetTempoCiclo(tempoCiclo int) {
	m.tempoCiclo = tempoCiclo
}
func (m *Maquina) SetCapacidade(capacidade int) {
	m.capacidade = capacidade
}
func (m *Maquina) SetInstrucaoUso(instrucaoUso string) {
	m.instrucaoUso = instrucaoUso
}
func (m *Maquina) SetTipoMaquina(tipoMaquina Tipo) {
	m.tipoMaquina = tipoMaquina
}

func (m Maquina) GetMaquinaId() int {
	return m.maquinaId
}

func (m Maquina) GetMarca() string {
	return m.marca
}

func (m Maquina) GetTempoCiclo() int {
	return m.tempoCiclo
}

func (m Maquina) GetCapacidade() int {
	return m.capacidade
}

func (m Maquina) GetInstrucaoUso() string {
	return m.instrucaoUso
}

func (m Maquina) GetTipoMaquina() Tipo {
	return m.tipoMaquina
}
