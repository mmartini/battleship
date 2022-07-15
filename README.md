Ideias:

 1 - Todo jogador usa um import da biblioteca do NavalBattle
 	1.1 - A troca de dados pode ocorrer via channel, teriamos channel para: Inicio de partida (preciso do seu board), qual sua jogada? (o tiro), resultado da jogada (não acertou ou acertou matando ou não o barco), fim de partida. Todas essas opções podem utilizar o mesmo channel apenas teriamos que ter um tipo na struct.

 	1.2 - A troca de dados pode ocorrer via funções, teriamos uma abstração que obriga a existencia de funções para cada um dos tipos de interações

 	1.3 - A troca de informações ocorre com um dos modulos acima, mas o nosso pacote implementa a parte de comunicação RPC, socket ou tcp/http (precisamos escolher uma)

 	1.4 - O servidor será um serviço ou teremos a execução de um binário para cada game ou ainda um binario para cada jogador?

 	1.5 - Go será compilado com quantidade de pacotes super restritas (testar possibilidades)

2 - Possiveis modos de jogo:

	2.1 - Peças são variaveis quantidade e tipos
	2.2 - Jogadas revelam a peça que foi acertada
	2.3 - Jogadas revelam quando um peça é acertada por inteiro
	2.4 - Opção de receber os tiros que o adversário dá, todas ou apenas quando afunda (destruição completa da peça)
	2.5 - Opção que permite peças encostadas uma nas outras

3 - TODO List

	3.1 - Criar um ranking de jogadores, pensando inclusive no controle de versão de cada um dos jogadores (quantidade), ranking de pontos por jogador que mais ganha, mais rápido em respostas (média), menos partidas perdidas por falha de programação (item 3.4), melhor distribuição de tabuleiro (quem tiver a menor media de tiros contra acertados), melhor artilharia (quem tiver a menor média de acertos por tiro), comparação de código fonte para evitar plágios
	3.2 - Criar controles anti má programação, tabuleiro demora para ser gerado, tiro demora para ser dado, panics do fonte
	3.3 - Jogadores escolhem os modos de jogo que ele aceita
	3.4 - Penalizar jogadores que perdem a partida por dar tiro com retorno de erro (partidas que não terminam com falha de configuração / limites de timeout).
	3.5 Pensar em modo de bonificação, tipo dar um super shot para o jogador, esse super shot nada mais é que um tiro que acerta mais de um ponto