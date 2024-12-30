// route/route.go
package route

import (
	"crypto-finance/src/presentation/controller"
	presentation_repository "crypto-finance/src/presentation/repository"
	"fmt"
)

func NewRoutes(
	router presentation_repository.HttpRepository,
	authController *controller.AuthController,
	statusController *controller.StatusController,
) error {
	authController.SetupRoutes(router)
	statusController.SetupRoutes(router)

	fmt.Println("Servidor iniciado em http://127.0.0.1:3000")
	return router.Start("127.0.0.1:3000")
}
