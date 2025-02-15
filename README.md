# Мой репозиторий для ведения проекта по запуску торговли на блокчейне Quant Falcon

## Структура репозитория

`solana-local` - поддиректория, позволяющая локально развертывать solana-test-validator 

`solana-local-calls` - поддиректория, содержащая примеры использования локально поднятого блокчейна (предполагается, что поднято на IP=15.236.48.228)

Ниже представлены основные команды, которые могут пригодиться при ведении проекта

## Установка Solana CLI

1.   ```sh -c "$(curl -sSfL https://release.anza.xyz/stable/install)"```
   
2. Добавьте Solana в PATH:

   ```export PATH="/home/ubuntu/.local/share/solana/install/active_release/bin:$PATH"```
   
## Основные команды Solana CLI ([дока](https://solana.com/ru/docs/intro/installation))

- `solana config set --url localhost` - переключиться на localhost

- `solana-keygen new -o /home/ubuntu/.config/solana/id.json` - создать новый кошелек

- `solana address` - показать публичный адрес текущего кошелька

- `solana balance` - показать баланс текущего кошелька