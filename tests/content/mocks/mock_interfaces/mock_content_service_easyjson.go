// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  mock_interfaces

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( MockContentBehavior ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* MockContentBehavior ) UnmarshalJSON([]byte) error { return nil }
func ( MockContentBehavior ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* MockContentBehavior ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_MockContentBehavior *MockContentBehavior

func ( MockContentBehaviorMockRecorder ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* MockContentBehaviorMockRecorder ) UnmarshalJSON([]byte) error { return nil }
func ( MockContentBehaviorMockRecorder ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* MockContentBehaviorMockRecorder ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_MockContentBehaviorMockRecorder *MockContentBehaviorMockRecorder