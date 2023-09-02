/*
 * File Created: Saturday, 2nd September 2023 2:00:18 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package ewalletaccount

// Link is struct for response Link E-Wallet Account API
type Link struct {
	WalletType     string `json:"wallet_type"`
	Mobile         string `json:"mobile"`
	RefID          string `json:"ref_id"`
	Status         string `json:"status"`
	AppRedirectURL string `json:"app_redirect_url"`
	Message        string `json:"message"`
}

// Unlink is struct for response Unlink E-Wallet Account API
type Unlink struct {
	WalletType string `json:"wallet_type"`
	Mobile     string `json:"mobile"`
	RefID      string `json:"ref_id"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// Detail is struct for response EWallet Account Details API
type Detail struct {
	WalletType string `json:"wallet_type"`
	RefID      string `json:"ref_id"`
	Status     string `json:"status"`
	Mobile     string `json:"mobile"`
	Balance    string `json:"balance"`
	Currency   string `json:"currency"`
	Toke       string `json:"token"`
}

// {
//     "data": {
//         "wallet_type": "GOPAY",
//         "mobile": "8888888888",
//         "ref_id": "7f125e70-095e-481d-8db8-241df9d5b86d",
//         "status": "pending",
//         "app_redirect_url": "https://simulator.sandbox.midtrans.com/gopay/partner/web/otp?id=14c95e30-0586-4270-961e-f3b0b3d3d2b0",
//         "message": "use redirection url to bind the account"
//     }
// }
