// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadAgentLagacyKernel310 returns the embedded CollectionSpec for AgentLagacyKernel310.
func LoadAgentLagacyKernel310() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_AgentLagacyKernel310Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load AgentLagacyKernel310: %w", err)
	}

	return spec, err
}

// LoadAgentLagacyKernel310Objects loads AgentLagacyKernel310 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*AgentLagacyKernel310Objects
//	*AgentLagacyKernel310Programs
//	*AgentLagacyKernel310Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadAgentLagacyKernel310Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadAgentLagacyKernel310()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// AgentLagacyKernel310Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentLagacyKernel310Specs struct {
	AgentLagacyKernel310ProgramSpecs
	AgentLagacyKernel310MapSpecs
}

// AgentLagacyKernel310Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentLagacyKernel310ProgramSpecs struct {
	DevHardStartXmit                     *ebpf.ProgramSpec `ebpf:"dev_hard_start_xmit"`
	DevQueueXmit                         *ebpf.ProgramSpec `ebpf:"dev_queue_xmit"`
	IpQueueXmit                          *ebpf.ProgramSpec `ebpf:"ip_queue_xmit"`
	IpQueueXmit2                         *ebpf.ProgramSpec `ebpf:"ip_queue_xmit2"`
	IpRcvCore                            *ebpf.ProgramSpec `ebpf:"ip_rcv_core"`
	KprobeNfNatManipPkt                  *ebpf.ProgramSpec `ebpf:"kprobe__nf_nat_manip_pkt"`
	KprobeNfNatPacket                    *ebpf.ProgramSpec `ebpf:"kprobe__nf_nat_packet"`
	SecuritySocketRecvmsgEnter           *ebpf.ProgramSpec `ebpf:"security_socket_recvmsg_enter"`
	SecuritySocketSendmsgEnter           *ebpf.ProgramSpec `ebpf:"security_socket_sendmsg_enter"`
	SkbCopyDatagramIovec                 *ebpf.ProgramSpec `ebpf:"skb_copy_datagram_iovec"`
	SkbCopyDatagramIter                  *ebpf.ProgramSpec `ebpf:"skb_copy_datagram_iter"`
	SockAllocRet                         *ebpf.ProgramSpec `ebpf:"sock_alloc_ret"`
	TcpDestroySock                       *ebpf.ProgramSpec `ebpf:"tcp_destroy_sock"`
	TcpQueueRcv                          *ebpf.ProgramSpec `ebpf:"tcp_queue_rcv"`
	TcpRcvEstablished                    *ebpf.ProgramSpec `ebpf:"tcp_rcv_established"`
	TcpV4DoRcv                           *ebpf.ProgramSpec `ebpf:"tcp_v4_do_rcv"`
	TcpV4Rcv                             *ebpf.ProgramSpec `ebpf:"tcp_v4_rcv"`
	TracepointNetifReceiveSkb            *ebpf.ProgramSpec `ebpf:"tracepoint__netif_receive_skb"`
	TracepointSchedSchedProcessExec      *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_exec"`
	TracepointSchedSchedProcessExit      *ebpf.ProgramSpec `ebpf:"tracepoint__sched__sched_process_exit"`
	TracepointSkbCopyDatagramIovec       *ebpf.ProgramSpec `ebpf:"tracepoint__skb_copy_datagram_iovec"`
	TracepointSyscallsSysEnterAccept4    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_accept4"`
	TracepointSyscallsSysEnterClose      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_close"`
	TracepointSyscallsSysEnterConnect    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_connect"`
	TracepointSyscallsSysEnterRead       *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_read"`
	TracepointSyscallsSysEnterReadv      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_readv"`
	TracepointSyscallsSysEnterRecvfrom   *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_recvfrom"`
	TracepointSyscallsSysEnterRecvmsg    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_recvmsg"`
	TracepointSyscallsSysEnterSendfile64 *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_sendfile64"`
	TracepointSyscallsSysEnterSendmsg    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_sendmsg"`
	TracepointSyscallsSysEnterSendto     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_sendto"`
	TracepointSyscallsSysEnterWrite      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_write"`
	TracepointSyscallsSysEnterWritev     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_enter_writev"`
	TracepointSyscallsSysExitAccept4     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_accept4"`
	TracepointSyscallsSysExitClose       *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_close"`
	TracepointSyscallsSysExitConnect     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_connect"`
	TracepointSyscallsSysExitRead        *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_read"`
	TracepointSyscallsSysExitReadv       *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_readv"`
	TracepointSyscallsSysExitRecvfrom    *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_recvfrom"`
	TracepointSyscallsSysExitRecvmsg     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_recvmsg"`
	TracepointSyscallsSysExitSendfile64  *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_sendfile64"`
	TracepointSyscallsSysExitSendmsg     *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_sendmsg"`
	TracepointSyscallsSysExitSendto      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_sendto"`
	TracepointSyscallsSysExitWrite       *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_write"`
	TracepointSyscallsSysExitWritev      *ebpf.ProgramSpec `ebpf:"tracepoint__syscalls__sys_exit_writev"`
	XdpProxy                             *ebpf.ProgramSpec `ebpf:"xdp_proxy"`
}

