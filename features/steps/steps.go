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

func RegisterSteps(ctx *godog.ScenarioContext, component *feature.HomePageComponent, uiFeature *componenttest.UIFeature) {
	goCtx := context.Background()

	// Custom steps
	ctx.Step(`^the census about link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-about'] > ul > li:nth-child(1) > a"))

	ctx.Step(`^the 1st census data link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-data'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd census data link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-data'] > ul > li:nth-child(2) > a"))

	ctx.Step(`^the 1st census data link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-data'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd census data link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-data'] > ul > li:nth-child(2) > a"))

	ctx.Step(`^the 1st census releases link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-releases'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd census releases link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-releases'] > ul > li:nth-child(2) > a"))

	ctx.Step(`^the census topics link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-topics'] > ul > li:nth-child(1) > a"))

	ctx.Step(`^the census dictionary link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-dictionary'] > ul > li:nth-child(1) > a"))

	ctx.Step(`^the census historic link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-historic'] > ul > li:nth-child(1) > a"))

	ctx.Step(`^the census planning link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-planning'] > ul > li:nth-child(1) > a"))

	ctx.Step(`^the 1st census contact link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-contact'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd census contact link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-contact'] > ul > li:nth-child(2) > a"))

	ctx.Step(`^the 1st census other link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-other'] > ul > li:nth-child(1) > a"))
	ctx.Step(`^the 2nd census other link href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-other'] > ul > li:nth-child(2) > a"))

	ctx.Step(`^the census href value should be "([^"]*)"`, selectedLinkShouldHaveHREF(uiFeature, "[data-test='census-link'] > a"))

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
