package models

type CutInfo struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	Duration string `json:"duration"`
	Mid      string `json:"mid"`
	Poster   string `json:"poster"`
	Path     string `orm:"pk";json:"path"`
	Movie    *Movie `orm:"null;rel(fk)"`
}

func GetCutInfoByPath(path string) *CutInfo {
	m := CutInfo{Path: path}
	err := ormOpr.Read(&m)
	if err != nil {
		return nil
	}
	return &m
}

func (c CutInfo) AddCutInfoByPath() error {
	_, e := ormOpr.Insert(&c)
	return e
}

func (c CutInfo) DelCutInfoByPath() error {
	_, e := ormOpr.Delete(&c)
	return e
}

func (m Movie) GetCutInfosByMId() ([]*CutInfo, error) {
	_, err := ormOpr.LoadRelated(&m, "CutInfos")
	if err != nil {
		return nil, err
	}
	return m.CutInfos, nil
}
