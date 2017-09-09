package index

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-catalog"
	"github.com/whosonfirst/go-whosonfirst-catalog/record"
	"github.com/whosonfirst/go-whosonfirst-pgis/client"
)

type PostGISIndex struct {
	catalog.Index
	client *pgis.PgisClient
}

func (e *PostGISIndex) GetById(id int64) (catalog.Record, error) {

	row, err := e.client.GetById(id)

	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("postgis://%d", id)

	return record.NewDefaultRecord("postgis", "postgis", id, uri, row)
}

func NewPostGISIndex(host string, port int, user string, password string, dbname string, maxconns int) (catalog.Index, error) {

	client, err := pgis.NewPgisClient(host, port, user, password, dbname, maxconns)

	if err != nil {
		return nil, err
	}

	e := PostGISIndex{
		client: client,
	}

	return &e, nil
}
