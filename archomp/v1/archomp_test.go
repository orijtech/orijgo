package archomp_test

import (
	"encoding/json"
	"testing"

	"github.com/orijtech/orijgo/archomp/v1"
)

func TestRequestValidation(t *testing.T) {
	tests := [...]struct {
		reqJSON string
		wantErr bool
	}{
		0: {
			reqJSON: `{}`,
			wantErr: true,
		},

		1: {
			reqJSON: `
			  {
			    "files": [
				{"url": "https://orijtech.com/favicon.ico"},
				{"url": "https://twitter.com/orijtech", "name": "twitterpage"}
			    ]
			  }`,
		},

		2: {
			wantErr: true, // no resources
			reqJSON: `
			  {
			    "files": [
				{},
				{}
			    ]
			  }`,
		},

		3: {
			wantErr: true, // no resources
			reqJSON: `
			  {
			    "files": [
				{"url": "  "},
				{"url": ""}
			    ]
			  }`,
		},

		4: {
			wantErr: true, // no resources
			reqJSON: `
			  {
			    "files": [
				{"url": "  ", "name": "foo"},
				{"url": "", "name": ""},
				{"url": "", "name": " "}
			    ]
			  }`,
		},

	}

	for i, tt := range tests {
		req := new(archomp.Request)
		if err := json.Unmarshal([]byte(tt.reqJSON), req); err != nil {
			t.Errorf("#%d: unmarshal err: %v; data: %q", i, err, tt.reqJSON)
			continue
		}

		err := req.Validate()
		gotErr := err != nil
		if tt.wantErr != gotErr {
			t.Errorf("#%d: gotErr: %v wantErr: %v req: %#v", i, gotErr, tt.wantErr, req)
		}
	}
}
