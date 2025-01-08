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

type GoTlsLagacyKernel310GoCommonSymaddrsT struct {
	InternalSyscallConn   int64
	TlsConn               int64
	NetTCPConn            int64
	FD_SysfdOffset        int32
	TlsConnConnOffset     int32
	SyscallConnConnOffset int32
	G_goidOffset          int32
	G_addrOffset          int32
	_                     [4]byte
}

type GoTlsLagacyKernel310GoTlsSymaddrsT struct {
	WriteC_loc      GoTlsLagacyKernel310LocationT
	WriteB_loc      GoTlsLagacyKernel310LocationT
	WriteRetval0Loc GoTlsLagacyKernel310LocationT
	WriteRetval1Loc GoTlsLagacyKernel310LocationT
	ReadC_loc       GoTlsLagacyKernel310LocationT
	ReadB_loc       GoTlsLagacyKernel310LocationT
	ReadRetval0Loc  GoTlsLagacyKernel310LocationT
	ReadRetval1Loc  GoTlsLagacyKernel310LocationT
}

type GoTlsLagacyKernel310LocationT struct {
	Type   GoTlsLagacyKernel310LocationTypeT
	Offset int32
}

type GoTlsLagacyKernel310LocationTypeT uint32

const (
	GoTlsLagacyKernel310LocationTypeTKLocationTypeInvalid   GoTlsLagacyKernel310LocationTypeT = 0
	GoTlsLagacyKernel310LocationTypeTKLocationTypeStack     GoTlsLagacyKernel310LocationTypeT = 1
	GoTlsLagacyKernel310LocationTypeTKLocationTypeRegisters GoTlsLagacyKernel310LocationTypeT = 2
)

// LoadGoTlsLagacyKernel310 returns the embedded CollectionSpec for GoTlsLagacyKernel310.
func LoadGoTlsLagacyKernel310() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_GoTlsLagacyKernel310Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load GoTlsLagacyKernel310: %w", err)
	}

	return spec, err
}

// LoadGoTlsLagacyKernel310Objects loads GoTlsLagacyKernel310 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*GoTlsLagacyKernel310Objects
//	*GoTlsLagacyKernel310Programs
//	*GoTlsLagacyKernel310Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadGoTlsLagacyKernel310Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadGoTlsLagacyKernel310()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// GoTlsLagacyKernel310Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type GoTlsLagacyKernel310Specs struct {
	GoTlsLagacyKernel310ProgramSpecs
	GoTlsLagacyKernel310MapSpecs
}

// GoTlsLagacyKernel310Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type GoTlsLagacyKernel310ProgramSpecs struct {
	ProbeEntryTlsConnRead   *ebpf.ProgramSpec `ebpf:"probe_entry_tls_conn_read"`
	ProbeEntryTlsConnWrite  *ebpf.ProgramSpec `ebpf:"probe_entry_tls_conn_write"`
	ProbeReturnTlsConnRead  *ebpf.ProgramSpec `ebpf:"probe_return_tls_conn_read"`
	ProbeReturnTlsConnWrite *ebpf.ProgramSpec `ebpf:"probe_return_tls_conn_write"`
}

// GoTlsLagacyKernel310MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type GoTlsLagacyKernel310MapSpecs struct {
	ActiveSslReadArgsMap  *ebpf.MapSpec `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.MapSpec `ebpf:"active_ssl_write_args_map"`
	ActiveTlsConnOpMap    *ebpf.MapSpec `ebpf:"active_tls_conn_op_map"`
	ConnEvtRb             *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.MapSpec `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.MapSpec `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.MapSpec `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.MapSpec `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.MapSpec `ebpf:"filter_pidns_map"`
	FirstPacketEvtMap     *ebpf.MapSpec `ebpf:"first_packet_evt_map"`
	FirstPacketRb         *ebpf.MapSpec `ebpf:"first_packet_rb"`
	GoCommonSymaddrsMap   *ebpf.MapSpec `ebpf:"go_common_symaddrs_map"`
	GoSslUserSpaceCallMap *ebpf.MapSpec `ebpf:"go_ssl_user_space_call_map"`
	GoTlsSymaddrsMap      *ebpf.MapSpec `ebpf:"go_tls_symaddrs_map"`
	Rb                    *ebpf.MapSpec `ebpf:"rb"`
	RegsHeap              *ebpf.MapSpec `ebpf:"regs_heap"`
	SslDataMap            *ebpf.MapSpec `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.MapSpec `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.MapSpec `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.MapSpec `ebpf:"syscall_rb"`
}

