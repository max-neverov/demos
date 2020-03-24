package repo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	user              = "jsonbnull"
	password          = "jsonbnull"
	dbname            = "jsonbnull"
	port     nat.Port = "5432/tcp"
)

/*
todo: why 2 containers are started?
todo: how "github.com/fortytw2/dockertest" handles waiting a container to start?
todo: how many dependencies testcontainers vs dockertest
*/

func Test_JsonB(t *testing.T) {
	db := startPostgres(t)

	q := `
select * 
from users;
`
	rows, err := db.Query(q)
	require.NoError(t, err)
	defer rows.Close()
	if rows.Next() {
		t.Fail()
	}
}

func startPostgres(t *testing.T) *sql.DB {
	var env = map[string]string{
		"POSTGRES_PASSWORD": password,
		"POSTGRES_USER":     user,
		"POSTGRES_DB":       dbname,
	}

	req := testcontainers.ContainerRequest{
		Image:        "postgres:12-alpine",
		ExposedPorts: []string{string(port)},
		Env:          env,
		WaitingFor:   wait.NewHostPortStrategy(port),
	}

	ctx := context.Background()

	pg, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	t.Cleanup(func() { pg.Terminate(ctx) })

	host, err := pg.Host(ctx)
	require.NoError(t, err)

	port, err := pg.MappedPort(ctx, port)
	require.NoError(t, err)

	dbURL := fmt.Sprintf(
		"user=%s password='%s' host=%s dbname=%s port=%s sslmode=disable",
		user, password, host, dbname, port.Port(),
	)
	db, err := sql.Open("postgres", dbURL)
	require.NoError(t, err)

	t.Cleanup(func() { db.Close() })

	err = db.Ping()
	require.NoError(t, err)

	err = initDB(db)
	require.NoError(t, err)

	return db
}

func initDB(db *sql.DB) error {
	q := `
create table users
(
    id        serial,
    name      text,
    age       int,
    user_json jsonb
);
`
	_, err := db.Exec(q)
	return err
}
