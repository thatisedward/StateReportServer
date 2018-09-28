package propertyplus

import (
	ls "miglib/propertyplus/jce/LogStat"
	"reflect"
	"sync"
	"tars/protocol/serializer"
)

func toSkey(keys []string) string {
	_os := serializer.NewOutputStream()
	_os.Write(reflect.ValueOf(&keys), 0)
	return string(_os.ToBytes())
}

func toVkey(skey string) []string {

	var vkey []string

	_is := serializer.NewInputStream([]byte(skey))
	_i, _ := _is.Read(reflect.TypeOf(vkey), 0, true)

	return _i.([]string)
}

func toString(content ls.StatContent) string {
	_os := serializer.NewOutputStream()
	content.WriteTo(_os)
	return string(_os.ToBytes())
}

type PropertyMsg struct {
	reportData map[string][]ls.StatValue
	mlock      *sync.Mutex
}

func (pm *PropertyMsg) init() {
	pm.mlock = new(sync.Mutex)
	pm.reportData = make(map[string][]ls.StatValue)
}

func (pm *PropertyMsg) Report(keys []string, values []float32) {

	var tempkeys []string
	tempkeys = append(tempkeys, "")
	tempkeys = append(tempkeys, "")
	tempkeys = append(tempkeys, keys...)
	pm.merge(tempkeys, values)
}

func (pm *PropertyMsg) merge(keys []string, values []float32) {
	pm.mlock.Lock()
	defer pm.mlock.Unlock()

	sKey := toSkey(keys)

	reportData, ok := pm.reportData[sKey]
	if ok {
		if len(reportData) != len(values) {
			TLOG.Error("lost value info, %s, %s", len(reportData), len(values), reportData)
			return
		}
		for k, v := range values {
			reportData[k].Value += v
		}
		pm.reportData[sKey] = reportData
	} else {
		for _, v := range values {
			var stStat ls.StatValue
			stStat.Value = v
			stStat.Policy = "Sum"
			stStat.Count = 1
			pm.reportData[sKey] = append(pm.reportData[sKey], stStat)
		}
	}
}

func (pm *PropertyMsg) doReport(propertyName string) ls.StatLog {

	pm.mlock.Lock()
	defer pm.mlock.Unlock()

	var stat ls.StatLog
	stat.Logname = propertyName

	for k, v := range pm.reportData {
		var content ls.StatContent
		content.Keys = toVkey(k)
		content.Values = v
		scontent := toString(content)
		stat.Content = append(stat.Content, scontent)
	}
	pm.reportData = make(map[string][]ls.StatValue)

	return stat
}
