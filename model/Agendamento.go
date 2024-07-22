package model

import (
	"elaundry/enum"
	"time"
)

type Agendamento struct {
	agendamentoId  int
	maquina        Maquina
	dataHoraInicio time.Time
	dataHoraFim    time.Time
	status         enum.StatusEnum
}

func AgendamentoFactory(agendamentoId int, maquina Maquina, dataHoraInicio time.Time, dataHoraFim time.Time, status enum.StatusEnum) Agendamento {
	return Agendamento{agendamentoId: agendamentoId,
		maquina:        maquina,
		dataHoraInicio: dataHoraInicio,
		dataHoraFim:    dataHoraFim,
		status:         status}
}

func (a *Agendamento) SetAgendamentoId(agendamentoId int) {
	a.agendamentoId = agendamentoId
}

func (a *Agendamento) SetMaquina(maquina Maquina) {
	a.maquina = maquina
}

func (a *Agendamento) SetDataHoraInicio(dataHoraInicio time.Time) {
	a.dataHoraInicio = dataHoraInicio
}

func (a *Agendamento) SetDataHoraFim(dataHoraFim time.Time) {
	a.dataHoraFim = dataHoraFim
}

func (a *Agendamento) SetStatus(status enum.StatusEnum) {
	a.status = status
}

func (a Agendamento) GetAgendamentoId() int {
	return a.agendamentoId
}

func (a Agendamento) GetMaquina() Maquina {
	return a.maquina
}

func (a Agendamento) GetDataHoraInicio() time.Time {
	return a.dataHoraInicio
}

func (a Agendamento) GetDataHoraFim() time.Time {
	return a.dataHoraFim
}

func (a Agendamento) GetStatus() enum.StatusEnum {
	return a.status
}
