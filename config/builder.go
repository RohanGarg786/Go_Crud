package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"errors"
)


type AppConfig struct {
	Port        string
	DatabaseDSN string
}

type AppConfigBuilder struct {
	port        string
	databaseDSN string
}

func NewAppConfigBuilder() *AppConfigBuilder {
	return &AppConfigBuilder{}
}

func (b *AppConfigBuilder) SetPort(port string) *AppConfigBuilder {
	b.port = port
	return b
}

func (b *AppConfigBuilder) SetDatabase(dsn string) *AppConfigBuilder {
	b.databaseDSN = dsn
	return b
}

func (b *AppConfigBuilder) Build() *AppConfig {
	return &AppConfig{
		Port:        b.port,
		DatabaseDSN: b.databaseDSN,
	}
}

// ConnectDB establishes a database connection
func (c *AppConfig) ConnectDB() error {
	// Implement database connection logic (e.g., PostgreSQL or SQLite)
	dbPool,err := pgxpool.New(context.Background(),c.DatabaseDSN)
	if err!=nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := dbPool.Ping(context.Background()); err != nil {
        return fmt.Errorf("failed to ping the database: %w", err)
    }
	fmt.Println("Successfully connected to the PostgreSQL database!")
	return nil
}


type User struct {
    Username string
    Password string
    Email    string
}

type UserBuilder struct {
    user User
}

func NewUserBuilder() *UserBuilder {
    return &UserBuilder{}
}

func (b *UserBuilder) SetUsername(username string) *UserBuilder {
    b.user.Username = username
    return b
}

func (b *UserBuilder) SetPassword(password string) *UserBuilder {
    b.user.Password = password
    return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
    b.user.Email = email
    return b
}

func (b *UserBuilder) Build() (User, error) {
    if b.user.Username == "" || b.user.Password == "" {
        return User{}, errors.New("username and password cannot be empty")
    }
    return b.user, nil
}
