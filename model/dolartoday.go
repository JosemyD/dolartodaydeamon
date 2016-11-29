package model

type DolarToday struct {
	Antibloqueo Antibloqueo `json:"_antibloqueo"`
	Labels      Labels      `json:"_labels"`
	Timestamp   Timestamp   `json:"_timestamp"`
	USD         Moneda      `json:"USD"`
	EUR         Moneda      `json:"EUR"`
	COL         Col         `json:"COL"`
	GOLD        Rate        `json:"GOLD"`
	USDVEF      Rate        `json:"USDVEF"`
	USDCOL      Usdcol      `json:"USDCOL"`
	EURUSD      Rate        `json:"EURUSD"`
	BCV         Bcv         `json:"BCV"`
	MISC        Misc        `json:"MISC"`
}

type Antibloqueo struct {
	Mobile                    string `json:"mobile"`
	Video                     string `json:"video"`
	Corto_alternativo         string `json:"corto_alternativo"`
	Enable_iads               string `json:"enable_iads"`
	Enable_admobbanners       string `json:"enable_admobbanners"`
	Enable_admobinterstitials string `json:"enable_admobinterstitials"`
	Alternativo               string `json:"alternativo"`
	Alternativo2              string `json:"alternativo2"`
	Notifications             string `json:"notifications"`
	Resource_id               string `json:"resource_id"`
}

type Labels struct {
	A  string `json:"a"`
	A1 string `json:"a1"`
	B  string `json:"b"`
	C  string `json:"c"`
	D  string `json:"d"`
	E  string `json:"e"`
}

type Timestamp struct {
	Epoch        string `json:"epoch"`
	Fecha        string `json:"fecha"`
	Fecha_corta  string `json:"fecha_corta"`
	Fecha_corta2 string `json:"fecha_corta2"`
	Fecha_nice   string `json:"fecha_nice"`
	Dia          string `json:"dia"`
	Dia_corta    string `json:"dia_corta"`
}

type Moneda struct {
	Transferencia   float32 `json:"transferencia"`
	Transfer_cucuta float32 `json:"transfer_cucuta"`
	Efectivo        float32 `json:"efectivo"`
	Efectivo_real   float32 `json:"efectivo_real"`
	Efectivo_cucuta float32 `json:"efectivo_cucuta"`
	Promedio        float32 `json:"promedio"`
	Promedio_real   float32 `json:"promedio_real"`
	Cencoex         float32 `json:"cencoex"`
	Sicad1          float32 `json:"sicad1"`
	Sicad2          float32 `json:"sicad2"`
	Dolartoday      float32 `json:"dolartoday"`
}

type Col struct {
	Efectivo float32 `json:"efectivo"`
	Transfer float32 `json:"transfer"`
	Compra   float32 `json:"compra"`
	Venta    float32 `json:"venta"`
}

type Rate struct {
	Rate float32 `json:"rate"`
}

type Usdcol struct {
	Setfxsell     float32 `json:"setfxsell"`
	Setfxbuy      float32 `json:"setfxbuy"`
	Rate          float32 `json:"rate"`
	Ratecash      float32 `json:"ratecash"`
	Ratetrm       float32 `json:"ratetrm"`
	Trmfactor     float32 `json:"trmfactor"`
	Trmfactorcash float32 `json:"trmfactorcash"`
}

type Bcv struct {
	Fecha      string `json:"fecha"`
	Fecha_nice string `json:"fecha_nice"`
	Liquidez   string `json:"liquidez"`
	Reservas   string `json:"reservas"`
}

type Misc struct {
	Petroleo string `json:"petroleo"`
	Reservas string `json:"reservas"`
}
