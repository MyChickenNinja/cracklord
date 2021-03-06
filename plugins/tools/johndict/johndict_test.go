package johndict

import (
	"bytes"
	"encoding/csv"
	"reflect"
	"strings"
	"testing"
)

func TestParsingFormats(t *testing.T) {
	var formats = `descrypt, bsdicrypt, md5crypt, bcrypt, scrypt, LM, AFS, tripcode, dummy,
dynamic_n, agilekeychain, aix-ssha1, aix-ssha256, aix-ssha512, asa-md5,
bfegg, Bitcoin, blackberry-es10, WoWSRP, Blockchain, chap, Clipperz,
cloudkeychain, cq, CRC32, sha1crypt, sha256crypt, sha512crypt, Citrix_NS10,
dahua, Django, django-scrypt, dmd5, dmg, dominosec, dragonfly3-32,
dragonfly3-64, dragonfly4-32, dragonfly4-64, Drupal7, eCryptfs, EFS, eigrp,
EncFS, EPI, EPiServer, fde, Fortigate, FormSpring, gost, gpg, HAVAL-128-4,
HAVAL-256-3, hdaa, HMAC-MD5, HMAC-SHA1, HMAC-SHA224, HMAC-SHA256,
HMAC-SHA384, HMAC-SHA512, hMailServer, hsrp, IKE, ipb2, KeePass, keychain,
keyring, keystore, known_hosts, krb4, krb5, krb5pa-sha1, krb5-18, kwallet,
lotus5, lotus85, LUKS, MD2, md4-gen, mdc2, MediaWiki, MongoDB, Mozilla,
mscash, mscash2, MSCHAPv2, mschapv2-naive, krb5pa-md5, mssql, mssql05,
mssql12, mysqlna, mysql-sha1, mysql, nethalflm, netlm, netlmv2, net-md5,
netntlmv2, netntlm, netntlm-naive, net-sha1, nk, md5ns, nsldap, NT, nt2,
o5logon, ODF, Office, oldoffice, OpenBSD-SoftRAID, openssl-enc, oracle,
oracle11, osc, Panama, PBKDF2-HMAC-SHA1, PBKDF2-HMAC-SHA256,
PBKDF2-HMAC-SHA512, PDF, PFX, phpass, PHPS, pix-md5, PKZIP, po, postgres,
PST, PuTTY, pwsafe, RACF, RAdmin, RAKP, rar, RAR5, Raw-SHA512, Raw-Blake2,
Raw-Keccak, Raw-Keccak-256, Raw-MD4, Raw-MD5, Raw-MD5u, Raw-SHA1,
Raw-SHA1-Linkedin, Raw-SHA224, Raw-SHA256, Raw-SHA256-ng, Raw-SHA384,
Raw-SHA512-ng, Raw-SHA, ripemd-128, ripemd-160, rsvp, Siemens-S7,
Salted-SHA1, SSHA512, sapb, sapg, 7z, sha1-gen, Raw-SHA1-ng, SIP, skein-256,
skein-512, skey, aix-smd5, Snefru-128, Snefru-256, LastPass, SSH, SSH-ng,
STRIP, SunMD5, sxc, sybasease, Sybase-PROP, tcp-md5, Tiger, tc_aes_xts,
tc_ripemd160, tc_sha512, tc_whirlpool, OpenVMS, VNC, vtp, wbb3, whirlpool,
whirlpool0, whirlpool1, wpapsk, xsha, xsha512, ZIP, crypt`

	var results = []string{"descrypt", "bsdicrypt", "md5crypt", "bcrypt", "scrypt", "LM", "AFS", "tripcode", "dummy", "dynamic_n", "agilekeychain", "aix-ssha1", "aix-ssha256", "aix-ssha512", "asa-md5", "bfegg", "Bitcoin", "blackberry-es10", "WoWSRP", "Blockchain", "chap", "Clipperz", "cloudkeychain", "cq", "CRC32", "sha1crypt", "sha256crypt", "sha512crypt", "Citrix_NS10", "dahua", "Django", "django-scrypt", "dmd5", "dmg", "dominosec", "dragonfly3-32", "dragonfly3-64", "dragonfly4-32", "dragonfly4-64", "Drupal7", "eCryptfs", "EFS", "eigrp", "EncFS", "EPI", "EPiServer", "fde", "Fortigate", "FormSpring", "gost", "gpg", "HAVAL-128-4", "HAVAL-256-3", "hdaa", "HMAC-MD5", "HMAC-SHA1", "HMAC-SHA224", "HMAC-SHA256", "HMAC-SHA384", "HMAC-SHA512", "hMailServer", "hsrp", "IKE", "ipb2", "KeePass", "keychain", "keyring", "keystore", "known_hosts", "krb4", "krb5", "krb5pa-sha1", "krb5-18", "kwallet", "lotus5", "lotus85", "LUKS", "MD2", "md4-gen", "mdc2", "MediaWiki", "MongoDB", "Mozilla", "mscash", "mscash2", "MSCHAPv2", "mschapv2-naive", "krb5pa-md5", "mssql", "mssql05", "mssql12", "mysqlna", "mysql-sha1", "mysql", "nethalflm", "netlm", "netlmv2", "net-md5", "netntlmv2", "netntlm", "netntlm-naive", "net-sha1", "nk", "md5ns", "nsldap", "NT", "nt2", "o5logon", "ODF", "Office", "oldoffice", "OpenBSD-SoftRAID", "openssl-enc", "oracle", "oracle11", "osc", "Panama", "PBKDF2-HMAC-SHA1", "PBKDF2-HMAC-SHA256", "PBKDF2-HMAC-SHA512", "PDF", "PFX", "phpass", "PHPS", "pix-md5", "PKZIP", "po", "postgres", "PST", "PuTTY", "pwsafe", "RACF", "RAdmin", "RAKP", "rar", "RAR5", "Raw-SHA512", "Raw-Blake2", "Raw-Keccak", "Raw-Keccak-256", "Raw-MD4", "Raw-MD5", "Raw-MD5u", "Raw-SHA1", "Raw-SHA1-Linkedin", "Raw-SHA224", "Raw-SHA256", "Raw-SHA256-ng", "Raw-SHA384", "Raw-SHA512-ng", "Raw-SHA", "ripemd-128", "ripemd-160", "rsvp", "Siemens-S7", "Salted-SHA1", "SSHA512", "sapb", "sapg", "7z", "sha1-gen", "Raw-SHA1-ng", "SIP", "skein-256", "skein-512", "skey", "aix-smd5", "Snefru-128", "Snefru-256", "LastPass", "SSH", "SSH-ng", "STRIP", "SunMD5", "sxc", "sybasease", "Sybase-PROP", "tcp-md5", "Tiger", "tc_aes_xts", "tc_ripemd160", "tc_sha512", "tc_whirlpool", "OpenVMS", "VNC", "vtp", "wbb3", "whirlpool", "whirlpool0", "whirlpool1", "wpapsk", "xsha", "xsha512", "ZIP", "crypt"}

	cleaned := strings.Replace(strings.Replace(formats, "\n", "", -1), " ", "", -1)
	buf := bytes.NewBufferString(cleaned)
	csvReader := csv.NewReader(buf)

	f, err := csvReader.Read()
	if err != nil {
		t.Error("Failed to parse." + err.Error())
	}

	if !reflect.DeepEqual(results, f) {
		t.Error("Parsed data is not equal to known value.")
	}
}

