package message

type Text struct {
	Content string
}

func NewText(content string) *Text {
	return &Text{
		Content: content,
	}
}
