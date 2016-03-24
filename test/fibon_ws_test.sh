#!/bin/bash

# To test a positive integer
if [[ $(curl --silent -X GET http://localhost:8080/7) = '[0,1,1,2,3,5,8]' ]]; then
  echo "Functional test for positive integer OK"
else
  echo "Functional test for positive integer failed"
fi

# To test a negative number
if [[ $(curl --silent -X GET http://localhost:8080/-1) = '{"Error":"Input must be a positive integer"}' ]]; then
  echo "Functional test for negative number OK"
else
  echo "Functional test for negativef number failed"
fi

# To test an invalid input
if [[ $(curl --silent -X GET http://localhost:8080/xyz) = '{"Error":"Input must be a positive integer"}' ]]; then
  echo "Functional test for invalid input OK"
else
  echo "Functional test for invalid input failed"
fi
