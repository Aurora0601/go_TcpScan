package main

import (
	"net"
	"strconv"
	"time"
)

type Custom_Field interface {
	Get() int
	Next()
	Isend() bool
	Size() int
}

type Interval struct {
	start int
	end   int
	cur   int
}

func (interval *Interval) Get() int {
	return interval.cur
}

func (interval *Interval) Next() {
	if interval.cur <= interval.end {
		interval.cur = interval.cur + 1
	}
}

func (interval *Interval) Isend() bool {
	return interval.cur == interval.end+1
}

func (interval *Interval) Size() int {
	return interval.end - interval.start + 1
}

type Dict struct {
	portdict []int
	idx      int
}

func newDict(dict []int) *Dict {
	var res Dict
	res.portdict = make([]int, len(dict))
	copy(res.portdict, dict)
	res.idx = 0
	return &res
}

func (dict *Dict) Get() int {
	return dict.portdict[dict.idx]
}

func (dict *Dict) Next() {
	if dict.idx != len(dict.portdict) {
		dict.idx = dict.idx + 1
	}
}

func (dict *Dict) Isend() bool {
	return dict.idx == len(dict.portdict)
}

func (dict *Dict) Size() int {
	return len(dict.portdict)
}

type Scanner struct {
	iP         string
	timeout    time.Duration
	container  []int
	port_range Custom_Field
}

// ScanPort 现在是 Scanner 的“成员方法”
func (s *Scanner) Scan_Port(port int) bool {
	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort(s.iP, strconv.Itoa(port)),
		s.timeout,
	)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func (s *Scanner) Set_Scan_Port_Range(start int, end int) {
	s.port_range = &Interval{start: start, end: end, cur: start}
}

func (s *Scanner) Set_Scan_Default_Port_Range() {
	s.port_range = &Interval{start: 1, end: 1024, cur: 0}
}

func (s *Scanner) Set_Scan_Timeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *Scanner) Set_Scan_Port_Dict(dict []int) {
	// 创建新切片并复制
	s.port_range = newDict(dict)
}

func (s *Scanner) Set_Scan_Ip(ip string) {
	s.iP = ip
}

func (s *Scanner) Scan_Ports_Single_Thread() {
	for s.port_range.Isend() {
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(s.iP, strconv.Itoa(s.port_range.Get())), s.timeout)
		if err != nil {
			continue
		}
		conn.Close()
		s.container = append(s.container, s.port_range.Get())
		s.port_range.Next()
	}
}

func (s *Scanner) Scan_Ports_Multi_Threads(maxConcurrency int) {
	//还剩用worker池多线程扫描端口没写完
}
