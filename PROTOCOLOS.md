# Definições de protocolos

As operações terão numeração continuada entre mensagens de controle e mensagens de dados.

A aplicação deve levar em conta que os bytes estão sendo enviados de maneira Big Endian.

## Controle

### Operações

- SUB (0): inscrever-se em tópico
- UNSUB (1): desinscrever-se de tópico
- ADD (2): criar tópico
- DEL (3): deletar tópico
- UPD_META (4): atualizar cache de metadados do cluster (servidor envia quando houver uma alteração)
- GET_META (5): atualizar cache de metadados do cluster (cliente envia quando precisar dos metadados sob demanda)
- PING (6): verificar se servidor/cliente ainda está ativo

### Decisões

- As mensagens (headers + payload) devem ser prefixadas por 4 bytes, delimitando o tamanho da mensagem
- Os headers das mensagens serão os seguintes (em ordem):
  - Operação - 1 byte
  - Tópico - 2 bytes (deve ser 0 em caso de a operação não precisar)
- Metadados do cluster devem ser enviados utilizando Protobuf, focando em eficiência de tamanho e velocidade, evitando grandes atrasos nas atualizações de metadados
- As seguintes operações devem retornar 0 para sucesso e 1 para falha:
  - SUB
  - UNSUB
  - ADD
  - DEL
  - UPD_META
  - PING
- SUB deve enviar também o nome do cliente fazendo a operação
  - Deve fazer isso mandando no payload o nome
- GET_META deve retornar e UPD_META deve enviar os seguintes dados:
  - Brokers ativos
    - Tópicos
      - Nome
      - Identificador
- O tamanho máximo de mensagem recebida pelo servidor deve ser de 1024B
- Cliente deve confiar no tamanho das respostas do servidor

## Dados

### Operações

- SEND (7): publicação de mensagem do publisher
- FORWARDING (8): encaminhamento de mensagem para consumers

### Decisões

- As mensagens serão lidas com 2048 bytes cada, o que passar disso será ignorado
- Os headers das mensagens serão os seguintes (em ordem):
  - Operação - 1 byte
  - Tópico - 2 bytes
