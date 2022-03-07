package useragent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cleanVersion(t *testing.T) {

	versions := []struct {
		text string
		want string
	}{
		{text: "v1.0",
			want: "1.0",
		},
		{text: "1_0",
			want: "1.0",
		},
		{text: "1_0_1__",
			want: "1.0.1",
		},
		{text: "__1_0_1",
			want: "1.0.1",
		},
		{text: "v 1_0_1",
			want: "1.0.1",
		},
		{text: "asdf",
			want: "",
		},
		{text: "_1.",
			want: "1",
		},
		{text: "",
			want: "",
		},
	}

	for _, version := range versions {
		assert.Equal(t, version.want, cleanVersion(version.text))

	}

}

func Test_getVersionString(t *testing.T) {

	versions := []struct {
		useragent string
		text      string
		want      string
	}{
		{
			useragent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
			text:      "windows nt ",
			want:      "10.0",
		},
		{
			useragent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0.3 Safari/605.1.15",
			text:      "mac os x ",
			want:      "10.15.6",
		},
		{
			useragent: "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
			text:      "windows nt ",
			want:      "6.1",
		},
		{
			useragent: "Mozilla/5.0 (Windows Mobile 10; Android 8.0.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36 Edge/40.15254.369",
			text:      "windows mobile ",
			want:      "10",
		},
		{ // test when there is no space after version
			useragent: "acer_S200 Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; Windows Phone 6.5)",
			text:      "windows phone ",
			want:      "6.5",
		},
		{
			useragent: "Mozilla/5.0 (Windows Mobile 10; Android 8.0.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36 Edge/40.15254.369",
			text:      "android ",
			want:      "8.0.0",
		},
		{
			useragent: "Mozilla/5.0 (Windows Mobile 10; Android 8.0.0; Microsoft; Lumia 950XL) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Mobile Safari/537.36 Edge/40.15254.369",
			text:      "it will be blank",
			want:      "",
		},
		{
			useragent: "Mozilla/5.0 (Windows Mobile )",
			text:      "windows mobile",
			want:      "",
		},
	}

	for _, version := range versions {
		ua := Parse(version.useragent)
		assert.Equal(t, version.want, ua.getVersionString(version.text))
	}

}
