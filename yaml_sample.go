package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	yaml "gopkg.in/yaml.v2"
)

// structたち
type Data struct {
	Setting []Sites `yaml:"setting"`
}

type Sites struct {
	Fcid   int    `yaml:"fcid"`
	FcInfo fcInfo `yaml:"fc_info"`
}

type fcInfo struct {
	Name   string `yaml:"name"`
	Status string `yaml:"status"`
}

func main() {
	// yamlを読み込む
	buf, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		panic(err)
	}

	// structにUnmasrshal
	var data Data
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data.Setting)
	fmt.Println(data.Setting[2].FcInfo.Name)
	fmt.Println(data.Setting[3].FcInfo.Status)

	// 読み込んだyamlファイルのステータスを取り出して配列に格納
	var status []string
	// var trun string
	for i := 0; i < len(data.Setting); i++ {
		// 配列に既に追加済みかをチェックする
		duplicate_judge := include(status, data.Setting[i].FcInfo.Status)
		// 一意のステータスを配列に格納する
		if !duplicate_judge {
			status = append(status, data.Setting[i].FcInfo.Status)
		}
	}
	fmt.Println(status)

	// Case文SQL作成処理
	sql_case := "case XXXXX"
	fmt.Println(sql_case)
	// 一意のステータスの数For文を回す(NoMonitaring...Attention...)
	for s := 0; s < len(status); s++ {
		sql_case = sql_case + " when " + status[s] + " then "
		fmt.Println(sql_case)
		// fcidの数だけFor文を回す(111...112...)
		for i := 0; i < len(data.Setting); i++ {
			// ステータスが一致している場合、SQLにfcidを追加する
			if data.Setting[i].FcInfo.Status == status[s] {
				sql_case = sql_case + " " + strconv.Itoa(data.Setting[i].Fcid) + " ,"
				fmt.Println(sql_case)
			}
		}
		fmt.Println(sql_case)
		sql_case = sql_case[:len(sql_case)-1]
		fmt.Println(sql_case)
	}
}

// 配列に任意の要素があるかを確認する関数
func include(slice []string, target string) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}
	return false
}
