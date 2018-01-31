//This file will show the framework of multi-chains model of PCHAIN


type PCHAIN interface {
    MainChain: MainChain
    SideChains: []SideChain
}

