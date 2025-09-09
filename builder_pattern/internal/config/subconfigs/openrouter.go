package subconfigs



type OpenRouter struct {
	Model   string
	Referer string
}


type OpenRouterBuilder struct {
	or OpenRouter
}

func NewOpenRouterBuilder() *OpenRouterBuilder {
	return &OpenRouterBuilder{}
}

func (b *OpenRouterBuilder) SetModel(model string) *OpenRouterBuilder {
	b.or.Model = model
	return b
}

func (b *OpenRouterBuilder) SetReferer(ref string) *OpenRouterBuilder {
	b.or.Referer = ref
	return b
}

func (b *OpenRouterBuilder) Build() OpenRouter {
	return b.or
}