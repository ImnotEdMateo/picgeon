package scanning

import "picgeon/utils"

type Scanner interface {
	Scan() ([]utils.Media, error)
}
