/*
 * File Created: Saturday, 2nd September 2023 3:27:26 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package durianpay

/*
Payloads
*/

/*
Options
*/

// SettlementOption is parameter for Fetch and Details API.
type SettlementOption struct {
	From  string `url:"from"`
	To    string `url:"to"`
	Skip  uint16 `url:"skip"`
	Limit uint16 `url:"limit"`
}
