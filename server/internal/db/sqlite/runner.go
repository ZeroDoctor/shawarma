package sqlite

import (
	"github.com/google/uuid"
	"github.com/zerodoctor/shawarma/pkg/model"
)

func (s *SqliteDB) SaveRunner(runner model.Runner) (model.Runner, error) {
	var err error
	runner.UUID, err = uuid.NewV7()
	if err != nil {
		return runner, err
	}

	insert := `INSERT INTO runners (
		uuid, "type", hostname, 
		created_at, modified_at
	) VALUES (
		:uuid, :type, :hostname, 
		:created_at, :modified_at
	) ON CONFLICT (hostname) DO UPDATE SET 
		uuid        = excluded.uuid, 
		"type"      = excluded.type, 
		hostname    = excluded.hostname, 
		created_at  = excluded.created_at, 
		modified_at = excluded.modified_at
	;`

	_, err = s.conn.NamedExec(insert, runner)
	return runner, err
}
