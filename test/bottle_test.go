package test

import (
    "testing"
    "water-bottle-api/internal/bottle"
)

func TestBottleValidate(t *testing.T) {
    cases := []struct {
        name string
        b    bottle.Bottle
        wantErr bool
    }{
        {"valid", bottle.Bottle{Brand: "Morshynska", Volume: 500, OwnerID: 1}, false},
        {"empty brand", bottle.Bottle{Brand: "", Volume: 500, OwnerID: 1}, true},
        {"small volume", bottle.Bottle{Brand: "Morshynska", Volume: 0, OwnerID: 1}, true},
        {"large volume", bottle.Bottle{Brand: "Morshynska", Volume: 3000, OwnerID: 1}, true},
        {"zero owner", bottle.Bottle{Brand: "Morshynska", Volume: 500, OwnerID: 0}, true},
    }

    for _, c := range cases {
        err := bottle.Validate(c.b)
        if c.wantErr && err == nil {
            t.Errorf("%s: expected error, got none", c.name)
        }
        if !c.wantErr && err != nil {
            t.Errorf("%s: did not expect error, got %v", c.name, err)
        }
    }
}
