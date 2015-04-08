// Copyright 2015 elliott@tkwcafe.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package password

import (
	"crypto/rand"
	"io"
)

var AlphaNum = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var Specials = []byte(`~!@#$%^&*()-+=[]{}:;,.?\|/"'<>`)

const (
	CHAR_MIN = 12
	OUT_MIN  = 3
)

func New(length int, incSpecials bool) (string, error) {
	res := make([]byte, length)
	buf := make([]byte, length+(length/4))

	chars := make([]byte, len(AlphaNum))
	copy(chars, AlphaNum)
	if incSpecials {
		chars = append(chars, Specials...)
	}

	clen := len(chars)
	maxbuf := 512 - (512 % clen)

	i := 0
	for {
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			return "", err
		}

		for _, dat := range buf {
			c := int(dat)
			if c > maxbuf {
				continue //  skip to avoid modulo bias
			}

			res[i] = chars[c%clen]
			i++

			if i == length {
				return string(res), nil
			}
		}
	}
}
