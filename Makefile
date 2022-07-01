rundev:
	@./scripts/run.sh dev

unittests:
	@./scripts/test.sh unit

unittestscover:
	@./scripts/test.sh unit cover

unittestscoverHTML:
	@./scripts/test.sh unit coverHTML

e2etests:
	@./scripts/test.sh e2e

tests:
	@./scripts/test.sh unit
	@./scripts/test.sh e2e

genopenapi:
	@./scripts/open-api.sh
