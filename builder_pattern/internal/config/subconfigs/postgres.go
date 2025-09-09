package subconfigs 




type PostgreSQL struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}




type PostgreSQLBuilder struct {
	db PostgreSQL
}

func NewPostgreSQLBuilder() *PostgreSQLBuilder {
	return &PostgreSQLBuilder{}
}

func (b *PostgreSQLBuilder) SetHost(host string) *PostgreSQLBuilder {
	b.db.Host = host
	return b
}


func (b *PostgreSQLBuilder) SetPort(port int) *PostgreSQLBuilder {
	b.db.Port = port
	return b
}


func (b *PostgreSQLBuilder) SetUser(user string) *PostgreSQLBuilder {
	b.db.User = user
	
	return b
}



func (b *PostgreSQLBuilder) SetPassword(password string) *PostgreSQLBuilder {
	b.db.Password = password
	return b
}

func (b *PostgreSQLBuilder) SetName(name string) *PostgreSQLBuilder {
	b.db.Name = name
	return b
}


func (b *PostgreSQLBuilder) Build() PostgreSQL {
	return b.db
}