package types

import "strings"

// Reference: https://www.ietf.org/rfc/rfc4120.txt
// Section: 5.2.2

type PrincipalName struct {
	NameType   int      `asn1:"explicit,tag:0"`
	NameString []string `asn1:"generalstring,explicit,tag:1"`
}

func (pn *PrincipalName) GetSalt(realm string) string {
	var sb []byte
	sb = append(sb, realm...)
	for _, n := range pn.NameString {
		sb = append(sb, n...)
	}
	return string(sb)
}

func (pn *PrincipalName) Equal(n PrincipalName) bool {
	if n.NameType != pn.NameType {
		return false
	}
	for i, s := range pn.NameString {
		if n.NameString[i] != s {
			return false
		}
	}
	return true
}

func (pn *PrincipalName) GetPrincipalNameString() string {
	return strings.Join(pn.NameString, "/")
}
