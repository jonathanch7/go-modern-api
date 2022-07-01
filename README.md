# Go Modern Api

Go Modern Api example whit: Clean architecture, Hexagonal architecture, CQRS

## Environment Variables

- **SERVICE_PORT**: Port for deployment.
- **DB_USER**: Database user.
- **DB_PASSWORD**: Database password.
- **DB_SCHEMA**: Database schema.

## Run

- For local development execution use the command ```make rundev```, this commands uses _.dev.env_ file.

## Test

To run the tests use the commands:

- Unit tests: ```make unittests```.
  - Unit tests with coverage: ```make unittestscover```.
  - Unit tests with coverage in HTML: ```make unittestscoverHTML```.
- End-to-End tests: ```make e2etests```, this test uses _.e2e.env_ file.
- All tests: ```make tests```.