// AgentLagacyKernel310MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type AgentLagacyKernel310MapSpecs struct {
	AcceptArgsMap         *ebpf.MapSpec `ebpf:"accept_args_map"`
	ActiveSendfileArgsMap *ebpf.MapSpec `ebpf:"active_sendfile_args_map"`
	ActiveSslReadArgsMap  *ebpf.MapSpec `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.MapSpec `ebpf:"active_ssl_write_args_map"`
	CloseArgsMap          *ebpf.MapSpec `ebpf:"close_args_map"`
	ConnEvtRb             *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.MapSpec `ebpf:"conn_info_map"`
	ConnInfoT_map         *ebpf.MapSpec `ebpf:"conn_info_t_map"`
	ConnectArgsMap        *ebpf.MapSpec `ebpf:"connect_args_map"`
	ControlValues         *ebpf.MapSpec `ebpf:"control_values"`
	EnabledLocalIpv4Map   *ebpf.MapSpec `ebpf:"enabled_local_ipv4_map"`
	EnabledLocalPortMap   *ebpf.MapSpec `ebpf:"enabled_local_port_map"`
	EnabledRemoteIpMap    *ebpf.MapSpec `ebpf:"enabled_remote_ip_map"`
	EnabledRemotePortMap  *ebpf.MapSpec `ebpf:"enabled_remote_port_map"`
	FilterMntnsMap        *ebpf.MapSpec `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.MapSpec `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.MapSpec `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.MapSpec `ebpf:"filter_pidns_map"`
	GoCommonSymaddrsMap   *ebpf.MapSpec `ebpf:"go_common_symaddrs_map"`
	GoSslUserSpaceCallMap *ebpf.MapSpec `ebpf:"go_ssl_user_space_call_map"`
	KernEvtT_map          *ebpf.MapSpec `ebpf:"kern_evt_t_map"`
	NatFlowMap            *ebpf.MapSpec `ebpf:"nat_flow_map"`
	ProcExecEvents        *ebpf.MapSpec `ebpf:"proc_exec_events"`
	ProcExitEvents        *ebpf.MapSpec `ebpf:"proc_exit_events"`
	Rb                    *ebpf.MapSpec `ebpf:"rb"`
	ReadArgsMap           *ebpf.MapSpec `ebpf:"read_args_map"`
	SockKeyConnIdMap      *ebpf.MapSpec `ebpf:"sock_key_conn_id_map"`
	SockXmitMap           *ebpf.MapSpec `ebpf:"sock_xmit_map"`
	SslDataMap            *ebpf.MapSpec `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.MapSpec `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.MapSpec `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.MapSpec `ebpf:"syscall_rb"`
	WriteArgsMap          *ebpf.MapSpec `ebpf:"write_args_map"`
}

// AgentLagacyKernel310Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadAgentLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type AgentLagacyKernel310Objects struct {
	AgentLagacyKernel310Programs
	AgentLagacyKernel310Maps
}

func (o *AgentLagacyKernel310Objects) Close() error {
	return _AgentLagacyKernel310Close(
		&o.AgentLagacyKernel310Programs,
		&o.AgentLagacyKernel310Maps,
	)
}

// AgentLagacyKernel310Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadAgentLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type AgentLagacyKernel310Maps struct {
	AcceptArgsMap         *ebpf.Map `ebpf:"accept_args_map"`
	ActiveSendfileArgsMap *ebpf.Map `ebpf:"active_sendfile_args_map"`
	ActiveSslReadArgsMap  *ebpf.Map `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.Map `ebpf:"active_ssl_write_args_map"`
	CloseArgsMap          *ebpf.Map `ebpf:"close_args_map"`
	ConnEvtRb             *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.Map `ebpf:"conn_info_map"`
	ConnInfoT_map         *ebpf.Map `ebpf:"conn_info_t_map"`
	ConnectArgsMap        *ebpf.Map `ebpf:"connect_args_map"`
	ControlValues         *ebpf.Map `ebpf:"control_values"`
	EnabledLocalIpv4Map   *ebpf.Map `ebpf:"enabled_local_ipv4_map"`
	EnabledLocalPortMap   *ebpf.Map `ebpf:"enabled_local_port_map"`
	EnabledRemoteIpMap    *ebpf.Map `ebpf:"enabled_remote_ip_map"`
	EnabledRemotePortMap  *ebpf.Map `ebpf:"enabled_remote_port_map"`
	FilterMntnsMap        *ebpf.Map `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.Map `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.Map `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.Map `ebpf:"filter_pidns_map"`
	GoCommonSymaddrsMap   *ebpf.Map `ebpf:"go_common_symaddrs_map"`
	GoSslUserSpaceCallMap *ebpf.Map `ebpf:"go_ssl_user_space_call_map"`
	KernEvtT_map          *ebpf.Map `ebpf:"kern_evt_t_map"`
	NatFlowMap            *ebpf.Map `ebpf:"nat_flow_map"`
	ProcExecEvents        *ebpf.Map `ebpf:"proc_exec_events"`
	ProcExitEvents        *ebpf.Map `ebpf:"proc_exit_events"`
	Rb                    *ebpf.Map `ebpf:"rb"`
	ReadArgsMap           *ebpf.Map `ebpf:"read_args_map"`
	SockKeyConnIdMap      *ebpf.Map `ebpf:"sock_key_conn_id_map"`
	SockXmitMap           *ebpf.Map `ebpf:"sock_xmit_map"`
	SslDataMap            *ebpf.Map `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.Map `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.Map `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.Map `ebpf:"syscall_rb"`
	WriteArgsMap          *ebpf.Map `ebpf:"write_args_map"`
}

