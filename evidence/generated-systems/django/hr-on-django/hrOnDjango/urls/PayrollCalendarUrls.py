from django.urls import path
from hrOnDjango.views import PayrollCalendarView

urlpatterns = [
    path('', PayrollCalendarView.index, name='index'),
	path('create', PayrollCalendarView.get, name='create'),
	path('get/<int:payrollCalendarId>/', PayrollCalendarView.get, name='get'),
	path('save', PayrollCalendarView.save, name='save'),
	path('getAll', PayrollCalendarView.getAll, name='getAll'),
	path('delete/<int:payrollCalendarId>/', PayrollCalendarView.delete, name='delete'),
	path('assignOrganization/<int:payrollCalendarId>/<int:OrganizationId>/', PayrollCalendarView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:payrollCalendarId>/', PayrollCalendarView.unassignOrganization, name='unassignOrganization'),
	path('addPayrollRuns/<int:payrollCalendarId>/<PayrollRunsIds>/', PayrollCalendarView.addPayrollRuns, name='addPayrollRuns'),
	path('removePayrollRuns/<int:payrollCalendarId>/<PayrollRunsIds>/', PayrollCalendarView.removePayrollRuns, name='removePayrollRuns'),
	path('addEmployees/<int:payrollCalendarId>/<EmployeesIds>/', PayrollCalendarView.addEmployees, name='addEmployees'),
	path('removeEmployees/<int:payrollCalendarId>/<EmployeesIds>/', PayrollCalendarView.removeEmployees, name='removeEmployees'),
]
