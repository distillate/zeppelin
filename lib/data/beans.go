package data

type Pack struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	//Schedule     string        `json:"schedule"`
	Suites     []Suite    `json:"suites"`
	Properties []Property `json:"properties"`
	//Integrations []Integration `json:"schedule"`
}

// Includes flags, tests in parallel
type Property struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Suite struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Artifact *Artifact `json:"artifact"`
}

// Artifact with tests like Maven or NPM artifact
type Artifact struct {
	Group    string `json:"group"`
	Artifact string `json:"artifact"`
	Version  string `json:"version"`
}

// Stores data for hooks and properties to integrate with external systems, link packs and so on
type Integration struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Properties []Property `json:"properties"`
}

type Launch struct {
	Id     string        `json:"id"`
	Pack   *Pack         `json:"pack,omitempty"`
	State  int           `json:"state"`
	Suites []SuiteLaunch `json:"suites"`
}

// Launch states
const (
	LAUNCH_QUEUED = iota
	LAUNCH_RUNNING
	LAUNCH_FINISHED
	LAUNCH_REVOKED
)

type SuiteLaunch struct {
	Id    string `json:"id"`
	Suite *Suite `json:"suite,omitempty"`
	State int    `json:"state"`
	Log   string `json:"log"` // log file relative url (changes from temporary runtime url to permanent storage url)
}

// Suite states
const (
	SUITE_QUEUED = iota
	SUITE_RUNNING
	SUITE_PASSED
	SUITE_FAILED
	SUITE_REVOKED
)

type ReportLaunch struct {
	Id     string  `json:"id"`
	Pack   *Pack   `json:"pack,omitempty"` // Used to get report specific properties from pack
	Launch *Launch `json:"launch,omitempty"`
	State  int     `json:"state"`
}

// Report states
const (
	REPORT_QUEUED = iota
	REPORT_RUNNING
	REPORT_FINISHED
	REPORT_FAILED
)
