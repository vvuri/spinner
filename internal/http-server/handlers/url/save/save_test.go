package save_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"spinner/internal/http-server/handlers/url/save"
	"spinner/internal/http-server/handlers/url/save/mocks"
	"testing"
)

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name      string
		alias     string
		url       string
		respError string
		mockError error
	}{
		{
			name:  "Success",
			alias: "test_alias",
			url:   "https://ya.ru",
		},
		{
			name:  "EmptyAlias",
			alias: "",
			url:   "https://ya.ru",
		},
		{
			name:      "Invalid URL",
			alias:     "",
			url:       "not a URL",
			respError: "field URL is not",
		},
		{
			name:      "SaveURL Error",
			alias:     "test_alias",
			url:       "https://ya.ru",
			respError: "failed to add url",
			mockError: errors.New("unexpected error"),
		},
	}

	// TODO: mock log
	log := slog.New(
		slog.NewTextHandler(os.Stdout, nil))

	for _, tc := range cases {
		urlSaverMock := mocks.NewURLSaver(t)

		if tc.respError == "" || tc.mockError != nil {
			urlSaverMock.On("SaverURL", tc.url, mock.AnythingOfType("string")).
				Return(int64(1), tc.mockError).
				Once()
		}

		handler := save.New(log, urlSaverMock, 0)

		input := fmt.Sprintf(`{"url":"#{tc.url}", "alias":"#{tc.alias}" }`)

		req, err := http.NewRequest(http.MethodPost, "/save", bytes.NewReader([]byte(input)))
		require.NoError(t, err) // assert + break

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, rr.Code, http.StatusOK)

		body := rr.Body.String()

		var resp save.Response
		require.NoError(t, json.Unmarshal([]byte(body), &resp))

		// TODO: Add check fields
	}
}
