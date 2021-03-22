package json

import (
	"encoding/json"

	"github.com/carlosh1989/hex-microservice/shortener"
	errs "github.com/pkg/errors"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Encoding")
	}
	return rawMsg, nil
}
