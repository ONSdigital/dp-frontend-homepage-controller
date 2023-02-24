Feature: Groups
  Scenario: GET / and checking the response status 200
    When I navigate to "/"
    Then the improve this page banner should be visible
    # heading section
    And the page should have the following content
    """
        {
            "[data-test='header'] > h1": "Main figures",
            "[data-test='header'] > span > a": "From our time series explorer"
        }
    """
    # employment section
    And the page should have the following content
    """
        {
            "[data-test='empl-sub']": "Employment rate",
            "[data-test='empl-period']": "Aged 16 to 64 seasonally adjusted (Jun - Aug 2022)",
            "[data-test='empl-value']": "75.5%",
            "[data-test='empl-trend']": "-0.9% on previous year",
            "[data-test='unempl-sub']": "Unemployment rate",
            "[data-test='unempl-period']": "Aged 16+ seasonally adjusted (Jun - Aug 2022)",
            "[data-test='unempl-value']": "3.5%",
            "[data-test='unempl-trend']": "1.2pp on previous year"
        }
    """
    # inflation section
    And the page should have the following content
    """
        {
            "[data-test='inflation-h2'] > span": "Inflation",
            "[data-test='inflation-period']": "CPIH 12-month rate",
            "[data-test='inflation-date']": "Aug 2022",
            "[data-test='inflation-value']": "8.6%",
            "[data-test='inflation-trend']": "-0.2pp on previous month"
        }
    """
    # GDP section
    And the page should have the following content
    """
        {
            "[data-test='gdp-h2'] > span": "GDP",
            "[data-test='gdp-period']": "Quarter on Quarter",
            "[data-test='gdp-date']": "Apr - Jun 2022",
            "[data-test='gdp-value']": "0.2%",
            "[data-test='gdp-trend']": "-0.5pp on previous quarter"
        }
    """
    # UK population
    And the page should have the following content
    """
        {
            "[data-test='ukpop-h2'] > span": "UK population",
            "[data-test='ukpop-period']": "Mid-year estimate (2020)",
            "[data-test='ukpop-value']": "67,081,000"
        }
    """
    # Census
    And the page should have the following content
    """
        {
            "[data-test='census'] > div": "Results from Census 2021 are out now. Find data and analysis from Census 2021.\n",
            "[data-test='census-link'] > a": "Find out more about census"
        }
    """
    And the census href value should be "/census"
