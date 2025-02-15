curl -X POST http://15.236.48.228:8899 \
-H "Content-Type: application/json" \
-d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "getBlockHeight",
  "params": []
}'
