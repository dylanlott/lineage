package datasource

import (
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// Check checks
func (s *Source) Check(fn CheckerFunc) (Result, error) {
	return Result{}, errs.New("not impl")
}

// Fix fixes
func (s *Source) Fix(fn FixFunc) (Result, error) {
	return Result{}, errs.New("not impl")
}

// Query queries
func (s *Source) Query(reader io.Reader) (Result, error) {
	return Result{}, errs.New("not impl")
}

// Register registers
func (s *Source) Register(src Datasource) (Result, error) {
	return Result{}, errs.New("not impl")
}

// Unregister unregisters
func (s *Source) Unregister(id uuid.UUID) error {
	return errs.New("not impl")
}

// Remove removes
func (s *Source) Remove(id uuid.UUID) error {
	return errs.New("not impl")
}

var _ Datasource = (&Source{})

// Datasource declares the interface
type Datasource interface {
	Check(fn CheckerFunc) (Result, error)
	Fix(fn FixFunc) (Result, error)
	Query(reader io.Reader) (Result, error)
	Register(src Datasource) (Result, error)
	Unregister(id uuid.UUID) error
	Remove(id uuid.UUID) error
}

// Status declares the interface
type Status interface {
	String() string
}

// Result returns a payload of bytes
type Result struct {
	Payload []byte
}

// CheckerFunc declares a function for checking a service
type CheckerFunc func() (Status, error)

// FixFunc declares a type for fxing a service
type FixFunc func() error

// Encryptor defines a simple encryption interface
type Encryptor interface {
	Encrypt() //
	Decrypt() //
}

// Source implements Datasource interface
type Source struct {
	ID          uuid.UUID
	URL         []byte // sensitive information
	LastContact time.Time
	Status      Status
	CheckRate   time.Duration
	Checkers    map[string]CheckerFunc
	Fixers      map[string]FixFunc
	Observers   []Observer
}

// Observer emits events
type Observer interface {
	Emit(status Status, payload map[string]interface{})
}
