package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to Cassandra cluster; default port :9042
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra connection successfuly.")
	defer session.Close()
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
