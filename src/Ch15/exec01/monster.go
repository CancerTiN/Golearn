package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type monster struct {
	Name  string
	Age   int
	Skill string
}

func (m *monster) Store() error {
	filename := `D:\Workspace\Golearn\src\Ch15\exec01\monster.json`
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Serialization failed error:")
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println("Write file error:")
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *monster) ReStore() error {
	filename := `D:\Workspace\Golearn\src\Ch15\exec01\monster.json`
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Read file error:")
		fmt.Println(err)
		return err
	}
	err = json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("Deserialization failed error:")
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	mw := &monster{
		Name:  "OxDevil",
		Age:   200,
		Skill: "BananaFan",
	}
	err := mw.Store()
	fmt.Println("Store error:", err)
	// var mr *monster // error: Receiver 'mr' may be 'nil' in call
	mr := &monster{}
	fmt.Println("mr =", mr)
	err = mr.ReStore()
	fmt.Println("ReStore error:", err)
	fmt.Println("mr =", mr)
}
