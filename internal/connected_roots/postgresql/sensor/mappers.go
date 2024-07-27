package sensor

import (
	"github.com/Kortivex/connected_roots/internal/connected_roots"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/crop_type"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/orchard"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/role"
	"github.com/Kortivex/connected_roots/internal/connected_roots/postgresql/user"
	"time"
)

func toDomain(sensor *Sensors) *connected_roots.Sensors {
	if sensor == nil {
		return nil
	}
	return &connected_roots.Sensors{
		ID:              sensor.ID,
		Name:            sensor.Name,
		Location:        sensor.Location,
		ModelNumber:     sensor.ModelNumber,
		Manufacturer:    sensor.Manufacturer,
		CalibrationDate: sensor.CalibrationDate,
		BatteryLife:     sensor.BatteryLife,
		SSID:            sensor.SSID,
		Channel:         sensor.Channel,
		DNS:             sensor.DNS,
		IP:              sensor.IP,
		Subnet:          sensor.Subnet,
		MAC:             sensor.MAC,
		Status:          sensor.Status,
		FirmwareVersion: sensor.FirmwareVersion,
		OrchardID:       sensor.OrchardID,
		Orchard:         toOrchardDomain(sensor.Orchard),
		CreatedAt:       sensor.CreatedAt,
		UpdatedAt:       sensor.UpdatedAt,
	}
}

func toOrchardDomain(orchard *orchard.Orchards) *connected_roots.Orchards {
	if orchard == nil {
		return nil
	}
	return &connected_roots.Orchards{
		ID:         orchard.ID,
		Name:       orchard.Name,
		Location:   orchard.Location,
		Size:       orchard.Size,
		Soil:       orchard.Soil,
		Fertilizer: orchard.Fertilizer,
		Composting: orchard.Composting,
		ImageURL:   orchard.ImageURL,
		UserID:     orchard.UserID,
		User:       toUserDomain(orchard.User),
		CropTypeID: orchard.CropTypeID,
		CropType:   toCropTypeDomain(orchard.CropType),
		CreatedAt:  orchard.BaseModel.CreatedAt,
		UpdatedAt:  orchard.BaseModel.UpdatedAt,
	}
}

func toUserDomain(usr *user.Users) *connected_roots.Users {
	if usr == nil {
		return nil
	}
	return &connected_roots.Users{
		ID:        usr.ID,
		Name:      usr.Name,
		Surname:   usr.Surname,
		Email:     usr.Email,
		Password:  usr.Password,
		Telephone: usr.Telephone,
		Language:  usr.Language,
		RoleID:    usr.RoleID,
		Role:      toRoleDomain(usr.Role),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func toRoleDomain(rl *role.Roles) *connected_roots.Roles {
	if rl == nil {
		return nil
	}
	return &connected_roots.Roles{
		ID:          rl.ID,
		Name:        rl.Name,
		Description: rl.Description,
		CreatedAt:   rl.CreatedAt,
		UpdatedAt:   rl.UpdatedAt,
	}
}

func toCropTypeDomain(cropType *crop_type.CropTypes) *connected_roots.CropTypes {
	if cropType == nil {
		return nil
	}
	return &connected_roots.CropTypes{
		ID:             cropType.ID,
		Name:           cropType.Name,
		ScientificName: cropType.ScientificName,
		LifeCycle:      cropType.LifeCycle,
		PlantingSeason: cropType.PlantingSeason,
		HarvestSeason:  cropType.HarvestSeason,
		Irrigation:     cropType.Irrigation,
		ImageURL:       cropType.ImageURL,
		Description:    cropType.Description,
		CreatedAt:      cropType.CreatedAt,
		UpdatedAt:      cropType.UpdatedAt,
	}
}

func toDomainSlice(sensors []*Sensors) []*connected_roots.Sensors {
	sensorsDomain := []*connected_roots.Sensors{}
	for _, sensor := range sensors {
		sensorDomain := toDomain(sensor)
		sensorsDomain = append(sensorsDomain, sensorDomain)
	}
	return sensorsDomain
}

func toDomainUserSlice(sensors []*Sensors) []*connected_roots.Sensors {
	sensorsDomain := []*connected_roots.Sensors{}
	for _, sensor := range sensors {
		if sensor.Orchard == nil {
			continue
		}
		sensorDomain := toDomain(sensor)
		sensorsDomain = append(sensorsDomain, sensorDomain)
	}
	return sensorsDomain
}

func toDB(sensor *connected_roots.Sensors, id string) *Sensors {
	return &Sensors{
		ID:              id,
		Name:            sensor.Name,
		Location:        sensor.Location,
		ModelNumber:     sensor.ModelNumber,
		Manufacturer:    sensor.Manufacturer,
		CalibrationDate: sensor.CalibrationDate,
		BatteryLife:     sensor.BatteryLife,
		SSID:            sensor.SSID,
		Channel:         sensor.Channel,
		DNS:             sensor.DNS,
		IP:              sensor.IP,
		Subnet:          sensor.Subnet,
		MAC:             sensor.MAC,
		Status:          sensor.Status,
		FirmwareVersion: sensor.FirmwareVersion,
		OrchardID:       sensor.OrchardID,
	}
}

func toDomainData(sensorData *SensorsData) *connected_roots.SensorsData {
	return &connected_roots.SensorsData{
		ID:             sensorData.ID,
		Voltage:        sensorData.Voltage,
		Battery:        sensorData.Battery,
		Soil:           sensorData.Soil,
		Salt:           sensorData.Salt,
		Light:          sensorData.Light,
		TemperatureIn:  sensorData.TemperatureIn,
		TemperatureOut: sensorData.TemperatureOut,
		HumidityIn:     sensorData.HumidityIn,
		HumidityOut:    sensorData.HumidityOut,
		Pressure:       sensorData.Pressure,
		Altitude:       sensorData.Altitude,
		SensorID:       sensorData.SensorID,
		Sensor:         toDomain(sensorData.Sensor),
		CreatedAt:      sensorData.CreatedAt,
	}
}

func toDomainDataSlice(sensorsData []*SensorsData) []*connected_roots.SensorsData {
	sensorsDataDomain := []*connected_roots.SensorsData{}
	for _, sensorData := range sensorsData {
		sensorDataDomain := toDomainData(sensorData)
		sensorsDataDomain = append(sensorsDataDomain, sensorDataDomain)
	}
	return sensorsDataDomain
}

func toDBData(sensorData *connected_roots.SensorsData, id string) *SensorsData {
	return &SensorsData{
		ID:             id,
		Voltage:        sensorData.Voltage,
		Battery:        sensorData.Battery,
		Soil:           sensorData.Soil,
		Salt:           sensorData.Salt,
		Light:          sensorData.Light,
		TemperatureIn:  sensorData.TemperatureIn,
		TemperatureOut: sensorData.TemperatureOut,
		HumidityIn:     sensorData.HumidityIn,
		HumidityOut:    sensorData.HumidityOut,
		Pressure:       sensorData.Pressure,
		Altitude:       sensorData.Altitude,
		SensorID:       sensorData.SensorID,
	}
}
