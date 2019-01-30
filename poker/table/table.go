package table

import (
    . "poker/client"
    . "poker/card"
    "sync"
    )

type Table struct {
    
    clients []Client
    deck []Card
    winingCombinations map[string]int

    mutex sync.RWMutex
}
