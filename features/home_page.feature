Feature: Groups
  Scenario: GET /census and checking the response status 200
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
            "[data-test='empl-rate-sub']": "Unemployment rate"
        }
    """
    # inflation section
    And the page should have the following content
    """
        {
            "[data-test='inflation-h2'] > span": "Inflation"
        }
    """
    # GDP section
    And the page should have the following content
    """
        {
            "[data-test='gdp-h2'] > span": "GDP"
        }
    """
    # UK population
    And the page should have the following content
    """
        {
            "[data-test='ukpop-h2'] > span": "UK population"
        }
    """
    # main figures section
    And the page should have the following content
    """
        {
            "[data-test='ukpop-h2'] > span": "UK population"
        }
    """
