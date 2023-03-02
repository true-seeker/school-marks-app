package input

type ClassTransferInput struct {
	NewAcademicYearId        uint
	NewSchoolClassId         uint
	NewStudentIds            []uint
	NotTransferredStudentIds []uint
}
