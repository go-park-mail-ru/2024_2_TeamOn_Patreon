// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  api

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( Route ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Route ) UnmarshalJSON([]byte) error { return nil }
func ( Route ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Route ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Route *Route