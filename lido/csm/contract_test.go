package csm

import (
	"regexp"
	"strings"
	"testing"
)

func TestOperatorNames(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"CSM keeps its historical tag", GetOperatorName(NodeOperatorCustom{Index: 7}), "csm_operator7_lido"},
		{"CMv2 operator", CMv2OperatorName(NodeOperatorCustom{Index: 8}), "lido_cmv2_operator8"},
		{"CSM 0x02 operator", CSM0x02OperatorName(NodeOperatorCustom{Index: 3}), "lido_csm0x02_operator3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}

// sqlLike reports whether s matches a SQL LIKE pattern, where '%' is any run
// of characters and '_' is exactly one character.
func sqlLike(s, pattern string) bool {
	var b strings.Builder
	b.WriteString("^")
	for _, r := range pattern {
		switch r {
		case '%':
			b.WriteString(".*")
		case '_':
			b.WriteString(".")
		default:
			b.WriteString(regexp.QuoteMeta(string(r)))
		}
	}
	b.WriteString("$")
	return regexp.MustCompile(b.String()).MatchString(s)
}

// Downstream consumers select Lido CSM with these patterns and classify pooled
// stake with '%lido%'. The new modules must stay out of the first and inside
// the second, or they get folded into CSM metrics or stop counting as pooled.
func TestNewModuleTagsStayOutOfCSMFilters(t *testing.T) {
	csmPatterns := []string{"csm_%", "csm_%_lido", "csm_operator%_lido"}

	if !sqlLike(GetOperatorName(NodeOperatorCustom{Index: 7}), "csm_%_lido") {
		t.Fatal("sqlLike is broken: a CSM tag must match csm_%_lido")
	}

	tags := []string{
		CMv2EntityTag,
		CSM0x02EntityTag,
		CMv2OperatorName(NodeOperatorCustom{Index: 12}),
		CSM0x02OperatorName(NodeOperatorCustom{Index: 12}),
	}
	for _, tag := range tags {
		if !sqlLike(tag, "%lido%") {
			t.Errorf("%q must match '%%lido%%'", tag)
		}
		for _, pattern := range csmPatterns {
			if sqlLike(tag, pattern) || sqlLike(strings.ToLower(tag), pattern) {
				t.Errorf("%q must not match %q", tag, pattern)
			}
		}
		// Unlike the older Lido tags, these do not end in "_lido" (LIKE
		// '%\_lido'). No consumer matches on that suffix: BindEth and the
		// website take Lido as '%lido%' and CSM as the patterns above. Checked
		// so that adopting the suffix is a decision, since it would put the
		// operator tags inside a future '%\_lido' filter.
		if strings.HasSuffix(tag, "_lido") {
			t.Errorf("%q ends in _lido: check the consumers before changing the convention", tag)
		}
	}
}
