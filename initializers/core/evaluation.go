package core

import (
	"fmt"
	_ "log"
	"time"

	"./benchmarking"
	"example.com/packages/performance"
	"example.com/packages/reporting"
)

type Performer struct {
	Name    string
	KPIs    []performance.KPI
	Manager performance.Manager
}

func getDatabaseConnection() {
	panic("unimplemented")
}

type EvaluationResult struct {
	Performer Performer
	Report    reporting.Report
}

func EvaluatePerformer(p Performer) (*EvaluationResult, error) {
	// Step 1: Set up the benchmarking process
	bm := benchmarking.NewBenchmarkingProcess(p.Manager)

	// Step 2: Run the benchmarking process
	scores, err := bm.RunBenchmarkingProcess(p.KPIs)
	if err != nil {
		return nil, fmt.Errorf("error running benchmarking process: %v", err)
	}

	// Step 3: Generate a report
	report := reporting.Report{
		PerformerName: p.Name,
		BenchmarkDate: time.Now(),
		Scores:        scores,
	}

	// Generate a console report
	consoleReporter := reporting.NewConsoleReporter()
	consoleReport, err := consoleReporter.GenerateReport(report)
	if err != nil {
		return nil, fmt.Errorf("error generating console report: %v", err)
	}
	fmt.Println(consoleReport)

	// Generate an email report
	emailReporter := reporting.NewEmailReporter(
		"framanreubinsten@gmail.com",
		"mndemefuraha00@gmail.com",
		"Benchmark Report",
		"smtp.example.com",
		"username",
		"password",
		587,
	)
	_, err = emailReporter.GenerateReport(report)
	if err != nil {
		return nil, fmt.Errorf("error generating email report: %v", err)
	}

	// Generate a PDF report
	pdfReporter := reporting.NewPDFReporter("report.pdf")
	_, err = pdfReporter.GenerateReport(report)
	if err != nil {
		return nil, fmt.Errorf("error generating PDF report: %v", err)
	}

	return &EvaluationResult{
		Performer: p,
		Report:    report,
	}, nil
}

func Evaluate() {
	// Create a performer to evaluate
	kpis := []performance.KPI{
		{Name: "Project Completion", Weight: 2},
		{Name: "Budget Management", Weight: 1},
		{Name: "Communication", Weight: 1},
		{Name: "Team Management", Weight: 1},
		{Name: "Stakeholder Management", Weight: 1},
	}
	manager := performance.NewMockManager()
	performer := Performer{
		Name:    "John Doe",
		KPIs:    kpis,
		Manager: manager,
	}

	// Evaluate the performer
	result, err := EvaluatePerformer(performer)
	if err != nil {
		fmt.Printf("Error evaluating performer: %v", err)
		return
	}

	fmt.Printf("Evaluation complete for performer %s:\n", result.Performer.Name)
	for k, v := range result.Report.Scores {
		fmt.Printf("- %s: %d\n", k, v)
	}
}
