/*
Copyright 2025 Rodolfo González González

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"golang.org/x/net/publicsuffix"
)

//-----------------------------------------------------------------------------

// Gets the whois server from a static list (which wmight need to be updated someday).
// TODO: pass the path to a list in a file using a CLI parameter.

type WhoisServers struct {
	WhoisServers map[string]string
}

// The big list of servers. Feel free to add or change any of them, and drop
// an issue at https://github.com/rgglez/nagios-check-domain/issues

func NewWhoisServers() *WhoisServers {
	return &WhoisServers{
		WhoisServers: map[string]string{
			"ac": "whois.nic.ac",
			"ad": "whois.ripe.net",
			"ae": "whois.aeda.net.ae",
			"aero": "whois.aero",
			"af": "whois.nic.af",
			"ag": "whois.nic.ag",
			"ai": "whois.ai",
			"al": "whois.ripe.net",
			"am": "whois.amnic.net",
			"ao": "whois.nic.ao",
			"aq": "whois.ripe.net",
			"ar": "whois.nic.ar",
			"arpa": "whois.iana.org",
			"as": "whois.nic.as",
			"asia": "whois.nic.asia",
			"at": "whois.nic.at",
			"au": "whois.aunic.net",
			"aw": "whois.nic.aw",
			"ax": "whois.ax",
			"az": "whois.ripe.net",
			"ba": "whois.ripe.net",
			"bb": "whois.bb",
			"bd": "whois.btcl.net.bd",
			"be": "whois.dns.be",
			"bf": "whois.nic.bf",
			"bg": "whois.register.bg",
			"bh": "whois.bh",
			"bi": "whois.nic.bi",
			"biz": "whois.biz",
			"bj": "whois.nic.bj",
			"bn": "whois.bn",
			"bo": "whois.nic.bo",
			"br": "whois.registro.br",
			"bt": "whois.netnames.net",
			"bw": "whois.nic.net.bw",
			"by": "whois.cctld.by",
			"bz": "whois.belizenic.bz",
			"ca": "whois.cira.ca",
			"cat": "whois.nic.cat",
			"cc": "whois.nic.cc",
			"cd": "whois.nic.cd",
			"cf": "whois.dot.cf",
			"cg": "whois.cg",
			"ch": "whois.nic.ch",
			"ci": "whois.nic.ci",
			"ck": "whois.nic.ck",
			"cl": "whois.nic.cl",
			"cm": "whois.cm",
			"cn": "whois.cnnic.net.cn",
			"co": "whois.nic.co",
			"com": "whois.verisign-grs.com",
			"coop": "whois.nic.coop",
			"cr": "whois.nic.cr",
			"cu": "whois.cu",
			"cv": "whois.dns.cv",
			"cw": "whois.cw",
			"cx": "whois.nic.cx",
			"cy": "whois.ripe.net",
			"cz": "whois.nic.cz",
			"de": "whois.denic.de",
			"dj": "whois.nic.dj",
			"dk": "whois.dk-hostmaster.dk",
			"dm": "whois.nic.dm",
			"do": "whois.nic.do",
			"dz": "whois.nic.dz",
			"ec": "whois.nic.ec",
			"edu": "whois.educause.edu",
			"ee": "whois.tld.ee",
			"eg": "whois.ripe.net",
			"es": "whois.nic.es",
			"eu": "whois.eu",
			"fi": "whois.fi",
			"fj": "whois.usp.ac.fj",
			"fm": "whois.nic.fm",
			"fr": "whois.nic.fr",
			"ga": "whois.nic.ga",
			"gd": "whois.nic.gd",
			"ge": "whois.nic.ge",
			"gg": "whois.gg",
			"gi": "whois.nic.gi",
			"gl": "whois.nic.gl",
			"gm": "whois.nic.gm",
			"gov": "whois.nic.gov",
			"gr": "whois.ripe.net",
			"gs": "whois.nic.gs",
			"gt": "whois.gt",
			"gu": "whois.gu",
			"gy": "whois.registry.gy",
			"hk": "whois.hkirc.hk",
			"hm": "whois.registry.hm",
			"hn": "whois.nic.hn",
			"hr": "whois.dns.hr",
			"ht": "whois.nic.ht",
			"hu": "whois.nic.hu",
			"id": "whois.id",
			"ie": "whois.domainregistry.ie",
			"il": "whois.isoc.org.il",
			"im": "whois.nic.im",
			"in": "whois.registry.in",
			"info": "whois.afilias.info",
			"int": "whois.iana.org",
			"io": "whois.nic.io",
			"iq": "whois.cmc.iq",
			"ir": "whois.nic.ir",
			"is": "whois.isnic.is",
			"it": "whois.nic.it",
			"je": "whois.je",
			"jm": "whois.jm",
			"jo": "whois.jo",
			"jobs": "jobswhois.verisign-grs.com",
			"jp": "whois.jprs.jp",
			"ke": "whois.kenic.or.ke",
			"kg": "whois.kg",
			"kr": "whois.kr",
			"kz": "whois.nic.kz",
			"la": "whois.nic.la",
			"lb": "whois.lb",
			"lc": "whois.nic.lc",
			"li": "whois.nic.li",
			"lk": "whois.lk",
			"lt": "whois.domreg.lt",
			"lu": "whois.dns.lu",
			"lv": "whois.nic.lv",
			"ly": "whois.nic.ly",
			"ma": "whois.registre.ma",
			"md": "whois.nic.md",
			"me": "whois.nic.me",
			"mg": "whois.nic.mg",
			"mil": "whois.nic.mil",
			"mk": "whois.marnet.mk",
			"ml": "whois.ml",
			"mm": "whois.registry.mm",
			"mn": "whois.nic.mn",
			"mo": "whois.monic.mo",
			"mobi": "whois.dotmobiregistry.net",
			"mp": "whois.mp",
			"mq": "whois.mq",
			"mr": "whois.nic.mr",
			"ms": "whois.nic.ms",
			"mt": "whois.nic.mt",
			"mu": "whois.nic.mu",
			"mv": "whois.mv",
			"mw": "whois.nic.mw",
			"mx": "whois.mx",
			"my": "whois.my",
			"mz": "whois.nic.mz",
			"na": "whois.na",
			"name": "whois.nic.name",
			"net": "whois.verisign-grs.com",
			"ng": "whois.nic.net.ng",
			"ni": "whois.nic.ni",
			"nl": "whois.domain-registry.nl",
			"no": "whois.norid.no",
			"np": "whois.np",
			"nr": "whois.nic.nr",
			"nu": "whois.nic.nu",
			"nz": "whois.srs.net.nz",
			"om": "whois.registry.om",
			"org": "whois.pir.org",
			"pa": "whois.nic.pa",
			"pe": "kero.yachay.pe",
			"pf": "whois.registry.pf",
			"ph": "whois.nic.ph",
			"pk": "whois.pknic.net.pk",
			"pl": "whois.dns.pl",
			"pm": "whois.nic.pm",
		},
	}
}

// GetWhoisServer finds the WHOIS server for a given domain
func (w *WhoisServers) GetWhoisServer(domain string) (string, bool) {
	// Get the public suffix for the given domain
	tld, _ := publicsuffix.PublicSuffix(domain)
	
	// Return the server, if any, from the list, or exists=false if 
	// it was not found
	server, exists := w.WhoisServers[tld]

	return server, exists
}