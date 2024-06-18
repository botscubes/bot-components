package components

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
)

type HTTPComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Url    *string `json:"url"`
		Body   *string `json:"body"`
		Method *string `json:"method"`
		Header *string `json:"header"`
	} `json:"data"`
}

func (c *HTTPComponent) GetOutputs() Outputs {
	return &c.Outputs
}
func (c *HTTPComponent) Execute(ctx *context.Context) (*any, error) {
	if c.Data.Method == nil {
		return nil, errors.New("Method not specified")
	}
	if c.Data.Url == nil {
		return nil, errors.New("URL not specified")
	}
	url, err := format.Format(*c.Data.Url, ctx)
	if err != nil {
		return nil, err
	}
	var body string = ""
	if c.Data.Body != nil {
		var err error
		body, err = format.Format(*c.Data.Body, ctx)
		if err != nil {
			return nil, err
		}
	}

	result := map[string]any{}
	client := &http.Client{}
	req, err := http.NewRequest(
		*c.Data.Method,
		url,
		bytes.NewReader([]byte(body)),
	)
	if err != nil {
		return nil, err
	}
	var m map[string]string
	if c.Data.Header != nil && *c.Data.Header != "" {
		err = json.Unmarshal([]byte(*c.Data.Header), &m)
		if err != nil {
			return nil, err
		}
		for k, v := range m {
			req.Header.Add(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result["body"] = bs
	result["statusCode"] = resp.StatusCode
	var a any = result
	return &a, nil
}
