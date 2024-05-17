package builder

type ExampleBuilder struct {
	baseBuilder
}

func NewExampleBuilder() *ExampleBuilder {
	return &ExampleBuilder{baseBuilder: newBaseBuilder()}
}

func (r *ExampleBuilder) WithID(id int) *ExampleBuilder {
	r.withEq("id", id)
	return r
}

func (r *ExampleBuilder) WithLtID(id int) *ExampleBuilder {
	r.withLt("id", id)
	return r
}

func (r *ExampleBuilder) OrderByIDDesc() *ExampleBuilder {
	r.withOrder("id desc")
	return r
}

func (r *ExampleBuilder) WithCursor(cursor int) *ExampleBuilder {
	if cursor > 0 {
		r.withLt("id", cursor)
	}
	r.withOrder("id desc")
	return r
}
