package data

import (
	"context"

	"usm/internal/conf"
	"usm/internal/data/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/lib/pq"
)

type Data struct {
	db *ent.Client
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	drv, err := sql.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		log.NewHelper(logger).Info(i...)
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: client,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
