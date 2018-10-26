package main

import "gomoku/gomoku"

func decisionCompleteOrBock(res searchResult) (trial uvector) {
	if res.resultType == searchResultHorizontal {
		if !res.blockedBefore {
			trial.x = res.x - 1
			trial.y = res.y
		} else {
			trial.x = res.x + res.size
			trial.y = res.y
		}
	} else if res.resultType == searchResultVectical {
		if !res.blockedBefore {
			trial.x = res.x
			trial.y = res.y - 1
		} else {
			trial.x = res.x
			trial.y = res.y + res.size
		}
	}
	return
}

func decisionTake(game *gomoku.Gomoku) uvector {
	resfoe := searchHorizontal(game.Board, gomoku.BoardCellFoe)
	tmp := searchVertical(game.Board, gomoku.BoardCellFoe)
	if tmp.size > resfoe.size {
		resfoe = tmp
	}
	resown := searchHorizontal(game.Board, gomoku.BoardCellOwn)
	tmp = searchVertical(game.Board, gomoku.BoardCellOwn)
	if tmp.size > resown.size {
		resown = tmp
	}

	if resfoe.size != 0 && resfoe.size >= resown.size {
		return decisionCompleteOrBock(resfoe)
	}
	if resown.size != 0 {
		return decisionCompleteOrBock(resown)
	}
	return uvector{0, 0}
}
