#!/usr/bin/env bash

./fabric.sh invoke -c '{"Args":["AddInvoices", "`cat ./fixtures/invoice_ex1.json`"]}'
./fabric.sh query -c  '{"Args":["RetrieveInvoice", "IN3201"]}'