package test

import (
    "testing"
    "water-bottle-api/internal/user"
)

func TestUserValidate(t *testing.T) {
    cases := []struct {
        name string
        u    user.User
        wantErr bool
    }{
        {"valid", user.User{Name: "Vasya", Age: 19}, false},
        {"empty name", user.User{Name: "", Age: 20}, true},
        {"young age", user.User{Name: "Vasya", Age: 17}, true},
    }

    for _, c := range cases {
        err := user.Validate(c.u)
        if c.wantErr && err == nil {
            t.Errorf("%s: expected error, got none", c.name)
        }
        if !c.wantErr && err != nil {
            t.Errorf("%s: did not expect error, got %v", c.name, err)
        }
    }
}
