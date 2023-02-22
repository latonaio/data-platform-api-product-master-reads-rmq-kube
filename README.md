# data-platform-api-product-master-reads-rmq-kube

data-platform-api-product-master-reads-rmq-kube は、周辺業務システム　を データ連携基盤 と統合することを目的に、API で品目マスタデータを登録するマイクロサービスです。  
https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 動作環境

data-platform-api-product-master-reads-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  


## 本レポジトリ が 対応する API サービス
data-platform-api-product-master-reads-rmq-kube が対応する APIサービス は、次のものです。

APIサービス URL: https://xxx.xxx.io/api/API_PRODUCT_MASTER_SRV/reads/

## 本レポジトリ に 含まれる API名
data-platform-api-product-master-reads-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_General（データ連携基盤 品目マスタ - 基本データ）
* A_Allergen（データ連携基盤 品目マスタ - アレルゲンデータ）
* A_NutritionalInfo（データ連携基盤 品目マスタ - 栄養成分データ）
* A_Calories（データ連携基盤 品目マスタ - 熱量データ）
* A_BusinessPartner（データ連携基盤 品目マスタ - 取引先データ）
* A_BPPlant（データ連携基盤 品目マスタ - 取引先プラントデータ）
* A_StorageLocation（データ連携基盤 品目マスタ - 保管場所データ）
* A_MRPArea（データ連携基盤 品目マスタ - MRPエリアデータ）
* A_WorkScheduling（データ連携基盤 品目マスタ - 作業計画データ）
* A_Accounting（データ連携基盤 品目マスタ - 会計データ）
* A_Tax（データ連携基盤 品目マスタ - 税データ）
* A_ProductDesc（データ連携基盤 品目マスタ - 品目テキストデータ）
* A_ProductDescByBP（データ連携基盤 品目マスタ - ビジネスパートナ品目テキストデータ）

 

## API への 値入力条件 の 初期値
data-platform-api-product-master-reads-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "DPFMProductMasterCreates",
	"accepter": ["General"],
	"product": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMProductMasterCreates",
	"accepter": ["All"],
	"product": "",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncProductMasterReads(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)

	var response interface{}
	// SQL処理
	response = c.readSqlProcess(nil, &mtx, input, output, accepter, &errs, log)

	return response, nil
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は 品目マスタ の 基本データ が取得された結果の JSON の例です。  
以下の項目のうち、"OrderID" ～ "PlusMinusFlag" は、/DPFM_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
    "connection_key": "request",
    "result": true,
    "redis_key": "abcdefg",
    "filepath": "/var/lib/aion/Data/rededge_sdc/abcdef.json",
    "api_status_code": 200,
    "runtime_session_id": "f982e32343b24ea39272c534547df545",
    "business_partner": 201,
    "service_label": "ORDERS",
    "message": {
        "Header": {
            "OrderID": 114,
            "OrderDate": "2022-11-22",
            "OrderType": "",
            "Buyer": 101,
            "Seller": 201,
            "CreationDate": null,
            "LastChangeDate": null,
            "ContractType": "",
            "ValidityStartDate": null,
            "ValidityEndDate": null,
            "InvoiceScheduleStartDate": null,
            "InvoiceScheduleEndDate": null,
            "TotalNetAmount": null,
            "TotalTaxAmount": null,
            "TotalGrossAmount": null,
            "OverallDeliveryStatus": "",
            "TotalBlockStatus": null,
            "OverallOrdReltdBillgStatus": "",
            "OverallDocReferenceStatus": "",
            "TransactionCurrency": "",
            "PricingDate": null,
            "PriceDetnExchangeRate": null,
            "RequestedDeliveryDate": null,
            "HeaderCompleteDeliveryIsDefined": null,
            "HeaderBillingBlockReason": null,
            "DeliveryBlockReason": null,
            "Incoterms": "CIF",
            "PaymentTerms": "0001",
            "PaymentMethod": "T",
            "ReferenceDocument": null,
            "ReferenceDocumentItem": null,
            "BPAccountAssignmentGroup": "01",
            "AccountingExchangeRate": null,
            "BillingDocumentDate": null,
            "IsExportImportDelivery": null,
            "HeaderText": ""
        },
        "HeaderPartner": [
            {
                "OrderID": 114,
                "PartnerFunction": "DELIVERTO",
                "BusinessPartner": 102,
                "BusinessPartnerFullName": "株式会社ABC虎ノ門店",
                "BusinessPartnerName": "ABC虎ノ門店",
                "Organization": "",
                "Country": "JP",
                "Language": "JA",
                "Currency": "JPY",
                "ExternalDocumentID": "",
                "AddressID": 200000
            },
            {
                "OrderID": 114,
                "PartnerFunction": "BUYER",
                "BusinessPartner": 101,
                "BusinessPartnerFullName": "株式会社ABC本社",
                "BusinessPartnerName": "ABC本社",
                "Organization": "",
                "Country": "JP",
                "Language": "JA",
                "Currency": "JPY",
                "ExternalDocumentID": "",
                "AddressID": 100000
            },
            {
                "OrderID": 114,
                "PartnerFunction": "SELLER",
                "BusinessPartner": 201,
                "BusinessPartnerFullName": "パン販売株式会社",
                "BusinessPartnerName": "パン販売",
                "Organization": "",
                "Country": "JP",
                "Language": "JA",
                "Currency": "JPY",
                "ExternalDocumentID": "",
                "AddressID": 300000
            }
        ],
        "HeaderPartnerPlant": [
            {
                "OrderID": 114,
                "PartnerFunction": "BUYER",
                "BusinessPartner": 101,
                "Plant": "AB01"
            },
            {
                "OrderID": 114,
                "PartnerFunction": "DELIVERTO",
                "BusinessPartner": 102,
                "Plant": "AB02"
            },
            {
                "OrderID": 114,
                "PartnerFunction": "SELLER",
                "BusinessPartner": 201,
                "Plant": "TE01"
            }
        ]
    },
    "api_schema": "DPFMOrdersCreates",
    "accepter": [
        "Header"
    ],
    "deleted": false,
    "sql_update_result": true,
    "sql_update_error": "",
    "subfunc_result": true,
    "subfunc_error": "",
    "exconf_result": true,
    "exconf_error": "",
    "api_processing_result": true,
    "api_processing_error": ""
}
```

