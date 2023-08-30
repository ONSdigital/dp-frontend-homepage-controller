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
            "[data-test='census-about'] > h2": "About the census",
            "[data-test='census-about'] > p": "Find out what the census is and why it's important for all of us.",
            "[data-test='census-about'] > ul > li:nth-child(1) > a": "About the census"
        }
    """
    And the census about link href value should be "census/aboutcensus"
    # search container section #2
    And the page should have the following content
    """
        {
            "[data-test='census-data'] > h2": "Census 2021 data",
            "[data-test='census-data'] > p": "Find data for Census 2021.",
            "[data-test='census-data'] > ul > li:nth-child(1) > a": "Get census data (England and Wales)",
            "[data-test='census-data'] > ul > li:nth-child(2) > a": "Get census data (Wales)"
        }
    """
    And the 1st census data link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimatesenglandandwalescensus2021"
    And the 2nd census data link href value should be "/peoplepopulationandcommunity/populationandmigration/populationestimates/datasets/populationandhouseholdestimateswalescensus2021"
    # search container section #3
    And the page should have the following content
    """
        {
            "[data-test='census-releases'] > h2": "Census releases",
            "[data-test='census-releases'] > p": "See what we've published, and our plans for the future.",
            "[data-test='census-releases'] > ul > li:nth-child(1) > a": "Release calendar",
            "[data-test='census-releases'] > ul > li:nth-child(2) > a": "Results and timeline"
        }
    """
    And the 1st census releases link href value should be "/releasecalendar?query=census&fromDateDay=&fromDateMonth=&fromDateYear=&toDateDay=&toDateMonth=&toDateYear=&view=upcoming"
    And the 2nd census releases link href value should be "/census/censustransformationprogramme/census2021outputs/releaseplans"
    # search container section #4
    And the page should have the following content
    """
        {
            "[data-test='census-topics'] > h2": "Census topics",
            "[data-test='census-topics'] > p": "Find census data and analysis using these topics.",
            "[data-test='census-topics'] > ul > li:nth-child(1) > a": "Topic summaries"
        }
    """
    And the census topics link href value should be "/census/aboutcensus/censusproducts/topicsummaries"
    # search container section #5
    And the page should have the following content
    """
        {
            "[data-test='census-dictionary'] > h2": "Census 2021 dictionary",
            "[data-test='census-dictionary'] > p": "Definitions, variables and classifications to help when using Census 2021 data.",
            "[data-test='census-dictionary'] > ul > li:nth-child(1) > a": "Census 2021 dictionary"
        }
    """
    And the census dictionary link href value should be "/census/census2021dictionary"
        # search container section #6
    And the page should have the following content
    """
        {
            "[data-test='census-historic'] > h2": "Historic census data",
            "[data-test='census-historic'] > p": "Find census data and analysis for 2011 and earlier.",
            "[data-test='census-historic'] > ul > li:nth-child(1) > a": "Get all historic census data"
        }
    """
    And the census historic link href value should be "/census/historiccensusdata"
        # search container section #7
    And the page should have the following content
    """
        {
            "[data-test='census-planning'] > h2": "Planning for Census 2021",
            "[data-test='census-planning'] > p": "How we researched, prepared and planned for Census 2021.",
            "[data-test='census-planning'] > ul > li:nth-child(1) > a": "Planning for Census 2021"
        }
    """
    And the census planning link href value should be "census/planningforcensus2021"
        # search container section #8
    And the page should have the following content
    """
        {
            "[data-test='census-contact'] > h2": "Contact us",
            "[data-test='census-contact'] > p": "If you need help, contact census customer services.",
            "[data-test='census-contact'] > ul > li:nth-child(1) > a": "Census customer services",
            "[data-test='census-contact'] > ul > li:nth-child(2) > a": "Request a 2011 census dataset"
        }
    """
    And the 1st census contact link href value should be "/census/censuscustomerservices"
    And the 2nd census contact link href value should be "/census/2011census/2011censusdata/2011censusadhoctables"
        # search container section #9
    And the page should have the following content
    """
        {
            "[data-test='census-other'] > h2": "Scotland and Northern Ireland censuses",
            "[data-test='census-other'] > p": "We are responsible for the census in England and Wales.",
            "[data-test='census-other'] > ul > li:nth-child(1) > a": "Scotland census ",
            "[data-test='census-other'] > ul > li:nth-child(2) > a": "Northern Ireland census "
        }
    """
    And the 1st census other link href value should be "https://www.scotlandscensus.gov.uk/"
    And the 2nd census other link href value should be "https://www.nisra.gov.uk/statistics/census"
 Scenario: GET /census and checking the response status 200
    When the census hub flags are enabled
    And I navigate to "/census"
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
            "[data-test='census-about'] > h2": "About the census",
            "[data-test='census-about'] > p": "Find out what the census is and why it's important for all of us.",
            "[data-test='census-about'] > ul > li:nth-child(1) > a": "About the census"
        }
    """
    And the census about link href value should be "census/aboutcensus"
    # search container section #2
    And the page should have the following content
    """
        {
            "[data-test='census-data'] > h2": "Census 2021 data",
            "[data-test='census-data'] > p": "Find data for Census 2021.",
            "[data-test='census-data'] > ul > li:nth-child(1) > a": "Get census data",
            "[data-test='census-data'] > ul > li:nth-child(2) > a": "Create a custom dataset",
            "[data-test='census-data'] > ul > li:nth-child(3) > a": "Census 2021 data on NOMIS"
        }
    """
    And the 1st census data link href value should be "/census/find-a-dataset"
    And the 2nd census data link href value should be "/datasets/create"
    And the 3rd census data link href value should be "https://www.nomisweb.co.uk/sources/census_2021"


