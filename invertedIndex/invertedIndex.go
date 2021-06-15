package invertedIndex

// Custom index datatype for storing the indexes
// Types used for store the pointer to the actual data(role or user) with index
type record interface{}
type key map[uint64]record
type hMaps map[uint64]key

// InvertedIndex contains a hash map to search strings
// maps[mapkey][datakey]data{}
type InvertedIndex struct {
	InvertedIndex hMaps
}

// NewInvertedIndex Creates a new inverted index, Empty map
func NewInvertedIndex() *InvertedIndex {

	invertedIndex := make(hMaps, 0)
	return &InvertedIndex{
		InvertedIndex: invertedIndex,
	}
}

// Find  Search the hashMap and the key from the current indexes
func (invertedIndex *InvertedIndex) Find(mapKey, dataKey uint64) (value interface{}) { // Index is always in lower
	if hashMap, found := invertedIndex.InvertedIndex[mapKey]; found { // The tag is found
		if val, found := hashMap[dataKey]; found { // search string is found
			return val
		}
	}
	return nil
}

// FindMaps Search the hashMap and return the key map list (map[uint64]record)
func (invertedIndex *InvertedIndex) FindMaps(mapKey uint64) (value key) { // Index is always in lower
	if hashMap, found := invertedIndex.InvertedIndex[mapKey]; found { // The tag is found
		return hashMap
	}
	return nil
}

// AddItem Add the key and data to the index
func (invertedIndex *InvertedIndex) AddItem(mapKey, dataKey uint64, data interface{}) {

	// Search map for the presence of mapKey
	if _, found := invertedIndex.InvertedIndex[mapKey]; !found {
		invertedIndex.InvertedIndex[mapKey] = make(map[uint64]record)
	}

	// Search map for the presence of data key
	if _, found := invertedIndex.InvertedIndex[mapKey][dataKey]; !found {
		invertedIndex.InvertedIndex[mapKey][dataKey] = data
	}
}
