// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  end_to_end

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( TestServer ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* TestServer ) UnmarshalJSON([]byte) error { return nil }
func ( TestServer ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* TestServer ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_TestServer *TestServer
