package checker

import (
	"github.com/pkg/errors"
	"net/http"
)

type UrlChecker struct {
	Size        uint
	ResourceUrl string
}

func NewChecker(resourceUrl string) *UrlChecker {
	urlChecker := &UrlChecker{ResourceUrl: resourceUrl}
	return urlChecker
}

func (uc *UrlChecker) setSize(size uint) error {
	uc.Size = size
	return nil
}

func (uc *UrlChecker) setResourceUrl(resourceUrl string) {
	uc.ResourceUrl = resourceUrl
}

func isRedirectUrl(responseUrl string, originalUrl string) bool {
	return responseUrl != originalUrl && responseUrl != ""
}

//set size field
func (uc *UrlChecker) CheckResourceSupportRangeAccess() error {
	res, err := http.Head(uc.ResourceUrl)
	if err != nil {
		return errors.Wrap(err, "failed to head request: "+uc.ResourceUrl)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		err := errors.Errorf("not supported range access: %s", uc.ResourceUrl)
		return err
	}
	responseUrl := res.Request.URL.String()
	if isRedirectUrl(responseUrl, uc.ResourceUrl) {
		//リダイレクトされた別のurlであれば、resourceUrlを更新する
		uc.setResourceUrl(responseUrl)
	}
	if res.ContentLength <= 0 {
		err := errors.New("invalid content length")
		return err
	}
	uc.setSize(uint(res.ContentLength))
	return nil
}
