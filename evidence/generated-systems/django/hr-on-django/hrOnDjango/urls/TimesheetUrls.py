from django.urls import path
from hrOnDjango.views import TimesheetView

urlpatterns = [
    path('', TimesheetView.index, name='index'),
	path('create', TimesheetView.get, name='create'),
	path('get/<int:timesheetId>/', TimesheetView.get, name='get'),
	path('save', TimesheetView.save, name='save'),
	path('getAll', TimesheetView.getAll, name='getAll'),
	path('delete/<int:timesheetId>/', TimesheetView.delete, name='delete'),
	path('assignEmployee/<int:timesheetId>/<int:EmployeeId>/', TimesheetView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:timesheetId>/', TimesheetView.unassignEmployee, name='unassignEmployee'),
	path('addTimeEntries/<int:timesheetId>/<TimeEntriesIds>/', TimesheetView.addTimeEntries, name='addTimeEntries'),
	path('removeTimeEntries/<int:timesheetId>/<TimeEntriesIds>/', TimesheetView.removeTimeEntries, name='removeTimeEntries'),
	path('addApprovals/<int:timesheetId>/<ApprovalsIds>/', TimesheetView.addApprovals, name='addApprovals'),
	path('removeApprovals/<int:timesheetId>/<ApprovalsIds>/', TimesheetView.removeApprovals, name='removeApprovals'),
]
