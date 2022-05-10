package builder

import (
	"fmt"
	"testing"
)

var selectStatement = fmt.Sprintf(`projects.*,
	(SELECT COUNT(*) FROM client.tasks WHERE projects.id = tasks.project_id) AS "task_stats.total",
	(SELECT COUNT(*) FROM client.tasks WHERE projects.id = tasks.project_id AND tasks.status = %d) AS "task_stats.pending",
	(SELECT COUNT(*) FROM client.tasks WHERE projects.id = tasks.project_id AND tasks.status = %d) AS "task_stats.processing",
	(SELECT COUNT(*) FROM client.tasks WHERE projects.id = tasks.project_id AND tasks.status = %d) AS "task_stats.completed"`,
	0,
	1,
	2,
)

func Test_Group(t *testing.T) {

	//q := New("postgres").
	//	SelectRaw(selectStatement).
	//	From("client.projects").
	//	LeftJoin("client.tasks", "tasks", "projects.id = tasks.project_id").
	//	GroupBy("projects.id").
	//	Where("projects.id", "=", 1).
	//	Limit(1)

	q := New("postgres").
		Update("test").
		Column("uuid", "?").
		Column("updated_at", "NOW()").
		Column("created_at", "NOW()")

	fmt.Println(q.Build())
}
