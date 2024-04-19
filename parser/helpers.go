package parser

import "fmt"

func (p *Parser) move() {
	for {
		select {
		case piece := <-p.channel:
			p.prev = p.piece
			p.piece = &piece
			if p.logEnabled {
				fmt.Println("moving", p.prev, p.piece)
			}
			return
		}

	}
}
