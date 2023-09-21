dp-frontend-homepage-controller
================

### Getting started

This project uses go modules, ensure that go version 1.12 or above is in use.
If on go 1.12 then ensure the project either resides outside of your `GOPATH` or `GO111MODULE` is set to true

1. To start the service use make, `make debug`

### Dependencies

### Configuration

| Config                              | Description                                                                            | Default                   |
|-------------------------------------|----------------------------------------------------------------------------------------|---------------------------|
| API_ROUTER_URL                      | The URL of the [dp-api-router](https://github.com/ONSdigital/dp-api-router)            | http://localhost:23200/v1 |
| BIND_ADDR                           | The Port to run on                                                                     | :24400                    |
| CACHE_CENSUS_TOPICS_UPDATE_INTERVAL | Duration for census topic cache updates                                                | 1 minute                  |
| CACHE_NAVIGATION_UPDATE_INTERVAL    | Duration for navigation cache updates                                                  | 1 minute                  |
| CACHE_UPDATE_INTERVAL               | Duration for homepage cache updation                                                   | 10 seconds                |
| CENSUS_TOPIC_ID                     | Root census id (for getting census topics)                                             | 4445                      |
| DATASET_FINDER_ENABLED              | Displays link to dataset finder                                                        | false                     |
| DEBUG                               |                                                                                        |                           |
| ENABLE_CENSUS_TOPIC_SUBSECTION      | Displays topics subsection                                                             | false                     |
| ENABLE_CUSTOM_DATASET               | Displays link to custom dataset builder                                                | false                     |
| ENABLE_FEEDBACK_API                 | Enable the feedback API for use in the feedback form (as opposed to the controller)    | false                     |
| ENABLE_GET_DATA_CARD                | Displays Get Data Card                                                                 | false                     |
| ENABLE_NEW_NAVBAR                   | Enables Topic API driven Nav bar                                                       | false                     |
| FEEDBACK_API_URL                    |                                                                                        |                           |
| GRACEFUL_SHUTDOWN_TIMEOUT           | Time to wait during graceful shutdown                                                  | 5 seconds                 |
| HEALTHCHECK_CRITICAL_TIMEOUT        | Amount of time to pass since last healthy health check to be deemed a critical failure | 90 seconds                |
| HEALTHCHECK_INTERVAL                | Interval between health checks                                                         | 30 seconds                |
| IS_PUBLISHING_MODE                  | Mode in which service is running                                                       | false                     |
| PATTERN_LIBRARY_ASSETS_PATH         |                                                                                        |                           |
| SERVICE_AUTH_TOKEN                  |                                                                                        |                           |
| SIXTEENS_VERSION                    | Homepage still uses a Sixteens version for styling                                     |                           |
| SITE_DOMAIN                         |                                                                                        | localhost                 |
| SUPPORTED_LANGUAGES                 | Languages which are supported separated by comma                                       | en,cy                     |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
