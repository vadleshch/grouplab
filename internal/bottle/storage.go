package bottle

import (
    "context"
    "errors"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
    db *pgxpool.Pool
}

func NewStorage(db *pgxpool.Pool) *Storage {
    return &Storage{db: db}
}

func (s *Storage) Create(ctx context.Context, b *Bottle) error {
    q := `INSERT INTO bottles (owner_id, brand, volume) VALUES ($1, $2, $3) RETURNING id`
    return s.db.QueryRow(ctx, q, b.OwnerID, b.Brand, b.Volume).Scan(&b.ID)
}

func (s *Storage) GetByID(ctx context.Context, id int) (*Bottle, error) {
    b := &Bottle{}
    q := `SELECT id, owner_id, brand, volume FROM bottles WHERE id=$1`
    err := s.db.QueryRow(ctx, q, id).Scan(&b.ID, &b.OwnerID, &b.Brand, &b.Volume)
    if err != nil {
        return nil, err
    }
    return b, nil
}

func (s *Storage) List(ctx context.Context) ([]*Bottle, error) {
    q := `SELECT id, owner_id, brand, volume FROM bottles`
    rows, err := s.db.Query(ctx, q)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var result []*Bottle
    for rows.Next() {
        b := &Bottle{}
        if err := rows.Scan(&b.ID, &b.OwnerID, &b.Brand, &b.Volume); err != nil {
            return nil, err
        }
        result = append(result, b)
    }
    return result, nil
}

func (s *Storage) Update(ctx context.Context, b *Bottle) error {
    q := `UPDATE bottles SET owner_id=$1, brand=$2, volume=$3 WHERE id=$4`
    cmd, err := s.db.Exec(ctx, q, b.OwnerID, b.Brand, b.Volume, b.ID)
    if err != nil {
        return err
    }
    if cmd.RowsAffected() == 0 {
        return errors.New("not found")
    }
    return nil
}

func (s *Storage) Delete(ctx context.Context, id int) error {
    q := `DELETE FROM bottles WHERE id=$1`
    cmd, err := s.db.Exec(ctx, q, id)
    if err != nil {
        return err
    }
    if cmd.RowsAffected() == 0 {
        return errors.New("not found")
    }
    return nil
}
