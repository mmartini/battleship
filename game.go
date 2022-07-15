package main

import (
    "errors"
)

type sGamePlay struct {
    playerId1   uint64
    board1      *sBoard
    shots1      *sBoard
    playerId2   uint64
    board2      *sBoard
    shots2      *sBoard
}

func (g *sGamePlay) startGame(board1, board2 *sBoard) {
    g.board1 = board1
    g.board2 = board2
    g.shots1 = newBoard(BOARDSIZE)
    g.shots2 = newBoard(BOARDSIZE)
}

// Retorna erro na jogada, numero da peça que acertou, se afundou, e se terminou o jogo
func (g *sGamePlay) shot(player, line, column uint8) (error, uint8, bool, bool) {
    var board [][]uint8
    var shots [][]uint8
    switch player {
    case 1:
        board = g.board1.board
        shots = g.shots1.board
    case 2:
        board = g.board2.board
        shots = g.shots2.board
    }
    
    if line < uint8(len(shots)) {
        if column < uint8(len(shots[line])) {
            if shots[line][column] == 1 {
                return errors.New("duplicated shot"), 0, false, false
            }
            shots[line][column] = 1
            switch player {
            case 1:
                g.shots1.board = shots
            case 2:
                g.shots2.board = shots
            }

            if board[line][column] == 0 {
                return nil, 0, false, false
            }
            partId := board[line][column]
            rp := g.restParts(board, shots)
            _, dontExterminator := rp[partId]
            finishGame := (len(rp) == 0)
            return nil, partId, !dontExterminator, finishGame
        }
    }
    
    return errors.New("out of board"), 0, false, false
}

func (g *sGamePlay) restParts (board, shots [][]uint8) (map[uint8]*uint8) {
    ret := make(map[uint8]*uint8)
    
    // contando o que resta
    for l, _ := range board {
        for c, _ := range board[l] {
            partId := board[l][c]
            shot := shots[l][c]
            // Tem barco é já não é um tiro
            if partId != 0 && shot == 0 {
                _, ok := ret[partId]
                if !ok {
                    total := uint8(0)
                    ret[partId] = &total
                }
                *ret[partId]++
            }
        }
    }
    
    return ret
}
