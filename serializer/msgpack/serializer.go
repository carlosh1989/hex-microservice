package msgpack

import (
	"github.com/carlosh1989/hex-microservice/shortener"
	errs "github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errs.Wrap(err, "serializer.Redirect.Encoding")
	}
	return rawMsg, nil
}