func (m *AgentLagacyKernel310Maps) Close() error {
	return _AgentLagacyKernel310Close(
		m.AcceptArgsMap,
		m.ActiveSendfileArgsMap,
		m.ActiveSslReadArgsMap,
		m.ActiveSslWriteArgsMap,
		m.CloseArgsMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.ConnInfoT_map,
		m.ConnectArgsMap,
		m.ControlValues,
		m.EnabledLocalIpv4Map,
		m.EnabledLocalPortMap,
		m.EnabledRemoteIpMap,
		m.EnabledRemotePortMap,
		m.FilterMntnsMap,
		m.FilterNetnsMap,
		m.FilterPidMap,
		m.FilterPidnsMap,
		m.GoCommonSymaddrsMap,
		m.GoSslUserSpaceCallMap,
		m.KernEvtT_map,
		m.NatFlowMap,
		m.ProcExecEvents,
		m.ProcExitEvents,
		m.Rb,
		m.ReadArgsMap,
		m.SockKeyConnIdMap,
		m.SockXmitMap,
		m.SslDataMap,
		m.SslRb,
		m.SslUserSpaceCallMap,
		m.SyscallDataMap,
		m.SyscallRb,
		m.WriteArgsMap,
	)
}

// AgentLagacyKernel310Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadAgentLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type AgentLagacyKernel310Programs struct {
	DevHardStartXmit                     *ebpf.Program `ebpf:"dev_hard_start_xmit"`
	DevQueueXmit                         *ebpf.Program `ebpf:"dev_queue_xmit"`
	IpQueueXmit                          *ebpf.Program `ebpf:"ip_queue_xmit"`
	IpQueueXmit2                         *ebpf.Program `ebpf:"ip_queue_xmit2"`
	IpRcvCore                            *ebpf.Program `ebpf:"ip_rcv_core"`
	KprobeNfNatManipPkt                  *ebpf.Program `ebpf:"kprobe__nf_nat_manip_pkt"`
	KprobeNfNatPacket                    *ebpf.Program `ebpf:"kprobe__nf_nat_packet"`
	SecuritySocketRecvmsgEnter           *ebpf.Program `ebpf:"security_socket_recvmsg_enter"`
	SecuritySocketSendmsgEnter           *ebpf.Program `ebpf:"security_socket_sendmsg_enter"`
	SkbCopyDatagramIovec                 *ebpf.Program `ebpf:"skb_copy_datagram_iovec"`
	SkbCopyDatagramIter                  *ebpf.Program `ebpf:"skb_copy_datagram_iter"`
	SockAllocRet                         *ebpf.Program `ebpf:"sock_alloc_ret"`
	TcpDestroySock                       *ebpf.Program `ebpf:"tcp_destroy_sock"`
	TcpQueueRcv                          *ebpf.Program `ebpf:"tcp_queue_rcv"`
	TcpRcvEstablished                    *ebpf.Program `ebpf:"tcp_rcv_established"`
	TcpV4DoRcv                           *ebpf.Program `ebpf:"tcp_v4_do_rcv"`
	TcpV4Rcv                             *ebpf.Program `ebpf:"tcp_v4_rcv"`
	TracepointNetifReceiveSkb            *ebpf.Program `ebpf:"tracepoint__netif_receive_skb"`
	TracepointSchedSchedProcessExec      *ebpf.Program `ebpf:"tracepoint__sched__sched_process_exec"`
	TracepointSchedSchedProcessExit      *ebpf.Program `ebpf:"tracepoint__sched__sched_process_exit"`
	TracepointSkbCopyDatagramIovec       *ebpf.Program `ebpf:"tracepoint__skb_copy_datagram_iovec"`
	TracepointSyscallsSysEnterAccept4    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_accept4"`
	TracepointSyscallsSysEnterClose      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_close"`
	TracepointSyscallsSysEnterConnect    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_connect"`
	TracepointSyscallsSysEnterRead       *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_read"`
	TracepointSyscallsSysEnterReadv      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_readv"`
	TracepointSyscallsSysEnterRecvfrom   *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_recvfrom"`
	TracepointSyscallsSysEnterRecvmsg    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_recvmsg"`
	TracepointSyscallsSysEnterSendfile64 *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_sendfile64"`
	TracepointSyscallsSysEnterSendmsg    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_sendmsg"`
	TracepointSyscallsSysEnterSendto     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_sendto"`
	TracepointSyscallsSysEnterWrite      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_write"`
	TracepointSyscallsSysEnterWritev     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_enter_writev"`
	TracepointSyscallsSysExitAccept4     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_accept4"`
	TracepointSyscallsSysExitClose       *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_close"`
	TracepointSyscallsSysExitConnect     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_connect"`
	TracepointSyscallsSysExitRead        *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_read"`
	TracepointSyscallsSysExitReadv       *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_readv"`
	TracepointSyscallsSysExitRecvfrom    *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_recvfrom"`
	TracepointSyscallsSysExitRecvmsg     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_recvmsg"`
	TracepointSyscallsSysExitSendfile64  *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_sendfile64"`
	TracepointSyscallsSysExitSendmsg     *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_sendmsg"`
	TracepointSyscallsSysExitSendto      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_sendto"`
	TracepointSyscallsSysExitWrite       *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_write"`
	TracepointSyscallsSysExitWritev      *ebpf.Program `ebpf:"tracepoint__syscalls__sys_exit_writev"`
	XdpProxy                             *ebpf.Program `ebpf:"xdp_proxy"`
}

