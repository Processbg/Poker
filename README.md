# Poker

## Instalation: 

You can install it using this go command go get "https://github.com/Processbg/Poker".

## Summary:

The project repesents an client-server web application for playing the game poker. The basic idea is that a client logs with his account name and password to the server where he can choose the type of poker he wants to play and after that he is atomatically put in a room where the selected poker game is played and on a table which has the least empty seats. If all the players on the table leave the table, he will be put on the next table with least empty seats at the given moment and if all the tables are full he will be put on a new table and wait for players to join him to start playing.

## Features:

A client can log in with his user name and password to the server which has a hash table where he checks if the client has an account and than checks the value of the password if it is the same using the user name as a key.

After the log-in, the server will ask for a string like command in order to choose the correct interface for the selected poker game.
Every different poker game interface has a diffrent slice with all the combinations of five cards sorted by the priority of each combination for the selected game.

After the client has decided the type of game he wants. He will be asked to choose the minimal waiger for a table he wants to bet if he has enough money for it. 

After there are at least two players on a table the selected game will start with a basic graphics user interface displaying the table viewing it from above and where are the seat of the players. Each player has two cards displayed in front of him face up but he can't see the other players cards beacuse they are face down. On the middle of the table there will be displayed the cards all the players can see face up and each new card after a round of the game until it is finished. A round is discribed by the agreement of all the players on the table for the playing amount of winnings. Each Player starts with a fixed sum from the first log-in if he loses his credits in the next log-in he will have the same amount as from the very beginning or the yearnings he has accumulated from his previous games.

When a player has his turn on the table he has four buttons to choose from "fold", "check", "call" or "raise". If he chooses "raise" he will be asked to enter the amount of credits to raise which can't be bigger than the amount he has accumulated on the table he is playing.
The amount for which all the players are playing will be diplayed under the cards in the midle of the table. If he chooses "fold" he won't play untill there is a winner of the current game but he can watch the other players. If he presses "check" he continues to the next round and if he presses "call" the amount of credits he agrees on is added to the hall amount for which all the players involved are playing and continues to the next round.   
