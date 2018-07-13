# Chaincode Template

Run all tests with:
```
$ ginkgo
```

Run in devmode:
```
$ ./fabric.sh up
$ ./fabric.sh startCC ccname v1
$ ./fabric.sh runCC ccname v1
```

## Create/Query Invoice E2E

All from Fixture:
```
#!/usr/bin/env bash

./fabric.sh invoke '{"Args":["AddInvoices","[{\"FabricKey\":\"CN_AtlasEurope68109A1235\",\"Seller\":\"A3\",\"Date\":\"2-Jan\",\"Ref\":\"68109\",\"Buyer\":\"A5\",\"PONum\":\"A1235\",\"SKU\":\"85412\",\"Qty\":\"200\",\"Curr\":\"EUR\",\"UnitCost\":\"200\",\"Amount\":\"40,000\"}]"]}'
./fabric.sh query '{"Args":["RetrieveInvoice","CN_AtlasEurope68109A1235"]}'
```


All from Fixture:
```
#!/usr/bin/env bash

export FILE=/Users/marek.bejda/go/src/github.com/FabricBros/chaincode/fixtures/invoice_ex1.json
export INPUT=/Users/marek.bejda/go/src/github.com/FabricBros/chaincode/fixtures/e2e_input.json

export CMD=`jq --argfile file $FILE '.Args[1] = ($file | tojson)' $INPUT | tr -d " \t\n\r"`
#jq --argfile file $FILE '.Records[0].Sns.Message = ($file | tojson)' input.json


./fabric.sh invoke defaultcc v1 $CMD
./fabric.sh query defaultcc v1 '{"Args":["RetrieveInvoice","IN3201"]}'
```
