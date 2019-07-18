package types

// DeviceProfile (s) are profiles that are individual to the device.
type DeviceProfile struct {
	ID                uint `gorm:"primary_key"`
	PayloadUUID       string
	PayloadIdentifier string
	HashedPayloadUUID string
	MobileconfigData  []byte
	MobileconfigHash  []byte
	DeviceUDID        string
	Installed         bool `gorm:"default:true"`
}

// SharedProfile (s) are profiles that go on every device.
type SharedProfile struct {
	ID                uint `gorm:"primary_key"`
	PayloadUUID       string
	HashedPayloadUUID string
	PayloadIdentifier string
	MobileconfigData  []byte
	MobileconfigHash  []byte
	Installed         bool `gorm:"default:true"`
}

// ProfilePayload - struct to unpack the payload sent to mdmdirector
type ProfilePayload struct {
	SerialNumbers []string `json:"serial_numbers,omitempty"`
	DeviceUDIDs   []string `json:"udids,omitempty"`
	Mobileconfigs []string `json:"profiles"`
}

type DeleteProfilePayload struct {
	SerialNumbers []string                     `json:"serial_numbers,omitempty"`
	DeviceUDIDs   []string                     `json:"udids,omitempty"`
	Mobileconfigs []DeletedMobileconfigPayload `json:"profiles"`
}

type DeletedMobileconfigPayload struct {
	UUID              string `json:"uuid"`
	PayloadIdentifier string `json:"payload_identifier"`
}

// ProfileList - returned data from the ProfileList MDM command
type ProfileListData struct {
	ProfileList []ProfileList
}

type ProfileList struct {
	HasRemovalPasscode       bool          `plist:"HasRemovalPasscode"`
	IsEncrypted              bool          `plist:"IsEncrypted"`
	PayloadContent           []interface{} `plist:"PaylodContent"`
	PayloadDescription       string        `plist:"PayloadDescription"`
	PayloadDisplayName       string        `plist:"PayloadDisplayName"`
	PayloadIdentifier        string        `plist:"PayloadIdentifier"`
	PayloadOrganization      string        `plist:"PayloadOrganization"`
	PayloadRemovalDisallowed bool          `plist:"PayloadRemovalDisallowed"`
	PayloadUUID              string        `plist:"PayloadUUID"`
	PayloadVersion           int           `plist:"PayloadVersion"`
	FullPayload              bool          `plist:"FullPayload"`
}

func (profile *DeviceProfile) AfterCreate() (err error) {
	BumpDeviceLastUpdated(profile.DeviceUDID)
	return
}

func (profile *DeviceProfile) AfterUpdate() (err error) {
	BumpDeviceLastUpdated(profile.DeviceUDID)
	return
}
