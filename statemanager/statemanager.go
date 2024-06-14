package statemanager

import (
	"fmt"

	cm "coursemanagement/common"
	"coursemanagement/mysqldbmodels"
	dbmodels "coursemanagement/mysqldbmodels"

	"gorm.io/gorm"
)

// Goal is to keep tracking of open positions
// Extension would be to support for multi-day strategy
type StateManager struct {
	db *dbmodels.DBClient
}

func InitStateManager() (*StateManager, error) {
	sm := &StateManager{}
	var err error
	if sm.db, err = dbmodels.InitializeDatabase(); err != nil {
		return sm, fmt.Errorf("error occured while creating database connection,err: %s", err.Error())
	}
	return sm, nil
}
func (sm *StateManager) GetDBConnection()  *gorm.DB {
	return sm.db.Conn

}

func (sm *StateManager) CreateCourseEntry(c1 cm.Course) error {
	return sm.db.CreateCourseRow(c1)
}

func (sm *StateManager) CreateStudentEntry(s1 cm.Student) error {
	return sm.db.CreateStudentRow(s1)
}

func (sm *StateManager) CreateTeacherEntry(t1 cm.Teacher) error {
	return sm.db.CreateTeacherRow(t1)
}

func (sm *StateManager) CreateEntrollmentEntry(e1 cm.Entrollment) error {
	return sm.db.CreateEntrollmentRow(e1)
}

func (sm *StateManager) GetAllCourses() ([]mysqldbmodels.Course, error) {
	return sm.db.GetCourseRaw()
}

func (sm *StateManager) GetAllStudents() ([]mysqldbmodels.Student, error) {
	return sm.db.GetStudentRaw()
}

func (sm *StateManager) GetAllTeachers() ([]mysqldbmodels.Teacher, error) {
	return sm.db.GetTeacherRaw()
}

func (sm *StateManager) GetAllEntrollments() ([]mysqldbmodels.Entrollment, error) {
	return sm.db.GetEntrollmentRaw()
}

func (sm *StateManager) UpdateAllCourses(ID int, Name string, TeacherID int) ([]mysqldbmodels.Course, error) {

	return sm.db.UpdateCourseRaw(ID, Name, TeacherID)
}

func (sm *StateManager) UpdateAllStudents(ID int, Name string) ([]mysqldbmodels.Student, error) {

	return sm.db.UpdateStudentRaw(ID, Name)
}

func (sm *StateManager) UpdateAllTeachers(ID int, Name string) ([]mysqldbmodels.Teacher, error) {

	return sm.db.UpdateTeacherRaw(ID, Name)
}

func (sm *StateManager) UpdateAllEntrollments(ID int, CourseID int, StudentID int) ([]mysqldbmodels.Entrollment, error) {

	return sm.db.UpdateEntrollmentRaw(ID, CourseID, StudentID)
}

func (sm *StateManager) DeleteCourseID(ID int) ([]mysqldbmodels.Course, error) {

	return sm.db.DeleteCourseRaw(ID)
}

func (sm *StateManager) DeleteStudentID(ID int) ([]mysqldbmodels.Student, error) {

	return sm.db.DeleteStudentRaw(ID)
}

func (sm *StateManager) DeleteEntrollmentID(ID int) ([]mysqldbmodels.Entrollment, error) {

	return sm.db.DeleteEntrollmentRaw(ID)
}

func (sm *StateManager) DeleteTeacherID(ID int) ([]mysqldbmodels.Teacher, error) {

	return sm.db.DeleteTeacherRaw(ID)
}
