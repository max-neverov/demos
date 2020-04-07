package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/max-neverov/demos/jsonbnull/model"
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

type UserResource struct {
	Name     string `db:"name"`
	Age      int    `db:"age"`
	SomeInfo []byte `db:"some_info"`
}

func Test_JsonB(t *testing.T) {
	db := startPostgres(t)

	expected := model.User{Name: "Name", Age: 42, SomeInfo: &model.SomeInfo{Whatever: "useful info"}}
	ur := UserResource{}

	// need returning part here, `no rows returned` otherwise;
	// cannot use `returning *` because `missing destination name id in *model.User`
	q := "insert into users(name,age,some_info) values($1, $2, $3) returning name, age, some_info"

	infoBytes, err := json.Marshal(expected.SomeInfo)
	require.NoError(t, err)

	err = db.Get(&ur, q, expected.Name, expected.Age, infoBytes)
	assert.NoError(t, err)

	info := model.SomeInfo{}
	err = json.Unmarshal(ur.SomeInfo, &info)
	assert.NoError(t, err)

	actual := model.User{Name: ur.Name, Age: ur.Age, SomeInfo: &info}

	assert.Equal(t, expected, actual)
	t.Logf("%#v", ur)
}

func startPostgres(t *testing.T) *sqlx.DB {
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

	db, err := sqlx.Connect("postgres", dbURL)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })
	db.Mapper = reflectx.NewMapper("db")

	err = initDB(db)
	require.NoError(t, err)

	return db
}

func initDB(db *sqlx.DB) error {
	q := `
create table users
(
    id        serial,
    name      text,
    age       int,
    some_info jsonb
);
`
	_, err := db.Exec(q)
	return err
}
