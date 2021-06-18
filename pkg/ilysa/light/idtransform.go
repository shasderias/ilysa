package light

type IDTransformer func(ID) IDSet

func IDDivide(divisor int) IDTransformer {
	return nil
}

func IDDivideIntoGroupsOf(groupCount int) IDTransformer {
	return nil
}

func IDFan(groupCount int) IDTransformer {
	return nil
}
