package structure

import (
	"context"

	"github.com/taubyte/go-interfaces/services/substrate"
	"github.com/taubyte/go-interfaces/services/tns"
	"github.com/taubyte/go-interfaces/vm"
	http "github.com/taubyte/http"
	httpMock "github.com/taubyte/http/mocks"
	"github.com/taubyte/odo/protocols/substrate/components/counters"
	"github.com/taubyte/p2p/peer"
)

var _ substrate.Service = &NodeService{}

type NodeService struct {
	node         peer.Node
	tns          tns.Client
	vm           vm.Service
	httpSrv      http.Service
	nodeSmartOps substrate.SmartOpsService
	nodeCounters substrate.CounterService

	branch string
	ctx    context.Context
}

func MockNodeService(node peer.Node, ctx context.Context) substrate.Service {
	s := &NodeService{
		node:         node,
		tns:          &TestClient{},
		vm:           &TestVm{},
		nodeSmartOps: &TestSmartOps{},
		branch:       "master",
		ctx:          ctx,
	}

	s.nodeCounters, _ = counters.New(s)

	s.httpSrv = httpMock.NewUnimplemented(ctx)

	return s
}

func (s *NodeService) Node() peer.Node {
	return s.node
}

func (s *NodeService) Close() error { return nil }

func (s *NodeService) Http() http.Service {
	return s.httpSrv
}

func (s *NodeService) Orbitals() []vm.Plugin {
	return nil
}

func (s *NodeService) Vm() vm.Service {
	return s.vm
}

func (s *NodeService) Tns() tns.Client {
	return s.tns
}

func (s *NodeService) SmartOps() substrate.SmartOpsService {
	return s.nodeSmartOps
}

func (s *NodeService) Branch() string {
	return s.branch
}

func (s *NodeService) Context() context.Context {
	return s.ctx
}

func (s *NodeService) Counter() substrate.CounterService {
	return s.nodeCounters
}

func (s *NodeService) Dev() bool {
	return true
}

func (s *NodeService) Verbose() bool {
	return false
}
