package main

import "fmt"

type Player struct {
	health int 
}

func takeDamageFromExplosion(player *Player){
	fmt.Printf("Player health before explosion: %v\n", player.health)
	player.health -= 20;
	fmt.Printf("Player health after explosion: %v\n", player.health)
}

func main() {
	player := Player{ health: 100}
	takeDamageFromExplosion(&player)

	fmt.Printf("Player health after explosion: %v\n", player.health)
}
