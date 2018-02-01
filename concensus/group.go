//This file define the group's functionalities of group

type Group interface {

  IsMember(pubKey string) bool
  SubmitBlock() bool
}

type GroupMember interface {

  Vote(hash string) string
  
}

