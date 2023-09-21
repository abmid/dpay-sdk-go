/*
 * File Created: Thursday, 21st September 2023 5:17:01 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package example

import (
	"fmt"
	"time"

	durianpay "github.com/abmid/dpay-sdk-go"
)

func PromoCreate() {
	payload := durianpay.PromoPayload{
		Type:     "card_promos",
		Label:    "SALE502022",
		Currency: "IDR",
		PromoDetails: durianpay.PromoDetails{
			BinList:   []int{424242},
			BankCodes: []string{},
		},
		DiscountType:       "percentage",
		Discount:           "10",
		StartsAt:           time.Now(),
		EndsAt:             time.Now(),
		SubType:            "direct_discount",
		LimitType:          "quota",
		LimitValue:         "100",
		PriceDeductionType: "total_price",
		Code:               "SALE2022",
	}

	res, err := c.Promo.Create(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}

func PromoFetchPromos() {
	res, err := c.Promo.FetchPromos(ctx)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
