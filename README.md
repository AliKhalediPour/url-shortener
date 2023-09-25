## Simple URL Shortener written in Golang


### Run Databases
> make db-run 

### Run URL Shortener service using docker
> make docker-run

#### To generate new short link:
`curl -L 'localhost:8000/v1/generate' --header 'Content-Type: application/json' --data '{ "long": "https://p30download.com"}'`

#### To use generated short link:
`curl -L 'localhost:8000/v1/url/7lVJ1g2Le'`