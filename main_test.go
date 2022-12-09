package main

import (
	"flag"
	"os"
	"sync"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"

	enableFlagSteps "github.com/ONSdigital/dp-frontend-homepage-controller/enableFlag_features/steps"
	legacySteps "github.com/ONSdigital/dp-frontend-homepage-controller/features/steps"
)

var componentFlag = flag.Bool("component", false, "perform component tests")

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})
}

func TestComponent(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if *componentFlag {
			status := 0

			var opts = godog.Options{
				Output: colors.Colored(os.Stdout),
				Format: "pretty",
				Paths:  flag.Args(),
				Tags:   "~@avoid",
			}

			status = godog.TestSuite{
				Name:                 "feature_tests",
				ScenarioInitializer:  legacySteps.InitializeScenario,
				TestSuiteInitializer: InitializeTestSuite,
				Options:              &opts,
			}.Run()

			if status > 0 {
				t.Fail()
			}
		} else {
			t.Skip("component flag required to run component tests")
		}
		wg.Done()
	}()
	go func() {
		if *componentFlag {
			status := 0

			var opts = godog.Options{
				Output: colors.Colored(os.Stdout),
				Format: "pretty",
				Paths:  flag.Args(),
				Tags:   "~@avoid",
			}

			status = godog.TestSuite{
				Name:                 "feature_tests",
				ScenarioInitializer:  enableFlagSteps.InitializeScenarioEnableFlag,
				TestSuiteInitializer: InitializeTestSuite,
				Options:              &opts,
			}.Run()

			if status > 0 {
				t.Fail()
			}
		} else {
			t.Skip("component flag required to run component tests")
		}
		wg.Done()
	}()
	wg.Wait()
}
