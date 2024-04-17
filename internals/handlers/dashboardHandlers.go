package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/utils"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/pages"
)

func (h *Handler) GetDashboardHome(c echo.Context) error {
	title := "Dashboard Home"
	page := pages.DashboardHomeComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardSessions(c echo.Context) error {
	title := "Dashboard Sessions"
	page := pages.DashboardSessionsComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardTerms(c echo.Context) error {
	title := "Dashboard Terms"
	page := pages.DashboardTermsComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardClasses(c echo.Context) error {
	title := "Dashboard Classes"
	page := pages.DashboardClassesComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardSubjects(c echo.Context) error {
	title := "Dashboard Subjects"
	page := pages.DashboardSubjectsComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardStudents(c echo.Context) error {
	title := "Dashboard Students"
	page := pages.DashboardStudentsComponent([]string{"First Name", "Middle Name", "Last Name", "Admission No", "Gender", "Password Change", "Active", "Current Class", "Joining Date", "view"})
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardStaff(c echo.Context) error {
	title := "Dashboard Staff"
	page := pages.DashboardStaffComponent([]string{"Display Name", "Admission No", "Gender", "Role", "Password Changed", "Active", "Current Class", "Joining Date", "View", "Update"})
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardResultUpload(c echo.Context) error {
	title := "Dashboard Result Upload"
	page := pages.DashboardResultUploadComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardPsychomotor(c echo.Context) error {
	title := "Dashboard Psychomotor"
	page := pages.DashboardPsychomotorComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardComments(c echo.Context) error {
	title := "Dashboard Comments"
	page := pages.DashboardCommentsComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}

func (h *Handler) GetDashboardViewResults(c echo.Context) error {
	title := "Dashboard View Results"
	page := pages.DashboardViewResultsComponent()
	return utils.RenderDashboardTemplate(c, page, title)
}
