package symbol_table

type env struct {
	table map[string]string
	prev  *env
}

func newEnv(prev *env) *env {
	return &env{
		make(map[string]string),
		prev,
	}
}

func (e *env) put(k string, v string) {
	e.table[k] = v
}

func (e *env) get(k string) string {
	for current := e; current != nil; current = current.prev {
		if v, ok := current.table[k]; ok {
			return v
		}
	}

	return ""
}
