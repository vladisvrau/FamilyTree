package entity

type Kinship struct {
	Parent int `json:"parent"`
	Child  int `json:"child"`
}

type KinshipCollection []*Kinship

func (c *KinshipCollection) Pop(k *Kinship) *Kinship {
	if c == nil {
		return nil
	}
	index := c.IndexOf(k)
	return c.PopIndex(index)

}

func (c *KinshipCollection) PopIndex(index int) *Kinship {
	if c == nil || index > len(*c)-1 || index < 0 {
		return nil
	}

	collection := *c

	k := collection[index]
	*c = append(collection[:index], collection[index+1:]...)
	return k

}

func (c *KinshipCollection) IndexOf(k *Kinship) int {
	if c == nil {
		return -1
	}
	for index, item := range *c {
		if item.Child == k.Child && item.Parent == k.Parent {
			return index
		}
	}

	return -1
}

func (c *KinshipCollection) Push(k *Kinship) int {
	if c == nil {
		return -1
	}
	collection := *c
	collection = append(collection, k)
	*c = collection

	return len(collection) - 1
}
