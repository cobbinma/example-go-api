package postgres

type postgres struct {
	dbClient DBClient
}

func NewPostgres(client DBClient) *postgres {
	return &postgres{dbClient: client}
}
