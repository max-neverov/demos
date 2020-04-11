package repo

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/max-neverov/demos/jsonbnull/model"
)

/*
   How many dependencies testcontainers vs dockertest?
   testcontainers: 74 dependencies
   dockertest: itself + github.com/lib/pq
   makefile: no dependencies
*/

type UserResource struct {
	Name     string `db:"name"`
	Age      int    `db:"age"`
	SomeInfo []byte `db:"some_info"`
}

func Test_JsonB(t *testing.T) {
	db := testDB(t)

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
}

func testDB(t *testing.T) *sqlx.DB {
	db, err := sql.Open("postgres", "user=jsonbnull password='jsonbnull' host=localhost dbname=jsonbnull port=5432 sslmode=disable")
	require.NoError(t, err)
	require.NoError(t, db.Ping())

	dbx := sqlx.NewDb(db, "postgres")

	q := `
create table if not exists users
(
    id        serial,
    name      text,
    age       int,
    some_info jsonb
);
`
	_, err = dbx.Exec(q)
	require.NoError(t, err)

	// no cleanup: need to show result in the DB
	//t.Cleanup(func() {
	//	q := `truncate users`
	//	_, err := dbx.Exec(q)
	//	if err != nil {
	//		t.Logf("fail to cleanup db: %+v", err)
	//	}
	//})

	return dbx
}
