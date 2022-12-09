Feature: Groups
  Scenario: GET /census and checking the response status 200
    When I navigate to "/census"
    Then the improve this page banner should be visible
    # hero content
    Then element "[data-test='hero-h1']" should be visible
    And element "[data-test='hero-p']" should be visible
    And the page should have the following content
    """
        {
            "[data-test='hero-h1']": "Census",
            "[data-test='hero-p']": "The census takes place every 10 years. It gives us a picture of all the people and households in England and Wales."
        }
    """
    # search container section #1
    And the page should have the following content
    """
        {
            "[data-test='search-1'] > h2": "About the census",
            "[data-test='search-1'] > p": "Find out what the census is and why it's important for all of us.",
            "[data-test='search-1'] > ul > li:nth-child(1) > a": "About the census"
        }
    """
    And the 1st link href value should be "census/aboutcensus"
    # search container section #2
    And the page should have the following content
    """
        {
            "[data-test='search-2'] > h2": "Census 2021 data",
            "[data-test='search-2'] > p": "Find data for Census 2021.",
            "[data-test='search-2'] > ul > li:nth-child(1) > a": "Get census data (England and Wales)",
            "[data-test='search-2'] > ul > li:nth-child(2) > a": "Get census data (Wales)"
        }
    """
    And the 2nd link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimatesenglandandwalescensus2021"
    And the 3rd link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimateswalescensus2021"
    # search container section #3
    And the page should have the following content
    """
        {
            "[data-test='search-3'] > h2": "Census releases",
            "[data-test='search-3'] > p": "See what we've published, and our plans for the future.",
            "[data-test='search-3'] > ul > li:nth-child(1) > a": "Release calendar",
            "[data-test='search-3'] > ul > li:nth-child(2) > a": "First results and timeline"
        }
    """
    And the 4th link href value should be "/releasecalendar?query=census&fromDateDay=&fromDateMonth=&fromDateYear=&toDateDay=&toDateMonth=&toDateYear=&view=upcoming"
    And the 5th link href value should be "/census/censustransformationprogramme/census2021outputs/releaseplans"
    # search container section #4
    And the page should have the following content
    """
        {
            "[data-test='search-4'] > h2": "Census topics",
            "[data-test='search-4'] > p": "Find census data and analysis using these topics.",
            "[data-test='search-4'] > ul > li:nth-child(1) > a": "Topic summaries"
        }
    """
    And the 6th link href value should be "/census/aboutcensus/censusproducts/topicsummaries"
    # search container section #5
    And the page should have the following content
    """
        {
            "[data-test='search-5'] > h2": "Historic census data",
            "[data-test='search-5'] > p": "Find census data and analysis for 2011 and earlier.",
            "[data-test='search-5'] > ul > li:nth-child(1) > a": "Get all historic census data"
        }
    """
    And the 7th link href value should be "/census/historiccensusdata"
        # search container section #6
    And the page should have the following content
    """
        {
            "[data-test='search-6'] > h2": "Planning for Census 2021",
            "[data-test='search-6'] > p": "How we researched, prepared and planned for Census 2021.",
            "[data-test='search-6'] > ul > li:nth-child(1) > a": "Planning for Census 2021"
        }
    """
    And the 8th link href value should be "census/planningforcensus2021"
        # search container section #7
    And the page should have the following content
    """
        {
            "[data-test='search-7'] > h2": "Contact us",
            "[data-test='search-7'] > p": "If you need help, contact census customer services.",
            "[data-test='search-7'] > ul > li:nth-child(1) > a": "Census customer services",
            "[data-test='search-7'] > ul > li:nth-child(2) > a": "Request a 2011 census dataset"
        }
    """
    And the 9th link href value should be "/census/censuscustomerservices"
    And the 10th link href value should be "/census/2011census/2011censusdata/2011censusadhoctables"
        # search container section #8
    And the page should have the following content
    """
        {
            "[data-test='search-8'] > h2": "Scotland and Northern Ireland censuses",
            "[data-test='search-8'] > p": "We are responsible for the census in England and Wales.",
            "[data-test='search-8'] > ul > li:nth-child(1) > a": "Scotland census ",
            "[data-test='search-8'] > ul > li:nth-child(2) > a": "Northern Ireland census "
        }
    """
    And the 11th link href value should be "https://www.scotlandscensus.gov.uk/"
    And the 12th link href value should be "https://www.nisra.gov.uk/statistics/census"

    Scenario: GET /census and checking the response status 200
        When I navigate to "/census"
        Then the improve this page banner should be visible
        # hero content
        Then element "[data-test='hero-h1']" should be visible
        And element "[data-test='hero-p']" should be visible
        And the page should have the following content
        """
            {
                "[data-test='hero-h1']": "Census",
                "[data-test='hero-p']": "The census takes place every 10 years. It gives us a picture of all the people and households in England and Wales."
            }
        """
        # search container section #1
        And the page should have the following content
        """
            {
                "[data-test='search-1'] > h2": "About the census",
                "[data-test='search-1'] > p": "Find out what the census is and why it's important for all of us.",
                "[data-test='search-1'] > ul > li:nth-child(1) > a": "About the census"
            }
        """
        And the 1st link href value should be "census/aboutcensus"
        # search container section #2
        And the page should have the following content
        """
            {
                "[data-test='search-2'] > h2": "Census 2021 data",
                "[data-test='search-2'] > p": "Find data for Census 2021.",
                "[data-test='search-2'] > ul > li:nth-child(1) > a": "Get census data (England and Wales)",
                "[data-test='search-2'] > ul > li:nth-child(2) > a": "Get census data (Wales)"
            }
        """
        And the 2nd link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimatesenglandandwalescensus2021"
        And the 3rd link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimateswalescensus2021"
        # search container section #3
        And the page should have the following content
        """
            {
                "[data-test='search-3'] > h2": "Census releases",
                "[data-test='search-3'] > p": "See what we've published, and our plans for the future.",
                "[data-test='search-3'] > ul > li:nth-child(1) > a": "Release calendar",
                "[data-test='search-3'] > ul > li:nth-child(2) > a": "First results and timeline"
            }
        """
        And the 4th link href value should be "/releasecalendar?query=census&fromDateDay=&fromDateMonth=&fromDateYear=&toDateDay=&toDateMonth=&toDateYear=&view=upcoming"
        And the 5th link href value should be "/census/censustransformationprogramme/census2021outputs/releaseplans"
        # search container section #4
        And the page should have the following content
        """
            {
                "[data-test='search-4'] > h2": "Census topics",
                "[data-test='search-4'] > p": "Find census data and analysis using these topics.",
                "[data-test='search-4'] > ul > li:nth-child(1) > a": "Topic summaries"
            }
        """
        And the 6th link href value should be "/census/aboutcensus/censusproducts/topicsummaries"
        # search container section #5
        And the page should have the following content
        """
            {
                "[data-test='search-5'] > h2": "Historic census data",
                "[data-test='search-5'] > p": "Find census data and analysis for 2011 and earlier.",
                "[data-test='search-5'] > ul > li:nth-child(1) > a": "Get all historic census data"
            }
        """
        And the 7th link href value should be "/census/historiccensusdata"
            # search container section #6
        And the page should have the following content
        """
            {
                "[data-test='search-6'] > h2": "Planning for Census 2021",
                "[data-test='search-6'] > p": "How we researched, prepared and planned for Census 2021.",
                "[data-test='search-6'] > ul > li:nth-child(1) > a": "Planning for Census 2021"
            }
        """
        And the 8th link href value should be "census/planningforcensus2021"
            # search container section #7
        And the page should have the following content
        """
            {
                "[data-test='search-7'] > h2": "Contact us",
                "[data-test='search-7'] > p": "If you need help, contact census customer services.",
                "[data-test='search-7'] > ul > li:nth-child(1) > a": "Census customer services",
                "[data-test='search-7'] > ul > li:nth-child(2) > a": "Request a 2011 census dataset"
            }
        """
        And the 9th link href value should be "/census/censuscustomerservices"
        And the 10th link href value should be "/census/2011census/2011censusdata/2011censusadhoctables"
            # search container section #8
        And the page should have the following content
        """
            {
                "[data-test='search-8'] > h2": "Scotland and Northern Ireland censuses",
                "[data-test='search-8'] > p": "We are responsible for the census in England and Wales.",
                "[data-test='search-8'] > ul > li:nth-child(1) > a": "Scotland census ",
                "[data-test='search-8'] > ul > li:nth-child(2) > a": "Northern Ireland census "
            }
        """
        And the 11th link href value should be "https://www.scotlandscensus.gov.uk/"
        And the 12th link href value should be "https://www.nisra.gov.uk/statistics/census"