from django.urls import path
from manufacturingOnDjango.views import ShiftAssignmentView

urlpatterns = [
    path('', ShiftAssignmentView.index, name='index'),
	path('create', ShiftAssignmentView.get, name='create'),
	path('get/<int:shiftAssignmentId>/', ShiftAssignmentView.get, name='get'),
	path('save', ShiftAssignmentView.save, name='save'),
	path('getAll', ShiftAssignmentView.getAll, name='getAll'),
	path('delete/<int:shiftAssignmentId>/', ShiftAssignmentView.delete, name='delete'),
	path('assignShift/<int:shiftAssignmentId>/<int:ShiftId>/', ShiftAssignmentView.assignShift, name='assignShift'),
	path('unassignShift/<int:shiftAssignmentId>/', ShiftAssignmentView.unassignShift, name='unassignShift'),
	path('assignEmployee/<int:shiftAssignmentId>/<int:EmployeeId>/', ShiftAssignmentView.assignEmployee, name='assignEmployee'),
	path('unassignEmployee/<int:shiftAssignmentId>/', ShiftAssignmentView.unassignEmployee, name='unassignEmployee'),
	path('assignWorkCenter/<int:shiftAssignmentId>/<int:WorkCenterId>/', ShiftAssignmentView.assignWorkCenter, name='assignWorkCenter'),
	path('unassignWorkCenter/<int:shiftAssignmentId>/', ShiftAssignmentView.unassignWorkCenter, name='unassignWorkCenter'),
]
