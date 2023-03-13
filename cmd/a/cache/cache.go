package main

import (
	"encoding/json"
	"fmt"
)

type UserProfile struct {
	UUID int              `json:"uuid"`
	Url  []string         `json:"url"`
	Data map[int]UserData `json:"data"`
}

type UserData struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Tel  []string `json:"tel"`
}

type CacheData struct {
	Cache map[int]UserProfile `json:"cache"`
}

type Cache struct {
	Data map[int][]byte `json:"cache"`
}

func main() {
	u := UserProfile{
		UUID: 1,
		Url:  []string{"111@com", "222@com", "333@com"},
		Data: map[int]UserData{
			1: {
				ID:   111,
				Name: "Name_111",
				Tel:  []string{"911", "922", "933"},
			},
			2: {
				ID:   222,
				Name: "Name_222",
				Tel:  []string{"955", "966", "977"},
			},
		},
	}

	//c := NewCacheData()
	//c.Set(u)
	//
	//fmt.Println(c.Get(1))
	//
	//u.UUID = 777
	//u.Url[0] = "BAD_BAD"
	//u.Data[1].Tel[0] = "BAD"
	//
	//fmt.Println(c.Get(1))

	cj := NewCache()
	cj.Set(u)
	fmt.Println(cj.Get(1))

	u.UUID = 999
	u.Url[0] = "BAD_BAD_BAD"
	u.Data[1].Tel[0] = "BAD_BAD"

	fmt.Println(cj.Get(1))

}

func NewCache() *Cache {
	return &Cache{
		Data: map[int][]byte{},
	}
}

func (c *Cache) Set(u UserProfile) {
	d, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("error json marshal, %s", err.Error())
		return
	}

	c.Data[u.UUID] = d
}

func (c *Cache) Get(UUID int) UserProfile {
	var u UserProfile
	d, ok := c.Data[UUID]
	if !ok {
		return u
	}

	err := json.Unmarshal(d, &u)
	if err != nil {
		fmt.Printf("error json unmarshal, %s", err.Error())
		return u
	}

	return u
}

//=================================================

func NewCacheData() *CacheData {
	return &CacheData{
		Cache: map[int]UserProfile{},
	}
}

func (c *CacheData) Set(u UserProfile) {
	tmpUrl := make([]string, len(u.Url))
	copy(tmpUrl, u.Url)

	tmpData := make(map[int]UserData, len(u.Data))
	for k, v := range u.Data {
		tmpTel := make([]string, len(v.Tel))
		copy(tmpTel, v.Tel)
		v.Tel = tmpTel
		tmpData[k] = v
	}

	d := UserProfile{
		UUID: u.UUID,
		Url:  tmpUrl,
		Data: tmpData,
	}

	c.Cache[d.UUID] = d
}

func (c *CacheData) Get(UUID int) UserProfile {
	user, ok := c.Cache[UUID]
	if !ok {
		return UserProfile{}
	}
	return user
}
