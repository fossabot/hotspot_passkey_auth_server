package wa

import (
	// "crypto/ecdsa"
	// "crypto/elliptic"
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	// "github.com/tstranex/u2f"
)

type User struct {
	ID                string
	Name              string                `json:"name"`
	DisplayName       string                `json:"display_name"`
	Icon              string                `json:"icon,omitempty"`
	Credentials       []webauthn.Credential `json:"credentials,omitempty"`
}

func (u User) WebAuthnHasU2F() bool {
	for _, cred := range u.Credentials {
		if cred.AttestationType == "fido-u2f" {
			return true
		}
	}

	return false
}

func (u User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u User) WebAuthnName() string {
	return u.Name
}

func (u User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u User) WebAuthnIcon() string {
	return u.Icon
}

func (u User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u User) WebAuthnCredentialDescriptors() (descriptors []protocol.CredentialDescriptor) {
	descriptors = make([]protocol.CredentialDescriptor, len(u.Credentials))

	for i, credential := range u.Credentials {
		descriptors[i] = credential.Descriptor()
	}

	return descriptors
}

// func (u User) U2FRegistrations() (registrations []u2f.Registration) {
// 	registrations = []u2f.Registration{}

// 	for _, credential := range u.Credentials {
// 		if credential.AttestationType != "fido-u2f" {
// 			continue
// 		}

// 		x, y := elliptic.Unmarshal(elliptic.P256(), credential.PublicKey)

// 		registrations = append(registrations, u2f.Registration{
// 			KeyHandle: credential.ID,
// 			PubKey: ecdsa.PublicKey{
// 				Curve: elliptic.P256(),
// 				X:     x,
// 				Y:     y,
// 			},
// 		})
// 	}

// 	return registrations
// }
