# gotest

## Synopsis

Houseanywhere.com's test job.

## Installation

1. `cd gotest`
2. To create table *listing* execute file *initdb.sql*. For example:
   
   `psql -U postgres haw.com/initdb.sql`
3. To parse xml feed and fill database use:
   
   `go install haw.com/process`
    
   `$GOPATH/bin/process feed.xml partner`

   Now only one partner is valid: xxx
4. To start web-server use:
    
   `go install haw.com/api/app`
    
   `$GOPATH/bin/app`
  
   List of listings will be available on http://localhost:8080/index
