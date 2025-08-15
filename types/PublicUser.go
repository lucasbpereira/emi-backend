package types

type PublicUser struct {
    ID       int    `db:"id" json:"id"`
    Name     string `db:"name" json:"name"`
    Role     string `db:"role" json:"role"`
    // outros campos que você quer exibir
}