package metadata

import "sync"

var metaDataOnce sync.Once
var metaDataInstance *CfgMetaData

func Instance() *CfgMetaData {
	metaDataOnce.Do(func() {
		metaDataInstance = newCfgMetaData()
		metaDataInstance.init()
	})
	return metaDataInstance
}

func newCfgMetaData() *CfgMetaData {
	return &CfgMetaData{metaMap: make(map[string]Item)}
}

type CfgMetaData struct {
	metaMap map[string]Item
}

func (c *CfgMetaData) init() {
	c.initFromArray(ConfigList)
}

func (c *CfgMetaData) initFromArray(items []Item) {
	c.metaMap = make(map[string]Item)
	for _, item := range items {
		c.metaMap[item.Name] = item
	}
}

// GetByName - Get current metadata of current name, if not defined, return false in second params
func (c *CfgMetaData) GetByName(name string) (*Item, bool) {
	if item, ok := c.metaMap[name]; ok {
		return &item, true
	}
	return nil, false
}

// GetAll - Get all metadata in current env
func (c *CfgMetaData) GetAll() []Item {
	metaDataList := make([]Item, 0)
	for _, value := range c.metaMap {
		metaDataList = append(metaDataList, value)
	}
	return metaDataList
}
