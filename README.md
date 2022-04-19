# clash-royale-clan-manage-tool
Help clan leader managing clan rules.\
Send promote member list and demote member list to leader.

Promote and demote conditions:\
if donations < 70 and isOldMember:\
    demote\
if donations >= 350 and role = "member":\
    promote to elder\
if lastDonations >= 1000 and donations >= 1000 and role = "elder":\
    promote to co-leader
