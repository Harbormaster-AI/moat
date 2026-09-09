from django.urls import path
from manufacturingOnDjango.views import EmployeeView

urlpatterns = [
    path('', EmployeeView.index, name='index'),
	path('create', EmployeeView.get, name='create'),
	path('get/<int:employeeId>/', EmployeeView.get, name='get'),
	path('save', EmployeeView.save, name='save'),
	path('getAll', EmployeeView.getAll, name='getAll'),
	path('delete/<int:employeeId>/', EmployeeView.delete, name='delete'),
	path('assignWorkCenter/<int:employeeId>/<int:WorkCenterId>/', EmployeeView.assignWorkCenter, name='assignWorkCenter'),
	path('unassignWorkCenter/<int:employeeId>/', EmployeeView.unassignWorkCenter, name='unassignWorkCenter'),
	path('addShiftAssignments/<int:employeeId>/<ShiftAssignmentsIds>/', EmployeeView.addShiftAssignments, name='addShiftAssignments'),
	path('removeShiftAssignments/<int:employeeId>/<ShiftAssignmentsIds>/', EmployeeView.removeShiftAssignments, name='removeShiftAssignments'),
	path('addCorrectiveActions/<int:employeeId>/<CorrectiveActionsIds>/', EmployeeView.addCorrectiveActions, name='addCorrectiveActions'),
	path('removeCorrectiveActions/<int:employeeId>/<CorrectiveActionsIds>/', EmployeeView.removeCorrectiveActions, name='removeCorrectiveActions'),
]
