package shortener

import (
	"errors"
	"time"

	errs "github.com/pkg/errors"
	"github.com/teris-io/shortid"
	"gopkg.in/validator.v2"
)

var (
	ErrRedirectNotFound = errors.New("redirect not found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type redirectService struct {
	redirectRepo RedirectRepository
}

func NewRedirectService(redirectRepo RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepo,
	}
}

func (r *redirectService) Find(code string) (*Redirect, error) {
	return r.redirectRepo.Find(code)
}

func (r *redirectService) Store(redirect *Redirect) error {
	if err := validator.Validate(redirect); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Stores")
	}

	redirect.Code = shortid.MustGenerate()
	redirect.CreatedAt = time.Now().UTC().Unix()
	return r.redirectRepo.Store(redirect)
}
