package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseCase_GenerateToken(t *testing.T) {
	type args struct {
		username string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "success",
			args: args{
				username: "test123",
			},
			want:    "test123",
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, gotError := GenerateToken(test.args.username)
			// not testing returned token because we not using depency injection for the generate function
			// assert.Equal(t, test.want, got)
			assert.Equal(t, test.wantErr, gotError)
		})
	}
}
