package main

import (
	dpfm_api_caller "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller"
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
	dpfm_api_input_reader "data-platform-api-product-master-reads-rmq-kube/DPFM_API_Input_Reader"
	"data-platform-api-product-master-reads-rmq-kube/config"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client"
	sap_api_time_value_converter "github.com/latonaio/sap-api-time-value-converter"
	"golang.org/x/xerrors"
)

func main() {
	l := logger.NewLogger()
	conf := config.NewConf()
	rmq, err := rabbitmq.NewRabbitmqClient(conf.RMQ.URL(), conf.RMQ.QueueFrom(), conf.RMQ.QueueTo())
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Close()
	caller := dpfm_api_caller.NewDPFMAPICaller(
		conf.RMQ.QueueTo(),
		rmq,
		l,
	)

	iter, err := rmq.Iterator()
	if err != nil {
		l.Fatal(err.Error())
	}
	defer rmq.Stop()

	for msg := range iter {
		err = callProcess(caller, msg)
		if err != nil {
			msg.Fail()
			l.Error(err)
			continue
		}
		msg.Success()
	}
}

func callProcess(caller *dpfm_api_caller.DPFMAPICaller, msg rabbitmq.RabbitmqMessage) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = xerrors.Errorf("error occurred: %w", e)
			return
		}
	}()
	metaData, product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl, language, productDescription, country, taxCategory := extractData(msg.Data())
	accepter := getAccepter(msg.Data())
	caller.AsyncGetProductMaster(metaData, product, plant, mrpArea, valuationArea, productSalesOrg, productDistributionChnl, language, productDescription, country, taxCategory, accepter)
	return nil
}

func extractData(data map[string]interface{}) (
	metaData map[string]interface{},
	general *requests.General,
	plant *requests.Plant,
	mrpArea *requests.MRPArea,
	procurement *requests.Procurement,
	workScheduling *requests.WorkScheduling,
	salesPlant *requests.SalesPlant,
	accounting *requests.Accounting,
	salesOrganization *requests.SalesOrganization,
	productDesc *requests.ProductDesc,
	quality *requests.Quality,
) {

	sdc := dpfm_api_input_reader.ConvertToSDC(data)
	sap_api_time_value_converter.ChangeTimeFormatToSAPFormatStruct(&sdc)
	metaData = sdc.MetaData
	general = sdc.ConvertToGeneral()
	plant = sdc.ConvertToPlant()
	mrpArea = sdc.ConvertToMRPArea()
	procurement = sdc.ConvertToProcurement()
	workScheduling = sdc.ConvertToWorkScheduling()
	salesPlant = sdc.ConvertToSalesPlant()
	accounting = sdc.ConvertToAccounting()
	salesOrganization = sdc.ConvertToSalesOrganization()
	productDesc = sdc.ConvertToProductDesc()
	quality = sdc.ConvertToQuality()
	return
}

func getAccepter(data map[string]interface{}) []string {
	sdc := dpfm_api_input_reader.ConvertToSDC(data)
	accepter := sdc.Accepter
	if len(sdc.Accepter) == 0 {
		accepter = []string{"All"}
	}

	if accepter[0] == "All" {
		accepter = []string{
			"General", "Plant", "MRPArea", "Procurement",
			"WorkScheduling", "SalesPlant",
			"Accounting", "SalesOrganization", "ProductDesc",
			"Quality",
		}
	}
	return accepter
}
