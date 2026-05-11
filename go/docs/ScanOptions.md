# ScanOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **string** |  | 
**Excludes** | Pointer to **string** |  | [optional] 
**ScanName** | Pointer to **string** |  | [optional] 
**ScanDescription** | Pointer to **string** | A description of the scan. | [optional] 
**ScanTemplate** | Pointer to **string** |  | [optional] 
**ScanFrequency** | Pointer to **string** | A string time duration value representing execution frequency, if scheduled to repeat. | [optional] 
**ScanStart** | Pointer to **string** | Unix timestamp value indicating when the template was created. | [optional] 
**ScanTags** | Pointer to **string** |  | [optional] 
**ScanGracePeriod** | Pointer to **string** |  | [optional] 
**Agent** | Pointer to **string** |  | [optional] 
**Explorer** | Pointer to **string** |  | [optional] 
**ExplorerGroupId** | Pointer to **string** |  | [optional] 
**HostedZoneId** | Pointer to **string** | The string &#39;auto&#39; will use any available hosted zone. Otherwise, provide the string name (hostedzone1) or UUID (\&quot;e77602e0-3fb8-4734-aef9-fbc6fdcb0fa8\&quot;) of a hosted zone. | [optional] 
**HostedZoneName** | Pointer to **string** | The string &#39;auto&#39; will use any available hosted zone. Otherwise, provide the string name (hostedzone1) of the hosted zone. | [optional] 
**Rate** | Pointer to **string** |  | [optional] 
**MaxHostRate** | Pointer to **string** |  | [optional] 
**Passes** | Pointer to **string** |  | [optional] 
**MaxAttempts** | Pointer to **string** |  | [optional] 
**MaxSockets** | Pointer to **string** |  | [optional] 
**MaxGroupSize** | Pointer to **string** |  | [optional] 
**MaxTtl** | Pointer to **string** |  | [optional] 
**Tos** | Pointer to **string** |  | [optional] 
**TcpPorts** | Pointer to **string** |  | [optional] 
**TcpExcludes** | Pointer to **string** |  | [optional] 
**Screenshots** | Pointer to **string** |  | [optional] 
**Nameservers** | Pointer to **string** |  | [optional] 
**SubnetPing** | Pointer to **string** |  | [optional] 
**SubnetPingNetSize** | Pointer to **string** |  | [optional] 
**SubnetPingProbes** | Pointer to **string** | Optional subnet ping probe list as comma separated strings. The example shows possibilities. | [optional] 
**SubnetPingSampleRate** | Pointer to **string** |  | [optional] 
**HostPing** | Pointer to **string** |  | [optional] 
**HostPingProbes** | Pointer to **string** | Optional host ping probe list as comma separated strings. The example shows possibilities. | [optional] 
**Probes** | Pointer to **string** | Optional probe list, otherwise all probes are used | [optional] 

## Methods

### NewScanOptions

`func NewScanOptions(targets string, ) *ScanOptions`

NewScanOptions instantiates a new ScanOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanOptionsWithDefaults

`func NewScanOptionsWithDefaults() *ScanOptions`

