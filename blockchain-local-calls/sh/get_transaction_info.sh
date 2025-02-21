curl -X POST http://15.236.48.228:8899 \
-H "Content-Type: application/json" \
-d '{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "getConfirmedTransaction",
  "params": ["3uiRb8SbbkLDusBeSE6RHuz4RsAzZoTztaBeZDyxUEjLVqQKm1PW7WE1ficjQksu4x4T15288eFDTqk2VS8ymW59"]
}'
