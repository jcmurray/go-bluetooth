package device



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "reflect"
  "github.com/fatih/structs"
  "github.com/muka/go-bluetooth/util"
  "github.com/godbus/dbus"
)

var Device1Interface = "org.bluez.Device1"


// NewDevice1 create a new instance of Device1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX
func NewDevice1(objectPath dbus.ObjectPath) (*Device1, error) {
	a := new(Device1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Device1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Device1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Device1 Device hierarchy

*/
type Device1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Device1Properties
}

// Device1Properties contains the exposed properties of an interface
type Device1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Appearance External appearance of device, as found on GAP service.
	*/
	Appearance uint16

	/*
	UUIDs List of 128-bit UUIDs that represents the available
  remote services.
	*/
	UUIDs []string

	/*
	Connected Indicates if the remote device is currently connected.
  A PropertiesChanged signal indicate changes to this
  status.
	*/
	Connected bool

	/*
	RSSI Received Signal Strength Indicator of the remote
  device (inquiry or advertising).
	*/
	RSSI int16

	/*
	ServicesResolved Indicate whether or not service discovery has been
  resolved.
	*/
	ServicesResolved bool

	/*
	AdvertisingFlags The Advertising Data Flags of the remote device.
	*/
	AdvertisingFlags []byte

	/*
	Alias The name alias for the remote device. The alias can
  be used to have a different friendly name for the
  remote device.
  In case no alias is set, it will return the remote
  device name. Setting an empty string as alias will
  convert it back to the remote device name.
  When resetting the alias with an empty string, the
  property will default back to the remote name.
	*/
	Alias string

	/*
	AdvertisingData The Advertising Data of the remote device. Keys are
  are 8 bits AD Type followed by data as byte array.
  Note: Only types considered safe to be handled by
  application are exposed.
  Possible values:
  <type> <byte array>
  ...
  Example:
  <Transport Discovery> <Organization Flags...>
  0x26                   0x01         0x01...
	*/
	AdvertisingData map[string]interface{}

	/*
	ManufacturerData Manufacturer specific advertisement data. Keys are
  16 bits Manufacturer ID followed by its byte array
  value.
	*/
	ManufacturerData map[uint16]interface{}

	/*
	ServiceData Service advertisement data. Keys are the UUIDs in
  string format followed by its byte array value.
	*/
	ServiceData map[string]interface{}

	/*
	Address The Bluetooth device address of the remote device.
	*/
	Address string

	/*
	Icon Proposed icon name according to the freedesktop.org
  icon naming specification.
	*/
	Icon string

	/*
	Class The Bluetooth class of device of the remote device.
	*/
	Class uint32

	/*
	Trusted Indicates if the remote is seen as trusted. This
  setting can be changed by the application.
	*/
	Trusted bool

	/*
	Blocked If set to true any incoming connections from the
  device will be immediately rejected. Any device
  drivers will also be removed and no new ones will
  be probed as long as the device is blocked.
	*/
	Blocked bool

	/*
	TxPower Advertised transmitted power level (inquiry or
  advertising).
	*/
	TxPower int16

	/*
	AddressType The Bluetooth device Address Type. For dual-mode and
  BR/EDR only devices this defaults to "public". Single
  mode LE devices may have either value. If remote device
  uses privacy than before pairing this represents address
  type used for connection and Identity Address after
  pairing.
  Possible values:
  "public" - Public address
  "random" - Random address
	*/
	AddressType string

	/*
	Name The Bluetooth remote name. This value can not be
  changed. Use the Alias property instead.
  This value is only present for completeness. It is
  better to always use the Alias property when
  displaying the devices name.
  If the Alias property is unset, it will reflect
  this value which makes it more convenient.
	*/
	Name string

	/*
	Paired Indicates if the remote device is paired.
	*/
	Paired bool

	/*
	Adapter The object path of the adapter the device belongs to.
	*/
	Adapter dbus.ObjectPath

	/*
	LegacyPairing Set to true if the device only supports the pre-2.1
  pairing mechanism. This property is useful during
  device discovery to anticipate whether legacy or
  simple pairing will occur if pairing is initiated.
  Note that this property can exhibit false-positives
  in the case of Bluetooth 2.1 (or newer) devices that
  have disabled Extended Inquiry Response support.
	*/
	LegacyPairing bool

	/*
	Modalias Remote Device ID information in modalias format
  used by the kernel and udev.
	*/
	Modalias string

}

