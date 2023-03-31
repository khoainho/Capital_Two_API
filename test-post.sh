#!/bin/bash

# Test POST endpoint

domain="localhost"
port=3000

create-account () {
	curl --header "Content-Type: application/json" \
 --request POST \
 --data '{"FirstName":"John", "LastName":"Doe"}' \
 http://${domain}:${port}/account

curl --header "Content-Type: application/json" \
 --request POST \
 --data '{"first_name":"John", "last_name":"Doe"}' \
 http://${domain}:${port}/account

 curl --header "Content-Type: application/json" \
 --request POST \
 --data '{"first_name":"Khoa", "last_name":"Ho"}' \
 http://${domain}:${port}/account

# This should give an error: First and last name are null
  curl --header "Content-Type: application/json" \
 --request POST \
 --data '{"first_name":, "last_name":}' \
 http://${domain}:${port}/account
}

main () {
	create-account
}

main