from django.urls import path
from hrOnDjango.views import PayrollItemView

urlpatterns = [
    path('', PayrollItemView.index, name='index'),
	path('create', PayrollItemView.get, name='create'),
	path('get/<int:payrollItemId>/', PayrollItemView.get, name='get'),
	path('save', PayrollItemView.save, name='save'),
	path('getAll', PayrollItemView.getAll, name='getAll'),
	path('delete/<int:payrollItemId>/', PayrollItemView.delete, name='delete'),
	path('assignPayrollRun/<int:payrollItemId>/<int:PayrollRunId>/', PayrollItemView.assignPayrollRun, name='assignPayrollRun'),
	path('unassignPayrollRun/<int:payrollItemId>/', PayrollItemView.unassignPayrollRun, name='unassignPayrollRun'),
	path('assignEmployee/<int:payrollItemId>/<int:EmployeeId>/', PayrollItemView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:payrollItemId>/', PayrollItemView.unassignEmployee, name='unassignEmployee'),
]
