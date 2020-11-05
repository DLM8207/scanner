package holesources

import (
	"scanner/bean"
)

/*
 *bug源
 * author:leon
 * date:2020-11-5
 */
type holesources interface {
	holdsList() []bean.Bug
}
