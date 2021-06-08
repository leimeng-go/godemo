package json

import (
	"bufio"
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"strings"
	"testing"
	"time"
)
var jsonStr1=``

type Transaction struct{
  
   TransactionTime int64 `json:"transactiontime"`
   TransactionAmount int `json:"transactionamount"`
   PaymentChannel string `json:"paymentchannel"`
   TenantName string `json:"tenantname"`
   LandlordName string `json:"landlordname"`
   BuildingName string `json:"buildingname"`
   Desc string `json:"desc"`
   
}
func TestJson(t *testing.T){
	newStr:=strings.ReplaceAll(jsonStr1,"NumberLong(","")
	newStr1:=strings.ReplaceAll(newStr,")","")
	br:=bufio.NewReader(strings.NewReader(newStr1))
	list:=make([]Transaction,0)
	for {
       l,e:=br.ReadBytes('\n')
	   if e==io.EOF{
		   break
	   }
	   element:=Transaction{}
	   err:=json.Unmarshal(l,&element)
	   if err!=nil{
		   t.Log(err.Error())
	   }
	   list=append(list,element)
	//    os.Stdout.Write(l)
	}
   
   for _,v:=range list{
	  fmt.Printf("房东名称: %s,租客名称: %s,支付渠道: %s,交易金额: %0.2f,交易时间: %s,交易描述: %s\n",
	  v.LandlordName,v.TenantName,v.PaymentChannel,
	  float64(v.TransactionAmount)/100,time.Unix(v.TransactionTime,0).Local().Format("2006-01-02 15:04:05"),v.Desc)
   }
}
func TestExcel(t *testing.T) {
	categories:=map[string]string{
		"A1": "房东名称", "B1": "楼栋名称", "C1": "租客名称", "D1": "支付渠道", "E1": "交易金额", "F1": "交易时间", "G1": "交易描述",
	}
	//values:=map[string]string{
	//	"A2":"测试","B2":"haha","C2":"测试1","D2":"232","E2":"323","F2":"32","G2":"2323",
	//}
    f:=excelize.NewFile()
    for k,v:=range categories {
    	f.SetCellValue("Sheet1",k, v)
	}
	data:=getFileList("F:\\data\\何小翠交易流水.txt")


	for i,v:=range data["刘"] {
		f.SetCellValue("Sheet1",fmt.Sprintf("A%d",i+2),v.LandlordName)
		f.SetCellValue("Sheet1",fmt.Sprintf("B%d",i+2),v.BuildingName)
		f.SetCellValue("Sheet1",fmt.Sprintf("C%d",i+2),v.TenantName)
		f.SetCellValue("Sheet1",fmt.Sprintf("D%d",i+2),v.PaymentChannel)
		f.SetCellValue("Sheet1",fmt.Sprintf("E%d",i+2),v.TransactionAmount)
		f.SetCellValue("Sheet1",fmt.Sprintf("F%d",i+2),v.TransactionTime)
		f.SetCellValue("Sheet1",fmt.Sprintf("G%d",i+2),v.Desc)

	}
	if err:=f.SaveAs("Book1.xlsx");err!=nil{
		t.Log(err.Error())
	}
}
func getFileList(dir string)map[string][]Transaction {
	data:=make(map[string][]Transaction,0)
	dirPath:=filepath.Dir(dir)
	filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			file, err:=os.Open(path)
			if err!=nil{
				log.Println(err.Error())
				return err
			}
			buf, err:=ioutil.ReadAll(file)
			newStr:=strings.ReplaceAll(string(buf),"NumberLong(","")
			newStr1:=strings.ReplaceAll(newStr,")","")
			br:=bufio.NewReader(strings.NewReader(newStr1))
			list:=make([]Transaction,0)
			for {
				l,e:=br.ReadBytes('\n')
				if e==io.EOF{
					break
				}
				element:=Transaction{}
				err:=json.Unmarshal(l,&element)
				if err!=nil{
					log.Println(err.Error())
					return err
				}
				list=append(list,element)

			}
			data[d.Name()]= list
			return nil
		}

		return nil
	})
	return data
}
