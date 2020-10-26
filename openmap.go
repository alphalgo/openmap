package openmap

const (
	ErrNotFound = Error("value cannot found")
	ErrExist    = Error("value already existed")
	ErrNotExist = Error("value not exist")
)

// Error defines an alias type
type Error string

// Error shows the error message
func (e Error) Error() string {
	return string(e)
}

// openmap struct
type Omap struct {
	m  map[interface{}]interface{}
	id interface{}
	v  interface{}
}

// Openmap defines the Omap interface
type Openmap interface {
	ilen(v ...interface{})
	Search(v ...interface{}) error
	MultiSearch(v ...interface{}) error
	Add(v ...interface{}) error
	MultiAdd(v ...interface{}) error
	Update(v ...interface{}) error
	MultiUpdate(v ...interface{}) error
	Delete(v ...interface{}) error
	MultiDelete(v ...interface{}) error
}

func (o *Omap) ilen(v ...interface{}) {
	switch len(v) {
	case 0:
		o.id = nil
		o.v = nil
	case 1:
		o.id = v[0]
		o.v = nil
	default:
		o.id = v[0]
		o.v = v[1]
	}
}

// Search finds the value of given id
func (o *Omap) Search(v ...interface{}) error {
	o.ilen(v)
	_, ok := o.m[o.v]
	if !ok {
		return ErrNotFound
	}
	return nil
}

// MultiSearch finds the value concurrently
func (o *Omap) MultiSearch(v ...interface{}) error {
	return nil
}

// Add adds the value into map
func (o *Omap) Add(v ...interface{}) error {
	err := o.Search(v)
	switch err {
	case ErrNotFound:
		o.m[o.id] = o.v
	case nil:
		return ErrExist
	default:
		return err
	}
	return nil
}

// MultiAdd adds the value concurrently
func (o *Omap) MultiAdd(v ...interface{}) error {
	return nil
}

// Update updates the value into map
func (o *Omap) Update(v ...interface{}) error {
	err := o.Search(v)
	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		o.m[o.id] = o.v
	default:
		return err
	}
	return nil
}

// MultiUpdate updates the value concurrently
func (o *Omap) MultiUpdate(v ...interface{}) error {
	return nil
}

// Delete deletes the value into map
func (o *Omap) Delete(v ...interface{}) error {
	err := o.Search(v)
	switch err {
	case ErrNotFound:
		return ErrNotExist
	case nil:
		delete(o.m, o.id)
	default:
		return err
	}
	return nil
}

// MultiDelete deletes the value concurrently
func (o *Omap) MultiDelete(v ...interface{}) error {
	return nil
}
