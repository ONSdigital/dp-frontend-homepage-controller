package steps

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"

	componenttest "github.com/ONSdigital/dp-component-test"
	feature "github.com/ONSdigital/dp-frontend-homepage-controller/features"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	goCtx := context.Background()
	component, _ := feature.New(goCtx)
	url := fmt.Sprintf("http://%s%s", component.Config.SiteDomain, component.Config.BindAddr)
	uiFeature := componenttest.NewUIFeature(url)

	RegisterSteps(ctx, component, uiFeature)
}

func InitializeScenarioWithGetDataCardEnabled(ctx *godog.ScenarioContext) {
	goCtx := context.Background()
	component, _ := feature.New(goCtx)
	component.Config.EnableGetDataCard = true // Update configuration for enabling changes to "Get data card"
	url := fmt.Sprintf("http://%s%s", component.Config.SiteDomain, component.Config.BindAddr)
	uiFeature := componenttest.NewUIFeature(url)

	RegisterSteps(ctx, component, uiFeature)
}

func RegisterSteps(ctx *godog.ScenarioContext, component *feature.HomePageComponent, uiFeature *componenttest.UIFeature) {
	goCtx := context.Background()

	// Custom steps
	ctx.Step(`^the 1st link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-1'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-2'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 3rd link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-2'] > ul > li:nth-child(2) > a"))
	ctx.Step(`^the 4th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-3'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 5th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-3'] > ul > li:nth-child(2) > a"))
	ctx.Step(`^the 6th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-4'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 7th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-5'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 8th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-6'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 9th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-7'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 10th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-7'] > ul > li:nth-child(2) > a"))
	ctx.Step(`^the 11th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-8'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 12th link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='search-8'] > ul > li:nth-child(2) > a"))

	ctx.BeforeScenario(func(*godog.Scenario) {
		uiFeature.Reset()
	})

	ctx.AfterScenario(func(*godog.Scenario, error) {
		uiFeature.Close()
		component.StopService(goCtx)
	})

	uiFeature.RegisterSteps(ctx)
}

func selectedLinkShouldHaveHREF(f *componenttest.UIFeature, elementSelector string) func(string) error {
	return func(expectedContent string) error {
		var actualContent []map[string]string
		err := chromedp.Run(f.Chrome.Ctx,
			f.RunWithTimeOut(f.WaitTimeOut, chromedp.Tasks{
				chromedp.AttributesAll(elementSelector, &actualContent),
			}),
		)
		if err != nil {
			return err
		}
		assert.EqualValues(f, expectedContent, actualContent[0]["href"])
		return f.StepError()
	}
}