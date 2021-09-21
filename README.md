# Lending API

## Tech Stack : 

```
   - gin-gonic 
   - Glide
   - Stretchr
```

## Why Need Those Tech Stack
   * Gin-Gonic  
   A http framework, gin gonic have a feature recovery (crash-free).  
   
   * Glide  
   A dependency management, will help a new developer to continue the project without get dependency one by one. glide install will get all  
   dependency based on glide.yaml file

   * Stretchr
   Stretchr have assert function to help developer when build tdd apps.

## Minimum infrastructure Requirement : 

   ```
   OS Linux (Centos 7)  
   RAM 1 GB  
   Storage 30GB  
   Processor 64 Bit  
   ```
## Directory structure

```
  + $GOPATH/
  |
  +--+ src/
  |  |
  |  +--+ github.com/
  |     |
  |     +--+ lending/
  |        |
  |        +--+ handler/
  |        |  |
  |        |  +--+ loan/
  |        |     |
  |        +--+ internal/
  |        |  |
  |        |  +--+ core/
  |        |     |
  |        |     +--+ domain/
  |        |     +--+ ports/
  |        |     +--+ queue/
  |        |     +--+ services/
  |        |  +--+ fsm/
  |        |  +--+ logger/
  |        +--+ repositories
  |        +--+ resources
  |        +--+ main.go
```

## Instructions

### How To Run Test File

    To run all test file, you can run using this command :
    ```
    go test
    ```

### Setup Instructions 
    Under `$GOPATH` directory, do the following command :    
     ```
     mkdir -p src/github.com/lending
     cd src/github.com/lending
     git clone <url>
     go mod download
     ``` 
     if you are using go >= 1.11 don't forget to turn on go module. GO111Module=on
### Application Deployment
    go to resources
    run : docker build -t lending:<tag> .
    run : docker run -d -p <your available port>:8080 --name go-app-container lending@<tag>
