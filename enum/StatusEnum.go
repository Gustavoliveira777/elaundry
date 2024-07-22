package enum

type StatusEnum int

const (
	AGENDADO StatusEnum = iota
	EM_ANDAMENTO
	FINALIZADO
	CANCELADO
)

func (s StatusEnum) String() string {
	return [...]string{"Agendado", "Em andamento", "Finalizado", "Cancelado"}[s]
}
