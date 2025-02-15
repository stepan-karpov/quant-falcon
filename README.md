# Мой репозиторий для ведения проекта по запуску торговли на блокчейне Quant Falcon

## Установка Solana CLI, запуск

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
