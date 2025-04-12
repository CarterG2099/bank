# bank
A common dice game to use for a bot fighting tournament.

## Rules
- Everybody starts with 0 points. 
- There are 20 rounds. 
- Every round begins with a pot of 0. 
- Every time you roll the dice, add whatever you get to the pot. 
- Rolling a 1 ends the round and empties the pot.
- Rolling a 2 doubles whatever is currently in the pot.  
- If your first roll is a 1 or a 2, add what you got to the pot instead of following the previous two rules.

Every time the dice is rolled you will be given a chance to "bank" (see title of game). Banking means you add the
current score of the pot to your personal score, but you can only bank once per round. The object of the game is to end
with more points than your opponents.

## Implementing

Implement the `bank.PlayerStrategy` interface found in `player.go`. See comments on that interface for details.

Create your bot in the example package, or any other package of your choice that isn't the bank package.
