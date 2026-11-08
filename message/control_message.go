package message

/*
acho que é bom modelar por enquanto tudo de controle menos a parte de metadata aqui
acho que protobuf brilha so nesse caso de metadata, de resto acho que é bem parecido com data messages
aí é bom fazer essa parte do protobuf em conjunto
upd e get meta sao muito diferentes das outras, tem que definir como delimitar cada broker, cada dado do broker, se vai mandar tudo
ou so o que mudou, etc
aqui é fazer uma coisa generica que funciona pra tudo, sem tratar muito os payloads e criar structs auxiliares para tipos especificos
por exemplo: add precisa tratar a string que vai ser o nome do struct pra nao ficar um bagulho estrnaho com \n e etc
mesma coisa com os de metadata se pa, o que vai permitir que a gente decidida como tratar essa parte depois (com ou sem protobuf)
*/
type ControlMessage struct{}
