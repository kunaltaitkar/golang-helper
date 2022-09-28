package cassandra

import (
	"sync"

	"github.com/gocql/gocql"
)

var _ Factory = Implementation{}

type Implementation struct{}

func (Implementation) GetService() Service {
	return &impl{}
}

type impl struct {
}

var once sync.Once

var session *gocql.Session

func (i impl) Connect(config Config) error {
	var err error
	once.Do(func() {
		session, err = connect(config)
	})
	return err
}

func connect(config Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(config.Cluster...)
	cluster.Keyspace = config.Keyspace
	return cluster.CreateSession()
}

func (i impl) Fetch(query string, values ...interface{}) []map[string]interface{} {
	iterator := session.Query(query).Iter()
	result := []map[string]interface{}{}
	m := map[string]interface{}{}

	for iterator.MapScan(m) {
		result = append(result, m)
		m = map[string]interface{}{}
	}
	return result
}

func (i impl) Insert(query string, values ...interface{}) error {
	return session.Query(query, values).Exec()
}
