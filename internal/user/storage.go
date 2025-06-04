package user

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

func (s *Storage) Create(ctx context.Context, u *User) error {
    q := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
    return s.db.QueryRow(ctx, q, u.Name, u.Age).Scan(&u.ID)
}

func (s *Storage) GetByID(ctx context.Context, id int) (*User, error) {
    u := &User{}
    q := `SELECT id, name, age FROM users WHERE id=$1`
    err := s.db.QueryRow(ctx, q, id).Scan(&u.ID, &u.Name, &u.Age)
    if err != nil {
        return nil, err
    }
    return u, nil
}

func (s *Storage) List(ctx context.Context) ([]*User, error) {
    q := `SELECT id, name, age FROM users`
    rows, err := s.db.Query(ctx, q)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    var users []*User
    for rows.Next() {
        u := &User{}
        if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}

func (s *Storage) Update(ctx context.Context, u *User) error {
    q := `UPDATE users SET name=$1, age=$2 WHERE id=$3`
    cmd, err := s.db.Exec(ctx, q, u.Name, u.Age, u.ID)
    if err != nil {
        return err
    }
    if cmd.RowsAffected() == 0 {
        return errors.New("not found")
    }
    return nil
}

func (s *Storage) Delete(ctx context.Context, id int) error {
    q := `DELETE FROM users WHERE id=$1`
    cmd, err := s.db.Exec(ctx, q, id)
    if err != nil {
        return err
    }
    if cmd.RowsAffected() == 0 {
        return errors.New("not found")
    }
    return nil
}
