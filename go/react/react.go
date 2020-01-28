package react

type MyReactor struct {
}

type MyCell struct {
	Value int
}

type MyComputeCell struct {

}

func (reactor MyReactor) CreateInput(value int) MyCell {
	return MyCell{Value: value}
}

func (reactor MyReactor) CreateCompute1(cell MyCell, func(int) int) MyComputeCell {

}

CreateCompute2(Cell, Cell, func(int, int) int) ComputeCell

func (cell MyCell) SetValue(value int) {
	cell.Value = value
}

func New() MyReactor {
	return MyReactor{}
}
