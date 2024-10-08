/*
MIT License

# Copyright (c) 2024 Aaron Saikovski

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package monitordetail

// InverterData - Struct to hold data returned from the Inverter Powerstation API
type MonitorData struct {
	Language string   `json:"language"`
	Function []string `json:"function"`
	HasError bool     `json:"hasError"`
	Msg      string   `json:"msg"`
	Code     string   `json:"code"`
	Data     struct {
		Kpi struct {
			MonthGeneration float64 `json:"month_generation"`
			Pac             float64 `json:"pac"`
			Power           float64 `json:"power"`
			TotalPower      float64 `json:"total_power"`
			DayIncome       float64 `json:"day_income"`
			TotalIncome     float64 `json:"total_income"`
			YieldRate       float64 `json:"yield_rate"`
			Currency        string  `json:"currency"`
		} `json:"kpi"`
		PowercontrolStatus int   `json:"powercontrol_status"`
		Images             []any `json:"images"`
		Inverter           []struct {
			IsStored   bool    `json:"is_stored"`
			InPac      float64 `json:"in_pac"`
			OutPac     float64 `json:"out_pac"`
			Eday       float64 `json:"eday"`
			Emonth     float64 `json:"emonth"`
			Etotal     float64 `json:"etotal"`
			Status     int     `json:"status"`
			TurnonTime string  `json:"turnon_time"`
			Type       string  `json:"type"`
			Capacity   float64 `json:"capacity"`
			D          struct {
				Capacity              string  `json:"capacity"`
				Model                 string  `json:"model"`
				OutputPower           string  `json:"output_power"`
				OutputCurrent         string  `json:"output_current"`
				GridVoltage           string  `json:"grid_voltage"`
				BackupOutput          string  `json:"backup_output"`
				Soc                   string  `json:"soc"`
				Soh                   string  `json:"soh"`
				LastRefreshTime       string  `json:"last_refresh_time"`
				WorkMode              string  `json:"work_mode"`
				DcInput1              string  `json:"dc_input1"`
				DcInput2              string  `json:"dc_input2"`
				Battery               string  `json:"battery"`
				BmsStatus             string  `json:"bms_status"`
				Warning               string  `json:"warning"`
				ChargeCurrentLimit    string  `json:"charge_current_limit"`
				DischargeCurrentLimit string  `json:"discharge_current_limit"`
				FirmwareVersion       float64 `json:"firmware_version"`
				CreationDate          string  `json:"creationDate"`
				EDay                  float64 `json:"eDay"`
				ETotal                float64 `json:"eTotal"`
				Pac                   float64 `json:"pac"`
				HTotal                float64 `json:"hTotal"`
				Vpv1                  float64 `json:"vpv1"`
				Vpv2                  float64 `json:"vpv2"`
				Vpv3                  float64 `json:"vpv3"`
				Vpv4                  float64 `json:"vpv4"`
				Ipv1                  float64 `json:"ipv1"`
				Ipv2                  float64 `json:"ipv2"`
				Ipv3                  float64 `json:"ipv3"`
				Ipv4                  float64 `json:"ipv4"`
				Vac1                  float64 `json:"vac1"`
				Vac2                  float64 `json:"vac2"`
				Vac3                  float64 `json:"vac3"`
				Iac1                  float64 `json:"iac1"`
				Iac2                  float64 `json:"iac2"`
				Iac3                  float64 `json:"iac3"`
				Fac1                  float64 `json:"fac1"`
				Fac2                  float64 `json:"fac2"`
				Fac3                  float64 `json:"fac3"`
				Istr1                 float64 `json:"istr1"`
				Istr2                 float64 `json:"istr2"`
				Istr3                 float64 `json:"istr3"`
				Istr4                 float64 `json:"istr4"`
				Istr5                 float64 `json:"istr5"`
				Istr6                 float64 `json:"istr6"`
				Istr7                 float64 `json:"istr7"`
				Istr8                 float64 `json:"istr8"`
				Istr9                 float64 `json:"istr9"`
				Istr10                float64 `json:"istr10"`
				Istr11                float64 `json:"istr11"`
				Istr12                float64 `json:"istr12"`
				Istr13                float64 `json:"istr13"`
				Istr14                float64 `json:"istr14"`
				Istr15                float64 `json:"istr15"`
				Istr16                float64 `json:"istr16"`
			} `json:"d"`
			ItChangeFlag bool    `json:"it_change_flag"`
			Temperature  float64 `json:"tempperature"`
			CheckCode    string  `json:"check_code"`
			Next         any     `json:"next"`
			Prev         any     `json:"prev"`
			NextDevice   struct {
				Sn       any  `json:"sn"`
				IsStored bool `json:"isStored"`
			} `json:"next_device"`
			PrevDevice struct {
				Sn       any  `json:"sn"`
				IsStored bool `json:"isStored"`
			} `json:"prev_device"`
			InvertFull struct {
				CtSolutionType          int     `json:"ct_solution_type"`
				Cts                     any     `json:"cts"`
				CheckCode               string  `json:"check_code"`
				ModelType               string  `json:"model_type"`
				ChangeType              int     `json:"change_type"`
				ChangeTime              int     `json:"change_time"`
				Capacity                float64 `json:"capacity"`
				Eday                    float64 `json:"eday"`
				Iday                    float64 `json:"iday"`
				Etotal                  float64 `json:"etotal"`
				Itotal                  float64 `json:"itotal"`
				HourTotal               float64 `json:"hour_total"`
				Status                  int     `json:"status"`
				TurnonTime              int64   `json:"turnon_time"`
				Pac                     float64 `json:"pac"`
				Tempperature            float64 `json:"tempperature"`
				Vpv1                    float64 `json:"vpv1"`
				Vpv2                    float64 `json:"vpv2"`
				Vpv3                    float64 `json:"vpv3"`
				Vpv4                    float64 `json:"vpv4"`
				Ipv1                    float64 `json:"ipv1"`
				Ipv2                    float64 `json:"ipv2"`
				Ipv3                    float64 `json:"ipv3"`
				Ipv4                    float64 `json:"ipv4"`
				Vac1                    float64 `json:"vac1"`
				Vac2                    float64 `json:"vac2"`
				Vac3                    float64 `json:"vac3"`
				Iac1                    float64 `json:"iac1"`
				Iac2                    float64 `json:"iac2"`
				Iac3                    float64 `json:"iac3"`
				Fac1                    float64 `json:"fac1"`
				Fac2                    float64 `json:"fac2"`
				Fac3                    float64 `json:"fac3"`
				Istr1                   float64 `json:"istr1"`
				Istr2                   float64 `json:"istr2"`
				Istr3                   float64 `json:"istr3"`
				Istr4                   float64 `json:"istr4"`
				Istr5                   float64 `json:"istr5"`
				Istr6                   float64 `json:"istr6"`
				Istr7                   float64 `json:"istr7"`
				Istr8                   float64 `json:"istr8"`
				Istr9                   float64 `json:"istr9"`
				Istr10                  float64 `json:"istr10"`
				Istr11                  float64 `json:"istr11"`
				Istr12                  float64 `json:"istr12"`
				Istr13                  float64 `json:"istr13"`
				Istr14                  float64 `json:"istr14"`
				Istr15                  float64 `json:"istr15"`
				Istr16                  float64 `json:"istr16"`
				LastTime                int64   `json:"last_time"`
				Vbattery1               float64 `json:"vbattery1"`
				Ibattery1               float64 `json:"ibattery1"`
				Pmeter                  float64 `json:"pmeter"`
				Soc                     float64 `json:"soc"`
				Soh                     float64 `json:"soh"`
				BmsDischargeIMax        any     `json:"bms_discharge_i_max"`
				BmsChargeIMax           float64 `json:"bms_charge_i_max"`
				BmsWarning              int     `json:"bms_warning"`
				BmsAlarm                int     `json:"bms_alarm"`
				BattaryWorkMode         int     `json:"battary_work_mode"`
				Workmode                int     `json:"workmode"`
				Vload                   float64 `json:"vload"`
				Iload                   float64 `json:"iload"`
				Firmwareversion         float64 `json:"firmwareversion"`
				Bmssoftwareversion      any     `json:"bmssoftwareversion"`
				Pbackup                 float64 `json:"pbackup"`
				Seller                  float64 `json:"seller"`
				Buy                     float64 `json:"buy"`
				Yesterdaybuytotal       float64 `json:"yesterdaybuytotal"`
				Yesterdaysellertotal    float64 `json:"yesterdaysellertotal"`
				Yesterdayct2Sellertotal any     `json:"yesterdayct2sellertotal"`
				Yesterdayetotal         any     `json:"yesterdayetotal"`
				Yesterdayetotalload     float64 `json:"yesterdayetotalload"`
				Yesterdaylastime        int     `json:"yesterdaylastime"`
				Thismonthetotle         float64 `json:"thismonthetotle"`
				Lastmonthetotle         float64 `json:"lastmonthetotle"`
				RAM                     float64 `json:"ram"`
				Outputpower             float64 `json:"outputpower"`
				FaultMessge             int     `json:"fault_messge"`
				Battery1Sn              any     `json:"battery1sn"`
				Battery2Sn              any     `json:"battery2sn"`
				Battery3Sn              any     `json:"battery3sn"`
				Battery4Sn              any     `json:"battery4sn"`
				Battery5Sn              any     `json:"battery5sn"`
				Battery6Sn              any     `json:"battery6sn"`
				Battery7Sn              any     `json:"battery7sn"`
				Battery8Sn              any     `json:"battery8sn"`
				Pf                      float64 `json:"pf"`
				PvPower                 float64 `json:"pv_power"`
				ReactivePower           float64 `json:"reactive_power"`
				LeakageCurrent          int     `json:"leakage_current"`
				IsoLimit                int     `json:"isoLimit"`
				Isbuettey               bool    `json:"isbuettey"`
				Isbuetteybps            bool    `json:"isbuetteybps"`
				Isbuetteybpu            bool    `json:"isbuetteybpu"`
				IsESUOREMU              bool    `json:"isESUOREMU"`
				BackUpPloadS            float64 `json:"backUpPload_S"`
				BackUpVloadS            float64 `json:"backUpVload_S"`
				BackUpIloadS            float64 `json:"backUpIload_S"`
				BackUpPloadT            float64 `json:"backUpPload_T"`
				BackUpVloadT            float64 `json:"backUpVload_T"`
				BackUpIloadT            float64 `json:"backUpIload_T"`
				ETotalBuy               any     `json:"eTotalBuy"`
				EDayBuy                 float64 `json:"eDayBuy"`
				EBatteryCharge          any     `json:"eBatteryCharge"`
				EChargeDay              float64 `json:"eChargeDay"`
				EBatteryDischarge       any     `json:"eBatteryDischarge"`
				EDischargeDay           float64 `json:"eDischargeDay"`
				BattStrings             any     `json:"battStrings"`
				MeterConnectStatus      any     `json:"meterConnectStatus"`
				MtActivepowerR          float64 `json:"mtActivepowerR"`
				MtActivepowerS          float64 `json:"mtActivepowerS"`
				MtActivepowerT          float64 `json:"mtActivepowerT"`
				EzProConnectStatus      any     `json:"ezPro_connect_status"`
				Dataloggersn            string  `json:"dataloggersn"`
				EquipmentName           any     `json:"equipment_name"`
				Hasmeter                bool    `json:"hasmeter"`
				MeterType               any     `json:"meter_type"`
				PreHourLasttotal        any     `json:"pre_hour_lasttotal"`
				PreHourTime             any     `json:"pre_hour_time"`
				CurrentHourPv           float64 `json:"current_hour_pv"`
				ExtendProperties        any     `json:"extend_properties"`
				EPConnectStatusHappen   any     `json:"eP_connect_status_happen"`
				EPConnectStatusRecover  any     `json:"eP_connect_status_recover"`
				TotalSell               float64 `json:"total_sell"`
				TotalBuy                float64 `json:"total_buy"`
				Errors                  []any   `json:"errors"`
				SafetyConutry           float64 `json:"safetyConutry"`
				DeratingMode            any     `json:"deratingMode"`
				Master                  any     `json:"master"`
				ParallelCode            any     `json:"parallel_code"`
				InverterType            int     `json:"inverter_type"`
				TotalPbattery           float64 `json:"total_pbattery"`
				BatteryCount            any     `json:"battery_count"`
				MoreBatterys            any     `json:"more_batterys"`
				GensetStartMode         any     `json:"genset_start_mode"`
				GensetMode              any     `json:"genset_mode"`
				GensetEtotal            any     `json:"genset_etotal"`
				GensetYeasterdayTotal   any     `json:"genset_yeasterdayTotal"`
				GensetPower             any     `json:"genset_power"`
				GensetEday              float64 `json:"genset_eday"`
				AllEday                 float64 `json:"all_eday"`
			} `json:"invert_full"`
			Time                     string  `json:"time"`
			Battery                  string  `json:"battery"`
			FirmwareVersion          float64 `json:"firmware_version"`
			WarningBms               string  `json:"warning_bms"`
			Soh                      string  `json:"soh"`
			DischargeCurrentLimitBms string  `json:"discharge_current_limit_bms"`
			ChargeCurrentLimitBms    string  `json:"charge_current_limit_bms"`
			Soc                      string  `json:"soc"`
			PvInput2                 string  `json:"pv_input_2"`
			PvInput1                 string  `json:"pv_input_1"`
			BackUpOutput             string  `json:"back_up_output"`
			OutputVoltage            string  `json:"output_voltage"`
			BackupVoltage            string  `json:"backup_voltage"`
			OutputCurrent            string  `json:"output_current"`
			OutputPower              string  `json:"output_power"`
			TotalGeneration          string  `json:"total_generation"`
			DailyGeneration          string  `json:"daily_generation"`
			BatteryCharging          string  `json:"battery_charging"`
			LastRefreshTime          string  `json:"last_refresh_time"`
			BmsStatus                string  `json:"bms_status"`
			FaultMessage             string  `json:"fault_message"`
			WarningCode              any     `json:"warning_code"`
			BatteryPower             float64 `json:"battery_power"`
			BackupPloadS             float64 `json:"backup_pload_s"`
			BackupVloadS             float64 `json:"backup_vload_s"`
			BackupIloadS             float64 `json:"backup_iload_s"`
			BackupPloadT             float64 `json:"backup_pload_t"`
			BackupVloadT             float64 `json:"backup_vload_t"`
			BackupIloadT             float64 `json:"backup_iload_t"`
			EtotalBuy                any     `json:"etotal_buy"`
			EdayBuy                  float64 `json:"eday_buy"`
			EbatteryCharge           any     `json:"ebattery_charge"`
			EchargeDay               float64 `json:"echarge_day"`
			EbatteryDischarge        any     `json:"ebattery_discharge"`
			EdischargeDay            float64 `json:"edischarge_day"`
			BattStrings              any     `json:"batt_strings"`
			MeterConnectStatus       any     `json:"meter_connect_status"`
			MtactivepowerR           float64 `json:"mtactivepower_r"`
			MtactivepowerS           float64 `json:"mtactivepower_s"`
			MtactivepowerT           float64 `json:"mtactivepower_t"`
			HasTigo                  bool    `json:"has_tigo"`
			CanStartIV               bool    `json:"canStartIV"`
			BatteryCount             any     `json:"battery_count"`
		} `json:"inverter"`
		Hjgx struct {
			Co2  float64 `json:"co2"`
			Tree float64 `json:"tree"`
			Coal float64 `json:"coal"`
		} `json:"hjgx"`
		HomKit struct {
			HomeKitLimit bool `json:"homeKitLimit"`
			Sn           any  `json:"sn"`
		} `json:"homKit"`
		IsTigo                 bool `json:"isTigo"`
		TigoIntervalTimeMinute int  `json:"tigoIntervalTimeMinute"`
		SmuggleInfo            struct {
			IsAllSmuggle    bool `json:"isAllSmuggle"`
			IsSmuggle       bool `json:"isSmuggle"`
			DescriptionText any  `json:"descriptionText"`
			Sns             any  `json:"sns"`
		} `json:"smuggleInfo"`
		HasPowerflow bool `json:"hasPowerflow"`
		Powerflow    struct {
			Battery                   string `json:"bettery"`
			BatteryStatus             int    `json:"betteryStatus"`
			BatteryStatusStr          any    `json:"betteryStatusStr"`
			Genset                    string `json:"genset"`
			GensetStatus              int    `json:"gensetStatus"`
			Grid                      string `json:"grid"`
			GridGensetStatus          int    `json:"gridGensetStatus"`
			GridStatus                int    `json:"gridStatus"`
			HasEquipment              bool   `json:"hasEquipment"`
			IsBpuAndInverterNoBattery bool   `json:"isBpuAndInverterNoBattery"`
			IsHomKit                  bool   `json:"isHomKit"`
			IsMoreBettery             bool   `json:"isMoreBettery"`
			Load                      string `json:"load"`
			LoadStatus                int    `json:"loadStatus"`
			Pv                        string `json:"pv"`
			PvStatus                  int    `json:"pvStatus"`
			Soc                       int    `json:"soc"`
			SocText                   string `json:"socText"`
		} `json:"powerflow"`
		HasGridLoad               bool `json:"hasGridLoad"`
		IsParallelInventers       bool `json:"isParallelInventers"`
		IsEvCharge                bool `json:"isEvCharge"`
		EvCharge                  any  `json:"evCharge"`
		HasEnergeStatisticsCharts bool `json:"hasEnergeStatisticsCharts"`
		EnergeStatisticsCharts    struct {
			ContributingRate  float64 `json:"contributingRate"`
			SelfUseRate       float64 `json:"selfUseRate"`
			Sum               float64 `json:"sum"`
			Buy               float64 `json:"buy"`
			BuyPercent        float64 `json:"buyPercent"`
			Sell              float64 `json:"sell"`
			SellPercent       float64 `json:"sellPercent"`
			SelfUseOfPv       float64 `json:"selfUseOfPv"`
			ConsumptionOfLoad float64 `json:"consumptionOfLoad"`
			ChartsType        int     `json:"chartsType"`
			HasPv             bool    `json:"hasPv"`
			HasCharge         bool    `json:"hasCharge"`
			Charge            float64 `json:"charge"`
			DisCharge         float64 `json:"disCharge"`
			GensetGen         float64 `json:"gensetGen"`
			HasGenset         bool    `json:"hasGenset"`
		} `json:"energeStatisticsCharts"`
		EnergeStatisticsTotals struct {
			ContributingRate  float64 `json:"contributingRate"`
			SelfUseRate       float64 `json:"selfUseRate"`
			Sum               float64 `json:"sum"`
			Buy               float64 `json:"buy"`
			BuyPercent        float64 `json:"buyPercent"`
			Sell              float64 `json:"sell"`
			SellPercent       float64 `json:"sellPercent"`
			SelfUseOfPv       float64 `json:"selfUseOfPv"`
			ConsumptionOfLoad float64 `json:"consumptionOfLoad"`
			ChartsType        int     `json:"chartsType"`
			HasPv             bool    `json:"hasPv"`
			HasCharge         bool    `json:"hasCharge"`
			Charge            float64 `json:"charge"`
			DisCharge         float64 `json:"disCharge"`
			GensetGen         float64 `json:"gensetGen"`
			HasGenset         bool    `json:"hasGenset"`
		} `json:"energeStatisticsTotals"`
		Soc struct {
			Power  int `json:"power"`
			Status int `json:"status"`
		} `json:"soc"`
		Environmental []any `json:"environmental"`
		Equipment     []struct {
			Type                 string `json:"type"`
			Status               int    `json:"status"`
			Model                any    `json:"model"`
			StatusText           any    `json:"statusText"`
			Capacity             any    `json:"capacity"`
			ActionThreshold      any    `json:"actionThreshold"`
			SubordinateEquipment string `json:"subordinateEquipment"`
			PowerGeneration      string `json:"powerGeneration"`
			Eday                 string `json:"eday"`
			Brand                string `json:"brand"`
			IsStored             bool   `json:"isStored"`
			Soc                  string `json:"soc"`
			IsChange             bool   `json:"isChange"`
			HasTigo              bool   `json:"has_tigo"`
			IsSec                bool   `json:"is_sec"`
			IsSecs               bool   `json:"is_secs"`
			TargetPF             any    `json:"targetPF"`
			ExportPowerlimit     any    `json:"exportPowerlimit"`
			TitleSn              any    `json:"titleSn"`
		} `json:"equipment"`
	} `json:"data"`
}

func NewMonitorData() *MonitorData {
	return &MonitorData{}
}
