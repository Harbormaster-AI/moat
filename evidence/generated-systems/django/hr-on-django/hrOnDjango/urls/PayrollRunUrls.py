from django.urls import path
from hrOnDjango.views import PayrollRunView

urlpatterns = [
    path('', PayrollRunView.index, name='index'),
	path('create', PayrollRunView.get, name='create'),
	path('get/<int:payrollRunId>/', PayrollRunView.get, name='get'),
	path('save', PayrollRunView.save, name='save'),
	path('getAll', PayrollRunView.getAll, name='getAll'),
	path('delete/<int:payrollRunId>/', PayrollRunView.delete, name='delete'),
	path('assignPayrollCalendar/<int:payrollRunId>/<int:PayrollCalendarId>/', PayrollRunView.assignPayrollCalendar, name='assignPayrollCalendar'),
	path('unassignPayrollCalendar/<int:payrollRunId>/', PayrollRunView.unassignPayrollCalendar, name='unassignPayrollCalendar'),
	path('addPayrollItems/<int:payrollRunId>/<PayrollItemsIds>/', PayrollRunView.addPayrollItems, name='addPayrollItems'),
	path('removePayrollItems/<int:payrollRunId>/<PayrollItemsIds>/', PayrollRunView.removePayrollItems, name='removePayrollItems'),
]
