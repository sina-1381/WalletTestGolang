package database

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Postgresql struct {
	Connection *pgxpool.Pool
	Config     PostgresqlConfig
}

type PostgresqlConfig struct {
	User           string
	Password       string
	Host           string
	Port           string
	Database       string
	SSLMode        string
	MigrationsPath string
}

func NewPostgresql(config PostgresqlConfig) (*Postgresql, error) {
	poolConfig, err := pgxpool.ParseConfig(config.Url())
	if err != nil {
		return nil, err
	}
	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute
	connection, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := connection.Ping(ctx); err != nil {
		connection.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}
	return &Postgresql{
		Connection: connection,
		Config:     config,
	}, nil
}

func (p *Postgresql) MigrationRestart() error {
	err := p.MigrationDown()
	if err != nil {
		return err
	}
	err = p.MigrationUp()
	if err != nil {
		return err
	}
	return nil
}

func (p *Postgresql) MigrationDown() error {
	migration, err := migrate.New(p.Config.MigrationsPath, p.Config.Url())
	if err != nil {
		return err
	}
	defer migration.Close()
	if err := migration.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (p *Postgresql) MigrationUp() error {
	migration, err := migrate.New(p.Config.MigrationsPath, p.Config.Url())
	if err != nil {
		return err
	}
	defer migration.Close()
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (p *PostgresqlConfig) Url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.Database, p.SSLMode)
}
