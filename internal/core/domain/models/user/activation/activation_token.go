package activation

import "github.com/google/uuid"

type Generator struct {
	genFunc GenFunc
}

func NewGenerator(gf GenFunc) *Generator {
	if gf == nil {
		gf = defaultGenFunc
	}
	return &Generator{
		genFunc: gf,
	}
}

func (g *Generator) Generate() (string, error) {
	return g.genFunc()
}

type GenFunc func() (string, error)

func defaultGenFunc() (string, error) {
	activationToken := uuid.New()
	return activationToken.String(), nil
}
