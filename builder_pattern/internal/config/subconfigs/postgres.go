package subconfigs 

// PostgreSQL holds database connection configuration.
type PostgreSQL struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}


// PostgreSQLBuilder provides a fluent way to build PostgreSQL configuration.
type PostgreSQLBuilder struct {
	db PostgreSQL
}

// NewPostgreSQLBuilder creates and returns a new instance of PostgreSQLBuilder
func NewPostgreSQLBuilder() *PostgreSQLBuilder {
	return &PostgreSQLBuilder{}
}


// then settts of particular fileds.
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


// Build - calls after setting needed fields, and finishes the PosgreSQL Configuration
func (b *PostgreSQLBuilder) Build() PostgreSQL {
	return b.db
}