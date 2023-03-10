package scraping

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScrapingAmazonURL(t *testing.T) {
	t.Parallel()

	type args struct {
		amazonUrl string
		imgUrl    string
	}
	tests := []struct {
		name    string
		args    args
		want    ExtractedData
		wantErr bool
	}{
		{
			name: "successful case",
			args: args{
				amazonUrl: "https://onl.bz/Xr9qS57",
			},
			want: ExtractedData{
				Title: "カオスエンジニアリング ―回復力のあるシステムの実践",
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			res, err := AmazonURL(tc.args.amazonUrl)
			fmt.Println(res)

			if diff := cmp.Diff(tc.wantErr, err != nil); diff != "" {
				t.Error(err)
			}
			if diff := cmp.Diff(tc.want.Title, res.Title); diff != "" {
				t.Error(err)
			}
		})
	}
}
