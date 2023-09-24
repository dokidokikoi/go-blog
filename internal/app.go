package application

import (
	"crypto/tls"
	"go-blog/internal/db/store/data"
	"go-blog/internal/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func initTask() {
	data.SetStoreDBFactory()
}

func Run() {
	initTask()
	e := gin.Default()
	e.Use(gin.Recovery())
	router.InstallAll(e)

	server := &http.Server{
		Addr:    ":18080",
		Handler: e,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{loadTLSCertificate()},
		},
	}

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}

func loadTLSCertificate() tls.Certificate {
	cert, err := tls.LoadX509KeyPair("./internal/conf/harukaze.top.crt", "./internal/conf/harukaze.top.key")
	if err != nil {
		panic(err)
	}
	return cert
}
