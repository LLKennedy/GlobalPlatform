package gpapdu

// Commands is the full client command list
// TODO: none of these have actual arguments/returns yet
type Commands interface {
	Delete(deleteRelatedObjects bool, cmd DeleteCommand, logicalChannel uint8) (confirmation []byte, err error)
	GetData()
	GetStatus()
	Install() // TODO: there's a LOT of weirdness for this one
	Load()
	PutKey()
	Select()
	SetStatus()
	StoreData()
}
