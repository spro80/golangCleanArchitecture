package infraestructure

import (
	"fmt"
	"os"
)

type InfrastructureInterface interface {
	GraceFullShutdown(sig os.Signal)
	SetupDependencies() error
}

type Infrastructure struct {
}

func NewInfrastructure() InfrastructureInterface {
	return &Infrastructure{}
}

func (i *Infrastructure) SetupDependencies() error {

	// setup modules

	// setup use cases
	/*
		useCase := registerUserUseCase.NewRegisterUserUseCase()

		// setup controllers
		controller := registerUserController.NewRegisterUserController(useCase)
	*/
	//web.InitRoutes(inputCheckoutOrderApi, inputPickedOrderApi)
	return nil
}

func (i *Infrastructure) GraceFullShutdown(sig os.Signal) {
	fmt.Printf("[App]: Init in gracefulShutdown")

	fmt.Printf("[App]: Shutdown process completed for signal: %v", sig)
}
