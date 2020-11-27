package qrcode

import (
	"io"

	"github.com/tuotoo/qrcode"
)

// QRCodeParse 二维码图片解析
func QRCodeParse(fi io.Reader) (string, error) {
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		return "", err
	}
	return qrmatrix.Content, nil
}
