package types

import "strings"

//move setupacl stuff here

func NewRegistryRole(name string, tokenType string, policies []Policy, localUse bool) RegistryRole {
	// to conform to the payload of the registry create role API,
	// we convert the slice of policies from type Policy to string and make it unique
	// as the policy name needs to be unique per API's requirement
	policyNames := make([]string, 0, len(policies))
	tempMap := make(map[string]bool)
	for _, policy := range policies {
		if _, exists := tempMap[policy.Name]; !exists {
			policyNames = append(policyNames, policy.Name)
		}
	}

	return RegistryRole{
		RoleName:    strings.TrimSpace(name),
		TokenType:   string(tokenType),
		PolicyNames: policyNames,
		Local:       localUse,
		// unlimited for now
		TimeToLive: "0s",
	}
}
