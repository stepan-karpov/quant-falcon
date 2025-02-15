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

  ```Hello```