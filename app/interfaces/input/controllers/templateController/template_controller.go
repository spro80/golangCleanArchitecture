package templateController

import (
	"fmt"

	"github.com/spro80/golangCleanArchitecture/app/application/useCase/templateUseCase"
)

type TemplateControllerInterface interface {
	HandlerTemplateController() error
}

type TemplateControllerHandler struct {
	useCase templateUseCase.TemplateUseCaseInterface
}

func NewTemplateController(useCase templateUseCase.TemplateUseCaseInterface) *TemplateControllerHandler {
	return &TemplateControllerHandler{useCase: useCase}
}

func (t *TemplateControllerHandler) HandlerTemplateController() error {
	fmt.Println("[template_controller] Init in TemplateController")
	fmt.Println("[template_controller] Calling templateUseCase")

	err := t.useCase.HandlerTemplateUseCase()
	if err != nil {
		fmt.Printf("[template_controller] Error: [%s]", err.Error())
	}

	fmt.Println("[template_controller] templateUseCase was called successfully")
	fmt.Println("[template_controller] End in TemplateController")

	return nil
}
