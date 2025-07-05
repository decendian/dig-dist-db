type Multihash []byte
type NodeId Multihash
type PublicKey []byte
type PrivateKey []byte

type Node struct{
    NodeId NodeID
    PubKey  PublicKey
    PriKey  PrivateKey
}

type IPFSRouting interface{
    FindPeer(node NodeId) // gets a particular peer’s network address
    SetValue(key []byte, value []byte) // stores a small metadata value in DHT
    GetValue(key []byte) // retrieves small metadata value from DHT
    ProvideValue(key Multihash) // announces this node can serve a large value
    FindValuePeers(key Multihash, min int) // gets a number of peers serving a large value
}

type Ledger struct{
	Owner NodeId


}



