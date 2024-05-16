package keycloak

type KeycloakConfig struct {
	Server Server
	Client ClientInformations
	Admin  AdminInformations
}

type Server struct {
	Host string
	Port string
}

type ClientInformations struct {
	ClientID     string
	ClientSecret string
	Realm        string
}

type AdminInformations struct {
	Username string
	Password string
	Realm    string
}

func NewKeycloakConfig(host, port, clientID, clientSecret, realm, adminUsername, adminPassword, adminRealm string) *KeycloakConfig {
	return &KeycloakConfig{
		Server: Server{
			Host: host,
			Port: port,
		},
		Client: ClientInformations{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Realm:        realm,
		},
		Admin: AdminInformations{
			Username: adminUsername,
			Password: adminPassword,
			Realm:    adminRealm,
		},
	}
}
