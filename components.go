package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenterror"
	"go.opentelemetry.io/collector/service/defaultcomponents"

	// exporters
	stackdriver "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stackdriverexporter"
)

func components() (component.Factories, error) {
	var errs []error
	factories, err := defaultcomponents.Components()
	if err != nil {
		return component.Factories{}, err
	}

	exporters := []component.ExporterFactory{
		stackdriver.NewFactory(),
	}
	for _, ep := range factories.Exporters {
		exporters = append(exporters, ep)
	}

	factories.Exporters, err = component.MakeExporterFactoryMap(exporters...)
	if err != nil {
		errs = append(errs, err)
	}

	return factories, componenterror.CombineErrors(errs)
}
