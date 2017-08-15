package buildinfo

type BuildInfo struct {
	TestCoverage float32 `json:"coverage"`
	ApiVersion   float32 `json:"apiversion"`
	SwaggerLink  string  `json:"swaggerlink"`
	BuildTime    float32 `json:"buildtime"`
}

type BuildInfos []BuildInfo
