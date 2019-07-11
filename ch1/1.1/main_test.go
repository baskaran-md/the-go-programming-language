package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainFn(t *testing.T) {
	for _, test := range []struct {
		desc       string
		args       []string
		wantOutput string
	}{
		{
			desc:       "Test with valid args - 1",
			args:       []string{"echo", "foo", "bar"},
			wantOutput: "All Args: foo,bar\n",
		},
		{
			desc:       "Test with valid args - 2",
			args:       []string{"echo", "1", "2", "3"},
			wantOutput: "All Args: 1,2,3\n",
		},
		{
			desc:       "Test with empty args",
			args:       []string{"echo"},
			wantOutput: "All Args: \n",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			out = &bytes.Buffer{}
			os.Args = test.args
			main()
			output := out.(*bytes.Buffer).String()
			if output != test.wantOutput {
				t.Fatalf("%s: got output %s, want %s", test.desc, output, test.wantOutput)
			}
		})
	}
}
