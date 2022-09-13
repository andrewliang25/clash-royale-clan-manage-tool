package main

import "time"

type Arena struct {
	Id   int
	Name string
}

type ClanMemberInfo struct {
	Tag              string
	Name             string
	Role             string
	LastSeen         string
	ExpLevel         int
	Trophies         int
	Arena            Arena
	ClanRank         int
	PreviousClanRank int
	Donations        int
	DonationsRecived int
	ClanChestPoints  int
}

type ClanMembersData struct {
	Reason  string
	Message string
	Items   []ClanMemberInfo
	Paging  interface{}
}

type DemoteAndPromoteList struct {
	DemoteList  []string
	PromoteList []string
	UpdateTime  time.Time
}
