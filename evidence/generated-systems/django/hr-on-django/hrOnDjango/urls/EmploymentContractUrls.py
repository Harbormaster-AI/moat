from django.urls import path
from hrOnDjango.views import EmploymentContractView

urlpatterns = [
    path('', EmploymentContractView.index, name='index'),
	path('create', EmploymentContractView.get, name='create'),
	path('get/<int:employmentContractId>/', EmploymentContractView.get, name='get'),
	path('save', EmploymentContractView.save, name='save'),
	path('getAll', EmploymentContractView.getAll, name='getAll'),
	path('delete/<int:employmentContractId>/', EmploymentContractView.delete, name='delete'),
	path('assignEmployee/<int:employmentContractId>/<int:EmployeeId>/', EmploymentContractView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:employmentContractId>/', EmploymentContractView.unassignEmployee, name='unassignEmployee'),
	path('assignCompensationPackage/<int:employmentContractId>/<int:CompensationPackageId>/', EmploymentContractView.assignCompensationPackage, name='assignCompensationPackage'),
	path('unassignCompensationPackage/<int:employmentContractId>/', EmploymentContractView.unassignCompensationPackage, name='unassignCompensationPackage'),
	path('assignWorkSchedule/<int:employmentContractId>/<int:WorkScheduleId>/', EmploymentContractView.assignWorkSchedule, name='assignWorkSchedule'),
	path('unassignWorkSchedule/<int:employmentContractId>/', EmploymentContractView.unassignWorkSchedule, name='unassignWorkSchedule'),
	path('assignLocation/<int:employmentContractId>/<int:LocationId>/', EmploymentContractView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:employmentContractId>/', EmploymentContractView.unassignLocation, name='unassignLocation'),
	path('assignPayrollCalendar/<int:employmentContractId>/<int:PayrollCalendarId>/', EmploymentContractView.assignPayrollCalendar, name='assignPayrollCalendar'),
	path('unassignPayrollCalendar/<int:employmentContractId>/', EmploymentContractView.unassignPayrollCalendar, name='unassignPayrollCalendar'),
]
