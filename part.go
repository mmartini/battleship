package main

type sPart struct {
    layout [][]uint8
    flip bool
}

var (
    Airplane        = newPart([][]uint8{[]uint8{0,1,0}, 
                                        []uint8{1,0,1}},
                              true)
    AircraftCarrier = newPart([][]uint8{[]uint8{1,1,1,1,1}},
                              true)
    Ship            = newPart([][]uint8{[]uint8{1,1}},
                              true)
    CargoShip       = newPart([][]uint8{[]uint8{1,1,1}}, 
                              true)
)

func newPart(part [][]uint8, flip bool) (p *sPart) {
    p = &sPart{}
    p.layout = part
    p.flip = flip
    
    return
}

func (p *sPart) find (b *sBoard, partId uint8) (finded bool) {
    // Faz um board separado so com o partId solicitado
    bAux := b.copy()
    for l, _ := range bAux.board {
        for c, _ := range bAux.board[l] {
            if bAux.board[l][c] != partId {
                bAux.board[l][c] = 0
            }
        }
    }
    
    part := newPart(bAux.Trim().board, true)
    return p.comparePart(part)
}

func (p *sPart) comparePart(comparePart *sPart) bool {
    equalParts := true
    if len(comparePart.layout) == len(p.layout) && len(comparePart.layout[0])== len(p.layout[0]) {
        for l, _ := range comparePart.layout {
            for c, _ := range comparePart.layout[l] {
                flag := comparePart.layout[l][c]
                if flag > 0 {
                    flag = 1
                }
                if p.layout[l][c] != flag {
                    equalParts = false
                }
            }
        }
        
        // Tamanho é igual e o layout também retorna true
        if equalParts {
            return true
        }
    }
    
    // Pode girar
    if p.flip {
        numberFlips := 1
        if len(comparePart.layout) > 1 && len(comparePart.layout[0]) > 1 {
            numberFlips = 3
        }
    
        for flip := 0; flip < numberFlips; flip++ {
            comparePart.flipPart()

            equalParts = true
            if len(comparePart.layout) == len(p.layout) && len(comparePart.layout[0])== len(p.layout[0]) {
                for l, _ := range comparePart.layout {
                    for c, _ := range comparePart.layout[l] {
                        flag := comparePart.layout[l][c]
                        if flag > 0 {
                            flag = 1
                        }
                        if p.layout[l][c] != flag {
                            equalParts = false
                        }
                    }
                }

                // Tamanho é igual e o layout também retorna true
                if equalParts {
                    return true
                }
            }
        }
    }
    
    return false
}

func (p *sPart) flipPart() {
    newLayout := [][]uint8{}
    for l, _ := range p.layout {
        for c, _ := range p.layout[l] {
            if c > len(newLayout)-1 {
                newLayout = append(newLayout, []uint8{})
            }
            if l > len(newLayout[c])-1 {
                newLayout[c] = append(newLayout[c], 0)
            }
            newLayout[c][l] = p.layout[len(p.layout)-l-1][c]
        }
    }
    p.layout = newLayout
}

