package oauth2

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/fosite"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandlerConsent(t *testing.T) {
	h := &Handler{
		L:             logrus.New(),
		ScopeStrategy: fosite.HierarchicScopeStrategy,
	}
	r := httprouter.New()
	h.SetRoutes(r)
	ts := httptest.NewServer(r)

	res, err := http.Get(ts.URL + DefaultConsentPath)
	assert.Nil(t, err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.NotEmpty(t, body)
}
