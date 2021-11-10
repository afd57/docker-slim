# Docker-Slim API Design Documentation



## 1. /v1/commands/build

Run `build` command in docker-slim via restapi.

`POST /v1/commands/build`

**Query parameters**:

- targetRef string,

- doPull bool,

- dockerConfigPath string,

- registryAccount string,

- registrySecret string,

- doShowPullLogs bool,

- composeFile string,

- targetComposeSvc string,

- composeSvcNoPorts bool,

- depExcludeComposeSvcAll bool,

- depIncludeComposeSvcDeps string,

- depIncludeTargetComposeSvcDeps bool,

- depIncludeComposeSvcs []string,

- depExcludeComposeSvcs []string,

- composeNets []string,

- composeEnvVars []string,

- composeEnvNoHost bool,

- composeWorkdir string,

- composeProjectName string,

- cbOpts *config.ContainerBuildOptions,

- crOpts *config.ContainerRunOptions,

- outputTags []string,

- doHTTPProbe bool,

- httpProbeCmds []config.HTTPProbeCmd,

- httpProbeRetryCount int,

- httpProbeRetryWait int,

- httpProbePorts []uint16,

- httpCrawlMaxDepth int,

- httpCrawlMaxPageCount int,

- httpCrawlConcurrency int,

- httpMaxConcurrentCrawlers int,

- doHTTPProbeFull bool,

- doHTTPProbeExitOnFailure bool,

- httpProbeAPISpecs []string,

- httpProbeAPISpecFiles []string,

- httpProbeApps []string,

- portBindings map[dockerapi.Port][]dockerapi.PortBinding,

- doPublishExposedPorts bool,

- doRmFileArtifacts bool,

- copyMetaArtifactsLocation string,

- doRunTargetAsUser bool,

- doShowContainerLogs bool,

- doShowBuildLogs bool,

- imageOverrideSelectors map[string]bool,

- overrides *config.ContainerOverrides,

- instructions *config.ImageNewInstructions,

- links []string,

- etcHostsMaps []string,

- dnsServers []string,

- dnsSearchDomains []string,

- explicitVolumeMounts map[string]config.VolumeMount,

- doKeepPerms bool,

- pathPerms map[string]*fsutil.AccessInfo,

- excludePatterns map[string]*fsutil.AccessInfo,

- preservePaths map[string]*fsutil.AccessInfo,

- includePaths map[string]*fsutil.AccessInfo,

- includeBins map[string]*fsutil.AccessInfo,

- includeExes map[string]*fsutil.AccessInfo,

- doIncludeShell bool,

- doIncludeCertAll bool,

- doIncludeCertBundles bool,

- doIncludeCertDirs bool,

- doIncludeCertPKAll bool,

- doIncludeCertPKDirs bool,

- doUseLocalMounts bool,

- doUseSensorVolume string,

- doKeepTmpArtifacts bool,

- continueAfter *config.ContinueAfter,

- execCmd string,

- execFileCmd string,

- deleteFatImage bool



**Example response**:

```json
  HTTP/1.1 201 Created
  Content-Type: application/json

  {
       "ExecId":"1",
       "Warnings":[],
    	 "Time": Time,
    	 "Message": 
  }

HTTP/1.1 400 BadRequest
Content-Type: application/json

  {
    "Q": [
      "TargetRef",
      "IncludePaths"
    ]
  }
```



## 2. /v1/commands/build/\<ExecId\>

Return result of `build/<ExecId>` command.

`GET /v1/commands/build/1`

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
  "FatImage": "",
  "FatImageSHA": "",
  "SlimImage": "",
  "SlimImageSHA": "",
  "ExecutionTime": "",
  "Message": ""
}
```