func TestParsingStatus(t *testing.T) {
	/* UNIX
	0g 0:00:00:08 16.01% (ETA: 15:51:27) 0g/s 29181p/s 29181c/s 29181C/s xp14589l..xxgswqas
	0g 0:00:00:18 33.17% (ETA: 15:51:32) 0g/s 29197p/s 29197c/s 29197C/s 958674123..96zinger
	0g 0:00:00:22 40.02% (ETA: 15:51:32) 0g/s 29195p/s 29195c/s 29195C/s 33538683..3441
	0g 0:00:00:44 77.70% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s 19511980..19730246
	0g 0:00:00:45 79.41% (ETA: 15:51:34) 0g/s 29208p/s 29208c/s 29208C/s a7malp1..aalhamed
	0g 0:00:00:46 81.11% (ETA: 15:51:34) 0g/s 29197p/s 29197c/s 29197C/s galqc..garrix
	0g 0:00:00:47 82.81% (ETA: 15:51:34) 0g/s 29203p/s 29203c/s 29203C/s niklas06..nit
	0g 0:00:00:49 86.28% (ETA: 15:51:34) 0g/s 29208p/s 29208c/s 29208C/s 753niky..76777777
	0g 0:00:00:50 88.00% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s cr20323v..crn37
	0g 0:00:00:51 89.70% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s laikinas..lancom
	0g 0:00:00:52 91.40% (ETA: 15:51:34) 0g/s 29212p/s 29212c/s 29212C/s site1122..skiing
	0g 0:00:00:53 93.14% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s 231103..235846
	0g 0:00:00:54 94.84% (ETA: 15:51:34) 0g/s 29215p/s 29215c/s 29215C/s amparo..andersen85
	0g 0:00:00:55 96.57% (ETA: 15:51:34) 0g/s 29212p/s 29212c/s 29212C/s hallol..hannah927
	0g 0:00:00:56 98.28% (ETA: 15:51:34) 0g/s 29220p/s 29220c/s 29220C/s pasmo..patkoe
	*/
	var unixTestOutput = []string{
		"0g 0:00:00:08 16.01% (ETA: 15:51:27) 0g/s 29181p/s 29181c/s 29181C/s xp14589l..xxgswqas",
		"0g 0:00:00:18 33.17% (ETA: 15:51:32) 0g/s 29197p/s 29197c/s 29197C/s 958674123..96zinger",
		"0g 0:00:00:22 40.02% (ETA: 15:51:32) 0g/s 29195p/s 29195c/s 29195C/s 33538683..3441",
		"0g 0:00:00:44 77.70% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s 19511980..19730246",
		"0g 0:00:00:45 79.41% (ETA: 15:51:34) 0g/s 29208p/s 29208c/s 29208C/s a7malp1..aalhamed",
		"0g 0:00:00:46 81.11% (ETA: 15:51:34) 0g/s 29197p/s 29197c/s 29197C/s galqc..garrix",
		"0g 0:00:00:47 82.81% (ETA: 15:51:34) 0g/s 29203p/s 29203c/s 29203C/s niklas06..nit",
		"0g 0:00:00:49 86.28% (ETA: 15:51:34) 0g/s 29208p/s 29208c/s 29208C/s 753niky..76777777",
		"0g 0:00:00:50 88.00% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s cr20323v..crn37",
		"0g 0:00:00:51 89.70% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s laikinas..lancom",
		"0g 0:00:00:52 91.40% (ETA: 15:51:34) 0g/s 29212p/s 29212c/s 29212C/s site1122..skiing",
		"0g 0:00:00:53 93.14% (ETA: 15:51:34) 0g/s 29213p/s 29213c/s 29213C/s 231103..235846",
		"0g 0:00:00:54 94.84% (ETA: 15:51:34) 0g/s 29215p/s 29215c/s 29215C/s amparo..andersen85",
		"0g 0:00:00:55 96.57% (ETA: 15:51:34) 0g/s 29212p/s 29212c/s 29212C/s hallol..hannah927",
		"0g 0:00:00:56 98.28% (ETA: 2019-04-11 21:51) 0g/s 29220p/s 29220c/s 29220C/s pasmo..patkoe",
	}

	// Loop through each example and match it
	for _, s := range unixTestOutput {
		found := regStatusLine.FindStringSubmatch(s)
		if len(found) != 7 {
			t.Errorf("Did not match enough fields, matched %d\n", len(found))
		}
	}

	/* Windows
	0g 0:00:02:01 1.00% (ETA: 01:23:58) 0g/s 3743Kp/s 3743Kc/s 3743KC/s
	0g 0:00:02:01 1.00% (ETA: 01:24:01) 0g/s 3743Kp/s 3743Kc/s 3743KC/s
	0g 0:00:03:01 2.00% (ETA: 00:33:06) 0g/s 3348Kp/s 3348Kc/s 3348KC/s
	0g 0:00:03:01 2.00% (ETA: 00:33:31) 0g/s 3348Kp/s 3348Kc/s 3348KC/s
	0g 0:00:03:01 2.00% (ETA: 00:33:37) 0g/s 3348Kp/s 3348Kc/s 3348KC/s
	0g 0:00:04:01 2.00% (ETA: 01:23:29) 0g/s 3089Kp/s 3089Kc/s 3089KC/s
	0g 0:00:07:01 5.00% (ETA: 00:22:58) 0g/s 2896Kp/s 2896Kc/s 2896KC/s
	0g 0:00:10:01 7.00% (ETA: 00:25:50) 0g/s 3208Kp/s 3208Kc/s 3208KC/s
	*/

	var windowsTestOutput = []string{
		"0g 0:00:02:01 1.00% (ETA: 01:23:58) 0g/s 3743Kp/s 3743Kc/s 3743KC/s",
		"0g 0:00:02:01 1.00% (ETA: 01:24:01) 0g/s 3743Kp/s 3743Kc/s 3743KC/s",
		"0g 0:00:03:01 2.00% (ETA: 00:33:06) 0g/s 3348Kp/s 3348Kc/s 3348KC/s",
		"0g 0:00:03:01 2.00% (ETA: 00:33:31) 0g/s 3348Kp/s 3348Kc/s 3348KC/s",
		"0g 0:00:03:01 2.00% (ETA: 00:33:37) 0g/s 3348Kp/s 3348Kc/s 3348KC/s",
		"0g 0:00:04:01 2.00% (ETA: 01:23:29) 0g/s 3089Kp/s 3089Kc/s 3089KC/s",
		"0g 0:00:07:01 5.00% (ETA: 00:22:58) 0g/s 2896Kp/s 2896Kc/s 2896KC/s",
		"0g 0:00:10:01 7.00% (ETA: 2019-04-11 21:51) 0g/s 3208Kp/s 3208Kc/s 3208KC/s",
	}

	// Loop through each example and match it
	for _, s := range windowsTestOutput {
		found := regStatusLine.FindStringSubmatch(s)
		if len(found) != 7 {
			t.Errorf("[Windows] Did not match enough fields, matched %d\n", len(found))
		}
	}
}

// Test John ETA format parsing

func TestParsingETA(t *testing.T) {
	var timeTestTable = []struct {
		in  string
		out string
	}{
		{"01:24:01", "1 hours, 24 minutes"},
		{"00:25:50", "25 minutes, 49 seconds"},
		{"15:51:34", "15 hours, 51 minutes"},
		// These will always change and should only be enabled during dev
		// {"2019-07-01 14:02", "1443 days, 13 hours"},
		// {"2019-04-11 21:51", "1362 days, 21 hours"},
	}

	for _, v := range timeTestTable {
		// We have a full year time
		nt, _ := parseJohnETA(v.in)

		if printTimeUntil(nt) != v.out {
			t.Errorf(printTimeUntil(nt) + " != " + v.out)
		}

		// d := time.Since(t)
		// println(d.String())
	}
}

var testingRoot = ""

func TestConfigSetup(t *testing.T) {
	err := Setup(testingRoot + "johndict.conf")
	if err != nil {
		t.Error(err.Error())
	}

	// TODO: Configure fails for empty fields
}
