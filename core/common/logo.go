package common

import (
	"fmt"
	"log"
	"os"
)

const (
	LOGO_PATH = "hack/ascii.txt"
)

type Logo string

type LogoConfigurator func(*Logo) error

func DisplayLogo(lc ...LogoConfigurator) error {

	var logo Logo = LOGO_PATH

	for _, c := range lc {
		err := c(&logo)

		if err != nil {
			return err
		}
	}

	ascii, err := os.ReadFile(logo.String())
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(string(ascii))
	return nil
}

func (l *Logo) String() string {
	return string(*l)
}

func WithLogoPath(l string) LogoConfigurator {
	return func(logo *Logo) error {
		*logo = Logo(l)
		return nil
	}
}
