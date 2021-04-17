package model

type Fixtures struct {
	Branding []Branding
	Contacts []Contact
	Keys     []Key
}

func InitFixtures() *Fixtures {
	// generate API keys
	keys, err := InitKey()
	if err != nil {
		panic("failed creating API keys")
	}

	contacts, err := InitContact()
	if err != nil {
		panic("failed creating contacts")
	}

	branding, err := InitBranding()
	if err != nil {
		panic("failed creating branding options")
	}

	return &Fixtures{
		Branding: branding,
		Contacts: contacts,
		Keys:     keys,
	}

	// TODO: envelope fixtures
	// TODO: settings fixtures
	// TODO: team fixtures
	// TODO: template fixtures
	// TODO: user fixtures
	// TODO: webhook fixtures
}