NewScanOptionsWithDefaults instantiates a new ScanOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ScanOptions) GetTargets() string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ScanOptions) GetTargetsOk() (*string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ScanOptions) SetTargets(v string)`

SetTargets sets Targets field to given value.


### GetExcludes

`func (o *ScanOptions) GetExcludes() string`

GetExcludes returns the Excludes field if non-nil, zero value otherwise.

### GetExcludesOk

`func (o *ScanOptions) GetExcludesOk() (*string, bool)`

GetExcludesOk returns a tuple with the Excludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludes

`func (o *ScanOptions) SetExcludes(v string)`

SetExcludes sets Excludes field to given value.

### HasExcludes

`func (o *ScanOptions) HasExcludes() bool`

HasExcludes returns a boolean if a field has been set.

### GetScanName

`func (o *ScanOptions) GetScanName() string`

GetScanName returns the ScanName field if non-nil, zero value otherwise.

### GetScanNameOk

`func (o *ScanOptions) GetScanNameOk() (*string, bool)`

GetScanNameOk returns a tuple with the ScanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanName

`func (o *ScanOptions) SetScanName(v string)`

SetScanName sets ScanName field to given value.

### HasScanName

`func (o *ScanOptions) HasScanName() bool`

HasScanName returns a boolean if a field has been set.

### GetScanDescription

`func (o *ScanOptions) GetScanDescription() string`

GetScanDescription returns the ScanDescription field if non-nil, zero value otherwise.

### GetScanDescriptionOk

`func (o *ScanOptions) GetScanDescriptionOk() (*string, bool)`

GetScanDescriptionOk returns a tuple with the ScanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDescription

`func (o *ScanOptions) SetScanDescription(v string)`

SetScanDescription sets ScanDescription field to given value.

### HasScanDescription

`func (o *ScanOptions) HasScanDescription() bool`

HasScanDescription returns a boolean if a field has been set.

### GetScanTemplate

`func (o *ScanOptions) GetScanTemplate() string`

GetScanTemplate returns the ScanTemplate field if non-nil, zero value otherwise.

### GetScanTemplateOk

`func (o *ScanOptions) GetScanTemplateOk() (*string, bool)`

GetScanTemplateOk returns a tuple with the ScanTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTemplate

`func (o *ScanOptions) SetScanTemplate(v string)`

SetScanTemplate sets ScanTemplate field to given value.

### HasScanTemplate

`func (o *ScanOptions) HasScanTemplate() bool`

HasScanTemplate returns a boolean if a field has been set.

### GetScanFrequency

`func (o *ScanOptions) GetScanFrequency() string`

GetScanFrequency returns the ScanFrequency field if non-nil, zero value otherwise.

### GetScanFrequencyOk

`func (o *ScanOptions) GetScanFrequencyOk() (*string, bool)`

GetScanFrequencyOk returns a tuple with the ScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFrequency

`func (o *ScanOptions) SetScanFrequency(v string)`

SetScanFrequency sets ScanFrequency field to given value.

### HasScanFrequency

`func (o *ScanOptions) HasScanFrequency() bool`

HasScanFrequency returns a boolean if a field has been set.

### GetScanStart

`func (o *ScanOptions) GetScanStart() string`

GetScanStart returns the ScanStart field if non-nil, zero value otherwise.

### GetScanStartOk

`func (o *ScanOptions) GetScanStartOk() (*string, bool)`

GetScanStartOk returns a tuple with the ScanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStart

`func (o *ScanOptions) SetScanStart(v string)`

SetScanStart sets ScanStart field to given value.

### HasScanStart

`func (o *ScanOptions) HasScanStart() bool`

HasScanStart returns a boolean if a field has been set.

### GetScanTags

`func (o *ScanOptions) GetScanTags() string`

GetScanTags returns the ScanTags field if non-nil, zero value otherwise.

### GetScanTagsOk

`func (o *ScanOptions) GetScanTagsOk() (*string, bool)`

GetScanTagsOk returns a tuple with the ScanTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTags

`func (o *ScanOptions) SetScanTags(v string)`

SetScanTags sets ScanTags field to given value.

### HasScanTags

`func (o *ScanOptions) HasScanTags() bool`

HasScanTags returns a boolean if a field has been set.

### GetScanGracePeriod

`func (o *ScanOptions) GetScanGracePeriod() string`

GetScanGracePeriod returns the ScanGracePeriod field if non-nil, zero value otherwise.

### GetScanGracePeriodOk

`func (o *ScanOptions) GetScanGracePeriodOk() (*string, bool)`

GetScanGracePeriodOk returns a tuple with the ScanGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanGracePeriod

`func (o *ScanOptions) SetScanGracePeriod(v string)`

SetScanGracePeriod sets ScanGracePeriod field to given value.

### HasScanGracePeriod

`func (o *ScanOptions) HasScanGracePeriod() bool`

HasScanGracePeriod returns a boolean if a field has been set.

### GetAgent

`func (o *ScanOptions) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *ScanOptions) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *ScanOptions) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *ScanOptions) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetExplorer

