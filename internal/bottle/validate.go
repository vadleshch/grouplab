package bottle

import "errors"

func Validate(b Bottle) error {
    if b.Brand == "" {
        return errors.New("brand is required")
    }
    if b.Volume <= 0 || b.Volume > 2000 {
        return errors.New("volume must be in 1-2000 ml")
    }
    if b.OwnerID <= 0 {
        return errors.New("owner_id is required and must be > 0")
    }
    return nil
}
