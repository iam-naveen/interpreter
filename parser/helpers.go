package parser

func (p *Parser) move() {
	for {
		select {
		case piece := <-p.channel:
			p.prev = p.piece
			p.piece = &piece
			// fmt.Println("moving -->", p)
			return
		}

	}
}
