package day2

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

type day2 struct {
	reports [][]int
}

func New(inputFilepath string) *day2 {
	d := &day2{}
	f, err := os.Open(inputFilepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := d.parseInput(f); err != nil {
		panic(err)
	}
	return d
}

func (d *day2) parseInput(reader io.Reader) error {
	contents, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	lines := bytes.Split(contents, []byte("\n"))
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	d.reports = make([][]int, len(lines))
	for i, line := range lines {
		fields := bytes.Fields(line)
		d.reports[i] = make([]int, len(fields))
		for j, field := range fields {
			n, err := strconv.Atoi(string(field))
			if err != nil {
				return fmt.Errorf("invalid input: %w", err)
			}
			d.reports[i][j] = n
		}
	}
	return nil
}

func (d day2) Part1() string {
	safeReports := 0
	for _, report := range d.reports {
		if !isReportSafe(report) {
			continue
		}
		safeReports++
	}
	return strconv.Itoa(safeReports)
}

func (d day2) Part2() string {
	safeReports := 0
	for _, report := range d.reports {
		if !isReportSafeWithDeletion(report) {
			continue
		}
		safeReports++
	}
	return strconv.Itoa(safeReports)
}

func isReportSafe(report []int) bool {
	if len(report) < 2 {
		return true
	}

	minchange := 1
	maxchange := 3
	if report[1]-report[0] < 0 {
		minchange = -3
		maxchange = -1
	}

	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] < minchange ||
			report[i]-report[i-1] > maxchange {
			return false
		}
	}
	return true
}

func isReportSafeWithDeletion(report []int) bool {
	if len(report) < 2 {
		return true
	}
	if isReportSafeWithoutIndex(report, 0) || isReportSafeWithoutIndex(report, 1) {
		return true
	}

	minchange := 1
	maxchange := 3
	if report[1]-report[0] < 0 {
		minchange = -3
		maxchange = -1
	}

	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] < minchange ||
			report[i]-report[i-1] > maxchange {
			return isReportSafeWithoutIndex(report, i-1) || isReportSafeWithoutIndex(report, i)
		}
	}
	return true
}

func isReportSafeWithoutIndex(report []int, index int) bool {
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)

	return isReportSafe(append(reportCopy[:index], reportCopy[index+1:]...))
}

func (d day2) Day() int {
	return 2
}
