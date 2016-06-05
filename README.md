# gotest

## Installation

1. `cd gotest`
2. To create table *listing* execute file *initdb.sql*. For example:
   
   `psql -U postgres haw.com/initdb.sql`
3. Parse xml feed and fill database:
   
   `go install haw.com/process`
    
   `$GOPATH/bin/process <path_to_feedfile.xml> <partner_name>`

   Now only one value for `partner_name` is allowed: `xxx`
4. Start web-server:
    
   `go install haw.com/api/app`
    
   `$GOPATH/bin/app`
  
   List of listings will be available on http://localhost:8080/index
