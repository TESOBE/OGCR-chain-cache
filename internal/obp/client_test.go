package obp

import "testing"

func TestURLsFollowTheSpace(t *testing.T) {
	bank := NewClient("http://obp.test/", "u", "p", "k", "ogcr")
	if got, want := bank.entityURL("chain_sync_status", "", nil), "http://obp.test/obp/dynamic-entity/banks/ogcr/chain_sync_status"; got != want {
		t.Errorf("bank entityURL = %q, want %q", got, want)
	}
	if got, want := bank.managementURL(), "http://obp.test/obp/v7.0.0/management/banks/ogcr/dynamic-entities"; got != want {
		t.Errorf("bank managementURL = %q, want %q", got, want)
	}

	sys := NewClient("http://obp.test", "u", "p", "k", "")
	if got, want := sys.entityURL("parcel_on_chain", "/r1", nil), "http://obp.test/obp/dynamic-entity/parcel_on_chain/r1"; got != want {
		t.Errorf("system entityURL = %q, want %q", got, want)
	}
	if got, want := sys.managementURL(), "http://obp.test/obp/v7.0.0/management/banks/SYS/dynamic-entities"; got != want {
		t.Errorf("system managementURL = %q, want %q", got, want)
	}
}
