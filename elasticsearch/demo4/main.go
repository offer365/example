package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/olivere/elastic"
	"odin.ren/endecrypt"
	"odin.ren/endecrypt/endeaesrsa"
	"odin.ren/hardware"
)

var rank_code = `mVT7iVjK2gHrMV7awKUXjATP4KXDlm4VODkKwbyNyzDRAR5Hfv3/6rUSpsthkvzStFHUVKndyLFyhFflEvu1NZ1N3TcEhXlYxWG0D8YdBEzL22cLJ+vveR4jG1gN1dYVR0/PSZEC9HADuSOx1RGIpj7p7QcFZHrIRXUV0hzLKV5iv0VsQtlUaKEBl/DHXoCzV4fAW8Oz8684QFQuAWaJ/hTVVeOq0EKRty2XSM0uQQROB6kKjnc120L/o6Wm+N2Hkifdp1pyTM4nseNxnQGGa/Ij5tr3u9m0/qbzyDTYZBPwaDH3M0eiynqONyHRfMPKSxFpe97fjfmforTjNs3qWqgxBErirS6lyl7neH/85c5n41qDqFdIVdu8ZrEeXjZdEr2puRM63Xx4tVnrqt9xZZL34CTjRpRHlvN/zLqbUeOdxZ2RrqxJ5BKWMsBrLt5qxCF0ap9LfQV/apAv1QokJs/PKC5wYiaZJtN3wCNoq7rcpMo379+R/iJDCns3JkGi3rNoiKSHLjXVZAYkIckdFlOB/u/xjTMj4rNO0X7c97t1TRua5N87xvShrFIRHGZ3AMQ+zxEkcR0vbwRsRM5WGAsZxgl0m/4iLKE/96bGtLhMT/T2Nb1fS2ggwJt61kpu8SbHiUv8fTGGp4lK8Eyuz60yIcaxw21OdhY9aXKMrex57agC+8tZLWufCNSZVUUiuXYDdlVOLnxnSKarw5vsuD80sYbFP3Vj8IvXbJt0SLFX1TYYDwZTbRBqJafToIsPCesAbEDJqaVQcuwpxGDdGk9Tf/MoQpDCytQSrgTQXnLaaa5BI3yucbswF05d36ME7waEUV7v+Fxo9lgF6Nv1nhDBpLfS02S9/6WtfHfNH04mrhnvzybgeE49NNEs6ziFvTA2E7wg+4oc1eqBcqhfXLnfWs0UEaAYwkYKVu6Zn5ZMPSvlIW/cYyQ1L6WKUIlSICu7EqgY6jtIYlQLn0zRCZ5jVkYNveLs/ftt7qCkZ6KS75y2gbccVKUOkrMqUwvFDJwyQnhKsJ1SEmVObea+cA==`
var index = "asgard"
var typ = "rank_code"

// 授权
type License struct {
	Lid          string  `json:"lid"`            // 授权码唯一uuid,用来甄别是否重复授权。
	Rid          string  `json:"rid"`            // 机器码的id, lid与rid 一一对应
	Nodes        []nodes `json:"nodes"`          // 所有成员的hardware
	Instances    int     `json:"instances"`      // 实例
	Timestamp    int64   `json:"timestamp"`      // 授权到期的时间戳
	UpdateTime   int64   `json:"update_time"`    // 生成或更新时间戳
	MaxLifeCycle uint64  `json:"max_life_cycle"` // 最大生存周期 (授权到期时间-生成授权时间)/周期时间60s
	LifeCycle    uint64  `json:"life_cycle"`     // 当前生存周期
}

// 生成密文
func (l *License) String() (text string) {
	var (
		byt []byte
		err error
	)
	if byt, err = json.Marshal(l); err != nil {
		return
	}
	// 私钥加密
	if text, err = endeaesrsa.PriEncrypt(byt, endecrypt.PirvatekeyCode, endecrypt.AesKeyCode); err != nil {
		return
	}
	return
}

// 序列号
type RankCode struct {
	Rid  string `json:"rid"` // 序列号唯一uuid，用来标识序列号，并与 授权码相互校验，一一对应。
	Hws  []*hws `json:"hws"`
	Time int64  `json:"time"`
}

type hws struct {
	Name string            `json:"name"`
	Hw   hardware.Hardware `json:"hw"`
}

type nodes struct {
	Name string `json:"name"`
	Hw   string `json:"hw"`
}

