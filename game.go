package main

const (
    Rows    = 6
    Columns = 7
)

type Game struct {
    Board  [Rows][Columns]int // 0 = vide, 1 = joueur1, 2 = joueur2
    Turn   int                // joueur courant
    Winner int                // 0 = pas fini, 1 ou 2 = gagnant, -1 = égalité
}

// Initialise une nouvelle partie
func NewGame() *Game {
    return &Game{
        Turn: 1,
    }
}

// Joue un coup dans une colonne (retourne vrai si réussi)
func (g *Game) Play(col int) bool {
    if col < 0 || col >= Columns || g.Winner != 0 {
        return false
    }

    for row := Rows - 1; row >= 0; row-- {
        if g.Board[row][col] == 0 {
            g.Board[row][col] = g.Turn
            if g.checkWin(row, col) {
                g.Winner = g.Turn
            } else if g.isDraw() {
                g.Winner = -1
            } else {
                g.Turn = 3 - g.Turn // alterne entre 1 et 2
            }
            return true
        }
    }
    return false
}

// Vérifie victoire
func (g *Game) checkWin(r, c int) bool {
    directions := [][2]int{
        {0, 1}, {1, 0}, {1, 1}, {1, -1},
    }
    for _, d := range directions {
        count := 1
        for _, sign := range []int{-1, 1} {
            dr, dc := d[0]*sign, d[1]*sign
            rr, cc := r+dr, c+dc
            for rr >= 0 && rr < Rows && cc >= 0 && cc < Columns && g.Board[rr][cc] == g.Turn {
                count++
                rr += dr
                cc += dc
            }
        }
        if count >= 4 {
            return true
        }
    }
    return false
}

// Vérifie égalité
func (g *Game) isDraw() bool {
    for c := 0; c < Columns; c++ {
        if g.Board[0][c] == 0 {
            return false
        }
    }
    return true
}
