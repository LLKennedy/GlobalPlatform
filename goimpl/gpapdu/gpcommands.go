package gpapdu

// Commands is the full client command list
// TODO: none of these have actual arguments/returns yet
type Commands interface {
	Delete()
	GetData()
	GetStatus()
	Install() // TODO: there's a LOT of weirdness for this one
	Load()
	PutKey()
	Select()
	SetStatus()
	StoreData()
}
