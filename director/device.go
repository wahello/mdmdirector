package director

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/grahamgilbert/mdmdirector/db"
	"github.com/grahamgilbert/mdmdirector/types"

	// sqlite
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func UpdateDevice(newDevice types.Device) *types.Device {
	var device types.Device
	var oldDevice types.Device

	if newDevice.UDID == "" {
		return &newDevice
	}

	if err := db.DB.Where("ud_id = ?", newDevice.UDID).First(&device).Scan(&oldDevice).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			db.DB.Create(&newDevice)
		}
	} else {
		// err := db.DB.Model(&device).Where("ud_id = ?", newDevice.UDID).First(&newDevice).Save(&newDevice).Error
		// if err != nil {
		// 	log.Print(err)
		// }

		err := db.DB.Model(&device).Where("ud_id = ?", newDevice.UDID).Assign(&newDevice).FirstOrCreate(&device).Error
		if err != nil {
			log.Print(err)
		}

		// The below needs more work
		// if utils.PushOnNewBuild() {
		// 	if !cmp.Equal(&oldDevice.BuildVersion, &newDevice.BuildVersion) {
		// 		log.Printf("Re-pushing all profiles to %v due to build number change", device.UDID)
		// 		InstallAllProfiles(device)
		// 	}
		// }

	}

	return &device
}

func GetDevice(udid string) types.Device {
	var device types.Device

	err := db.DB.Model(device).Where("ud_id = ?", udid).First(&device).Scan(&device).Error
	if err != nil {
		fmt.Println(err)
		log.Print("Couldn't scan to Device model")
	}
	return device
}

func GetDeviceSerial(serial string) types.Device {
	var device types.Device

	err := db.DB.Model(device).Where("serial_number = ?", serial).First(&device).Scan(&device).Error
	if err != nil {
		fmt.Println(err)
		log.Print("Couldn't scan to Device model")
	}
	return device
}

func GetAllDevices() []types.Device {
	// var device types.Device
	var devices []types.Device

	err := db.DB.Find(&devices).Scan(&devices).Error
	if err != nil {
		fmt.Println(err)
		log.Print("Couldn't scan to Device model")
	}
	return devices
}

func GetAllDevicesAndAssociations() *[]types.Device {
	var devices []types.Device

	err := db.DB.Preload("OSUpdateSettings").Preload("SecurityInfo").Preload("SecurityInfo.FirmwarePasswordStatus").Preload("SecurityInfo.ManagementStatus").Find(&devices).Error
	if err != nil {
		fmt.Println(err)
		log.Print("Couldn't scan to Device model")
	}

	return &devices
}

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	devices := GetAllDevicesAndAssociations()

	output, err := json.MarshalIndent(&devices, "", "    ")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(output)

}

func RequestDeviceInformation(device types.Device) {
	var requestType = "DeviceInformation"
	log.Printf("Requesting Device Info for %v", device.UDID)
	var payload types.CommandPayload
	payload.UDID = device.UDID
	payload.RequestType = requestType
	payload.Queries = types.DeviceInformationQueries
	SendCommand(payload)
}

func SetTokenUpdate(device types.Device) {
	var deviceModel types.Device
	log.Printf("TokenUpdate received for %v", device.UDID)
	err := db.DB.Model(&deviceModel).Where("ud_id = ?", device.UDID).Update(map[string]interface{}{"token_update_recieved": true, "authenticate_recieved": true}).Error
	if err != nil {
		log.Print(err)
	}
}