// GoTlsLagacyKernel310Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadGoTlsLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type GoTlsLagacyKernel310Objects struct {
	GoTlsLagacyKernel310Programs
	GoTlsLagacyKernel310Maps
}

func (o *GoTlsLagacyKernel310Objects) Close() error {
	return _GoTlsLagacyKernel310Close(
		&o.GoTlsLagacyKernel310Programs,
		&o.GoTlsLagacyKernel310Maps,
	)
}

// GoTlsLagacyKernel310Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadGoTlsLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type GoTlsLagacyKernel310Maps struct {
	ActiveSslReadArgsMap  *ebpf.Map `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.Map `ebpf:"active_ssl_write_args_map"`
	ActiveTlsConnOpMap    *ebpf.Map `ebpf:"active_tls_conn_op_map"`
	ConnEvtRb             *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.Map `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.Map `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.Map `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.Map `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.Map `ebpf:"filter_pidns_map"`
	FirstPacketEvtMap     *ebpf.Map `ebpf:"first_packet_evt_map"`
	FirstPacketRb         *ebpf.Map `ebpf:"first_packet_rb"`
	GoCommonSymaddrsMap   *ebpf.Map `ebpf:"go_common_symaddrs_map"`
	GoSslUserSpaceCallMap *ebpf.Map `ebpf:"go_ssl_user_space_call_map"`
	GoTlsSymaddrsMap      *ebpf.Map `ebpf:"go_tls_symaddrs_map"`
	Rb                    *ebpf.Map `ebpf:"rb"`
	RegsHeap              *ebpf.Map `ebpf:"regs_heap"`
	SslDataMap            *ebpf.Map `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.Map `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.Map `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.Map `ebpf:"syscall_rb"`
}

func (m *GoTlsLagacyKernel310Maps) Close() error {
	return _GoTlsLagacyKernel310Close(
		m.ActiveSslReadArgsMap,
		m.ActiveSslWriteArgsMap,
		m.ActiveTlsConnOpMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.FilterMntnsMap,
		m.FilterNetnsMap,
		m.FilterPidMap,
		m.FilterPidnsMap,
		m.FirstPacketEvtMap,
		m.FirstPacketRb,
		m.GoCommonSymaddrsMap,
		m.GoSslUserSpaceCallMap,
		m.GoTlsSymaddrsMap,
		m.Rb,
		m.RegsHeap,
		m.SslDataMap,
		m.SslRb,
		m.SslUserSpaceCallMap,
		m.SyscallDataMap,
		m.SyscallRb,
	)
}

// GoTlsLagacyKernel310Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadGoTlsLagacyKernel310Objects or ebpf.CollectionSpec.LoadAndAssign.
type GoTlsLagacyKernel310Programs struct {
	ProbeEntryTlsConnRead   *ebpf.Program `ebpf:"probe_entry_tls_conn_read"`
	ProbeEntryTlsConnWrite  *ebpf.Program `ebpf:"probe_entry_tls_conn_write"`
	ProbeReturnTlsConnRead  *ebpf.Program `ebpf:"probe_return_tls_conn_read"`
	ProbeReturnTlsConnWrite *ebpf.Program `ebpf:"probe_return_tls_conn_write"`
}

func (p *GoTlsLagacyKernel310Programs) Close() error {
	return _GoTlsLagacyKernel310Close(
		p.ProbeEntryTlsConnRead,
		p.ProbeEntryTlsConnWrite,
		p.ProbeReturnTlsConnRead,
		p.ProbeReturnTlsConnWrite,
	)
}

func _GoTlsLagacyKernel310Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed gotlslagacykernel310_x86_bpfel.o
var _GoTlsLagacyKernel310Bytes []byte
