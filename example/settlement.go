/*
 * File Created: Thursday, 21st September 2023 5:26:32 pm
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

func SettlementFetch() {
	options := durianpay.SettlementOption{
		From:  time.Now().Format("2006-01-02"),
		To:    time.Now().Add(60 * time.Second).Format("2006-01-02"),
		Skip:  10,
		Limit: 10,
	}

	res, err := c.Settlement.FetchSettlements(ctx, options)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
