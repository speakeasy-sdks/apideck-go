// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package apidecksamplesdk

type ats struct {
	Applicants   *atsApplicants
	Applications *atsApplications
	Jobs         *atsJobs

	sdkConfiguration sdkConfiguration
}

func newAts(sdkConfig sdkConfiguration) *ats {
	return &ats{
		sdkConfiguration: sdkConfig,
		Applicants:       newAtsApplicants(sdkConfig),
		Applications:     newAtsApplications(sdkConfig),
		Jobs:             newAtsJobs(sdkConfig),
	}
}