package transactions

import (
	"fmt"

	"github.com/andutruel/go_transacciones/pkg/store"
)

type Transaccion struct {
	Id                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion" binding:"required"`
	Moneda            string  `json:"moneda" binding:"required"`
	Monto             float64 `json:"monto" binding:"required"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor" binding:"required"`
	FechaTransaccion  string  `json:"fecha_transaccion" binding:"required"`
}

var ts []Transaccion
var lastId int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
	LastID() (int, error)
	Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Transaccion, error) {
	r.db.Read(&ts)
	return ts, nil
}

func (r *repository) LastID() (int, error) {
	if err := r.db.Write(ts); err != nil {
		return 0, err
	}

	return lastId, nil
}

func (r *repository) Store(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	r.db.Read(&ts)

	t := Transaccion{id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion}
	lastId = t.Id
	ts = append(ts, t)

	if err := r.db.Write(ts); err != nil {
		return Transaccion{}, err
	}

	return t, nil
}

func (r *repository) Update(id int, codigo_transaccion string, moneda string, monto float64, emisor string, receptor string, fecha_transaccion string) (Transaccion, error) {
	t := Transaccion{id, codigo_transaccion, moneda, monto, emisor, receptor, fecha_transaccion}
	updated := false
	for i := range ts {
		if ts[i].Id == id {
			t.Id = id
			ts[i] = t
			updated = true
		}
	}

	if !updated {
		return Transaccion{}, fmt.Errorf("la transacción %d no fue encontrada", id)
	}

	return t, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
