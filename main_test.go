package main

import (
	"flag"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	steps "github.com/ONSdigital/dp-frontend-homepage-controller/features/steps"
)

var componentFlag = flag.Bool("component", false, "perform component tests")

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})
}

func TestComponent(t *testing.T) {
	if *componentFlag {
		status := 0

		var opts = godog.Options{
			Output: colors.Colored(os.Stdout),
			Format: "pretty",
			Paths:  []string{"features/census", "features/homepage"},
			Tags:   "~@avoid",
		}

		status = godog.TestSuite{
			Name:                 "feature_tests",
			ScenarioInitializer:  steps.InitializeScenario,
			TestSuiteInitializer: InitializeTestSuite,
			Options:              &opts,
		}.Run()

		if status > 0 {
			t.Fail()
		}
	} else {
		t.Skip("component flag required to run component tests")
	}
}
