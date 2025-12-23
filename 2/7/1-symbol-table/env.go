package symbol_table

type Env struct {
	table map[string]string
	prev  *Env
}

func NewEnv(prev *Env) *Env {
	return &Env{
		make(map[string]string),
		prev,
	}
}

func (e *Env) Put(k string, v string) {
	e.table[k] = v
}

func (e *Env) Get(k string) string {
	for current := e; current != nil; current = current.prev {
		if v, ok := current.table[k]; ok {
			return v
		}
	}

	return ""
}
