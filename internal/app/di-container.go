package app

import (
	"os"

	"github.com/nzb3/slogmanager"
)

type serviceProvider struct {
	slogManager slogmanager.Manager
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) Logger() *serviceProvider {
	manager := slogmanager.New()
	manager.AddWriter("stdout", slogmanager.NewWriter(os.Stdout, slogmanager.WithTextFormat()))
	return sp
}
