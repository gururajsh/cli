package v7

import (
	"fmt"

	"code.cloudfoundry.org/cli/api/uaa/constant"
	"code.cloudfoundry.org/cli/api/uaa/uaaversion"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/command/v7/shared"
)

type AuthCommand struct {
	BaseCommand

	RequiredArgs      flag.Authentication `positional-args:"yes"`
	ClientCredentials bool                `long:"client-credentials" description:"Use (non-user) service account (also called client credentials)"`
	Origin            string              `long:"origin" description:"Indicates the identity provider to be used for authentication"`
	Assertion         string              `long:"assertion" description:"Token based authentication with assertion (user) or in combination with client-credentials (non-user)"`
	usage             interface{}         `usage:"CF_NAME auth USERNAME PASSWORD\n   CF_NAME auth USERNAME PASSWORD --origin ORIGIN\n   CF_NAME auth CLIENT_ID CLIENT_SECRET --client-credentials\n   CF_NAME auth CLIENT_ID CLIENT_SECRET --assertion ID-TOKEN\n   CF_NAME auth CLIENT_ID --client-credentials --assertion ACCESS-TOKEN\n\nENVIRONMENT VARIABLES:\n   CF_USERNAME=user          Authenticating user. Overridden if USERNAME argument is provided.\n   CF_PASSWORD=password      Password associated with user. Overridden if PASSWORD argument is provided.\n\nWARNING:\n   Providing your password as a command line option is highly discouraged\n   Your password may be visible to others and may be recorded in your shell history\n   Consider using the CF_PASSWORD environment variable instead\n\nEXAMPLES:\n   CF_NAME auth name@example.com \"my password\" (use quotes for passwords with a space)\n   CF_NAME auth name@example.com \"\\\"password\\\"\" (escape quotes if used in password)"`
	relatedCommands   interface{}         `related_commands:"api, login, target"`
}

func (cmd AuthCommand) Execute(args []string) error {
	if len(cmd.Origin) > 0 {
		uaaVersion, err := cmd.Actor.GetUAAAPIVersion()
		if err != nil {
			return err
		}

		err = command.MinimumUAAAPIVersionCheck(uaaVersion, uaaversion.MinUAAClientVersion, "Option '--origin'")
		if err != nil {
			return err
		}
	}

	if cmd.ClientCredentials && cmd.Origin != "" {
		return translatableerror.ArgumentCombinationError{
			Args: []string{"--client-credentials", "--origin"},
		}
	}

	username, password, err := cmd.getUsernamePassword()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor(
		"API endpoint: {{.Endpoint}}",
		map[string]interface{}{
			"Endpoint": cmd.Config.Target(),
		})

	versionWarning, err := shared.CheckCCAPIVersion(cmd.Config.APIVersion())
	if err != nil {
		cmd.UI.DisplayWarning("Warning: unable to determine whether targeted API's version meets minimum supported.")
	}
	if versionWarning != "" {
		cmd.UI.DisplayWarning(versionWarning)
	}

	if !cmd.ClientCredentials {
		if cmd.Config.UAAGrantType() == string(constant.GrantTypeClientCredentials) {
			return translatableerror.PasswordGrantTypeLogoutRequiredError{}
		} else if (cmd.Assertion == "" && cmd.Config.UAAOAuthClient() != "cf") || cmd.Config.UAAOAuthClientSecret() != "" {
			return translatableerror.ManualClientCredentialsError{}
		}
	}

	cmd.UI.DisplayNewline()

	cmd.UI.DisplayText("Authenticating...")

	credentials := make(map[string]string)
	grantType := constant.GrantTypePassword
	if cmd.ClientCredentials {
		grantType = constant.GrantTypeClientCredentials
		if cmd.Assertion != "" {
			if username == "" {
				username = cmd.Config.UAAOAuthClient()
			}
			// use assertion as client_assertion - replacing client_secret - but stay with client credentials grant type
			credentials["client_assertion"] = cmd.Assertion
			credentials["client_assertion_type"] = "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"
		} else {
			credentials["client_secret"] = password
		}
		credentials["client_id"] = username
	} else {
		if cmd.Assertion != "" {
			// use assertion as user authentication using JWT bearer grant type
			grantType = constant.GrantTypeJwtBearer
			credentials = map[string]string{
				"assertion": cmd.Assertion,
			}
			if username != "" {
				credentials["client_id"] = username
			}
			if password != "" {
				credentials["client_secret"] = password
			}
		} else {
			credentials = map[string]string{
				"username": username,
				"password": password,
			}
		}
	}

	err = cmd.Actor.Authenticate(credentials, cmd.Origin, grantType)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	cmd.UI.DisplayTextWithFlavor(
		"Use '{{.Command}}' to view or set your target org and space.",
		map[string]interface{}{
			"Command": fmt.Sprintf("%s target", cmd.Config.BinaryName()),
		})

	return nil
}

func (cmd AuthCommand) getUsernamePassword() (string, string, error) {
	var (
		userMissing     bool
		passwordMissing bool
	)

	username := cmd.RequiredArgs.Username
	if username == "" {
		if envUser := cmd.Config.CFUsername(); envUser != "" {
			username = envUser
		} else {
			userMissing = true
		}
	}

	password := cmd.RequiredArgs.Password
	if password == "" {
		if envPassword := cmd.Config.CFPassword(); envPassword != "" {
			password = envPassword
		} else {
			passwordMissing = true
		}
	}

	if cmd.Config.IsCFOnK8s() {
		if !passwordMissing {
			cmd.UI.DisplayWarning("Warning: password is ignored when authenticating against Kubernetes.")
		}
		passwordMissing = false
	}

	if cmd.Assertion == "" && (userMissing || passwordMissing) {
		return "", "", translatableerror.MissingCredentialsError{
			MissingUsername: userMissing,
			MissingPassword: passwordMissing,
		}
	}

	return username, password, nil
}
