package datasource

import (
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/zeebo/errs"
)

// Check checks
func (s *Source) Check(fn CheckerFunc) (string, error) {
	// run checker func
	v, err := fn(string(s.URL))
	if err != nil {
		go s.tryToFix()
		return "", errs.New("service down, attempting to fix")
		// switch v {
		// case "unauthorized":
		// 	return "requires new credentials", errs.Wrap(err)
		// case "down":
		// 	// run down fixer
		// 	restart, err := s.Fixers["restart"].Fixer()
		// 	if err != nil {
		// 		return "unable to retsart service, pge worker required", errs.Wrap(err)
		// 	}

		// 	return restart, nil
		// default:
		// 	return "no fixer available, hard restart required", errs.Wrap(err)
		// }
	}

	return v, nil
}

func (s *Source) tryToFix() error {
	//
	return errs.New("not impl")
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
	// Get() (*Source, error)
	Check(fn CheckerFunc) (string, error)
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
type CheckerFunc func(url string) (string, error)

// FixFunc declares a type for fxing a service
type FixFunc func() (string, error)

// Encryptor defines a simple encryption interface
type Encryptor interface {
	Encrypt() //
	Decrypt() //
}

// Source implements Datasource interface
type Source struct {
	ID          uuid.UUID
	URL         []byte // sensitive information, should be encrypted
	LastContact time.Time
	Name        string // display name
	Status      Status // last known status
	CheckRate   time.Duration
	Checkers    map[string]CheckerFunc

	Fixers map[string]Fixer

	Observers []Observer
}

// Fixer holds data for fixes
type Fixer struct {
	Priority int // 0 is lowest, ascending in priority
	Fixer    FixFunc
}

// Observer emits events
type Observer interface {
	Emit(status Status, payload map[string]interface{})
}
