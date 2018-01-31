//This file defines the apis of main chain


type SideChain interface {
    
    GetPCHAIN() PCHAIN
    GetService(name string) Service
    Name() string
}
