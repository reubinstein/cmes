package perform

import (
	"fmt"
	_ "fmt"
	"time"
)

type PerformanceStandard struct {
	KPI         string
	GoodPerf    string
	Measurement string
}

func GetPerformanceStandards() []PerformanceStandard {
	return []PerformanceStandard{
		{"Bills proposed", "Proposing a high number of bills that address key societal challenges in the constituency.", "Count the number of bills proposed by the MP or local counselor during a given period."},
		{"Bills passed", "A high percentage of the bills proposed by the MP or local counselor being passed into law.", "Calculate the percentage of bills proposed by the MP or local counselor that are passed into law during a given period."},
		{"Funding secured", "Securing a high amount of funding for projects aimed at addressing societal challenges in the constituency.", "Calculate the amount of funding secured by the MP or local counselor for projects during a given period."},
		{"Projects completed", "Completing a high number of projects aimed at addressing societal challenges in the constituency.", "Count the number of projects completed by the MP or local counselor during a given period."},
		{"Public engagement", "Achieving a high level of public engagement through various means such as town hall meetings, social media, petitions, .", "Measure the level of engagement achieved by the MP or local counselor through metrics such as the number of town hall meetings held, the number of constituents reached through social media, or the number of petitions signed."},
		{"Adherence to CCM policies", "Adhering to the policies and guidelines of the political party or coalition they belong to.", "Count the number of times an MP or local counselor votes in line with the party's position on a given issue or the number of times they attend party meetings."},
	}
}

type PerformanceData struct {
	KPI         string
	Value       int
	Measurement string
	Date        time.Time
}

type Performer interface {
	GetPerformance() []PerformanceData
}

type MP struct {
	Name         string
	Constituency string
}

func (mp *MP) GetPerformance() []PerformanceData {
	return []PerformanceData{
		{KPI: "Bills proposed", Value: 25, Measurement: "Count"},
		{KPI: "Bills passed", Value: 20, Measurement: "Percentage"},
		{KPI: "Funding secured", Value: 5000000, Measurement: "Amount"},
		{KPI: "Projects completed", Value: 10, Measurement: "Count"},
		{KPI: "Public engagement", Value: 2000, Measurement: "Count"},
		{KPI: "Adherence to CCM policies", Value: 80, Measurement: "Percentage"},
	}
}

type LocalCounselor struct {
	Name string
	Ward string
}

func (counselor *LocalCounselor) GetPerformance() []PerformanceData {
	return []PerformanceData{
		{KPI: "Bills proposed", Value: 10, Measurement: "Count"},
		{KPI: "Bills passed", Value: 5, Measurement: "Percentage"},
		{KPI: "Funding secured", Value: 1000000, Measurement: "Amount"},
		{KPI: "Projects completed", Value: 5, Measurement: "Count"},
		{KPI: "Public engagement", Value: 500, Measurement: "Count"},
		{KPI: "Adherence to CCM policies", Value: 90, Measurement: "Percentage"},
	}
}

type counselor1 struct {
	name string
	ward string
}

type mp1 struct {
}
type mp2 struct {
}
type counselor2 struct {
	name string
	ward string
}

func PS() {
	// Get performance standards
	performanceStandards := GetPerformanceStandards()

	// Create MPs and local counselors
	mp1 := &MP{Name: "John Doe", Constituency: "Nairobi West"}
	mp2 := &MP{Name: "Jane Doe", Constituency: "Nairobi East"}
	counselor1 := &LocalCounselor{Name: "James Smith", Ward: "Kawangware"}
	counselor2 := &LocalCounselor{Name: "Susan Brown", Ward: "Mbale"}

	// Loop through MPs and local counselors and print their performance data
	for _, performer := range []Performer{mp1, mp2, counselor1, counselor2} {
		fmt.Printf("%s Performance Data:\n", performer.(*MP).Name)
		fmt.Printf("%-25s %-10s %-10s %-20s\n", "KPI", "Value", "Measurement", "Date")
		for _, data := range performer.GetPerformance() {
			fmt.Printf("%-25s %-10d %-10s %-20s\n", data.KPI, data.Value, data.Measurement, data.Date.Format("2006-01-02 15:04:05"))
		}
		fmt.Println()
	}

}
