#!/bin/sh

curl -X POST -H "Content-Type: application/json" -d @createExampleProduct.json http://localhost:8080/api/products