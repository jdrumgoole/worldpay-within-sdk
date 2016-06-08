package fsm
import (
	"github.com/ryanfaerman/fsm"
)

type FSM struct {

	State fsm.State

	Machine *fsm.Machine
}

func Init(initialState fsm.State) (*FSM, error) {

	result := &FSM{}
	result.State = initialState

	result.Machine = &fsm.Machine{Subject: result}

	sdkRules := &fsm.Ruleset{}

	// Device setup
	sdkRules.AddTransition(fsm.T{ DEV_NOT_READY , DEV_READY })
	sdkRules.AddTransition(fsm.T{ DEV_READY , DEV_READY })

	// Producer setup
	sdkRules.AddTransition(fsm.T{ DEV_READY, PRO_READY })
	sdkRules.AddTransition(fsm.T{ PRO_READY, PRO_BROADCAST })
	sdkRules.AddTransition(fsm.T{ PRO_BROADCAST, PRO_READY })
	sdkRules.AddTransition(fsm.T{ PRO_READY, PRO_READY })

	// Consumer setup
	sdkRules.AddTransition(fsm.T{ DEV_READY, CON_DISCOVER_DEV })
	sdkRules.AddTransition(fsm.T{ CON_DISCOVER_DEV, DEV_READY})
	sdkRules.AddTransition(fsm.T{ CON_DISCOVER_DEV, CON_DEV_AVAILABLE})
	sdkRules.AddTransition(fsm.T{ CON_DEV_AVAILABLE, CON_DISCOVER_DEV })
	sdkRules.AddTransition(fsm.T{ CON_DEV_AVAILABLE, CON_READY })
	sdkRules.AddTransition(fsm.T{ CON_READY, CON_REQ_SVC })
	sdkRules.AddTransition(fsm.T{ CON_REQ_SVC, CON_READY })
	sdkRules.AddTransition(fsm.T{ CON_REQ_SVC, CON_SVC_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_DEV_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_SEL_SVC })
	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_SVC_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_SEL_SVC, CON_AWAIT_PAYMENT })
	sdkRules.AddTransition(fsm.T{ CON_SEL_SVC, CON_SVC_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_AWAIT_PAYMENT, CON_SVC_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_AWAIT_PAYMENT, CON_PAYMENT_ACCEPTED })
	sdkRules.AddTransition(fsm.T{ CON_PAYMENT_ACCEPTED, CON_SVC_AVAILABLE })
	sdkRules.AddTransition(fsm.T{ CON_PAYMENT_ACCEPTED, CON_DELIVERING_SERVICE })
	sdkRules.AddTransition(fsm.T{ CON_DELIVERING_SERVICE, CON_SVC_AVAILABLE })

	result.Machine.Rules = sdkRules

	return result, nil
}

// Add methods to comply with the fsm.Stater interface
func (t *FSM) CurrentState() fsm.State {

	return t.State
}
func (t *FSM) SetState(s fsm.State) {

	t.State = s
}

// A wrapper to the machines Transition method
func (t *FSM) Transition(goal fsm.State) error {

	return t.Machine.Transition(goal)
}