package utils

import (
	"math/rand/v2"
	"net/url"
)

// public api nodes
var PublicNodeUrls = []*url.URL{
	{Host: "https://api.hive.blog"},
	{Host: "https://api.openhive.network"},
	{Host: "https://anyx.io"},
	{Host: "rpc.mahdiyari.info"},
	{Host: "https://techcoderx.com/"},
	{Host: "https://api.deathwing.me/"},
	{Host: "https://api.c0ff33a.uk/"},
	{Host: "https://hive-api.3speak.tv/"},
}

func GetRandomApiUrlFromPublicNode() string {
	dest := make([]*url.URL, len(PublicNodeUrls))
	perm := rand.Perm(len(PublicNodeUrls))
	for i, v := range perm {
		dest[v] = PublicNodeUrls[i]
	}

	return dest[0].Host
}
