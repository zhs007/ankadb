package ankadb

import (
	"context"
)

var (
	// EventOnStarted - onStarted
	EventOnStarted = "onstarted"
)

// FuncAnkaDBEvent - func event
type FuncAnkaDBEvent func(ctx context.Context, anka AnkaDB) error

// eventMgr event manager
type eventMgr struct {
	mapAnkaDBEvent map[string]([]FuncAnkaDBEvent)
	anka           AnkaDB
}

func newEventMgr(anka AnkaDB) *eventMgr {
	mgr := &eventMgr{
		mapAnkaDBEvent: make(map[string]([]FuncAnkaDBEvent)),
		anka:           anka,
	}

	return mgr
}

func (mgr *eventMgr) checkAnkaDBEvent(event string) bool {
	return event == EventOnStarted
}

func (mgr *eventMgr) regAnkaDBEventFunc(event string, eventfunc FuncAnkaDBEvent) error {
	if !mgr.checkAnkaDBEvent(event) {
		return ErrInvalidEvent
	}

	mgr.mapAnkaDBEvent[event] = append(mgr.mapAnkaDBEvent[event], eventfunc)

	return nil
}

func (mgr *eventMgr) onAnkaDBEvent(ctx context.Context, event string) error {
	lst, ok := mgr.mapAnkaDBEvent[event]
	if !ok {
		return nil
	}

	for i := range lst {
		lst[i](ctx, mgr.anka)
	}

	return nil
}
