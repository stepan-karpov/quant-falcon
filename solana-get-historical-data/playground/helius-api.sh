curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"]}'

# result in next.json

curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN", {"limit": 1000, "before": "51bSAVjC2MDJukD5tBx4Lqgw5UvKGT7s143qcqM35YVuHbFPcxvQqSyvAsep3QyxwfjTSoVwRkwA84FfrowBiUif"}]}'

# result in previous.json