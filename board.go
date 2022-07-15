package main

import (
    "fmt"
)

type sBoard struct {
    board [][]uint8
    dimension uint8
}

func newBoard(dimension uint8) (b *sBoard) {
    b = &sBoard{}
    b.dimension = dimension
    for l := uint8(0); l < dimension; l++ {
        b.board = append(b.board, []uint8{})
        for c := uint8(0); c < dimension; c++ {
            b.board[l] = append(b.board[l], 0)
        }
    }
    return
}

func (b *sBoard) validate (parts map[uint8]*sPart) bool {
    // Valida o tamanho
    if b.dimension != uint8(len(b.board)) {
        return false
    }
    for l, _ := range b.board {
        if b.dimension != uint8(len(b.board[l])) {
            return false
        }
    }

    passParts := 0
    countParts := 0
    for partId, partInfo := range parts {
        countParts++
        if partInfo.find(b, partId) {
            fmt.Println("Validou partId:", partId)
            passParts++
        } else {
            fmt.Println("Não validou partId:", partId)
        }
    }
    
    // Procurar por números todos os números citados
    findedNumbers := make(map[uint8]bool)
    for l, _ := range b.board {
        for c, _ := range b.board[l] {
            if b.board[l][c] != 0 {
                findedNumbers[b.board[l][c]] = true
            }
        }
    }
    
    for number, _ := range findedNumbers {
        if _, ok := parts[number]; !ok {
            return false
        }
    }

    if countParts == passParts {
        return true
    }
    
    return false
}

func (b *sBoard) findSideBySide() bool {
    for l, _ := range b.board {
        limitTop := 0
        limitBottom := l
        if l > 0 {
            limitTop = l-1
        }
        if l < len(b.board)-1 {
            limitBottom = l+1
        }
        for c, _ := range b.board[l] {
            limitLeft := 0
            limitRight := c
            if c > 0 {
                limitLeft = c-1
            }
            if c < len(b.board[l])-1 {
                limitRight = c+1
            }

            partUnit := b.board[l][c]
            if partUnit != 0 {
                for ll := limitTop; ll <= limitBottom; ll++ {
                    for cc := limitLeft; cc <= limitRight; cc++ {
                        if b.board[ll][cc] != 0 && b.board[ll][cc] != partUnit {
                            return false
                        }
                    }
                }
            }
        }
    }
    
    return true
}

func (b *sBoard) Trim () (bRet *sBoard) {
    bRet = b.copy()
    
    // Trim de linhas do topo
    lineTrim := 0
    for l, _ := range bRet.board {
        if bRet.verifyClearLine(l) {
            lineTrim++
            continue
        }
        
        break
    }

    for l := 0; l < lineTrim; l++ {
        bRet.removeLine(0)
    }    
    
    // Trim de linhas de baixo
    lineTrim = 0
    for l, _ := range bRet.board {
        if bRet.verifyClearLine(len(bRet.board)-1-l) {
            lineTrim++
            continue
        }
        
        break
    }

    for l := 0; l < lineTrim; l++ {
        bRet.removeLine(len(bRet.board)-1)
    }
    
    if len(bRet.board) == 0 {
        return
    }
    
    // Trim de colunas a esquerda
    columnTrim := 0
    for c, _ := range bRet.board[0] {
        if bRet.verifyClearColumn(c) {
            columnTrim++
            continue
        }
        
        break
    }

    for c := 0; c < columnTrim; c++ {
        bRet.removeColumn(0)
    }
    
    if len(bRet.board) == 0 {
        return
    }    

    // Trim de colunas a direita
    columnTrim = 0
    for c, _ := range bRet.board[0] {
        if bRet.verifyClearColumn(len(bRet.board[0])-1-c) {
            columnTrim++
            continue
        }
        
        break
    }

    for c := 0; c < columnTrim; c++ {
        bRet.removeColumn(len(bRet.board[0])-1)
    }
    
    return
}

func (b *sBoard) verifyClearLine(line int) bool {
    findedPart := false
    for c, _ := range b.board[line] {
        if b.board[line][c] != 0 {
            findedPart = true
        }
    }
        
    return (!findedPart)
}

func (b *sBoard) verifyClearColumn(column int) bool {
    findedPart := false
    for l, _ := range b.board {
        if b.board[l][column] != 0 {
            findedPart = true
        }
    }

    return (!findedPart)
}

func (b *sBoard) removeLine(line int) {
    ret := [][]uint8{}
    for l, _ := range b.board {
        if l == line {
            continue
        }
        ret = append(ret, b.board[l])
    }
    
    b.board = ret
}

func (b *sBoard) removeColumn(column int) {
    ret := [][]uint8{}
    for l, _ := range b.board {
        ret = append(ret, append(b.board[l][0:column], b.board[l][column+1:]...))
    }
    
    b.board = ret
}

func (b *sBoard) copy() (ret *sBoard) {
    ret = newBoard(b.dimension)
    for l, _ := range b.board {
        ret.board = append(ret.board, []uint8{})
        for c, _ := range b.board[l] {
            ret.board[l] = append(ret.board[l], b.board[l][c])
        }
    }
    return
}