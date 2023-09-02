/*
 * File Created: Saturday, 2nd September 2023 1:53:00 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

// EwalletAccountLinkPayload is payload for Link E-Wallet Account API.
type EwalletAccountLinkPayload struct {
	Mobile      string `json:"mobile"`
	WalletType  string `json:"wallet_type"`
	RedirectURL string `json:"redirect_url"`
}

/*
Options
*/
