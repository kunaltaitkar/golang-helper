package cassandra

type Service interface {
	Connect(config Config) error
	Insert(query string, values ...interface{}) error
	Fetch(query string, values ...interface{}) []map[string]interface{}
}

type Factory interface {
	GetService() Service
}

type Config struct {
	Cluster  []string
	Keyspace string
}
