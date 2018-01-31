//This file defines the apis of main chain


type MainChain interface {
    
    GetPCHAIN() PCHAIN
    GetService(name string) Service
    RegisterSideChain(sideChain SideChain) error
}

