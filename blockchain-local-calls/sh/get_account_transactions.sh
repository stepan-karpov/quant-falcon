curl -X POST http://15.236.48.228:8899 \
-H "Content-Type: application/json" \
-d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "getConfirmedSignaturesForAddress2",
  "params": ["954m3369W174DASH1gyMGFEZaYF74eyxuQbyVMjVatxQ"]
}'