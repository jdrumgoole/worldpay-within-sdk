package fsm
import (
	"errors"
)

type State string
type Context string

type WPFSM interface {

	Permitted(input Context) bool
	Transition(input Context) error
}

//type Transition struct {
//
//	Input Context
//	Origin State
//	Exit State
//}

type FSM struct {

	State State
	rules map[State]map[Input]State
}

func Init(initialState State) (WPFSM, error) {

	result := &FSM{}
	result.State = initialState
	result.rules = make(map[State]map[Input]State, 0)

	// Device setup

	result.rules[DEV_NOT_READY] = make(map[Input]State, 0)
	result.rules[DEV_NOT_READY][CTX_INIT_DEVICE] = DEV_READY //Transition{ Input: CTX_INIT_DEVICE, Origin: DEV_NOT_READY, Exit: DEV_READY}

	// Producer setup
	result.rules[DEV_READY][] = DEV_READY
	result.rules[PRO_READY][] =
	result.rules[PRO_READY][] =
	result.rules[PRO_BROADCAST][] =

	sdkRules.AddTransition(fsm.T{ DEV_READY, PRO_READY })
	sdkRules.AddTransition(fsm.T{ PRO_READY, PRO_BROADCAST })
	sdkRules.AddTransition(fsm.T{ PRO_BROADCAST, PRO_READY })
	sdkRules.AddTransition(fsm.T{ PRO_READY, PRO_READY })
//
//	// Consumer setup
//	sdkRules.AddTransition(fsm.T{ DEV_READY, CON_DISCOVER_DEV })
//	sdkRules.AddTransition(fsm.T{ CON_DISCOVER_DEV, DEV_READY})
//	sdkRules.AddTransition(fsm.T{ CON_DISCOVER_DEV, CON_DEV_AVAILABLE})
//	sdkRules.AddTransition(fsm.T{ CON_DEV_AVAILABLE, CON_DISCOVER_DEV })
//	sdkRules.AddTransition(fsm.T{ CON_DEV_AVAILABLE, CON_READY })
//	sdkRules.AddTransition(fsm.T{ CON_READY, CON_REQ_SVC })
//	sdkRules.AddTransition(fsm.T{ CON_REQ_SVC, CON_READY })
//	sdkRules.AddTransition(fsm.T{ CON_REQ_SVC, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_DEV_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_SEL_SVC })
//	sdkRules.AddTransition(fsm.T{ CON_SVC_AVAILABLE, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_SEL_SVC, CON_AWAIT_PAYMENT })
//	sdkRules.AddTransition(fsm.T{ CON_SEL_SVC, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_AWAIT_PAYMENT, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_AWAIT_PAYMENT, CON_PROC_PAYMENT })
//	sdkRules.AddTransition(fsm.T{ CON_PROC_PAYMENT, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_PROC_PAYMENT, CON_PAYMENT_ACCEPTED })
//	sdkRules.AddTransition(fsm.T{ CON_PAYMENT_ACCEPTED, CON_SVC_AVAILABLE })
//	sdkRules.AddTransition(fsm.T{ CON_PAYMENT_ACCEPTED, CON_DELIVERING_SERVICE })
//	sdkRules.AddTransition(fsm.T{ CON_DELIVERING_SERVICE, CON_SVC_AVAILABLE })

	return result, nil
}

//// Add methods to comply with the fsm.Stater interface
//func (t *FSM) CurrentState() fsm.State {
//
//	return t.State
//}
//func (t *FSM) SetState(s fsm.State) {
//
//	t.State = s
//}

func (t *FSM) Permitted(input Context) bool {

	return false
}

func (t *FSM) Transition(input Context) error {

	return errors.New("Not implemented..")
}