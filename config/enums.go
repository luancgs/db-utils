package config

type DatabaseDriversEnum struct {
	Postgres string
	Sqlite   string
}

func (en *DatabaseDriversEnum) ToArray() []string {
	return []string{en.Postgres, en.Sqlite}
}

func GetDatabaseDriversEnum() *DatabaseDriversEnum {
	return &DatabaseDriversEnum{
		"postgres",
		"sqlite3",
	}
}

type ConnectionsEnum struct {
	ConnectionURL  string
	ConnectionForm string
}

func (en *ConnectionsEnum) ToArray() []string {
	return []string{en.ConnectionURL, en.ConnectionForm}
}

func GetConnectionsEnum() *ConnectionsEnum {
	return &ConnectionsEnum{
		"Connection URL",
		"Connection Form",
	}
}

type CommandsEnum struct {
	Clone string
	Dump string
	Query string
	Restore   string
}

func (en *CommandsEnum) ToArray() []string {
	return []string{en.Clone, en.Dump, en.Query, en.Restore}
}

func GetCommandsEnum() *CommandsEnum {
	return &CommandsEnum{
		"Clone",
		"Dump",
		"Query",
		"Restore",
	}
}