func (p *AgentLagacyKernel310Programs) Close() error {
	return _AgentLagacyKernel310Close(
		p.DevHardStartXmit,
		p.DevQueueXmit,
		p.IpQueueXmit,
		p.IpQueueXmit2,
		p.IpRcvCore,
		p.KprobeNfNatManipPkt,
		p.KprobeNfNatPacket,
		p.SecuritySocketRecvmsgEnter,
		p.SecuritySocketSendmsgEnter,
		p.SkbCopyDatagramIovec,
		p.SkbCopyDatagramIter,
		p.SockAllocRet,
		p.TcpDestroySock,
		p.TcpQueueRcv,
		p.TcpRcvEstablished,
		p.TcpV4DoRcv,
		p.TcpV4Rcv,
		p.TracepointNetifReceiveSkb,
		p.TracepointSchedSchedProcessExec,
		p.TracepointSchedSchedProcessExit,
		p.TracepointSkbCopyDatagramIovec,
		p.TracepointSyscallsSysEnterAccept4,
		p.TracepointSyscallsSysEnterClose,
		p.TracepointSyscallsSysEnterConnect,
		p.TracepointSyscallsSysEnterRead,
		p.TracepointSyscallsSysEnterReadv,
		p.TracepointSyscallsSysEnterRecvfrom,
		p.TracepointSyscallsSysEnterRecvmsg,
		p.TracepointSyscallsSysEnterSendfile64,
		p.TracepointSyscallsSysEnterSendmsg,
		p.TracepointSyscallsSysEnterSendto,
		p.TracepointSyscallsSysEnterWrite,
		p.TracepointSyscallsSysEnterWritev,
		p.TracepointSyscallsSysExitAccept4,
		p.TracepointSyscallsSysExitClose,
		p.TracepointSyscallsSysExitConnect,
		p.TracepointSyscallsSysExitRead,
		p.TracepointSyscallsSysExitReadv,
		p.TracepointSyscallsSysExitRecvfrom,
		p.TracepointSyscallsSysExitRecvmsg,
		p.TracepointSyscallsSysExitSendfile64,
		p.TracepointSyscallsSysExitSendmsg,
		p.TracepointSyscallsSysExitSendto,
		p.TracepointSyscallsSysExitWrite,
		p.TracepointSyscallsSysExitWritev,
		p.XdpProxy,
	)
}

func _AgentLagacyKernel310Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed agentlagacykernel310_x86_bpfel.o
var _AgentLagacyKernel310Bytes []byte
