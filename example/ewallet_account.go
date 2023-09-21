/*
 * File Created: Thursday, 21st September 2023 5:32:03 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package example

import (
	"fmt"

	durianpay "github.com/abmid/dpay-sdk-go"
)

func EWalletAccountLink() {
	payload := durianpay.EwalletAccountLinkPayload{
		Mobile:      "8888888888",
		WalletType:  "GOPAY",
		RedirectURL: "https://redirect_url.com/",
	}

	res, err := c.EWalletAccount.Link(ctx, payload)
	if err != nil {
		// Handle error
	}

	fmt.Println(res)
}
