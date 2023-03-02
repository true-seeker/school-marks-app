package db

type ClassTransferInput struct {
	NewAcademicYearId        uint
	NewSchoolClassId         uint
	NewStudentIds            []uint
	NotTransferredStudentIds []uint
}
