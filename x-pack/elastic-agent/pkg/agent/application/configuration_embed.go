// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildfleetcfg/buildfleetcfg.go - DO NOT EDIT.

package application

import "github.com/JitendraKSahu/beats/v7/x-pack/elastic-agent/pkg/packer"

// DefaultAgentFleetConfig is the content of the default configuration when we enroll a beat, the elastic-agent.yml
// will be replaced with this variables.
var DefaultAgentFleetConfig []byte

func init() {
	// Packed File
	// _meta/elastic-agent.fleet.yml
	unpacked := packer.MustUnpack("eJycVUuTo7oV3udnzD43PJqpIVVZWHQjxLTpGF9LSJsUkjwytoSp2GCLVP57SuC2u29Npabugo2QzuN7nPOfL/8y23P9t62uT+dG/LVW2/b82w+93Z5/s0Z/+fsXZBfxP3/3fu1bg4yRSAuTntDzL755fN+RBQkjqScs6IQFex5EpibSFwYfUIZ7BrFFL+zECPZeExDxcKNokPYIlpoZ3bM1OPOw8FBWaJmVHTdyTJqlYm0+8DXwarhR60A3lEQ7lOL8tQHsHiPDFwZTj0339Igg6zjcKAm/KRrEPTO6lVWuUVYe2RqMrCq9mkStsO5+6rk81P2H0SATMMoKXHiYe4ysFG0PqjRXzapVn5i5/jd1VAjinpL8xMjqK8pKK8nm+xSHSM3bVc9gHHJzHShZfXXnKFnsUVYMssr3bA0G1sw90Sr3asJ2NCzH1wRYVqV+XeVaWKA5TEcJ9R7Ba1cHGyVCl6fwpl5g2rME9JT4HTdC8YAqCXcawbzjJrU1weM9L9z0zC4GGeipLmH9p7dn2m+Nvsx3wMhDbGmA8eZw+IqS/Ciz8iLG4/AapFZCbSgpPGGje02vphheg3KQQXTiQXoQNm5m/r81j37BTkI18TDXVnbC4L2Esd2uJ414lOjecYHgI5bDXcB0Xwdpy6rld2QXCmXlThipZRofHD40LAfRulrRYRtOeXqnAVp5wwP7ckTquEfJQiHHNfE1D7GHIBtEAw48lD2H8Y5B3TMLLpQU/3b83TC5339rFuMyu9WRLBUz+sTW4EKr8oiyckDwRYkMN3zi6qy3a+CLAB9ueQJW5R1r9cTTh9o0S8BOtGXHTLqXH+LXhCpuYu9dx3W1VCLLtQhwLxOQ3nXxkraMRB7KZMfhRcksj+Z6Jt39sSYjTHxGEJ9oVXg1WfY0+KZkgBsRaI+tn244g1bMeX+OLbnuRFh21EYXFi4eXEPdi7DccXj9IWDq1Qk4sKrYC6MbVi1v/F8HGqReTeIewdh8eOOwuOMvbHwRJm6FSc/CPql1tZowoOQ6cjtj+8g7YWwZSQ+TjjJgeSAtJZ4SIdA00KYmxU5CPfB22d+47XhbeJRcT9Xc3097rUk0SpieeII+eVtk+UADPIrxeOeMTzr2dzyZZweCkc9JrkUDDA+Rq2Ung2jqU5h4z6pivGsb+h3X8fTurVlcl8+Ly3Jx1+2ZVruOV/iMMjCIdqV4i8/UYHvrfRQwDfk0Z13vhRZhocVlfs9J+lSF07uv6Hl5cbEf/iw7TvAgq5ViJrYIFg5Hb/ZjpKX9VGuPMnlk5End+59nfc/DlRIh3k9nAfM5vEbvsd5nqIS7Tnzw4IeZ6LkZiGB8YVU+89Tcd8hYk1JPOwS+KGbSjmfYstXc2w3THxLqMyOxL5+PajmCMXGeh7M3XttcO1+8PeaA5iZtOHQxSy0hto67SaPBTjudTnqqSjejXB0XHhZj7Wb9+kmVEBta4ZNMpn/GafLunZ/shHd9CIM9YbQWfrzn4eTjUUC8rwnrXO4Hp7FB2VRzN2H+4vZj6vNs9sDSPt1jblzM9uB2nM+dLgKqlmtwZiTtHTcyAWFNricE3e67xSTuXqRn7jeKEuaxCs1nkO1qcvWF2bz7ZMLB7ZIqKAZJomkWPrxXun0Qoay4MFJ0zPVnwcHNCIcdb8BZWNCwqgwZwf3/7cvxNuec3/9+/HDmcvm7bRq/x/6MHUmfauL7fP1ncy8un/wSzH29NQufr+66SbcZGLjBvYTa+fzAKqdNt3/xOHmjKo8OSzbN2MMUW2a5ptVt1lugP8d4cIkD7dXJtGfu7ymJDuj5g3ZGh8nihJ7p5TVZXJf7D/j8oSdm0pMINu+6fBIw7p0fapKe3pqbblfHf3z571/+FwAA//9N0n8p")
	raw, ok := unpacked["_meta/elastic-agent.fleet.yml"]
	if !ok {
		// ensure we have something loaded.
		panic("elastic-agent.fleet.yml is not included in the binary")
	}
	DefaultAgentFleetConfig = raw
}
