package obex



import (
  "sync"
  "github.com/muka/go-bluetooth/bluez"
  "reflect"
  "github.com/fatih/structs"
  "github.com/muka/go-bluetooth/util"
  "github.com/godbus/dbus"
)

var PhonebookAccess1Interface = "org.bluez.obex.PhonebookAccess1"


// NewPhonebookAccess1 create a new instance of PhonebookAccess1
//
// Args:
// - objectPath: [Session object path]
func NewPhonebookAccess1(objectPath dbus.ObjectPath) (*PhonebookAccess1, error) {
	a := new(PhonebookAccess1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez.obex",
			Iface: PhonebookAccess1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(PhonebookAccess1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
PhonebookAccess1 Phonebook Access hierarchy

*/
type PhonebookAccess1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*PhonebookAccess1Properties
}

// PhonebookAccess1Properties contains the exposed properties of an interface
type PhonebookAccess1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	FixedImageSize Indicate support for fixed image size.
  Possible values: True if image is JPEG 300x300 pixels
  otherwise False.
	*/
	FixedImageSize bool

	/*
	Folder Current folder.
	*/
	Folder string

	/*
	DatabaseIdentifier 128 bits persistent database identifier.
  Possible values: 32-character hexadecimal such
  as A1A2A3A4B1B2C1C2D1D2E1E2E3E4E5E6
	*/
	DatabaseIdentifier string

	/*
	PrimaryCounter 128 bits primary version counter.
  Possible values: 32-character hexadecimal such
  as A1A2A3A4B1B2C1C2D1D2E1E2E3E4E5E6
	*/
	PrimaryCounter string

	/*
	SecondaryCounter 128 bits secondary version counter.
  Possible values: 32-character hexadecimal such
  as A1A2A3A4B1B2C1C2D1D2E1E2E3E4E5E6
	*/
	SecondaryCounter string

}

//Lock access to properties
func (p *PhonebookAccess1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *PhonebookAccess1Properties) Unlock() {
	p.lock.Unlock()
}


// SetFixedImageSize set FixedImageSize value
func (a *PhonebookAccess1) SetFixedImageSize(v bool) error {
	return a.SetProperty("FixedImageSize", v)
}

// GetFixedImageSize get FixedImageSize value
func (a *PhonebookAccess1) GetFixedImageSize() (bool, error) {
	v, err := a.GetProperty("FixedImageSize")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}

// SetFolder set Folder value
func (a *PhonebookAccess1) SetFolder(v string) error {
	return a.SetProperty("Folder", v)
}

// GetFolder get Folder value
func (a *PhonebookAccess1) GetFolder() (string, error) {
	v, err := a.GetProperty("Folder")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetDatabaseIdentifier set DatabaseIdentifier value
func (a *PhonebookAccess1) SetDatabaseIdentifier(v string) error {
	return a.SetProperty("DatabaseIdentifier", v)
}

// GetDatabaseIdentifier get DatabaseIdentifier value
func (a *PhonebookAccess1) GetDatabaseIdentifier() (string, error) {
	v, err := a.GetProperty("DatabaseIdentifier")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetPrimaryCounter set PrimaryCounter value
func (a *PhonebookAccess1) SetPrimaryCounter(v string) error {
	return a.SetProperty("PrimaryCounter", v)
}

// GetPrimaryCounter get PrimaryCounter value
func (a *PhonebookAccess1) GetPrimaryCounter() (string, error) {
	v, err := a.GetProperty("PrimaryCounter")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}

// SetSecondaryCounter set SecondaryCounter value
func (a *PhonebookAccess1) SetSecondaryCounter(v string) error {
	return a.SetProperty("SecondaryCounter", v)
}

// GetSecondaryCounter get SecondaryCounter value
func (a *PhonebookAccess1) GetSecondaryCounter() (string, error) {
	v, err := a.GetProperty("SecondaryCounter")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}


// Close the connection
func (a *PhonebookAccess1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return PhonebookAccess1 object path
func (a *PhonebookAccess1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Interface return PhonebookAccess1 interface
func (a *PhonebookAccess1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *PhonebookAccess1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a PhonebookAccess1Properties to map
func (a *PhonebookAccess1Properties) ToMap() (map[string]interface{}, error) {
	return structs.Map(a), nil
}

// FromMap convert a map to an PhonebookAccess1Properties
func (a *PhonebookAccess1Properties) FromMap(props map[string]interface{}) (*PhonebookAccess1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an PhonebookAccess1Properties
func (a *PhonebookAccess1Properties) FromDBusMap(props map[string]dbus.Variant) (*PhonebookAccess1Properties, error) {
	s := new(PhonebookAccess1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// GetProperties load all available properties
func (a *PhonebookAccess1) GetProperties() (*PhonebookAccess1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *PhonebookAccess1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *PhonebookAccess1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *PhonebookAccess1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *PhonebookAccess1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *PhonebookAccess1) WatchProperties() (chan *bluez.PropertyChanged, error) {

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

func (a *PhonebookAccess1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	ch <- nil
	close(ch)
	return nil
}




/*
Select Select the phonebook object for other operations. Should
be call before all the other operations.
location : Where the phonebook is stored, possible
inputs :
"int" ( "internal" which is default )
"sim" ( "sim1" )
"sim2"
...
phonebook : Possible inputs :
"pb" :	phonebook for the saved contacts
"ich":	incoming call history
"och":	outgoing call history
"mch":	missing call history
"cch":	combination of ich och mch
"spd":	speed dials entry ( only for "internal" )
"fav":	favorites entry ( only for "internal" )
Possible errors: org.bluez.obex.Error.InvalidArguments
org.bluez.obex.Error.Failed
*/
func (a *PhonebookAccess1) Select(location string, phonebook string) error {
	
	return a.client.Call("Select", 0, location, phonebook).Store()
	
}

/*
PullAll Return the entire phonebook object from the PSE server
in plain string with vcard format, and store it in
a local file.
If an empty target file is given, a name will be
automatically calculated for the temporary file.
The returned path represents the newly created transfer,
which should be used to find out if the content has been
successfully transferred or if the operation fails.
The properties of this transfer are also returned along
with the object path, to avoid a call to GetProperties.
Possible filters: Format, Order, Offset, MaxCount and
Fields
Possible errors: org.bluez.obex.Error.InvalidArguments
org.bluez.obex.Forbidden
*/
func (a *PhonebookAccess1) PullAll(targetfile string, filters map[string]interface{}) (dbus.ObjectPath, map[string]interface{}, error) {
	
	var val0 dbus.ObjectPath
  var val1 map[string]interface{}
	err := a.client.Call("PullAll", 0, targetfile, filters).Store(&val0, &val1)
	return val0, val1, err	
}

/*
Pull Given a vcard handle, retrieve the vcard in the current
phonebook object and store it in a local file.
If an empty target file is given, a name will be
automatically calculated for the temporary file.
The returned path represents the newly created transfer,
which should be used to find out if the content has been
successfully transferred or if the operation fails.
The properties of this transfer are also returned along
with the object path, to avoid a call to GetProperties.
Possbile filters: Format and Fields
Possible errors: org.bluez.obex.Error.InvalidArguments
org.bluez.obex.Error.Forbidden
org.bluez.obex.Error.Failed
*/
func (a *PhonebookAccess1) Pull(vcard string, targetfile string, filters map[string]interface{}) error {
	
	return a.client.Call("Pull", 0, vcard, targetfile, filters).Store()
	
}

/*
Search Search for entries matching the given condition and
return an array of vcard-listing data where every entry
consists of a pair of strings containing the vcard
handle and the contact name.
vcard : name paired string match the search condition.
field : the field in the vcard to search with
{ "name" (default) | "number" | "sound" }
value : the string value to search for
Possible filters: Order, Offset and MaxCount
Possible errors: org.bluez.obex.Error.InvalidArguments
org.bluez.obex.Error.Forbidden
org.bluez.obex.Error.Failed
*/
func (a *PhonebookAccess1) Search(field string, value string, filters map[string]interface{}) error {
	
	return a.client.Call("Search", 0, field, value, filters).Store()
	
}

/*
GetSize Return the number of entries in the selected phonebook
object that are actually used (i.e. indexes that
correspond to non-NULL entries).
Possible errors: org.bluez.obex.Error.Forbidden
org.bluez.obex.Error.Failed
*/
func (a *PhonebookAccess1) GetSize() (uint16, error) {
	
	var val0 uint16
	err := a.client.Call("GetSize", 0, ).Store(&val0)
	return val0, err	
}

/*
UpdateVersion Attempt to update PrimaryCounter and SecondaryCounter.
Possible errors: org.bluez.obex.Error.NotSupported
org.bluez.obex.Error.Forbidden
org.bluez.obex.Error.Failed
*/
func (a *PhonebookAccess1) UpdateVersion() error {
	
	return a.client.Call("UpdateVersion", 0, ).Store()
	
}

/*
ListFilterFields Return All Available fields that can be used in Fields
filter.
Possible errors: None
Filter:		string Format:
Items vcard format
Possible values: "vcard21" (default) or "vcard30"
string Order:
Items order
Possible values: "indexed" (default), "alphanumeric" or
"phonetic"
uint16 Offset:
Offset of the first item, default is 0
uint16 MaxCount:
Maximum number of items, default is unlimited (65535)
array{string} Fields:
Item vcard fields, default is all values.
Possible values can be query with ListFilterFields.
array{string} FilterAll:
Filter items by fields using AND logic, cannot be used
together with FilterAny.
Possible values can be query with ListFilterFields.
array{string} FilterAny:
Filter items by fields using OR logic, cannot be used
together with FilterAll.
Possible values can be query with ListFilterFields.
bool ResetNewMissedCalls
Reset new the missed calls items, shall only be used
for folders mch and cch.
*/
func (a *PhonebookAccess1) ListFilterFields() ([]string, error) {
	
	var val0 []string
	err := a.client.Call("ListFilterFields", 0, ).Store(&val0)
	return val0, err	
}

