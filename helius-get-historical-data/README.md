# Получение исторических данных блокчейна

Зададимся целью получить список транзакций определенного token'а на Solana

## Получение адреса token'а

С помощью solscan найдем адрес кошелька token'а. Например:

https://solscan.io/token/6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN

Адрес: 6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN

## Получим список последних транзакций этого токена

```
curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"]}'
```

Отлично, мы получили список последних 1000 транзакций этого токена. В моем случае список транзакций заканчивался на ["51bSAVjC2MDJukD5tBx4Lqgw5UvKGT7s143qcqM35YVuHbFPcxvQqSyvAsep3QyxwfjTSoVwRkwA84FfrowBiUif", "3Q3sNziAYjrM7EXZQTXeBUEmkt9z3k3BFaaQiKLfnB7yQjJYKHF9fQcApfFAKyjSDjjrewwdGu9S2b7ED1FZBmC3"]. Для того, чтобы получить предыдущие от этого транзакции, сделаем запрос

```
curl -L \
  --request POST \
  --url 'https://mainnet.helius-rpc.com/?api-key=<paste-your-api-key>' \
  --header 'Content-Type: application/json' \
  --data '{"id":"1","jsonrpc":"2.0","method":"getSignaturesForAddress","params":["6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN", {"limit": 1000, "before": "51bSAVjC2MDJukD5tBx4Lqgw5UvKGT7s143qcqM35YVuHbFPcxvQqSyvAsep3QyxwfjTSoVwRkwA84FfrowBiUif"}]}'
```

## Получим последние 10'000 транзакций через код

```
cd solana-get-historical-data
cd get_transactions
go run get_transactions.go
```
