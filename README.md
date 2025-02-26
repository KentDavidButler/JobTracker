# JobTracker

Under Development as of 02/2025

Job Tracker is intended to reduce my need for a google sheet to track jobs, and instead store all of the same data within the APP itself. 

The first version is a simple app that runs locally and stores data to a Postgres DB. The way to send and get data is via the `/jobpostings` api endpoint using simple get and post commands.

Example of getting all jobs:
`curl localhost:8080/jobpostings`

Example of getting a specific job:
`curl localhost:8080/jobpostings/26984446-374a-4b76-a875-63ccb82247ac`

Example of adding a job to the app:
````
curl -d '{"id":"123", "companyName":"some_Company2", "positionLink":"www.some_company2.com"}' -H "Content-Type: application/json" -X POST http://localhost:8080/jobpostings
```

# Intent
Hopefully some day the full intent would be to have a lightweight page to interact with, that'll allow a way to easily add, view, and update jobs that I've applied to.

# How to run
To run, you'll need to add a .env file after you've cloned this. In the .env please add values of your db
`POSTGRES_USERNAME=example`
`POSTGRES_PASSWORD=examplePassword`
`POSTGRES_DBNAME=postgres`
And of course, a locally running postgres database running with your DB name in the env vars file


# Special Considerations
When working with Go and Postgress I found it is easiest to make sure all DB values are defaulted to something. Allowing Null values from the DB was clunky with a lot of additional checking of the values when pulling and pushing to the DB.
I'm trying to use the GOLANG standard library for GO 1.24, that said I've been unable to pull env vars using the standard library and had to use `joho/godotenv` since the standard library 