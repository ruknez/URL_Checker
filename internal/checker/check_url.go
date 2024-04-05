package checker

import (
	"context"

	"github.com/pkg/errors"
	"url-checker/internal/common"
	"url-checker/internal/repository/entity"
)

func (c *checker) checkAllUrls(ctx context.Context) {
	for _, urls := range c.urlRepo.GetAllUrls(ctx) {
		u := urls
		go func() {
			status, err := c.checkUrl(ctx, u)
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "c.checkUrl: "+u))
				return
			}

			err = c.saveStatus(ctx, u, status)
			if err != nil {
				c.logger.Error(ctx, errors.Wrap(err, "c.saveStatus: "+u))
				return
			}
		}()
	}

}

func (c *checker) checkUrl(ctx context.Context, url string) (entity.Status, error) {
	status, err := c.statuserService.GetUrlStatus(ctx, url)
	if err != nil {
		return entity.Status(status), errors.Wrap(err, "statuserService.GetUrlStatus")
	}

	return common.ConvertStatus(status), nil
}

func (c *checker) saveStatus(ctx context.Context, url string, status entity.Status) error {
	return c.urlRepo.UpdateStatus(ctx, url, status)
}
