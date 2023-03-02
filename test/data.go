package test

import "ddd-to-do-list/internal/aggregate"

var (
	Activity   = aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")
	Activities = func() aggregate.Activities {
		return aggregate.Activities{aggregate.RebuildActivity(1, "bagus@bagus.com", "kerja bro")}
	}

	Todo, _ = aggregate.RebuildTodos(1, 1, "kerja bro", 1, "high")
	Todos   = func() aggregate.Todos {
		t, _ := aggregate.RebuildTodos(1, 1, "kerja bro", 1, "high")
		s := aggregate.Todos{t}
		return s
	}
)
