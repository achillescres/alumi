package application

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/utils"
	"golang.org/x/sync/errgroup"
	"itamconnect/internal/controller/api"
)

type Application struct {
	api *api.API
}

func New(api *api.API) *Application {
	return &Application{api: api}
}

func (app *Application) Run(ctx context.Context) error {
	ew := utils.NewErrorWrapper("Application - Run")

	grp, grpctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return fmt.Errorf("api stopped: %w", app.api.Run(grpctx))
	})

	return ew(fmt.Errorf("errgroup stopped: %w", grp.Wait()))

}
