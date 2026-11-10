# Decisões

## Tamanhos máximos

- TCP: ???
- UDP: 2048 (contando headers)

## Ping

Acho que uma boa ideia para o ping é deixar uma go routine rodando no cliente fazendo pings e usando um mutex no canal TCP por exemplo pra evitar mensagens quebradas.

## Como identificar um cliente UDP
Criar struct com um id, UDPAddr e outros atributos
