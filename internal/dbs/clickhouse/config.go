package clickhouse

import (
	"context"
	"fmt"

	click "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/pkg/errors"
)

type Config struct {
	Addr     string
	Port     uint16
	User     string
	Password string
	DB       string
}

func New(cfg Config) (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = click.Open(&click.Options{
			Addr: []string{fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)},
			Auth: click.Auth{
				Database: cfg.DB,
				Username: cfg.User,
				Password: cfg.Password,
			},
			ClientInfo: click.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: "an-example-go-client", Version: "0.1"},
				},
			},

			Debugf: func(format string, v ...interface{}) {
				fmt.Printf(format, v)
			},
		})
	)

	if err != nil {
		return nil, errors.Wrap(err, "clickhouse-connect")
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*click.Exception); ok {
			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return nil, errors.Wrap(err, "clickhouse-ping failed")
	}
	return conn, nil
}
