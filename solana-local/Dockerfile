FROM solanalabs/solana:v1.9.24

RUN apt-get update && apt-get install -y \
    curl \
    unzip \
    && rm -rf /var/lib/apt/lists/* \
    procps


WORKDIR /solana

EXPOSE 8899
CMD solana-test-validator  --ledger /home/test-ledger
