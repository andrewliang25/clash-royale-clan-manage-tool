package main

import "time"

type arena struct {
	Id   int
	Name string
}

type clanMemberInfo struct {
	Tag              string
	Name             string
	Role             string
	LastSeen         string
	ExpLevel         int
	Trophies         int
	Arena            arena
	ClanRank         int
	PreviousClanRank int
	Donations        int
	DonationsRecived int
	ClanChestPoints  int
}

type clanMembersData struct {
	Reason  string
	Message string
	Items   []clanMemberInfo
	Paging  interface{}
}

type demoteAndPromoteList struct {
	DemoteList  []string
	PromoteList []string
	UpdateTime  time.Time
}