const mapping = `
{
    "settings":{
        "number_of_shards":1,
        "number_of_replicas":0
    },
    "mappings":{
        "rank_code":{
            "properties":{
                "rid":{
                    "type":"keyword"
                },
                "hws":{
                    "properties":{
                        "name":{
                            "type":"keyword"
                        },
                        "hw":{
                            "properties":{
                                "cpu":{
                                    "properties":{
                                        "model":{
                                            "type":"text"
                                        },
                                        "num":{
                                            "type":"byte"
                                        }
                                    }
                                },
                                "mac":{
                                    "type":"text"
                                },
                                "dmi":{
                                    "type":"text"
                                }
                            }
                        }
                    }
                },
                "time":{
                    "type":"long"
                }
            }
        }
    }
}`

var ctx = context.Background()

type ESC struct {
	Client *elastic.Client
}

var example = `{
	User:"admin", 
	Lid:"82d1cc82-c6d2-4624-b36b-0ad444f8641d", 
	Rid:"f9d4bf25-8945-4e11-a6b5-f913018a6a13", 
	Ctime:1554711040, 
	Create:1554781954, 
	RankCode:"waoYnpIEvu7a3yNeGGWGTCVLYk0NTr4P76FXkUR9Y3ACP3/qsJTggDKefLoG+KqsHMuI4ale5QZOTPUUQzJvT1RxL2sUSS0xX/RF8+TkFQUCw/0e1v9ysMNqr+v+ds9+W7KdEhDeA9Sv6DTmU/3UUKeyTAAxeMbSABdDaBy32hyZZ92xkLdFj10n+wPWhXrLu3sSnXogTnXUoyTA+KomnFideOCMl4vMcDXJB/0uCyNL+q1MJixhwDh8on1P+waYMdsSZhyn5mq/hEv16gbfUDg23aAlx4+ePo+2QBRSnyyPweAUDJUPulwKpmT2hJYFHIgVfWsMplnA0L6y6Q9AJ8OyMT+0fc/iUm4wiw3MYqZVEGHD7QxhcTnZQKMZBQdV6YPZ1g9+BbajMG4NeAJvQ2zCVeGrdqAF/sH5cU2Aq2u9nD6+VSMPCMRrzMMmtWx3dzPbdA2D63x5lpgDjV2GpugYRhMtmarUSW2N2LWcwjxmDjWoUPHkp8VPGpA3g+dVU4e/nemQjQf3Vix/nk0QA2ex6F8AzLigoZRFis22WANciTpfnq9tfh1zpx6YKkkenMWjqOquzfNnAzfODJc3JLr5rz2X7llwZHb9yHzZYuoSlC4iwTvfgkH0aCCTvay3JadqsmcTXsy1Tvzo2McOA7i/U+EZmoMcyYJrZkMtdCoc2+kb4cTzOYOfMr4+aI1O5nBgyBSx+nyqQgDDBotRMVrN9mHbabGlxfn9mEht47qdS43jXXgdHyTmwAU4F+D601FjoayCLa7jPbz11fikUFCiL8NDL8Uo107uhNrGNnZJuEsGF5u+XIBroqToY+WR/F9xkR8TrkkZQDFpgkZj2UGIxU9+jUcpTg4bdCC/Mqx5MUfBNffu/89JK1Ft+OfsFZXhtHFCW3/7JohOCoStVdV6KygGHJIgl4+esfj0AEIhOn6J8BTbhZ9lBMQ+E5+kLs+alOzQO/2I6sE644mUFO7N1tng+ZjkmqGxdNLXmRzD9F+7L3Y8ivlaDTDssva415d3SWdX7rcpwNBQbdyEjw==", 
	License:"3dXk3eiKvJPsbKlCt8/zM7aOzK9JiJAtGXDWn7xySUlJRaIN7NlY/Au8uT+uQ0HdIFk/gxOiL9rZEoT2RFs8cncpnsUaoMTFpGcTNBgRR3jzAnYYbW0jwqJxnheKZVBsJ0pyP7+kT84a4tInvEC8XMf6uHSk77PRIrbso/Xj7P3GQhPlvpjDto9LFfkEg10+OFbHlh0DV/PHlDNcXhGcgUavZgISywoe1pRldTC338hdram+y0aY7LwqMFjr6xtuXswn1A+RxuqhvKq2v2g9gqpFf6/ZUq+CT9Is8Xz0VnBs5EmRn6EWm5lLJnLL43PwOix0pkuv65yE+I+3eNI4Fi2ogEgbsqH5pY/dmD4B4E62oALur43a9dsi7NCrzlu5e2mhMJxwu+AB71YfRsnk38mPAiX4jwH62cc2uEarpq3bR6hhyrQ30LQ9ZhA+kIUimVvUu7hFv0vqEfuF8f4IIXADbobGgwV/xW8MWEmvByu2VpJjtv/lUf3J6IVASxYOXo0h3ElY6OlCS9tam5koiY1gO3oYgwb2LcVHyHyliGF4B7cMXxAu//Fet1n7r682jiEBEnc0z5BMUL4XJfHb9zfnjwnlJZH1IRLaoVGbt9oa8VyXFatC9YbNniC6HisOLCRQlu9Gd+N/A5pSAXYV3RYQRI2Ua5gb7OxnI+vuYDQ2bG0Ioeyd4mfzHQsSDU4zH1GDvgUPr5UShtt6t7kCQai5DODy37+Qm/IXI/c3v/3B6WbDUrGPcmY5TadyiJZMXcNpNHmMIoktQPWVLuJ0KHtxdXVjrNqCY9/x21o5xkeE3ZEIzfgVk+iQA8p6FUDABl3IoQeEwa5f/7G1cqR2ya6iHQiad3v3GdmbUjDUS8IBHcHjOxSnI6CxlUE/JtNJRei2FuWSIGiXWlg7j+78QfTAn99DCvX8ilteDMkTU5TwCvDKah3GkLOB/8AVI3bmHN6Dzi+uPDVn+w4SzcCebLzfosVReAwpw767iadksKa70nnFUOfjqeVbkpc79gmmtkJgZ1L1rHuwWWsdXu9gw2vDAajE6l6iklMIDTeAHRr9/eEp4HWg1//DX1xPFLCKw0zvf2rTTbD87J6OojbWNM0gWeWQU1oGRxmprQCvu+czwzNDm+epMP8SQqDSC63cHS15XL5OkiCdELKQedWI6ldhGjAq1aEippYstntERjbZ8uQ8M1wZjbu5keYyiNPQR1eCwVDBVdSua64brTCVDoymoFCIz9yjfQFx063wyfslRne7QVitGCFV8JR8agjY2UjnJ+ot1MTHlIao1ebvFRnEjH0xcVTwaa8bSYGx31mbu0Y0GhhPdWDirKXNouJPxAQgjUMdhUlbYXu6o5XpwYaFLdHFYClQCAggRfEmyt6wcn+f5irowalpARNeLWUnXBeY67Q3gVyLj5S4t8BNj4e7ubZjli4aeISuatJ6TuAw7QxrWf5iEV+P5kC32IBu/KGdl2S+ENJiWiP4kyvxiL/yApmDNTYqgOJCsoP6mJrRLTLlmYsE37zHMH25pTLXefO6UVz4lqViVsHzzer6rBKuTUZGRzN+jUzeOPekLUOBeEqOEwvrEntU2XX6JAq9DnkaikVkAEPglatH6Iq9B/97fVudX1His/P6E/j5Y6r6aVgHwlYVqRiGeNl3vOpJ1V6vw190Z7I91r+1IkzsWQhs0WHFa7p8xHf+DI3x0x2BhKi3bw5E7UblYnIs+HJ/", 
	Instance:100, 
	Timestamp:1556582400, 
	Hws:{
		"odin0":{
				Cpu:{"Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz":2}, 
				Mac:{"00:0c:29:3b:26:90"}, 
				Dmi:"VMware, Inc. VMware Virtual Platform/440BX Desktop Reference Platform, BIOS 6.00 07/02/2015"}, 
		"odin1":{
				Cpu:{"Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz":2}, 
				Mac:{"00:0c:29:f2:0a:1d"}, 
				Dmi:"VMware, Inc. VMware Virtual Platform/440BX Desktop Reference Platform, BIOS 6.00 07/02/2015"}, 
		"odin2":{
				Cpu:{"Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz":2}, 
				Mac:{"00:0c:29:a3:94:dc"}, 
				Dmi:"VMware, Inc. VMware Virtual Platform/440BX Desktop Reference Platform, BIOS 6.00 07/02/2015"}}, 
	Nodes:{
		"odin0":"27f9877dab69cdbb3c2662da4041e3b8553d9e2ba03a78fa6184b0b4791f702a",
		"odin1":"41b411b2edeeb4ba16eeaac14793976bcb39275fed4321446d245e21d3dc9980", 
		"odin2":"c2690a1b1cea5902d885c0e361710477f6a4ea17b6968cd2c57671fd20693073"}, 
	Sum:"2c4b5c891ef3a16790c639ac77738cad69da7ea20108e27cf180ea1343d9e223"
}
`

