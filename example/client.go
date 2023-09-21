/*
 * File Created: Thursday, 21st September 2023 5:08:28 pm
 * Author: Abdul Hamid (abdul.surel@gmail.com)
 *
 * Copyright (c) 2023 Author
 */
package example

import (
	"context"

	"github.com/abmid/dpay-sdk-go/client"
)

var ctx = context.Background()

var c = client.NewClient(client.Options{
	ServerKey: "XXX-XXX",
})
