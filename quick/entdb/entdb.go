package entdb

import (
	"app/gen/ent"
	"app/quick/db"
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

type EntDB struct {
	client *ent.Client
}

type Config struct {
	DB db.DB
}

func (e EntDB) Client() *ent.Client {
	return e.client
}

func New(config Config) EntDB {
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(config.DB.Dialect, config.DB.DB)

	// client := ent.NewClient(ent.Driver(drv), ent.Debug())
	opts := []ent.Option{
		ent.Driver(&CustomDriver{drv}),
	}

	debug := viper.GetBool("database.debug")
	if debug {
		opts = append(opts, ent.Debug())
	}

	client := ent.NewClient(opts...)

	client.Intercept(
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
				count, err := next.Query(ctx, query)
				return count, err
			})
		}),
	)

	// ent.InterceptFunc(func(next ent.Querier) ent.Querier {
	// 	return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
	// 		// Do something before the query execution.
	// 		value, err := next.Query(ctx, query)
	// 		// Do something after the query execution.
	// 		return value, err
	// 	})
	// })

	return EntDB{
		client: client,
	}
}

type CustomDriver struct {
	*entsql.Driver
}

func (d *CustomDriver) Query(ctx context.Context, query string, args, v any) error {
	err := d.Driver.Query(ctx, query, args, v)
	// fmt.Println("Custom ERROR", err)
	debug := viper.GetBool("database.debug")
	if debug {
		FileWithLineNum()
	}

	return err
}

// FileWithLineNum return the file name and line number of the current file
func FileWithLineNum() string {
	// the second caller usually from gorm internal, so set i start from 2
	for i := 1; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.Contains(file, "/vendor") && !strings.Contains(file, "/ent") && !strings.Contains(file, "generated.go")) {
			// fmt.Println(file, line, ok)
			fmt.Println(time.Now().Format(time.RFC3339), color.GreenString(file))
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}

	return ""
}
