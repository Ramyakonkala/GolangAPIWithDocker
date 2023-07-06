 # Fetch Receipt Processor API
 ## Command to run the docker containerized app
`docker compose up`

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