package model

import "time"

type Agendamento struct {
	agendamentoId  int
	maquina        Maquina
	dataHoraInicio time.Time
	dataHoraFim    time.Time
}