`func (o *ScanOptions) GetExplorer() string`

GetExplorer returns the Explorer field if non-nil, zero value otherwise.

### GetExplorerOk

`func (o *ScanOptions) GetExplorerOk() (*string, bool)`

GetExplorerOk returns a tuple with the Explorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorer

`func (o *ScanOptions) SetExplorer(v string)`

SetExplorer sets Explorer field to given value.

### HasExplorer

`func (o *ScanOptions) HasExplorer() bool`

HasExplorer returns a boolean if a field has been set.

### GetExplorerGroupId

`func (o *ScanOptions) GetExplorerGroupId() string`

GetExplorerGroupId returns the ExplorerGroupId field if non-nil, zero value otherwise.

### GetExplorerGroupIdOk

`func (o *ScanOptions) GetExplorerGroupIdOk() (*string, bool)`

GetExplorerGroupIdOk returns a tuple with the ExplorerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplorerGroupId

`func (o *ScanOptions) SetExplorerGroupId(v string)`

SetExplorerGroupId sets ExplorerGroupId field to given value.

### HasExplorerGroupId

`func (o *ScanOptions) HasExplorerGroupId() bool`

HasExplorerGroupId returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *ScanOptions) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ScanOptions) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ScanOptions) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *ScanOptions) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetHostedZoneName

`func (o *ScanOptions) GetHostedZoneName() string`

GetHostedZoneName returns the HostedZoneName field if non-nil, zero value otherwise.

### GetHostedZoneNameOk

`func (o *ScanOptions) GetHostedZoneNameOk() (*string, bool)`

GetHostedZoneNameOk returns a tuple with the HostedZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneName

`func (o *ScanOptions) SetHostedZoneName(v string)`

SetHostedZoneName sets HostedZoneName field to given value.

### HasHostedZoneName

`func (o *ScanOptions) HasHostedZoneName() bool`

HasHostedZoneName returns a boolean if a field has been set.

### GetRate

