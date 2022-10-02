package avatar

import (
	"fmt"
	"os"
	"path"
	"time"

	"apihut-server/config"
	"apihut-server/logic/consts"
	"apihut-server/models"

	"github.com/nullrocks/identicon"
)

func NewAvatar(req *models.AvatarReq) (string, error) {
	gen, err := identicon.New(req.GetNamespace(), req.GetBlock(), req.GetDensity())
	if err != nil {
		return "", err
	}
	ident, err := gen.Draw(req.GetHash())
	if err != nil {
		return "", err
	}

	filePath := getFilePath(req)

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if err != nil {
		return "", err
	}

	switch req.GetOutput() {
	case consts.JPG, consts.JPEG:
		err = ident.Jpeg(req.GetSize(), req.GetQuality(), f)
	case consts.SVG:
		err = ident.Svg(req.GetSize(), f)
	default:
		err = ident.Png(req.GetSize(), f)
	}
	if err != nil {
		return "", err
	}

	return filePath, nil
}

func getFilePath(req *models.AvatarReq) string {
	ext := req.GetOutput().String()
	switch req.GetOutput() {
	case consts.JPG, consts.JPEG:
		ext = consts.JPEG.String()
	case consts.SVG:
		ext = consts.SVG.String()
	default:
		ext = consts.PNG.String()
	}

	return path.Join(config.Share.File.Avatar, fmt.Sprintf(
		"%s-%d-%d.%s",
		req.GetHash(),
		req.GetSize(),
		time.Now().UnixNano(),
		ext,
	))
}
