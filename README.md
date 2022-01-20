# tech-task
As part of the recruitment process we want to know how you think, code and structure your work. In order to do that, we're going to ask you to complete this coding challenge.  


There is also a "hello world" go application in cmd/ and a docker-compose.yml for running Amazon postgres locally.

We need you to:

Provide a Go implementation of the http service in the cmd/ directory of this repo.
Implement a postgres based store for this HTTP service
Provide adequate test coverage for this simple service

#Challange 
- Create the Http/GRPC services to perform following operation 
- Create Bank 
    - BANK NAME 
    - IFSC CODE 
    - BRANCH NAME 
    - BANK ID 
    - etc 
- List the Banks - Good to have  support of filerting with Branch Name 
- Delete Banks 
- Get Banks With ID 
- Update Bank details 

- Create the Account in the Bank
  - Account ID 
  - Bank Name 
  - Branch Name 
  - Accound Holder Name 
  - Idenetiy ID 
  - First Name 
  - Last Name 
  - Address 
- Get the Account By ID 


## for database migration use 
- https://github.com/golang-migrate/migrate 

## Must Have 
- Unit/Integration Testing 
- Logging 
- Error Handling 
- API Working as Expected

## Nice to have 
- Configuration & Secret Management 
- Tracing & Monitoring

