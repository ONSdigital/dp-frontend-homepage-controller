module github.com/ONSdigital/dp-frontend-homepage-controller

go 1.13

replace github.com/ONSdigital/dp-api-clients-go => /Users/jon/dev/dp-api-clients-go

replace github.com/ONSdigital/dp-frontend-models => /Users/jon/dev/dp-frontend-models

require (
	github.com/ONSdigital/dp-api-clients-go v1.9.0
	github.com/ONSdigital/dp-frontend-models v1.5.0
	github.com/ONSdigital/dp-healthcheck v1.0.3
	github.com/ONSdigital/dp-rchttp v1.0.0
	github.com/ONSdigital/go-ns v0.0.0-20200205115900-a11716f93bad
	github.com/ONSdigital/log.go v1.0.0
	github.com/gorilla/mux v1.7.4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/pkg/errors v0.9.1
	github.com/smartystreets/goconvey v1.6.4
	golang.org/x/net v0.0.0-20200425230154-ff2c4b7c35a0 // indirect
	golang.org/x/sys v0.0.0-20200427175716-29b57079015a // indirect
)
