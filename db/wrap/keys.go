package wrap

var (
	wrapEventPrefix = []byte{0}

	// We hold here the account block height on which the orchestrator updated the blocks
	lastUpdatePrefix = []byte{1}
)
