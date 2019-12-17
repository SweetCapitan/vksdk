package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_UtilsCheckLink(t *testing.T) {
	needUserToken(t)

	res, err := vkUser.UtilsCheckLink(Params{
		"url": "http://google.ru",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Link)
	assert.NotEmpty(t, res.Status)
}

func TestVK_UtilsGetShortLink(t *testing.T) {
	needUserToken(t)

	shortLink, err := vkUser.UtilsGetShortLink(Params{
		"url":     "http://google.ru",
		"private": 1,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, shortLink)

	res, err := vkUser.UtilsGetLastShortenedLinks(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	del, err := vkUser.UtilsDeleteFromLastShortened(Params{
		"key": shortLink.Key,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, del)
}

func TestVK_UtilsGetLinkStats(t *testing.T) {
	needGroupToken(t)

	params := Params{
		"key":             "8TDuIz",
		"interval":        "month",
		"intervals_count": 12,
	}

	res, err := vkGroup.UtilsGetLinkStats(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, res.Key)
	assert.NotEmpty(t, res.Stats)

	resEx, err := vkGroup.UtilsGetLinkStatsExtended(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, resEx.Key)
	assert.NotEmpty(t, resEx.Stats)
}

func TestVK_UtilsGetServerTime(t *testing.T) {
	needGroupToken(t)

	res, err := vkGroup.UtilsGetServerTime(Params{})
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_UtilsResolveScreenName(t *testing.T) {
	needGroupToken(t)

	f := func(name string) {
		t.Helper()

		_, err := vkGroup.UtilsResolveScreenName(Params{
			"screen_name": name,
		})
		if err != nil {
			t.Errorf("VK.UtilsResolveScreenName() err = %v", err)
		}
	}

	f("durov")
	f("api")
	f("app6991405")
	f("z")
}
