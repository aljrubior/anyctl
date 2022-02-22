package response

type DeploymentLogMessageResponse struct {
	DocId     string `json:"docId"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	ReplicaId string `json:"replicaId"`
	LogLevel  string `json:"logLevel"`
}
