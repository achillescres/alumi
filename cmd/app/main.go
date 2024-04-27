package app

import (
	"context"
	"fmt"
	"github.com/achillescres/pkg/security/ajwt"
	"github.com/achillescres/pkg/security/passlib"
	"github.com/sirupsen/logrus"
	"itamconnect/cmd/app/config"
	"itamconnect/ent"
	"itamconnect/internal/application"
	"itamconnect/internal/controller/api"
	"itamconnect/internal/service/auth_serv"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	logger := logrus.New()

	cfg := config.Config{}
	err := cfg.Read()
	if err != nil {
		logger.Fatalf("read config: %s", err)
	}

	var sslmode string
	if !cfg.Postgres.UseCA {
		sslmode = "disable"
		cfg.Postgres.CaAbsPath = ""
	} else {
		sslmode = "verify-full"
	}
	entClient, err := ent.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s sslrootcert=%s",
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.Username,
			cfg.Postgres.DbName,
			cfg.Postgres.Password,
			sslmode,
			cfg.Postgres.CaAbsPath,
		),
	)
	if err != nil {
		logger.Fatalf("connect ent to postgres: %s\n", err)
	}

	hasher := passlib.NewHashManager("salt1")
	jwtManager := ajwt.NewJWTManager(hasher, "superSecretKey", time.Minute*60*2, time.Minute*60*12)

	authService := auth_serv.New(entClient, hasher, jwtManager)

	app := application.New(api.New("localhost:8090", logger.WithField("controller", "api"), authService))
	err = app.Run(ctx)
	cancel()
	logger.Fatal(err)
}
