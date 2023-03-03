package cognito

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_SignUp(t *testing.T) {
	t.Parallel()

	if _, err := os.Stat("../../test/private/cognito"); err != nil {
		return
	}

	read := func(pth string) string {
		s, err := os.ReadFile(pth)
		if err != nil {
			t.Fatal(err)
		}
		return string(s)
	}

	type args struct {
		id       string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successful case",
			args: args{
				id:       read("../../test/private/cognito/1.id"),
				password: read("../../test/private/cognito/1.password"),
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	client := New(
		read("../../test/private/cognito/1.clientid"),
		read("../../test/private/cognito/1.userpoolid"),
	)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			err := client.SignUp(ctx, tc.args.id, tc.args.password)

			if diff := cmp.Diff(tc.wantErr, err != nil); diff != "" {
				t.Error(err)
			}
		})
	}
}

func TestClient_Login(t *testing.T) {
	t.Parallel()

	if _, err := os.Stat("../../test/private/cognito"); err != nil {
		return
	}

	read := func(pth string) string {
		s, err := os.ReadFile(pth)
		if err != nil {
			t.Fatal(err)
		}
		return string(s)
	}

	type args struct {
		id       string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "successful case",
			args: args{
				id:       read("../../test/private/cognito/1.id"),
				password: read("../../test/private/cognito/1.password"),
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	client := New(
		read("../../test/private/cognito/1.clientid"),
		read("../../test/private/cognito/1.userpoolid"),
	)

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := client.Login(ctx, tc.args.id, tc.args.password)

			if diff := cmp.Diff(tc.wantErr, err != nil); diff != "" {
				t.Error(err)
			}
		})
	}
}
