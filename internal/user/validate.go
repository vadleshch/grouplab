package user

import "errors"

func Validate(u User) error {
    if u.Name == "" {
        return errors.New("name is required")
    }
    if u.Age < 18 {
        return errors.New("age must be 18 or older")
    }
    return nil
}
