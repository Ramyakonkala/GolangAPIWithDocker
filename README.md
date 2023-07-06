# Fetch Receipt Processor API
 ## Command to run the docker containerized app
 Git clone the application to your system and open the terminal inside the application. Remove if any image with fetch_api exists.
 Run `docker compose up` command
 Use Postman or your preferred way to call the endpoints

## Implemented endpoint

1. Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.

2. Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

3. Get all Process Receipts

* Path: `/receipts/process`
* Method: `GET`
* Response: List of receipts JSON objects.