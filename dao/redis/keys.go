package redis

const (
	Prefix           = "linkux:"
	KeyPostTimeZset  = "post:time"
	KeyPostScoreZset = "post:score"

	KeyPostVotedZsetPF = "post:voted:"
	KeyLabelSetPF      = "label:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
