package pipe_filter

type Input interface{}
type Output interface{}
type filter interface {
	Progress(inputData Input) (outputData Output, errors error)
}
