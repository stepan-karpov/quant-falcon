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

curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"getTransaction","params":["4naZGV5Fg327oTG4YgFokLR7CCKkrNDZgjd6Bgpv6jZcWHQHGs8SruMNaLY7fjoKTNiPg5ZHimKQAEyUmniLQqaf","json"]}'

# result in transaction_info.json

curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"jsonrpc":"2.0","id":1,"method":"getTransaction","params":["2WMpxfAX991GVEhGDrSRaSCUsXHhTUGhwYa9F4q6odTUncKhenFmfZJRTuZHfKnXs9DCWANh6EX2iZTncd2KH54G", {"maxSupportedTransactionVersion": 0}]}'

# result in trump_transaction_info.json