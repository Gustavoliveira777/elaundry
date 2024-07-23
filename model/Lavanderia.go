package model

import "time"

type Lavanderia struct {
	EmpresaId         int
	HorarioAbertura   time.Time
	HorarioFechamento time.Time
	DiasFuncionamento []string
	Maquinas          []Maquina
	Usuarios          []Usuario
	Aberto            bool
}

/* func LavanderiaFactory(empresaId int, horarioAbertura time.Time, horarioFechamento time.Time, diasFuncionamento []string, maquinas []Maquina, usuarios []Usuario, aberto bool) Lavanderia {
	return Lavanderia{
		empresaId:         empresaId,
		horarioAbertura:   horarioAbertura,
		horarioFechamento: horarioFechamento,
		diasFuncionamento: diasFuncionamento,
		maquinas:          maquinas,
		usuarios:          usuarios,
		aberto:            aberto}
}

func (l *Lavanderia) SetEmpresaId(id int) {
	l.empresaId = id
}

func (l *Lavanderia) SetHorarioAbertura(horarioAbertura time.Time) {
	l.horarioAbertura = horarioAbertura
}

func (l *Lavanderia) SetHorarioFechamento(horarioFechamento time.Time) {
	l.horarioFechamento = horarioFechamento
}

func (l *Lavanderia) SetDiasFuncionamento(diasFuncionamento []string) {
	l.diasFuncionamento = diasFuncionamento
}

func (l *Lavanderia) SetMaquinas(maquinas []Maquina) {
	l.maquinas = maquinas
}

func (l *Lavanderia) SetUsuario(usuarios []Usuario) {
	l.usuarios = usuarios
}

func (l *Lavanderia) SetAberto(aberto bool) {
	l.aberto = aberto
}

func (l Lavanderia) GetEmpresaId() int {
	return l.empresaId
}

func (l Lavanderia) GetHorarioAbertura() time.Time {
	return l.horarioAbertura
}

func (l Lavanderia) GetHorarioFechamento() time.Time {
	return l.horarioFechamento
}

func (l Lavanderia) GetDiasFuncionamento() []string {
	return l.diasFuncionamento
}

func (l Lavanderia) GetMaquinas() []Maquina {
	return l.maquinas
}

func (l Lavanderia) GetUsuario() []Usuario {
	return l.usuarios
}

func (l Lavanderia) IsAberto() bool {
	return l.aberto
} */
