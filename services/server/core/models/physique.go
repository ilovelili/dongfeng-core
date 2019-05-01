package models

import (
	"fmt"
	"math"
	"time"

	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// Physique physique entity
type Physique struct {
	ID            int64   `dapper:"id,primarykey,autoincrement,table=physiques"`
	Year          string  `dapper:"year"`
	Class         string  `dapper:"class"`
	Name          string  `dapper:"name"`
	Gender        int64   `dapper:"gender"`
	BirthDate     string  `dapper:"birth_date"`
	ExamDate      string  `dapper:"exam_date"`
	Age           string  `dapper:"age"`
	AgeComparison float64 `dapper:"age_cmp"`
	Height        float64 `dapper:"height"`
	HeightP       string  `dapper:"height_p"` // Height P Zone
	Weight        float64 `dapper:"weight"`
	WeightP       string  `dapper:"weight_p"` // Weight P Zone
	BMI           float64 `dapper:"bmi"`
	CreatedBy     string  `dapper:"created_by"`
}

// PMaster p standard master
type PMaster struct {
	ID             int64   `dapper:"id,primarykey,autoincrement,table=physique_p_master"`
	HeightOrWeight string  `dapper:"h_w"`
	Gender         int64   `dapper:"gender"`
	AgeMin         float64 `dapper:"age_min"`
	AgeMax         float64 `dapper:"age_max"`
	P3             float64 `dapper:"p3"`
	P10            float64 `dapper:"p10"`
	P20            float64 `dapper:"p20"`
	P50            float64 `dapper:"p50"`
	P97            float64 `dapper:"p97"`
}

// resolveAge diff by birth date and exam date
func (p *Physique) resolveAge() {
	birthdate, _ := time.Parse("2006-01-02", p.BirthDate)
	examdate, _ := time.Parse("2006-01-02", p.ExamDate)

	year, month, _, _, _, _ := sharedlib.Diff(birthdate, examdate)
	p.Age = fmt.Sprintf("%d年%d月", year, month)
	cmp := float64(year) + float64(month)/12.0
	p.AgeComparison = math.Round(cmp*100) / 100
}

// resolveBMI kg/m^2
func (p *Physique) resolveBMI() (err error) {
	bmi := p.Weight / (p.Height * p.Height)
	p.BMI = math.Round(bmi*100) / 100
	return
}

// resolveHeightP get the corresponding p zone
func (p *Physique) resolveHeightP(pmasters []*PMaster) (err error) {
	for _, m := range pmasters {
		if m.AgeMin <= p.AgeComparison && m.AgeMax >= p.AgeComparison && m.HeightOrWeight == "h" && m.Gender == p.Gender {
			if p.Height <= m.P3 {
				p.HeightP = "<P3"
				return
			}

			if p.Height <= m.P10 {
				p.HeightP = "P3~P10"
				return
			}

			if p.Height <= m.P20 {
				p.HeightP = "P10~P20"
				return
			}

			if p.Height <= m.P50 {
				p.HeightP = "P20~P50"
				return
			}

			if p.Height <= m.P97 {
				p.HeightP = "P50~P97"
				return
			}

			p.HeightP = ">P97"
			return
		}
	}

	return fmt.Errorf("P standard master data not found")
}
