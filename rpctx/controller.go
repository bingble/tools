package main

type dispatchType int

const (
	s2sType dispatchType = 1 << iota
	s2mType
	m2sType
	n2mType		// not to realise at current time
)

func dispatch() {
	spendableCount := len(input)
	if spendableCount == 0 {
		log.Error("There is no spendable transaction.")
	}

	// whether to create transaction recursively
	recursionConf := conf.DefaultBool("recursion", DefaultRecursion)

	if getDispatchType(m2sType) {
		for  !isEmpty() {
			if len(input) < InputLimit {
				break
			}
			m2sTx(recursionConf)
		}
	}

	if getDispatchType(s2mType) {
		for !isEmpty() {
			s2mTx(recursionConf)
		}
	}

	if getDispatchType(s2sType) {
		for !isEmpty() {
			s2sTx(recursionConf)
		}
	}

	// output tip message
	log.Info("Create Transactions Complete!\n")
}

func isEmpty() bool {
	// stop if no input data
	return len(input) == 0
}

func getDispatchType(t dispatchType) bool {
	dispatch := must(conf.Int("dispatch::type"))
	return dispatchType(dispatch.(int))&t == t
}
