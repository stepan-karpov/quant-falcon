FROM solanafoundation/solana-verifiable-build

RUN apt-get update && apt-get install -y python3 python3-pip

WORKDIR /app

# Копируем файл helloworld.py в контейнер
COPY helloworld.py /app/

# Запускаем helloworld.py
CMD ["python3", "helloworld.py"]
