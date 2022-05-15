package seimei

import (
	"errors"
	"fmt"

	// Using embed.
	_ "embed"

	"github.com/spf13/cobra"
)

var (
	//go:embed version.txt
	VersionText           string
	ErrEmptyName          = errors.New("provide name is empty (ex. 田中太郎)")
	ErrInvalidName        = errors.New("provide path is invalid")
	ErrEmptyPath          = errors.New("provide path is empty (ex. /tmp/foo.csv)")
	ErrInvalidPath        = errors.New("provide path is invalid")
	ErrInvalidParseString = errors.New("provide parse string is invalid")
)

type CmdMode string

func (c CmdMode) String() string {
	return string(c)
}

const (
	NameCmd     CmdMode = "name"
	FileCmd     CmdMode = "file"
	ParseOption string  = "parse"
)

func BuildMainCmd() *cobra.Command {
	c := cobra.Command{
		Use: "seimei",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Usage()
		},
		Version: fmt.Sprintf(VersionText),
	}
	cobra.EnableCommandSorting = false
	c.AddCommand(BuildNameCmd())
	c.AddCommand(BuildFileCmd())
	return &c
}

func BuildNameCmd() *cobra.Command {
	c := cobra.Command{
		Use:   "name",
		Short: "It parse single full name.",
		Long: `It parse single full name.
Provide the full name to be parsed with the required flag (--name).
`,
		Example: "seimei name --name 田中太郎",
		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := detectFlagForName(cmd)
			if err != nil {
				return fmt.Errorf("flag parse error: %w", err)
			}
			p, err := detectFlagParseString(cmd)
			if err != nil {
				return fmt.Errorf("flag parse error: %w", err)
			}
			return ParseName(cmd.OutOrStdout(), cmd.OutOrStderr(), n, p)
		},
	}
	c.Flags().SortFlags = false
	c.Flags().StringP(NameCmd.String(), "n", "", "田中太郎")
	err := c.MarkFlagRequired(NameCmd.String())
	// since name flag is set on above, it raise panic without returning an error.
	if err != nil {
		panic(err)
	}
	c.Flags().StringP(ParseOption, "p", " ", " ")
	return &c
}

func BuildFileCmd() *cobra.Command {
	c := cobra.Command{
		Use:   "file",
		Short: "It bulk parse full name lit in the file.",
		Long: `It bulk parse full name lit in the file.
Provide the file path with full name list to the required flag (--file).
`,
		Example: "seimei file --file /path/to/dir/foo.csv",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, err := detectFlagForFile(cmd)
			if err != nil {
				return fmt.Errorf("flag parse error: %w", err)
			}
			p, err := detectFlagParseString(cmd)
			if err != nil {
				return fmt.Errorf("flag parse error: %w", err)
			}
			return ParseFile(cmd.OutOrStdout(), cmd.ErrOrStderr(), f, p)
		},
	}
	c.Flags().SortFlags = false
	c.Flags().StringP(FileCmd.String(), "f", "", "/path/to/dir/foo.csv")
	err := c.MarkFlagRequired(FileCmd.String())
	// since file flag is set on above, it raise panic without returning an error.
	if err != nil {
		panic(err)
	}
	c.Flags().StringP(ParseOption, "p", " ", " ")
	return &c
}

func Run() error {
	cmd := BuildMainCmd()
	return cmd.Execute()
}

func detectFlagForName(cmd *cobra.Command) (Name, error) {
	n, err := cmd.Flags().GetString(NameCmd.String())
	if err != nil {
		return "", ErrInvalidName
	}
	if n == "" {
		return "", ErrEmptyName
	}

	return Name(n), nil
}

func detectFlagForFile(cmd *cobra.Command) (Path, error) {
	n, err := cmd.Flags().GetString(FileCmd.String())
	if err != nil {
		return "", ErrInvalidPath
	}
	if n == "" {
		return "", ErrEmptyPath
	}

	return Path(n), nil
}

func detectFlagParseString(cmd *cobra.Command) (ParseString, error) {
	p, err := cmd.Flags().GetString(ParseOption)
	if err != nil {
		return "", ErrInvalidParseString
	}
	return ParseString(p), nil
}
