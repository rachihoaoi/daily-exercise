package sort

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	array := []int64{1, 2, 4, 7, 4, 2, 4, 6, 7, 2, 34, 56, 7}
	BubbleSort(array)
	fmt.Println(array)
}

func TestIstio(t *testing.T) {
	client := &http.Client{}
	host := os.Getenv("HOST")
	//grpcHost := os.Getenv("GRPC_HOST")
	if len(host) < 1 {
		host = "https://t-api3.cls.cn/api/video/437562?app=stib&channel=0&cuid=194FEB5F-4E94-41E2-BDF7-F2FC4A5FD358&mb=iPhone11%2C8&net=2&os=ios&ov=13.4.1&platform=iphone&province_code=3101&sign=5775de1037a71b5240b6ee79a7e3bc0a&sv=1.4.1"
	}
	//if len(grpcHost) < 1 {
	//	grpcHost = "127.0.0.1:5001"
	//}
	reqUrl := host
	fmt.Println("REST HOST", reqUrl)
	//fmt.Println("GRPC HOST", grpcHost)
	for {
		time.Sleep(time.Millisecond * 10)
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			fmt.Println("new request failed.", err)
			break
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("do request failed.", err)
			break
		}
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			break
		}
		resp.Body.Close()
		fmt.Printf("REST Server sayhello[GET]: %s \n", string(bodyByte))
		//CallGrpc(grpcHost)
		fmt.Println("===============================================================")
	}
}
