package argparse

import (
	"testing"
)

var expectedTokens = map[string]*ContextFreeTokenType{
	"--option=value":      &CfGnuExplicitAssignment,
	"--option=12":         &CfGnuExplicitAssignment,
	"--long-option=value": &CfGnuExplicitAssignment,
	"--long_option=value": &CfGnuExplicitAssignment,
	"--long_option=12":    &CfGnuExplicitAssignment,
	"--po=TOTO_to":        &CfGnuExplicitAssignment,
	"-option=value":       &CfX2lktExplicitAssignment,
	"-option=12":          &CfX2lktExplicitAssignment,
	"-long-option=value":  &CfX2lktExplicitAssignment,
	"-long_option=value":  &CfX2lktExplicitAssignment,
	"+option":             &CfX2lktReverseSwitch,
	"+long-option":        &CfX2lktReverseSwitch,
	"+long_option":        &CfX2lktReverseSwitch,
	"--":                  &CfEndOfOptions,
	"-o":                  &CfOneDashLetter,
	"-ns.flag":            &CfOneDashWord, // go cli style namespaced flags
	"-n3":                 &CfOneDashWord, // Typical stick value assignment
	"-n12":                &CfOneDashWord,
	"--option":            &CfTwoDashWord,
	"--long-option":       &CfTwoDashWord,
	"-_not_an_option":     &CfWord,
	"--_not_an_option":    &CfWord,
	"word":                &CfWord,
	"word with spaces":    &CfWord,
}

func TestParseArgument(t *testing.T) {
	for arg, expectedType := range expectedTokens {
		foundType := *ParseArgument(arg)
		if expectedType.Name() != foundType.Name() {
			t.Errorf("ParseArgument didn't convert argument %v to expected type %v but to type %v", arg, expectedType, foundType)
		}
	}
}