`func (o *ScanOptions) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ScanOptions) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ScanOptions) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ScanOptions) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetMaxHostRate

`func (o *ScanOptions) GetMaxHostRate() string`

GetMaxHostRate returns the MaxHostRate field if non-nil, zero value otherwise.

### GetMaxHostRateOk

`func (o *ScanOptions) GetMaxHostRateOk() (*string, bool)`

GetMaxHostRateOk returns a tuple with the MaxHostRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHostRate

`func (o *ScanOptions) SetMaxHostRate(v string)`

SetMaxHostRate sets MaxHostRate field to given value.

### HasMaxHostRate

`func (o *ScanOptions) HasMaxHostRate() bool`

HasMaxHostRate returns a boolean if a field has been set.

### GetPasses

`func (o *ScanOptions) GetPasses() string`

GetPasses returns the Passes field if non-nil, zero value otherwise.

### GetPassesOk

`func (o *ScanOptions) GetPassesOk() (*string, bool)`

GetPassesOk returns a tuple with the Passes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasses

`func (o *ScanOptions) SetPasses(v string)`

SetPasses sets Passes field to given value.

### HasPasses

`func (o *ScanOptions) HasPasses() bool`

HasPasses returns a boolean if a field has been set.

### GetMaxAttempts

`func (o *ScanOptions) GetMaxAttempts() string`

GetMaxAttempts returns the MaxAttempts field if non-nil, zero value otherwise.

### GetMaxAttemptsOk

`func (o *ScanOptions) GetMaxAttemptsOk() (*string, bool)`

GetMaxAttemptsOk returns a tuple with the MaxAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAttempts

`func (o *ScanOptions) SetMaxAttempts(v string)`

SetMaxAttempts sets MaxAttempts field to given value.

### HasMaxAttempts

`func (o *ScanOptions) HasMaxAttempts() bool`

HasMaxAttempts returns a boolean if a field has been set.

### GetMaxSockets

`func (o *ScanOptions) GetMaxSockets() string`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *ScanOptions) GetMaxSocketsOk() (*string, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *ScanOptions) SetMaxSockets(v string)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *ScanOptions) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### GetMaxGroupSize

`func (o *ScanOptions) GetMaxGroupSize() string`

GetMaxGroupSize returns the MaxGroupSize field if non-nil, zero value otherwise.

### GetMaxGroupSizeOk

`func (o *ScanOptions) GetMaxGroupSizeOk() (*string, bool)`

GetMaxGroupSizeOk returns a tuple with the MaxGroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGroupSize

`func (o *ScanOptions) SetMaxGroupSize(v string)`

SetMaxGroupSize sets MaxGroupSize field to given value.

### HasMaxGroupSize

`func (o *ScanOptions) HasMaxGroupSize() bool`

HasMaxGroupSize returns a boolean if a field has been set.

### GetMaxTtl

`func (o *ScanOptions) GetMaxTtl() string`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *ScanOptions) GetMaxTtlOk() (*string, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *ScanOptions) SetMaxTtl(v string)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *ScanOptions) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetTos

`func (o *ScanOptions) GetTos() string`

GetTos returns the Tos field if non-nil, zero value otherwise.

### GetTosOk

`func (o *ScanOptions) GetTosOk() (*string, bool)`

GetTosOk returns a tuple with the Tos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTos

`func (o *ScanOptions) SetTos(v string)`

SetTos sets Tos field to given value.

### HasTos

`func (o *ScanOptions) HasTos() bool`

HasTos returns a boolean if a field has been set.

### GetTcpPorts

`func (o *ScanOptions) GetTcpPorts() string`

GetTcpPorts returns the TcpPorts field if non-nil, zero value otherwise.

### GetTcpPortsOk

`func (o *ScanOptions) GetTcpPortsOk() (*string, bool)`

GetTcpPortsOk returns a tuple with the TcpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPorts

`func (o *ScanOptions) SetTcpPorts(v string)`

SetTcpPorts sets TcpPorts field to given value.

### HasTcpPorts

`func (o *ScanOptions) HasTcpPorts() bool`

HasTcpPorts returns a boolean if a field has been set.

### GetTcpExcludes

`func (o *ScanOptions) GetTcpExcludes() string`

GetTcpExcludes returns the TcpExcludes field if non-nil, zero value otherwise.

### GetTcpExcludesOk

`func (o *ScanOptions) GetTcpExcludesOk() (*string, bool)`

GetTcpExcludesOk returns a tuple with the TcpExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpExcludes

`func (o *ScanOptions) SetTcpExcludes(v string)`

SetTcpExcludes sets TcpExcludes field to given value.

### HasTcpExcludes

`func (o *ScanOptions) HasTcpExcludes() bool`

HasTcpExcludes returns a boolean if a field has been set.

### GetScreenshots

`func (o *ScanOptions) GetScreenshots() string`

GetScreenshots returns the Screenshots field if non-nil, zero value otherwise.

### GetScreenshotsOk

`func (o *ScanOptions) GetScreenshotsOk() (*string, bool)`

GetScreenshotsOk returns a tuple with the Screenshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshots

`func (o *ScanOptions) SetScreenshots(v string)`

SetScreenshots sets Screenshots field to given value.

### HasScreenshots

`func (o *ScanOptions) HasScreenshots() bool`

HasScreenshots returns a boolean if a field has been set.

### GetNameservers

`func (o *ScanOptions) GetNameservers() string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *ScanOptions) GetNameserversOk() (*string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *ScanOptions) SetNameservers(v string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *ScanOptions) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetSubnetPing

`func (o *ScanOptions) GetSubnetPing() string`

GetSubnetPing returns the SubnetPing field if non-nil, zero value otherwise.

### GetSubnetPingOk

`func (o *ScanOptions) GetSubnetPingOk() (*string, bool)`

GetSubnetPingOk returns a tuple with the SubnetPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPing

