package viewmodel

// NameVM ...
type NameVM struct {
	ID string `json:"id"`
	EN string `json:"en"`
}

// RedisStringValueVM ...
type RedisStringValueVM struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
