package newton

import (
	. "math"

	. "github.com/AndreyZWorkAccount/XIRR/numMethods"
)

type Method struct {
	initialGuess float64
}

func NewMethod(guess float64) Method{
	return Method{initialGuess:guess}
}


// NumericMethodUsingDerivative interface implementation
func (nm *Method) Calculate(F INumericFunc, derivativeF INumericFunc, methodParams *NumericMethodParams) (float64, NumericResultType, *NumericMethodError) {

	xCurrent := nm.initialGuess

	var iterationPassed  uint64 = 0
	for iterationPassed < methodParams.MaxIterationsCount {

		fValue := F.ApplyTo(xCurrent)
		fDerivativeValue := derivativeF.ApplyTo(xCurrent)

		xNext := xCurrent - fValue/fDerivativeValue

		dx := Abs(xNext-xCurrent)
		if dx <= methodParams.Epsilon {
			return SolutionFound(xNext)
		}
		xCurrent = xNext

		iterationPassed++
	}

	return NoSolutionFound()
}