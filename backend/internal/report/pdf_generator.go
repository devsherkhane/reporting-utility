package report

import (
	"fmt"
	"reporting-utility/internal/models"
	"time"

	"github.com/signintech/gopdf"
)

func GenerateStudentsPDF(students []models.Student) (*gopdf.GoPdf, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4Landscape})

	err := pdf.AddTTFFont("arial", "C:\\Windows\\Fonts\\arial.ttf")
	if err != nil {
		err = pdf.AddTTFFont("arial", "assets/arial.ttf")
		if err != nil {
			return nil, fmt.Errorf("font file missing: please put arial.ttf in assets/ folder")
		}
	}

	headers := []string{"Sr No", "Name", "Email", "Mobile", "Address", "Gender", "DOB", "Blood"}
	colWidths := []float64{40, 100, 130, 85, 140, 60, 80, 50}

	drawHeader := func() {
		pdf.Image("assets/kk-wagh-logo.png", 30, 20, &gopdf.Rect{W: 100, H: 50})
		pdf.SetFont("arial", "", 16)
		pdf.SetXY(250, 35)
		pdf.Text("K.K Wagh Institute Of Engineering and Research")

		pdf.SetFont("arial", "", 10)
		pdf.SetTextColor(0, 0, 255)
		pdf.SetXY(250, 55)
		pdf.Text("https://www.kkwagh.edu.in/")
		pdf.SetTextColor(0, 0, 0)

		pdf.SetFont("arial", "", 14)
		pdf.SetXY(30, 95)
		pdf.Text("Student Records")

		pdf.SetFont("arial", "", 10)
		currX := 30.0
		for i, h := range headers {
			pdf.RectFromUpperLeftWithStyle(currX, 115, colWidths[i], 25, "D")
			pdf.SetXY(currX+5, 130)
			pdf.Text(h)
			currX += colWidths[i]
		}
	}

	drawFooter := func(pageNo int) {
		pdf.SetLineWidth(0.5)
		pdf.Line(30, 560, 810, 560)
		pdf.SetFont("arial", "", 8)
		pdf.SetXY(30, 575)
		pdf.Text("Created by Dev | " + time.Now().Format("02 Jan 2006"))
		pdf.SetXY(780, 575)
		pdf.Text(fmt.Sprintf("Page %d", pageNo))
	}

	pdf.AddPage()
	currentPage := 1
	drawHeader()
	drawFooter(currentPage)

	y := 140.0
	rowHeight := 20.0

	// --- ROW SPAN LOGIC PREPARATION ---
	// spanCount tracks how many rows the current email should span
	spanCount := make([]int, len(students))
	for i := 0; i < len(students); {
		count := 1
		for j := i + 1; j < len(students) && students[j].Email == students[i].Email; j++ {
			count++
		}
		spanCount[i] = count
		i += count
	}

	for i, s := range students {
		// Handle Page Breaks
		if y > 530 {
			pdf.AddPage()
			currentPage++
			drawHeader()
			drawFooter(currentPage)
			y = 140.0
		}

		data := []string{fmt.Sprintf("%d", i+1), s.StudentName, s.Email, s.MobileNumber, s.Address, s.Gender, s.DOB, s.BloodGroup}
		currX := 30.0
		pdf.SetFont("arial", "", 9)

		for j, txt := range data {
			// Column 2 is "Email"
			if j == 2 {
				if spanCount[i] > 0 {
					// Calculate the total height of the spanned cell
					totalSpanHeight := float64(spanCount[i]) * rowHeight

					// If the span goes off the current page, clip it to the page bottom
					remainingPageSpace := 530.0 - y + rowHeight
					cellHeight := totalSpanHeight
					if cellHeight > remainingPageSpace {
						cellHeight = remainingPageSpace
					}

					// Draw the spanned box
					pdf.RectFromUpperLeftWithStyle(currX, y, colWidths[j], cellHeight, "D")

					// Draw the text only once for the span
					pdf.SetXY(currX+5, y+(cellHeight/2)+4) // Center text vertically in the span
					limit := int(colWidths[j] / 5.5)
					if len(txt) > limit && limit > 3 {
						txt = txt[:limit-3] + "..."
					}
					pdf.Text(txt)
				}
				// If spanCount[i] is 0, it's a "covered" row; we skip drawing to maintain the span look
			} else {
				// Standard cell drawing for all other columns
				pdf.RectFromUpperLeftWithStyle(currX, y, colWidths[j], rowHeight, "D")
				pdf.SetXY(currX+5, y+13)
				limit := int(colWidths[j] / 5.5)
				if len(txt) > limit && limit > 3 {
					txt = txt[:limit-3] + "..."
				}
				pdf.Text(txt)
			}
			currX += colWidths[j]
		}
		y += rowHeight
	}

	return &pdf, nil
}
