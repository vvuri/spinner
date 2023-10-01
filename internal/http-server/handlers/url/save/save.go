package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	resp "spinner/internal/lib/api/response"
	"spinner/internal/lib/random"
	"spinner/internal/storage"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
}

func New(log *slog.Logger, urlSaver URLSaver, aliasLength int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const ap = "handlers.url.save.New"

		log = log.With(
			slog.String("ap", ap),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to parse request body", err)

			render.JSON(w, r, resp.Error("failed to parse request"))

			return
		}

		log.Info("request body parsed", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("failed to validate request", err)

			render.JSON(w, r, resp.Error("failed to validate request"))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = random.NewRandomString(aliasLength)
			log.Info("alias for", urlSaver, "is", alias)
		}

		err = urlSaver.SaveURL(req.URL, alias)
		if errors.Is(err, storage.ErrURLExists) {
			log.Info("url already exist", alias)

			render.JSON(w, r, resp.Error("url already exist"))

			return
		}
		if err != nil {
			log.Error("failed to add url", err)

			render.JSON(w, r, resp.Error("failed to add url"))

			return
		}

		log.Info("url added to DB")

		render.JSON(w, r, Response{
			Response: resp.Ok(),
			Alias:    alias,
		})
	}
}
