package templateUseCase

import "fmt"

type TemplateUseCaseInterface interface {
	HandlerTemplateUseCase() error
}

type TemplateUseCaseHandler struct {
}

func NewTemplateUseCase() *TemplateUseCaseHandler {
	return &TemplateUseCaseHandler{}
}

func (t *TemplateUseCaseHandler) HandlerTemplateUseCase() error {
	fmt.Println("[template_use_case] Init in HandlerTemplateUseCase")

	fmt.Println("[template_use_case] End in HandlerTemplateUseCase")
	return nil
}