`func (o *ScanOptions) SetSubnetPing(v string)`

SetSubnetPing sets SubnetPing field to given value.

### HasSubnetPing

`func (o *ScanOptions) HasSubnetPing() bool`

HasSubnetPing returns a boolean if a field has been set.

### GetSubnetPingNetSize

`func (o *ScanOptions) GetSubnetPingNetSize() string`

GetSubnetPingNetSize returns the SubnetPingNetSize field if non-nil, zero value otherwise.

### GetSubnetPingNetSizeOk

`func (o *ScanOptions) GetSubnetPingNetSizeOk() (*string, bool)`

GetSubnetPingNetSizeOk returns a tuple with the SubnetPingNetSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPingNetSize

`func (o *ScanOptions) SetSubnetPingNetSize(v string)`

SetSubnetPingNetSize sets SubnetPingNetSize field to given value.

### HasSubnetPingNetSize

`func (o *ScanOptions) HasSubnetPingNetSize() bool`

HasSubnetPingNetSize returns a boolean if a field has been set.

### GetSubnetPingProbes

`func (o *ScanOptions) GetSubnetPingProbes() string`

GetSubnetPingProbes returns the SubnetPingProbes field if non-nil, zero value otherwise.

### GetSubnetPingProbesOk

`func (o *ScanOptions) GetSubnetPingProbesOk() (*string, bool)`

GetSubnetPingProbesOk returns a tuple with the SubnetPingProbes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPingProbes

`func (o *ScanOptions) SetSubnetPingProbes(v string)`

SetSubnetPingProbes sets SubnetPingProbes field to given value.

### HasSubnetPingProbes

`func (o *ScanOptions) HasSubnetPingProbes() bool`

HasSubnetPingProbes returns a boolean if a field has been set.

### GetSubnetPingSampleRate

`func (o *ScanOptions) GetSubnetPingSampleRate() string`

GetSubnetPingSampleRate returns the SubnetPingSampleRate field if non-nil, zero value otherwise.

### GetSubnetPingSampleRateOk

`func (o *ScanOptions) GetSubnetPingSampleRateOk() (*string, bool)`

GetSubnetPingSampleRateOk returns a tuple with the SubnetPingSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetPingSampleRate

`func (o *ScanOptions) SetSubnetPingSampleRate(v string)`

SetSubnetPingSampleRate sets SubnetPingSampleRate field to given value.

### HasSubnetPingSampleRate

`func (o *ScanOptions) HasSubnetPingSampleRate() bool`

HasSubnetPingSampleRate returns a boolean if a field has been set.

### GetHostPing

`func (o *ScanOptions) GetHostPing() string`

GetHostPing returns the HostPing field if non-nil, zero value otherwise.

### GetHostPingOk

`func (o *ScanOptions) GetHostPingOk() (*string, bool)`

GetHostPingOk returns a tuple with the HostPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPing

`func (o *ScanOptions) SetHostPing(v string)`

SetHostPing sets HostPing field to given value.

### HasHostPing

`func (o *ScanOptions) HasHostPing() bool`

HasHostPing returns a boolean if a field has been set.

### GetHostPingProbes

`func (o *ScanOptions) GetHostPingProbes() string`

GetHostPingProbes returns the HostPingProbes field if non-nil, zero value otherwise.

### GetHostPingProbesOk

`func (o *ScanOptions) GetHostPingProbesOk() (*string, bool)`

GetHostPingProbesOk returns a tuple with the HostPingProbes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPingProbes

`func (o *ScanOptions) SetHostPingProbes(v string)`

SetHostPingProbes sets HostPingProbes field to given value.

### HasHostPingProbes

`func (o *ScanOptions) HasHostPingProbes() bool`

HasHostPingProbes returns a boolean if a field has been set.

### GetProbes

`func (o *ScanOptions) GetProbes() string`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *ScanOptions) GetProbesOk() (*string, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *ScanOptions) SetProbes(v string)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *ScanOptions) HasProbes() bool`

HasProbes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


