#!/usr/bin/env bash

# https://developers.google.com/places/supported_types#table1

URL="https://maps.googleapis.com/maps/api/place/nearbysearch/json"
URL="${URL}?key=AIzaSyBDymvvFxn3wwAtr-xFm3K9IzaU9R-iB6A"
URL="${URL}&location=$1,$2"
URL="${URL}&radius=$3"
URL="${URL}&rankby=prominence" #prominence
URL="${URL}&type=prominence" #prominence


curl -v "${URL}"

