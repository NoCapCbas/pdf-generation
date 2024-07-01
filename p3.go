package main

import (
  "fmt"
	"github.com/jung-kurt/gofpdf"
)

func main() {
  // orientaion, measurement, document type, font directory 
  pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
  w, h := pdf.GetPageSize()
  fmt.Printf("width=%v, height=%v\n", w, h)
  pdf.AddPage()

  const (
    bannerHt = 91.0
    xIndent = 40.0
  )

  // Header Polygon
  pdf.SetFillColor(26, 198, 235)
  pdf.Polygon([]gofpdf.PointType{
    {0, 0},
    {w, 0},
    {w, bannerHt},
    {0, bannerHt*0.9},
  }, "F")
  // Footer Polygon 
  pdf.Polygon([]gofpdf.PointType{
    {0, h},
    {0, h - (bannerHt*0.2)},
    {w, h - (bannerHt*0.1)},
    {w, h},
  }, "F")
  // Header INVOICE text
  pdf.SetFont("arial", "B", 40)
  pdf.SetTextColor(255, 255, 255)
  _, lineHt := pdf.GetFontSize()
  pdf.Text(xIndent, bannerHt-(bannerHt/2.0)+lineHt/3.1, "INVOICE") 

  // Header Phone, email, domain
  pdf.SetFont("arial", "B", 12)
  pdf.SetTextColor(255, 255, 255)
  _, lineHt = pdf.GetFontSize()
  pdf.MoveTo(w-xIndent-2.0*150.0, (bannerHt - (lineHt*1.5*3.0))/2.0)
  pdf.MultiCell(150.0, lineHt*1.5, "(123) 456-7890\ncompany@email.com\ncompany.com", gofpdf.BorderNone, gofpdf.AlignRight, false)

  // Header Address
  pdf.SetFont("arial", "B", 12)
  pdf.SetTextColor(255, 255, 255)
  _, lineHt = pdf.GetFontSize()
  pdf.MoveTo(w-xIndent-150.0, (bannerHt - (lineHt*1.5*3.0))/2.0)
  pdf.MultiCell(150.0, lineHt*1.5, "123 Fake St\nSome Town, NA\n12345", gofpdf.BorderNone, gofpdf.AlignRight, false)

  // Summary Billed To
  pdf.SetFont("times", "", 14)
  pdf.SetTextColor(180, 180, 180)
  _, lineHt = pdf.GetFontSize()
  x, y := xIndent, bannerHt+lineHt*3.0
  pdf.Text(x, y, "Billed To")
  pdf.SetTextColor(50, 50, 50)

  y = y + lineHt*1.5
  pdf.Text(x, y, "Client Name")

  y = y + lineHt*1.25
  pdf.Text(x, y, "123 Client Adress")

  y = y + lineHt*1.25
  pdf.Text(x, y, "City, State, Country")

  y = y + lineHt*1.25
  pdf.Text(x, y, "Postal Code")

  // Grid 
  drawGrid(pdf)


  err := pdf.OutputFileAndClose("p3.pdf") 
  if err != nil {
    panic(err)
  }
}

func summaryBlock(pdf *gofpdf.Fpdf) (x, y float64) {

}

func drawGrid(pdf *gofpdf.Fpdf) {
  w, h := pdf.GetPageSize()
  pdf.SetFont("courier", "", 12)
  pdf.SetTextColor(80, 80, 80)
  pdf.SetDrawColor(200, 200, 200)

  for x := 0.0; x < w; x = x + (w/20.0) {
    pdf.Line(x, 0, x, h)
    _, lineHt :=  pdf.GetFontSize()
    pdf.Text(x, lineHt, fmt.Sprintf("%d", int(x)))
  }

  for y := 0.0; y < h; y = y + (w/20.0) {
    pdf.Line(0, y, w, y)
    pdf.Text(0, y, fmt.Sprintf("%d", int(y)))
  }
}







