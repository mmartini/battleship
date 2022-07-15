package main

import (
    "fmt"
)

const (
    BOARDSIZE = 15
    PERMITSIDEBYSIDE = false
)

func main () {
    parts := make(map[uint8]*sPart)

    parts[1] = newPart(Airplane.layout, Airplane.flip)
    parts[2] = newPart(AircraftCarrier.layout, AircraftCarrier.flip)
    parts[3] = newPart(Ship.layout, Ship.flip)
    parts[4] = newPart(CargoShip.layout, CargoShip.flip)
    parts[5] = newPart(Ship.layout, Ship.flip)
    
    b := newBoard(BOARDSIZE)
    
    b.board[0]  = []uint8{0,0,0,0,0,0,1,0,0,0,0,0,0,0,0}
    b.board[1]  = []uint8{0,0,0,0,0,0,0,1,0,0,0,0,0,0,0}
    b.board[2]  = []uint8{0,0,0,0,0,0,1,0,0,0,0,0,0,0,0}
    b.board[3]  = []uint8{0,0,0,0,0,0,0,0,5,5,0,0,0,0,0}
    b.board[4]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,2,0,0}
    b.board[5]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,2,0,0}
    b.board[6]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,2,0,0}
    b.board[7]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,2,0,0}
    b.board[8]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,2,0,0}
    b.board[9]  = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    b.board[10] = []uint8{0,0,3,3,0,0,0,0,0,0,0,0,0,0,0}
    b.board[11] = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    b.board[12] = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    b.board[13] = []uint8{0,0,0,0,0,0,0,0,0,4,4,4,0,0,0}
    b.board[14] = []uint8{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
    
    if !PERMITSIDEBYSIDE {
        fmt.Println(b.findSideBySide())
    }
    fmt.Println(b.validate(parts))
    
    game := &sGamePlay{}
    
    game.startGame(b, b)
    err, partId, finishPart, finish := game.shot(1, 3, 3)
    fmt.Println(err, partId, finishPart, finish)
    err, partId, finishPart, finish  = game.shot(1, 17, 3)
    fmt.Println(err, partId, finishPart, finish)
    
    // Mata um avião
    err, partId, finishPart, finish  = game.shot(2, 2, 6)
    fmt.Println(err, partId, finishPart, finish)
    err, partId, finishPart, finish  = game.shot(2, 0, 6)
    fmt.Println(err, partId, finishPart, finish)
    err, partId, finishPart, finish  = game.shot(2, 1, 7)
    fmt.Println(err, partId, finishPart, finish)
    
    err, partId, finishPart, finish  = game.shot(2, 3, 3)
    fmt.Println(err, partId, finishPart, finish)
    err, partId, finishPart, finish  = game.shot(1, 2, 6)
    fmt.Println(err, partId, finishPart, finish)
    err, partId, finishPart, finish  = game.shot(1, 3, 3)
    fmt.Println(err, partId, finishPart, finish)
    
    KillPlayer2(b, game)
}

func KillPlayer2(b *sBoard, game *sGamePlay) {
    // Mata o player2
    for l, _ := range b.board {
        for c, _ := range b.board[l] {
            err, partId, finishPart, finish := game.shot(2, uint8(l), uint8(c))
            fmt.Println(err, partId, finishPart, finish)
            if finish {
                return
            }
        }
    }
}

