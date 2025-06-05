package api

type FfmpegBuilder struct{}

func (this *FfmpegBuilder) WithGlobalFlag(
	flag string,
) *FfmpegBuilder {
}

func (this *FfmpegBuilder) WithGlobalOption(
	option, value string,
) *FfmpegBuilder {
}

func (this *FfmpegBuilder) With
