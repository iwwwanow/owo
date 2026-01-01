package interfaces

type PageStruct struct {
	Title     string
	Meta      MetaStruct
	Resource  ResourceInfoStruct
	Resources []ResourceInfoStruct
}
