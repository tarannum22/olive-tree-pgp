package jobs

type Job struct {
	Name        string
	JobType     JobType
	Keyfile     string
	Passphrase  string
	Source      string
	Destination string
	Schedule    string
}

type JobType int

const (
	ENCRYPTION JobType = 1
	DECRYPTION JobType = 0
)

func InitJobs() {

}
