package model

type Usuario struct {
	UsuarioId    int `gorm:"primaryKey;autoIncrement"`
	NomeCompleto string
	Email        string `gorm:"uniqueIndex"`
	Senha        string
	Funcionario  bool
}

/* func UsuarioFactory(usuarioId int, nomeCompleto string, email string, senha string, funcionario bool) Usuario {
	return Usuario{usuarioId: usuarioId, nomeCompleto: nomeCompleto, email: email, senha: senha, funcionario: funcionario}
}

func (u *Usuario) SetUsuarioId(id int) {
	u.usuarioId = id
}
func (u *Usuario) SetNomeCompleto(nomeCompleto string) {
	u.nomeCompleto = nomeCompleto
}
func (u *Usuario) SetEmail(email string) {
	u.email = email
}
func (u *Usuario) SetSenha(senha string) {
	u.senha = senha
}
func (u *Usuario) SetFuncionario(funcionario bool) {
	u.funcionario = funcionario
}

func (u Usuario) GetUsuarioId() int {
	return u.usuarioId
}

func (u Usuario) GetNomeCompleto() string {
	return u.nomeCompleto
}

func (u Usuario) GetEmail() string {
	return u.email
}
func (u Usuario) GetSenha() string {
	return u.senha
}
func (u Usuario) IsFuncionario() bool {
	return u.funcionario
}
*/
