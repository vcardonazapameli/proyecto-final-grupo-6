package seller

import "errors"

var ExistingCIdError error = errors.New("A Seller with that CId already exists")
var ErrorNotFound error = errors.New("Not found entity with given ID")
