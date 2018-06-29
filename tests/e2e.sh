#!/usr/bin/env bash

export FILE=/Users/marek.bejda/go/src/github.com/FabricBros/chaincode/fixtures/invoice_ex1.json
export CMD=`jq --argfile file $FILE '.Args[1] = ($file | tojson)' input.json | tr -d " \t\n\r"`
#jq --argfile file $FILE '.Records[0].Sns.Message = ($file | tojson)' input.json


./fabric.sh invoke defaultcc v1 $CMD
./fabric.sh query defaultcc v1 '{"Args":["RetrieveInvoice", "IN3201"]}'