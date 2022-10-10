# web-scraper 

## run the project


- download the Project and and navigate to cmd\web-scraper  on command prompt
- to build the project simply run go build -o .bin/web-scraper
- to run the project simply run go run main.go. The project is configured to run on port 8081
- To run tests navigate back to root folder and run go test ./...


## docker build

- The build the docker image run docker build -t web-scraper
- To run the build image on docker with a specified port docker run -p  8081:8081 web-scraper


## Improvements to be made
- Limited tests were added. needs to add more test to cover all scenarios .
        - Needs to add tests to the handlers as well
- Front End tests were not added .
- Adding make file configurations to build the api and app in a easy way 
- Refactor the code and adehring to standard folder structure of Go .
- Adding more Exception handling and Fault Tolerance machanizems.
- Improvements in warning ,Information and error logs
-cors is configured in a default manner allowing all regions .needs  improve to allow only known headers and origin.
-adding Dcoker images to a docker registry 
-
