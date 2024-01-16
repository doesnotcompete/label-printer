package main

import (
	"errors"
	"github.com/datumbrain/label-printer/tag"
)

type Request struct {
	Text   string `json:"text"`
	CodeText string `json:"code_text"`
	CodeType tag.CodeType `json:"code_type"`
}

func (r Request) Validate() error {
	if r.Text == "" {
		return errors.New("`text` must be specified")
	} else if r.CodeText == "" {
		return errors.New("`code_text` must be specified")
	}

	return nil
}

type Response struct {
	Status string `json:"status"`
}
