package resources

import (
	api "github.com/antonmaltsev/cabstore/api"
	log "gopkg.in/inconshreveable/log15.v2"
	config "github.com/antonmaltsev/cabstore/cfg"
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"strconv"
	"math"
)

type CabifyResource struct {
	Cfg config.Config
}

func (tr *CabifyResource) OrderSum(c *gin.Context) {
	var data api.Order
	var arr []string 
	var summ float32

	summ = 0
	log.Info("Started")

	err := c.Bind(&data)
	if err != nil {
		c.JSON(400, api.NewError("problem decoding body"))
		return
	}

	log.Info("Items parcing, data is: " + data.Items)

	arr = strings.Split(data.Items, ",")
	if(arr == nil){
		log.Error("Empty order list, please fill it")
		}

	var VOUCHER_cnt = strings.Count(data.Items, api.VOUCHER)
	var TSHIRT_cnt = strings.Count(data.Items, api.TSHIRT)
	var MUG_cnt = strings.Count(data.Items, api.MUG)


	log.Info("The number of " + api.VOUCHER + " occurence = " + strconv.Itoa(int(VOUCHER_cnt)))
	log.Info("The number of "+ api.TSHIRT +" occurence = " + strconv.Itoa(int(TSHIRT_cnt)))
	log.Info("The number of " + api.MUG + " occurence = " + strconv.Itoa(int(MUG_cnt)))

	if(TSHIRT_cnt >= 3){
		summ = summ + float32(TSHIRT_cnt*api.CFO_DISCONT_TSHIRT_PRICE)
	}else if((TSHIRT_cnt < 3) && (TSHIRT_cnt > 0)){
		summ = summ + api.TSHIRT_PRICE
	}

	if(VOUCHER_cnt > 0){				
		summ = summ + api.VOUCHER_PRICE

		if(math.Mod(float64(MUG_cnt),2) == 0){
			summ = summ + float32(int(VOUCHER_cnt/2))*api.VOUCHER_PRICE
		}
	}

	summ = summ + float32(MUG_cnt)*api.MUG_PRICE
	
	c.JSON(200, gin.H{
        "Order Summ": fmt.Sprintf("%.2f", summ),
        "Status": "Success",
    })
}



    

    
