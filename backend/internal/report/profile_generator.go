package report

import (
	"fmt"
	"os"
	"path/filepath"
	"reporting-utility/internal/models"
	"time"

	"github.com/signintech/gopdf"
)

func GenerateStudentProfilesPDF(students []models.Student) (*gopdf.GoPdf, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	

	err := pdf.AddTTFFont("arial", "C:\\Windows\\Fonts\\arial.ttf")
	if err != nil {
		err = pdf.AddTTFFont("arial", "assets/arial.ttf")
		if err != nil {
			return nil, fmt.Errorf("font file missing: please put arial.ttf in assets/ folder")
		}
	}

	for i, s := range students {
		pdf.AddPage()

		// 1. Main Card Border
		pdf.SetLineWidth(0.8)
		pdf.SetStrokeColor(44, 62, 80) // FIXED: Use SetStrokeColor for the border
		pdf.RectFromUpperLeftWithStyle(30, 30, 535, 780, "D")

		// 2. Header
		pdf.Image("assets/kk-wagh-logo.png", 45, 45, &gopdf.Rect{W: 70, H: 35})
		pdf.SetFont("arial", "", 14)
		pdf.SetTextColor(0, 0, 0) // Ensure black text
		pdf.SetXY(130, 55)
		pdf.Text("K.K Wagh Institute Of Engineering and Research")
		
		pdf.SetLineWidth(0.5)
		pdf.SetStrokeColor(200, 200, 200) // Light grey line
		pdf.Line(45, 90, 550, 90)

		// 3. Photo Section
		photoPath := filepath.Join("uploads", s.Photo)
		pdf.SetLineWidth(1.5)
		pdf.SetStrokeColor(44, 62, 80)
		if _, err := os.Stat(photoPath); err == nil && s.Photo != "" {
			pdf.Image(photoPath, 45, 110, &gopdf.Rect{W: 110, H: 110})
			pdf.RectFromUpperLeftWithStyle(45, 110, 110, 110, "D") 
		} else {
			pdf.RectFromUpperLeftWithStyle(45, 110, 110, 110, "D")
			pdf.SetXY(65, 160)
			pdf.SetFont("arial", "", 10)
			pdf.Text("NO PHOTO")
		}

		// 4. Primary Info
		pdf.SetFont("arial", "", 20)
		pdf.SetXY(170, 125)
		pdf.Text(s.StudentName)

		pdf.SetFont("arial", "", 11)
		pdf.SetXY(170, 155)
		pdf.SetTextColor(100, 100, 100) // FIXED: Use SetTextColor for grey sub-text
		pdf.Text(s.Address)
		pdf.SetXY(170, 175)
		pdf.Text(fmt.Sprintf("%s, %s, %s", s.Taluka, s.District, s.State))
		pdf.SetTextColor(0, 0, 0) 

		// 5. Detailed Information Card
		labels := []string{"Email ID", "Mobile No", "Gender", "Date of Birth", "Blood Group", "Handicapped"}
		hcText := "No"
		if s.Handicapped { hcText = "Yes" }
		values := []string{s.Email, s.MobileNumber, s.Gender, s.DOB, s.BloodGroup, hcText}

		startY := 250.0
		for idx, label := range labels {
			if idx % 2 == 0 {
				pdf.SetFillColor(245, 247, 250) // Light background
				pdf.RectFromUpperLeftWithStyle(45, startY, 505, 35, "F")
			}

			pdf.SetFont("arial", "", 10)
			pdf.SetTextColor(80, 80, 80)
			pdf.SetXY(60, startY + 22)
			pdf.Text(label + " :")

			pdf.SetFont("arial", "", 11)
			pdf.SetTextColor(0, 0, 0)
			pdf.SetXY(180, startY + 22)
			pdf.Text(values[idx])

			startY += 35
		}

		// 6. Footer
		pdf.SetLineWidth(0.3)
		pdf.SetStrokeColor(200, 200, 200)
		pdf.Line(45, 780, 550, 780)

		pdf.SetFont("arial", "", 8)
		pdf.SetTextColor(120, 120, 120)
		pdf.SetXY(45, 795)
		pdf.Text(fmt.Sprintf("Official Student Record | Page %d", i+1))
		pdf.SetXY(410, 795)
		pdf.Text("Generated: " + time.Now().Format("02 Jan 2006 15:04"))
	}

	return &pdf, nil
}