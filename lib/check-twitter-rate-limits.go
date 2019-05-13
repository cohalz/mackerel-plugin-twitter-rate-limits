package twitterratelimits

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/cohalz/anaconda"
)

// Do the plugin
func Do() {

	api := anaconda.NewTwitterApiWithCredentials(

		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
	)

	resources := strings.Split(os.Getenv("API_RESOURCES"), ",")

	families := make([]string, len(resources))

	for i := range resources {
		families[i] = strings.Split(resources[i], "/")[1]
	}

	response, _ := api.GetRateLimits(families)
	now := time.Now().Unix()

	for i := range resources {
		resource := response.Resources[families[i]][resources[i]]

		r := strings.NewReplacer("/", "_", ":", "")
		metricName := fmt.Sprintf("twitter.rate_limits.%s", r.Replace(resources[i][1:]))

		fmt.Printf("%s\t%d\t%d\n", metricName, resource.Remaining, now)
	}

}