//Lock access to properties
func (p *Device1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Device1Properties) Unlock() {
	p.lock.Unlock()
}


// SetAppearance set Appearance value
func (a *Device1) SetAppearance(v uint16) error {
	return a.SetProperty("Appearance", v)
}

// GetAppearance get Appearance value
func (a *Device1) GetAppearance() (uint16, error) {
	v, err := a.GetProperty("Appearance")
	if err != nil {
		return uint16(0), err
	}
	return v.Value().(uint16), nil
}

// SetUUIDs set UUIDs value
func (a *Device1) SetUUIDs(v []string) error {
	return a.SetProperty("UUIDs", v)
}

// GetUUIDs get UUIDs value
func (a *Device1) GetUUIDs() ([]string, error) {
	v, err := a.GetProperty("UUIDs")
	if err != nil {
		return []string{}, err
	}
	return v.Value().([]string), nil
}

// SetConnected set Connected value
func (a *Device1) SetConnected(v bool) error {
	return a.SetProperty("Connected", v)
}

// GetConnected get Connected value
func (a *Device1) GetConnected() (bool, error) {
	v, err := a.GetProperty("Connected")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetRSSI set RSSI value
func (a *Device1) SetRSSI(v int16) error {
	return a.SetProperty("RSSI", v)
}

// GetRSSI get RSSI value
func (a *Device1) GetRSSI() (int16, error) {
	v, err := a.GetProperty("RSSI")
	if err != nil {
		return int16(0), err
	}
	return v.Value().(int16), nil
}

// SetServicesResolved set ServicesResolved value
func (a *Device1) SetServicesResolved(v bool) error {
	return a.SetProperty("ServicesResolved", v)
}

// GetServicesResolved get ServicesResolved value
func (a *Device1) GetServicesResolved() (bool, error) {
	v, err := a.GetProperty("ServicesResolved")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetAdvertisingFlags set AdvertisingFlags value
func (a *Device1) SetAdvertisingFlags(v []byte) error {
	return a.SetProperty("AdvertisingFlags", v)
}

// GetAdvertisingFlags get AdvertisingFlags value
func (a *Device1) GetAdvertisingFlags() ([]byte, error) {
	v, err := a.GetProperty("AdvertisingFlags")
	if err != nil {
		return []byte{}, err
	}
	return v.Value().([]byte), nil
}

// SetAlias set Alias value
func (a *Device1) SetAlias(v string) error {
	return a.SetProperty("Alias", v)
}

// GetAlias get Alias value
func (a *Device1) GetAlias() (string, error) {
	v, err := a.GetProperty("Alias")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetAdvertisingData set AdvertisingData value
func (a *Device1) SetAdvertisingData(v map[string]interface{}) error {
	return a.SetProperty("AdvertisingData", v)
}

// GetAdvertisingData get AdvertisingData value
func (a *Device1) GetAdvertisingData() (map[string]interface{}, error) {
	v, err := a.GetProperty("AdvertisingData")
	if err != nil {
		return map[string]interface{}{}, err
	}
	return v.Value().(map[string]interface{}), nil
}

// SetManufacturerData set ManufacturerData value
func (a *Device1) SetManufacturerData(v map[string]interface{}) error {
	return a.SetProperty("ManufacturerData", v)
}

// GetManufacturerData get ManufacturerData value
func (a *Device1) GetManufacturerData() (map[string]interface{}, error) {
	v, err := a.GetProperty("ManufacturerData")
	if err != nil {
		return map[string]interface{}{}, err
	}
	return v.Value().(map[string]interface{}), nil
}

// SetServiceData set ServiceData value
func (a *Device1) SetServiceData(v map[string]interface{}) error {
	return a.SetProperty("ServiceData", v)
}

// GetServiceData get ServiceData value
func (a *Device1) GetServiceData() (map[string]interface{}, error) {
	v, err := a.GetProperty("ServiceData")
	if err != nil {
		return map[string]interface{}{}, err
	}
	return v.Value().(map[string]interface{}), nil
}

// SetAddress set Address value
func (a *Device1) SetAddress(v string) error {
	return a.SetProperty("Address", v)
}

// GetAddress get Address value
func (a *Device1) GetAddress() (string, error) {
	v, err := a.GetProperty("Address")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetIcon set Icon value
func (a *Device1) SetIcon(v string) error {
	return a.SetProperty("Icon", v)
}

// GetIcon get Icon value
func (a *Device1) GetIcon() (string, error) {
	v, err := a.GetProperty("Icon")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetClass set Class value
func (a *Device1) SetClass(v uint32) error {
	return a.SetProperty("Class", v)
}

// GetClass get Class value
func (a *Device1) GetClass() (uint32, error) {
	v, err := a.GetProperty("Class")
	if err != nil {
		return uint32(0), err
	}
	return v.Value().(uint32), nil
}

// SetTrusted set Trusted value
func (a *Device1) SetTrusted(v bool) error {
	return a.SetProperty("Trusted", v)
}

// GetTrusted get Trusted value
func (a *Device1) GetTrusted() (bool, error) {
	v, err := a.GetProperty("Trusted")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetBlocked set Blocked value
func (a *Device1) SetBlocked(v bool) error {
	return a.SetProperty("Blocked", v)
}

// GetBlocked get Blocked value
func (a *Device1) GetBlocked() (bool, error) {
	v, err := a.GetProperty("Blocked")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetTxPower set TxPower value
func (a *Device1) SetTxPower(v int16) error {
	return a.SetProperty("TxPower", v)
}

// GetTxPower get TxPower value
func (a *Device1) GetTxPower() (int16, error) {
	v, err := a.GetProperty("TxPower")
	if err != nil {
		return int16(0), err
	}
	return v.Value().(int16), nil
}

// SetAddressType set AddressType value
func (a *Device1) SetAddressType(v string) error {
	return a.SetProperty("AddressType", v)
}

// GetAddressType get AddressType value
func (a *Device1) GetAddressType() (string, error) {
	v, err := a.GetProperty("AddressType")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetName set Name value
func (a *Device1) SetName(v string) error {
	return a.SetProperty("Name", v)
}

// GetName get Name value
func (a *Device1) GetName() (string, error) {
	v, err := a.GetProperty("Name")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetPaired set Paired value
func (a *Device1) SetPaired(v bool) error {
	return a.SetProperty("Paired", v)
}

// GetPaired get Paired value
func (a *Device1) GetPaired() (bool, error) {
	v, err := a.GetProperty("Paired")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetAdapter set Adapter value
func (a *Device1) SetAdapter(v dbus.ObjectPath) error {
	return a.SetProperty("Adapter", v)
}

// GetAdapter get Adapter value
func (a *Device1) GetAdapter() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Adapter")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}

// SetLegacyPairing set LegacyPairing value
func (a *Device1) SetLegacyPairing(v bool) error {
	return a.SetProperty("LegacyPairing", v)
}

// GetLegacyPairing get LegacyPairing value
func (a *Device1) GetLegacyPairing() (bool, error) {
	v, err := a.GetProperty("LegacyPairing")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetModalias set Modalias value
func (a *Device1) SetModalias(v string) error {
	return a.SetProperty("Modalias", v)
}

// GetModalias get Modalias value
func (a *Device1) GetModalias() (string, error) {
	v, err := a.GetProperty("Modalias")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}


// Close the connection
func (a *Device1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Device1 object path
func (a *Device1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return Device1 interface
func (a *Device1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Device1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Device1Properties to map
func (a *Device1Properties) ToMap() (map[string]interface{}, error) {
	return structs.Map(a), nil
}

// FromMap convert a map to an Device1Properties
func (a *Device1Properties) FromMap(props map[string]interface{}) (*Device1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Device1Properties
func (a *Device1Properties) FromDBusMap(props map[string]dbus.Variant) (*Device1Properties, error) {
	s := new(Device1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// GetProperties load all available properties
func (a *Device1) GetProperties() (*Device1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Device1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Device1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Device1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Device1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Device1) WatchProperties() (chan *bluez.PropertyChanged, error) {

	channel, err := a.client.Register(a.Path(), a.Interface())
	if err != nil {
		return nil, err
	}

	ch := make(chan *bluez.PropertyChanged)

	go (func() {
		for {

			if channel == nil {
				break
			}

			sig := <-channel

			if sig == nil {
				return
			}

			if sig.Name != bluez.PropertiesChanged {
				continue
			}
			if sig.Path != a.Path() {
				continue
			}

			iface := sig.Body[0].(string)
			changes := sig.Body[1].(map[string]dbus.Variant)

			for field, val := range changes {

				// updates [*]Properties struct when a property change
				s := reflect.ValueOf(a.Properties).Elem()
				// exported field
				f := s.FieldByName(field)
				if f.IsValid() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.
					if f.CanSet() {
						x := reflect.ValueOf(val.Value())
						a.Properties.Lock()
						f.Set(x)
						a.Properties.Unlock()
					}
				}

				propChanged := &bluez.PropertyChanged{
					Interface: iface,
					Name:      field,
					Value:     val.Value(),
				}
				ch <- propChanged
			}

		}
	})()

	return ch, nil
}

func (a *Device1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	ch <- nil
	close(ch)
	return nil
}




/*
Connect This is a generic method to connect any profiles
the remote device supports that can be connected
to and have been flagged as auto-connectable on
our side. If only subset of profiles is already
connected it will try to connect currently disconnected
ones.
If at least one profile was connected successfully this
method will indicate success.
For dual-mode devices only one bearer is connected at
time, the conditions are in the following order:
1. Connect the disconnected bearer if already
connected.
2. Connect first the bonded bearer. If no
bearers are bonded or both are skip and check
latest seen bearer.
3. Connect last seen bearer, in case the
timestamps are the same BR/EDR takes
precedence.
Possible errors: org.bluez.Error.NotReady
org.bluez.Error.Failed
org.bluez.Error.InProgress
org.bluez.Error.AlreadyConnected
*/
func (a *Device1) Connect() error {
	
	return a.client.Call("Connect", 0, ).Store()
	
}

/*
Disconnect This method gracefully disconnects all connected
profiles and then terminates low-level ACL connection.
ACL connection will be terminated even if some profiles
were not disconnected properly e.g. due to misbehaving
device.
This method can be also used to cancel a preceding
Connect call before a reply to it has been received.
For non-trusted devices connected over LE bearer calling
this method will disable incoming connections until
Connect method is called again.
Possible errors: org.bluez.Error.NotConnected
*/
func (a *Device1) Disconnect() error {
	
	return a.client.Call("Disconnect", 0, ).Store()
	
}

/*
ConnectProfile This method connects a specific profile of this
device. The UUID provided is the remote service
UUID for the profile.
Possible errors: org.bluez.Error.Failed
org.bluez.Error.InProgress
org.bluez.Error.InvalidArguments
org.bluez.Error.NotAvailable
org.bluez.Error.NotReady
*/
func (a *Device1) ConnectProfile(uuid string) error {
	
	return a.client.Call("ConnectProfile", 0, uuid).Store()
	
}

/*
DisconnectProfile This method disconnects a specific profile of
this device. The profile needs to be registered
client profile.
There is no connection tracking for a profile, so
as long as the profile is registered this will always
succeed.
Possible errors: org.bluez.Error.Failed
org.bluez.Error.InProgress
org.bluez.Error.InvalidArguments
org.bluez.Error.NotSupported
*/
func (a *Device1) DisconnectProfile(uuid string) error {
	
	return a.client.Call("DisconnectProfile", 0, uuid).Store()
	
}

/*
Pair This method will connect to the remote device,
*/
func (a *Device1) Pair() error {
	
	return a.client.Call("Pair", 0, ).Store()
	
}

/*
CancelPairing This method can be used to cancel a pairing
operation initiated by the Pair method.
Possible errors: org.bluez.Error.DoesNotExist
org.bluez.Error.Failed
*/
func (a *Device1) CancelPairing() error {
	
	return a.client.Call("CancelPairing", 0, ).Store()
	
}

