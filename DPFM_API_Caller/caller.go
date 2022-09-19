package dpfm_api_caller

import (
	"data-platform-api-product-master-reads-rmq-kube/DPFM_API_Caller/requests"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

type RMQOutputter interface {
	Send(sendQueue string, payload map[string]interface{}) error
}

type DPFMAPICaller struct {
	baseURL      string
	outputQueues []string
	outputter    RMQOutputter
	log          *logger.Logger
}

func NewDPFMAPICaller(outputQueueTo []string, outputter RMQOutputter, l *logger.Logger) *DPFMAPICaller {
	return &DPFMAPICaller{
		outputQueues: outputQueueTo,
		outputter:    outputter,
		log:          l,
	}
}

func (c *DPFMAPICaller) AsyncGetProductMaster(
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
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(metaData, general)
				wg.Done()
			}()
		case "Plant":
			func() {
				c.Plant(metaData, plant)
				wg.Done()
			}()
		case "MRPArea":
			func() {
				c.MRPArea(metaData, mrpArea)
				wg.Done()
			}()
		case "Procurement":
			func() {
				c.Procurement(metaData, procurement)
				wg.Done()
			}()
		case "WorkScheduling":
			func() {
				c.WorkScheduling(metaData, workScheduling)
				wg.Done()
			}()
		case "SalesPlant":
			func() {
				c.SalesPlant(metaData, salesPlant)
				wg.Done()
			}()
		case "Accounting":
			func() {
				c.Accounting(metaData, accounting)
				wg.Done()
			}()
		case "SalesOrganization":
			func() {
				c.SalesOrganization(metaData, salesOrganization)
				wg.Done()
			}()
		// case "ProductDescByProduct":
		// 	func() {
		// 		c.ProductDescByProduct(metaData, language)
		// 		wg.Done()
		// 	}()
		// case "ProductDescByDesc":
		// 	func() {
		// 		c.ProductDescByDesc(metaData, language, productDescription)
		// 		wg.Done()
		// 	}()
		case "Quality":
			func() {
				c.Quality(metaData, quality)
				wg.Done()
			}()
		// case "SalesTax":
		// 	func() {
		// 		c.SalesTax(metaData)
		// 		wg.Done()
		// 	}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *DPFMAPICaller) General(metaData map[string]interface{}, general *requests.General) {
	metaData["function"] = "ProductMasterGeneral"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": general, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master general get request to SQLReader")
}

func (c *DPFMAPICaller) Plant(metaData map[string]interface{}, plant *requests.Plant) {
	metaData["function"] = "ProductMasterPlant"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": plant, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master plant get request to SQLReader")
}

func (c *DPFMAPICaller) MRPArea(metaData map[string]interface{}, mrpArea *requests.MRPArea) {
	metaData["function"] = "ProductMasterMRPArea"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": mrpArea, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master mrp area get request to SQLReader")
}

func (c *DPFMAPICaller) Procurement(metaData map[string]interface{}, procurement *requests.Procurement) {
	metaData["function"] = "ProductMasterProcurement"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": procurement, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master Procurement get request to SQLReader")
}

func (c *DPFMAPICaller) WorkScheduling(metaData map[string]interface{}, workScheduling *requests.WorkScheduling) {
	metaData["function"] = "ProductMasterWorkScheduling"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": workScheduling, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master WorkScheduling get request to SQLReader")
}

func (c *DPFMAPICaller) SalesPlant(metaData map[string]interface{}, salesPlant *requests.SalesPlant) {
	metaData["function"] = "ProductMasterSalesPlant"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": salesPlant, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master SalesPlant get request to SQLReader")
}

func (c *DPFMAPICaller) Accounting(metaData map[string]interface{}, accounting *requests.Accounting) {
	metaData["function"] = "ProductMasterAccounting"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": accounting, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master Accounting get request to SQLReader")
}

func (c *DPFMAPICaller) SalesOrganization(metaData map[string]interface{}, salesOrganization *requests.SalesOrganization) {
	metaData["function"] = "ProductMasterSalesOrganization"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": salesOrganization, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master SalesOrganization get request to SQLReader")
}

func (c *DPFMAPICaller) ProductDescByProduct(metaData map[string]interface{}, language string) {
	metaData["function"] = "ProductMasterProductDescription"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": language, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master ProductDescByProduct get request to SQLReader")
}

func (c *DPFMAPICaller) ProductDescByDesc(metaData map[string]interface{}, productDesc *requests.ProductDesc) {
	metaData["function"] = "ProductMasterProductDescription"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": productDesc, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master ProductDescByDesc get request to SQLReader")
}

func (c *DPFMAPICaller) Quality(metaData map[string]interface{}, quality *requests.Quality) {
	metaData["function"] = "ProductMasterQuality"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": quality, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master Quality get request to SQLReader")
}

func (c *DPFMAPICaller) SalesTax(metaData map[string]interface{}, country, taxCategory string) {
	metaData["function"] = "ProductMasterSalesTax"
	err := c.outputter.Send(c.outputQueues[0], map[string]interface{}{"message": country, "metaData": metaData})
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info("Successfully send product master SalesTax get request to SQLReader")
}
