package subconfigs

// OpenRouter holds configuration for the OpenRouter API
type OpenRouter struct {
	Model   string
	Referer string
}

// OpenRouterBuilder provides a comfortable way to build an OpenRouter configuration
type OpenRouterBuilder struct {
	or OpenRouter
}

// NewOpenRouterBuilder returns a new OpenRouterBuilder
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

// Build returns the configured OpenRouter instance
func (b *OpenRouterBuilder) Build() OpenRouter {
	return b.or
}
