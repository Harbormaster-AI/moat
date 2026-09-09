from django.urls import path
from hrOnDjango.views import TimeEntryView

urlpatterns = [
    path('', TimeEntryView.index, name='index'),
	path('create', TimeEntryView.get, name='create'),
	path('get/<int:timeEntryId>/', TimeEntryView.get, name='get'),
	path('save', TimeEntryView.save, name='save'),
	path('getAll', TimeEntryView.getAll, name='getAll'),
	path('delete/<int:timeEntryId>/', TimeEntryView.delete, name='delete'),
	path('assignTimesheet/<int:timeEntryId>/<int:TimesheetId>/', TimeEntryView.assignTimesheet, name='assignTimesheet'),
	path('unassignTimesheet/<int:timeEntryId>/', TimeEntryView.unassignTimesheet, name='unassignTimesheet'),
	path('assignEmployee/<int:timeEntryId>/<int:EmployeeId>/', TimeEntryView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:timeEntryId>/', TimeEntryView.unassignEmployee, name='unassignEmployee'),
	path('assignCostCenter/<int:timeEntryId>/<int:CostCenterId>/', TimeEntryView.assignCostCenter, name='assignCostCenter'),
	path('unassignCostCenter/<int:timeEntryId>/', TimeEntryView.unassignCostCenter, name='unassignCostCenter'),
]
