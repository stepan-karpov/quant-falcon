# Запуск тестового валидатора на удаленном хосте (или локально)

## Установка Solana CLI, запуск тестового валидатора (вручную)

1. Установите Solana CLI:

   ```sh -c "$(curl -sSfL https://release.anza.xyz/stable/install)"```
   
2. Добавьте Solana в PATH:

   ```export PATH="/home/ubuntu/.local/share/solana/install/active_release/bin:$PATH"```
   
3. Запустите тестовый валидатор Solana:
   ```solana-test-validator```
   
   В случае ошибки можно посмотреть лог командой:
   ```solana-test-validator --log```
   
### Проблемы, с которыми я столкнулся

- Мой процессор не поддерживался `solana-test-validator` (ошибка: `Incompatible CPU detected: missing AVX support. Please build from source on the target`).
- Недостаточно оперативной памяти (`solana-test-validator` по умолчанию пытается аллоцировать 1Гб оперативной памяти). Попытался выделить swap, но даже с ним у меня не получилось, пересел на машинку получше.

В случае закрытия терминала с запущенным процессом solana-test-validator сам валидатор остановится, поэтому рекомендуется запускать его в Dockerfile как все стандартные приложения

## Установка Solana CLI, запуск тестового валидатора (из Dockerfile)

1. Установка Docker на удаленном хосте (если еще не установлен)

  ```curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -```

  ```sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"```

  ```sudo apt update```

  ```sudo apt install docker-ce```

  ```sudo systemctl status docker```

2. Запуск и сборка Docker-образа

  ```cd solana-local```

  ```sudo docker build -t solana-validator .```

  ```sudo docker images```

  ```sudo docker run -d -d -p 8899:8899 --name solana-validator solana-validator```

  Проверить, что контейнер поднялся: ```sudo docker ps```

  Для подключения к контейнеру: ```sudo docker exec -it solana-validator /bin/bash```

  Для остановки: ```sudo docker stop solana-validator```

  Для удаления: ```sudo docker rm solana-validator```

## Проверка работоспособности

```
curl -X POST http://15.236.48.228:8899 \
     -H "Content-Type: application/json" \
     -d '{
           "jsonrpc": "2.0",
           "id": 1,
           "method": "getSlot",
           "params": []
         }'
```

или

```
curl -X POST http://localhost:8899 \
     -H "Content-Type: application/json" \
     -d '{
           "jsonrpc": "2.0",
           "id": 1,
           "method": "getSlot",
           "params": []
         }'
```

## Подключение с клиента

```solana config set --url http://15.236.48.228:8899```

Проверить что работает: ```solana balance 954m3369W174DASH1gyMGFEZaYF74eyxuQbyVMjVatxQ``` - должен быть очень большой т.к. с этого адреса запускался solana-test-validator

## Создание токена (TODO)

```sudo apt install rustup```

```cargo install spl-token-cli```
