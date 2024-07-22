package model

type Pedido struct {
	pedidoId          int
	agendamentos      []Agendamento
	valorTotal        float32
	base64Comprovante string
	pagamentoAprovado bool
	usuarioRequerente Usuario
	usuarioAprovador  Usuario
}

func PedidoFactory(pedidoId int, agendamentos []Agendamento, valorTotal float32, base64Comprovante string, pagamentoAprovado bool, usuarioRequerente Usuario, usuarioAprovador Usuario) Pedido {
	return Pedido{
		pedidoId:          pedidoId,
		agendamentos:      agendamentos,
		valorTotal:        valorTotal,
		base64Comprovante: base64Comprovante,
		pagamentoAprovado: pagamentoAprovado,
		usuarioRequerente: usuarioRequerente,
		usuarioAprovador:  usuarioAprovador}
}

func (p *Pedido) SetPedidoId(id int) {
	p.pedidoId = id
}

func (p *Pedido) SetAgendamentos(agendamentos []Agendamento) {
	p.agendamentos = agendamentos
}

func (p *Pedido) SetValorTotal(valorTotal float32) {
	p.valorTotal = valorTotal
}

func (p *Pedido) SetBase64Comprovante(base64Comprovante string) {
	p.base64Comprovante = base64Comprovante
}

func (p *Pedido) SetPagamentoAprovado(pagamentoAprovado bool) {
	p.pagamentoAprovado = pagamentoAprovado
}

func (p *Pedido) SetUsuarioRequerente(usuarioRequerente Usuario) {
	p.usuarioRequerente = usuarioRequerente
}

func (p *Pedido) SetUsuarioAprovador(usuarioAprovador Usuario) {
	p.usuarioAprovador = usuarioAprovador
}

func (p Pedido) GetPedidoId() int {
	return p.pedidoId
}

func (p Pedido) GetAgendamentos() []Agendamento {
	return p.agendamentos
}

func (p Pedido) GetValorTotal() float32 {
	return p.valorTotal
}

func (p Pedido) GetBase64Comprovante() string {
	return p.base64Comprovante
}

func (p Pedido) IsPagamentoAprovado() bool {
	return p.pagamentoAprovado
}

func (p Pedido) GetUsuarioRequerente() Usuario {
	return p.usuarioRequerente
}

func (p Pedido) GetUsuarioAprovador() Usuario {
	return p.usuarioAprovador
}
