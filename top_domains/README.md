# Top Domain API
This Go API will take a JSON request with a url, and track the number of times a certain domain was hit. It stores this data in a NoSQL MongoDB database. Also, it returns the top 5 domains that were hit. 

## Endpoints
- `/ping`: health check
- `/domains`: GET all domains stored in MongoDB collection `top_domains`
- `/url`: POST a url to the server 
    - Example Request: `{"url": "http://nytimes.com/f82jd0203k349dk"}`
- `/top`: GET top 5 most hit domains

## How to run
- Reference [Makefile](./Makefile)
- `make run`