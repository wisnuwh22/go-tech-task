__Language :__ English | [Bahasa Indonesia](README_ID.md)

# Go Technical Task
Suggested recipes for lunch API

## Time management
There is no deadline to do this tech task. It's up to you how you manage your time to accomplish at least the requirements.

## Assessment

Our assessment criteria will pay attention on:
- How the application is structured.
- Code quality (Clean code).
- Quality of tests.
- Interpretation of the problem.
- Use of `git`.
- Implementation and final execution.
- Commits, as this will allow us to understand some of the decisions you make throughout the process.

## User Story
As a User I would like to make a request to an API that will determine from a set of recipes what I can have for lunch today based on the contents of my fridge, so that I quickly decide what I’ll be having.

__Acceptance Criteria__
- Given that I have made a request to the `/lunch` endpoint I should receive a JSON response of the recipes 
that I can prepare based on the availability of ingredients in my fridge.
- Given that an ingredient is past its `use-by` date (inclusive), I should not receive recipes containing this ingredient.
- Given that an ingredient is past its `best-before` date (inclusive), but is still within its `use-by` date (inclusive), any recipe containing the oldest (less fresh) ingredient should placed at the bottom of the response object.

__Additional Criteria__
- The application MUST contains unit / integration tests (e.g. using `testing` package).
- The application MUST be completed using an `OOP` approach.
- The application MUST get A+ score from [Go Report Card](https://goreportcard.com/) website.
- Any dependencies MUST be installed using `Go Modules` (no need to commit dependencies, the
`go.sum` file will be sufficient).
- Use `Go 1.11` or above.
- Any installation, build steps, testing and usage instructions MUST be provided in a `README.md` file in the root of the application.

## Framework
We recommend you to build this application using the default package of Go.

## Application Data
For the purpose of this task, the application should simply read data from 2 x JSON files. The contents for these files can be found [here](ingredient/data.json) and [here](recipe/data.json).
 
## Submission
The application should be committed to a __public repository__ on `GitHub` or `BitBucket` (`<lastname>-<firstname>-techtask-go`) and simply send us a link to the repository.

## Bonus
Configure a `Docker` environment so that we can test and run the application quickly. The application should be installed with a single command.
