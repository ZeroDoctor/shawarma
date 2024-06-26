package sqlite

import (
	"time"

	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (s *SqliteDB) QueryUserByName(name string) (model.User, error) {
	var user model.User
	var err error

	query := `SELECT * FROM users WHERE "name" = ?;`
	rows, err := s.conn.Queryx(query, name)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	var userMap map[string]interface{}
	for rows.Next() {
		if err := rows.MapScan(userMap); err != nil {
			return user, err
		}
	}

	var ok bool
	user, ok = convertModel(userMap, user).(model.User)
	if !ok {
		return user, ErrModelConvert
	}

	return user, err
}

func (s *SqliteDB) SaveUser(user model.User) (model.User, error) {
	var err error

	session, err := uuid.NewV7()
	if err != nil {
		return user, err
	}
	user.Session = session.String()

	now := time.Now()
	user.CreatedAt = now
	user.ModifiedAt = now

	insert := `INSERT INTO users (
		"name", "session",
		created_at, modified_at
	) VALUES (
		:name, :session,
		:created_at, :modified_at
	) ON CONFLICT("name") DO UPDATE SET
		session     = excluded.session,
		modified_at = excluded.modified_at,
		created_at  = excluded.created_at
	;`

	_, err = s.conn.NamedExec(insert, convertNamedSqlite(user))
	return user, err
}
