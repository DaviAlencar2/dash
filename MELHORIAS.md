# Melhorias

## Mensagem de controle

- Os 4 bytes de tamanho devem ser lidos fora, não entram no decode da mensagem (é necessário conhecer esse tamanho antes de ler a mensagem de fato)
- Na declaração de bodySize em encode o 3 que são os headers devem usar a const pra evitar retrabalho caso o tamanho dos headers mude
- Padronizar a lingua usada no codigo pra ingles
- No decode não precisa verificar se o tamanho lido é diferente do tamanho passado, quem precisa garantir isso é o cliente, ele vai ler os proximos X bytes com base nos 4 bytes de tamanho, nem mais, nem menos, nao tem como vir maior, ele vai ler no maximo aquilo (pode até ser feito, para evitar mal uso do cliente pela gente, mas deve vir certo já de fora)
