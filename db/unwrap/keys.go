package unwrap

var (
	// events that do not have a send block associated
	unwrapRequestPrefix = []byte{0}

	// We hold here the evm chain block height on which the orchestrator updated the blocks
	lastUpdatePrefix = []byte{1}
)
