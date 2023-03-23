dp-frontend-homepage-controller
================

### Getting started

This project uses go modules, ensure that go version 1.12 or above is in use.
If on go 1.12 then ensure the project either resides outside of your `GOPATH` or `GO111MODULE` is set to true

1. To start the service use make, `make debug`

### Dependencies


### Configuration

| Config                      | Description                                                                            | Default    |
|-----------------------------|----------------------------------------------------------------------------------------| -----------|
| BindAddr                    | The Port to run on                                                                     | :24400     |
| GracefulShutdownTimeout     | Time to wait during graceful shutdown                                                  | 5 seconds  |
| HealthCheckInterval         | Interval between health checks                                                         | 30 seconds |
| HealthCheckCriticalTimeout  | Amount of time to pass since last healthy health check to be deemed a critical failure | 90 seconds |
| CacheUpdateInterval         | Duration for homepage cache updation                                                   | 10 seconds |
| IsPublishingMode            | Mode in which service is running                                                       | false      |
| Languages                   | Languages which are supported separated by comma                                       | en,cy      |
| EnableCensusTopicSubsection | Displays topics subsection                                                             | false      |
| EnableCustomDataset         | Displays link to custom dataset builder                                                | false      |
| EnableGetDataCard           | Displays Get Data Card                                                                 | false      |
| EnableNewNavBar             | Enables Topic API driven Nav bar                                                       | false      |
| DatasetFinderEnabled        | Displays link to dataset finder                                                        | false      |

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details.

### License

Copyright © 2021, Office for National Statistics (https://www.ons.gov.uk)

Released under MIT license, see [LICENSE](LICENSE.md) for details.