func NewESC(host string) (esc *ESC, err error) {
	esc = new(ESC)
	esc.Client, err = elastic.NewClient(elastic.SetURL(host), elastic.SetSniff(false))
	return
}

func (E *ESC) CreateIndex() {

	// Use the IndexExists service to check if a specified index exists.
	exists, err := E.Client.IndexExists(index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := E.Client.CreateIndex(index).BodyString(mapping).Do(ctx)
		// createIndex, err := E.Client.CreateIndex(index).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func (E *ESC) Create(id string, r *RankCode) (err error) {
	res, err := E.Client.Index().Index(index).Type(typ).Id(id).BodyJson(r).Do(context.Background())
	res = res
	return
}

func (E *ESC) Get(id string) (err error) {
	res, err := E.Client.Get().Index(index).Type(typ).Id(id).Do(context.Background())
	res = res
	return
}

func (E *ESC) Search() (err error) {
	termQuery := elastic.NewTermQuery("_id", 2)
	res, err := E.Client.Search().
		Index(index).
		Type(typ).
		Query(termQuery).
		Sort("_id", true).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

	fmt.Printf("Query took %d milliseconds\n", res.TookInMillis)
	if res.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)

		for _, hit := range res.Hits.Hits {

			var rc RankCode
			err := json.Unmarshal(*hit.Source, &rc) // 另外一种取数据的方法
			if err != nil {
				fmt.Println("Deserialization failed")
			}
			fmt.Printf("%+v\n", rc)
		}
	} else {
		fmt.Printf("Found no Employee \n")
	}
	return
}

