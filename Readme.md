
# A simple scanner for mysql

this is part of the datasec challange interview

  

## About Architecture

I tried to follow as much as possible the clean architecture principles, so each service will have an usecase, services that needs to access any kind of data will also have a repository.
The architecture is far from being a settled theme here, so this format could be different, or even better.
  
  ## How it works
 -  The user will create a datastorage using the POST `/api/v1/database`
   endpoint   
- Then when needed, the user can trigger a new scan of that   
   datastorage using the POST `/api/v1/database/scan/{datastoreId}` endpoint
 - The system will use a set of regex rules that are stored at their   
   main database to classify all tables in the target database
- Finally, the user can check the scan result with the endpoint `api/v1/database/scan/{datastoreId}`

## Endpoints

### Post `/api/v1/database`
 body:
 
    {
    	"user":string,
    	"host":string,
    	"password":string,
    	"port": number,
    	"database":string
    }
   response:

    {
    	"id": number, #the datastorageId, used to request scans and get results
    	"host": string,
    	"classes": []
    }


### Get `api/v1/database/scan/{datastorageId}`
response:

    {
    	"id": number,
    	"host": string,
    	"classes": []
    }

### Post `/api/v1/database/scan/{datastorageId}`
response: 

    [{
    "Table": string,
    "Column": string,
    "Class": string # This is the data classification, like USERNAME,EMAIL_ADDRESS and so on
    }]

## Configuration

 - Create the config.json  file at the meliScanner folder, do it like
   this:

    {
      "database": {
        "url": "user:password@tcp(localhost:3306)/meliScannerDB"
      },
      "server": {
        "port": "3000"
      }
    }

 - At the folder migrations, you can find the sql file to create the
   tables, and a sample of how to create rules, the system uses the
   SQL's regexp expression to filter and sort the column names, anything
   that works with SQL's regexp should work here

## Logs
all requests are directed to the stdout using the apache log format.

## Improvements
here are a list of things I would like to have done, but i couldn't do to the time constraint:

 - [ ] Encrypt the database password using an AES key stored in some secret vault(probably AWS secrets manager)
 - [ ] Implement JWT authentication
 - [ ] A route to create new rules
 - [ ] Better way to implement new datastorages(to easily scan things like S3, Athena, or any other type of storage) 
 - [ ] Refactor the code so it can properly adhere to the clean architecture
