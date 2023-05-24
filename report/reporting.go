package report

import (
	"fmt"
	_ "io/ioutil"
	_ "os"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	"gopkg.in/gomail.v2"
)

type Report struct {
	PerformerName string
	BenchmarkDate time.Time
	Scores        map[string]int
}

type Reporter interface {
	GenerateReport(report Report) (string, error)
}

type ConsoleReporter struct{}

func NewConsoleReporter() *ConsoleReporter {
	return &ConsoleReporter{}
}

func (cr *ConsoleReporter) GenerateReport(report Report) (string, error) {
	sb := strings.Builder{}

	sb.WriteString(fmt.Sprintf("Benchmark report for %s on %s:\n", report.PerformerName, report.BenchmarkDate.Format("2006-01-02")))
	sb.WriteString("KPI scores:\n")
	for k, v := range report.Scores {
		sb.WriteString(fmt.Sprintf("- %s: %d\n", k, v))
	}

	return sb.String(), nil
}

type EmailReporter struct {
	From     string
	To       string
	Subject  string
	SMTPHost string
	SMTPPort int
	Username string
	Password string
}

func NewEmailReporter(from, to, subject, smtpHost, username, password string, smtpPort int) *EmailReporter {
	return &EmailReporter{
		From:     from,
		To:       to,
		Subject:  subject,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
		Username: username,
		Password: password,
	}
}

func (er *EmailReporter) GenerateReport(report Report) (string, error) {
	body, err := NewConsoleReporter().GenerateReport(report)
	if err != nil {
		return "", err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", er.From)
	m.SetHeader("To", er.To)
	m.SetHeader("Subject", er.Subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(er.SMTPHost, er.SMTPPort, er.Username, er.Password)

	if err := d.DialAndSend(m); err != nil {
		return "", err
	}

	return body, nil
}

type PDFReporter struct {
	FilePath string
}

func NewPDFReporter(filePath string) *PDFReporter {
	return &PDFReporter{
		FilePath: filePath,
	}
}

func (pr *PDFReporter) GenerateReport(report Report) (string, error) {
	body, err := NewConsoleReporter().GenerateReport(report)
	if err != nil {
		return "", err
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Benchmark Report")
	pdf.Ln(20)

	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(40, 10, "Performer Name:", "", 0, "", false, 0, "")
	pdf.CellFormat(0, 10, report.PerformerName, "", 1, "", false, 0, "")
	pdf.CellFormat(40, 10, "Benchmark Date:", "", 0, "", false, 0, "")
	pdf.CellFormat(0, 10, report.BenchmarkDate.Format("2006-01-02"), "", 1, "", false, 0, "")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 10, "KPI Scores:")
	return body, nil

}
