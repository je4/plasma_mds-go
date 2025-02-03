package service

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/je4/plasma_mds-go/pkg/grodata"
	"github.com/je4/utils/v2/pkg/zLogger"
	"github.com/nao1215/markdown"
	"io"
	"net/http"
	"net/url"
)

func NewDocServer(addr string, extAddr string, logger zLogger.ZLogger) *DocServer {
	mux := gin.Default()
	srv := &DocServer{
		addr:    addr,
		extAddr: extAddr,
		logger:  logger,
		srv: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
	mux.GET("/doc", srv.doc)
	return srv
}

type DocServer struct {
	addr    string
	extAddr string
	logger  zLogger.ZLogger
	srv     *http.Server
}

func (s *DocServer) Run() error {
	s.logger.Info().Str("addr", s.extAddr+"/doc").Msg("starting server")
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				s.logger.Error().Err(err).Msg("error running server")
			}
		}
	}()
	return nil
}

func (s *DocServer) Stop(ctx context.Context) error {
	return errors.WithStack(s.srv.Shutdown(ctx))
}

type dataSetName struct {
	DOI       string `form:"doi"`
	PlasmaMDS string `form:"plasma_mds"`
}

func (s *DocServer) doc(c *gin.Context) {
	var d = dataSetName{}
	if err := c.ShouldBind(&d); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if d.DOI == "" {
		c.JSON(400, gin.H{
			"error": "doi must be set",
		})
		return
	}

	grodataurl := "https://data.goettingen-research-online.de/api/datasets/export?exporter=dataverse_json&persistentId=" + url.QueryEscape(d.DOI)
	resp, err := http.Get(grodataurl)
	if err != nil {
		c.JSON(500, gin.H{
			"url":   grodataurl,
			"error": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	dataBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"url":   grodataurl,
			"error": err.Error(),
		})
		return
	}
	var data grodata.Grodata
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		c.JSON(500, gin.H{
			"url":   grodataurl,
			"error": err.Error(),
		})
		return
	}
	c.Header("Content-Type", "text/plain; charset=UTF-8")
	md := markdown.NewMarkdown(c.Writer)
	md.H1(data.GetTitle())
	var authors = []string{}
	for _, a := range data.GetAuthors() {
		authors = append(authors, fmt.Sprintf("%s (%s) - %s %s ", a.AuthorName, a.AuthorAffiliation, a.AuthorIdentifierScheme, a.AuthorIdentifier))
	}
	md.BulletList(authors...)
	md.Blockquote(data.GetDescription())

	md.Build()
}
