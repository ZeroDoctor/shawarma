package sqlite

import (
	"time"

	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (s *SqliteDB) InsertUser(user model.User) (model.User, error) {
	var err error

	session, err := uuid.NewV7()
	if err != nil {
		return user, err
	}
	user.Session = session.String()

	now := time.Now()
	user.CreatedAt = model.Time(now)
	user.ModifiedAt = model.Time(now)

	insert := `INSERT INTO users (
		"name", "session", github_token,
		github_user_id, created_at, modified_at
	) VALUES (
		:name, :session, :github_token,
		:github_user_id, :created_at, :modified_at
	) ON CONFLICT("name") DO UPDATE SET
		session        = excluded.session,
		github_token   = excluded.github_token,
		github_user_id = excluded.github_user_id,
		modified_at    = excluded.modified_at
	;`

	_, err = s.conn.NamedExec(insert, user)
	return user, err
}