func (E *ESC) Update() (err error) {
	update, err := E.Client.Update().Index("twitter").Type("tweet").Id("1").
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)
	return

}

func (E *ESC) Delete() (err error) {
	// Delete an index.
	resp, err := E.Client.Delete().Index(index).Type("employee").Id("1").Do(ctx)
	fmt.Println(resp, err)
	return
}

// //修改
// func update() {
//	res, err := client.Update().
//		Index("megacorp").
//		Type("employee").
//		Id("2").
//		Doc(map[string]interface{}{"age": 88}).
//		Do(context.Background())
//	if err != nil {
//		println(err.Error())
//	}
//	fmt.Printf("update age %s\n", res.Result)
//
// }
//
// //查找
// func gets() {
//	//通过id查找
//	get1, err := client.Get().Index("megacorp").Type("employee").Id("2").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	if get1.Found {
//		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
//	}
// }
//
// //搜索
// func query() {
//	var res *elastic.SearchResult
//	var err error
//	//取所有
//	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
//	printEmployee(res, err)
//
//	//字段相等
//	q := elastic.NewQueryStringQuery("last_name:Smith")
//	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//	if err != nil {
//		println(err.Error())
//	}
//	printEmployee(res, err)
//
//	if res.Hits.TotalHits > 0 {
//		fmt.Printf("Found a total of %d Employee \n", res.Hits.TotalHits)
//
//		for _, hit := range res.Hits.Hits {
//
//			var t Employee
//			err := json.Unmarshal(*hit.Source, &t) //另外一种取数据的方法
//			if err != nil {
//				fmt.Println("Deserialization failed")
//			}
//
//			fmt.Printf("Employee name %s : %s\n", t.FirstName, t.LastName)
//		}
//	} else {
//		fmt.Printf("Found no Employee \n")
//	}
//
//	//条件查询
//	//年龄大于30岁的
//	boolQ := elastic.NewBoolQuery()
//	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
//	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
//	res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
//	printEmployee(res, err)
//
//	//短语搜索 搜索about字段中有 rock climbing
//	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
//	res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
//	printEmployee(res, err)
//
//	//分析 interests
//	aggs := elastic.NewTermsAggregation().Field("interests")
//	res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
//	printEmployee(res, err)
//
// }
//
// //简单分页
// func list(size, page int) {
//	if size < 0 || page < 1 {
//		fmt.Printf("param error")
//		return
//	}
//	res, err := client.Search("megacorp").
//		Type("employee").
//		Size(size).
//		From((page - 1) * size).
//		Do(context.Background())
//	printEmployee(res, err)
//
// }
//
// //打印查询到的Employee
// func printEmployee(res *elastic.SearchResult, err error) {
//	if err != nil {
//		print(err.Error())
//		return
//	}
//	var typ Employee
//	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
//		t := item.(Employee)
//		fmt.Printf("%#v\n", t)
//	}
// }

func main() {
	var Host = "http://10.0.0.220:9200"
	esc, err := NewESC(Host)
	fmt.Println(err)
	esc.CreateIndex()
	rc, err := Decrypt(rank_code)
	// fmt.Println(rc)
	for i := 0; i < 10; i++ {
		esc.Create(strconv.Itoa(i), rc)
	}
	esc.Search()
}

// 解密序列号
func Decrypt(src string) (r *RankCode, err error) {
	r = new(RankCode)
	r.Hws = make([]*hws, 0)
	Err := func() (*RankCode, error) {
		return r, nil
	}
	defer func() {
		if err := recover(); err != nil {
			Err()
		}
	}()
	// 私钥解密
	if src, err = endeaesrsa.PirDecrypt(src, endecrypt.PirvatekeyCode, endecrypt.AesKeyCode); err != nil {
		fmt.Println("解密失败"+"", err.Error())
		return
	}
	if err = json.Unmarshal([]byte(src), r); err != nil {
		return
	}
	return
}
