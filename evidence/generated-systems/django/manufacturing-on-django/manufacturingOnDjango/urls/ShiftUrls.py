from django.urls import path
from manufacturingOnDjango.views import ShiftView

urlpatterns = [
    path('', ShiftView.index, name='index'),
	path('create', ShiftView.get, name='create'),
	path('get/<int:shiftId>/', ShiftView.get, name='get'),
	path('save', ShiftView.save, name='save'),
	path('getAll', ShiftView.getAll, name='getAll'),
	path('delete/<int:shiftId>/', ShiftView.delete, name='delete'),
	path('assignPlant/<int:shiftId>/<int:PlantId>/', ShiftView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:shiftId>/', ShiftView.unassignPlant, name='unassignPlant'),
	path('addAssignments/<int:shiftId>/<AssignmentsIds>/', ShiftView.addAssignments, name='addAssignments'),
	path('removeAssignments/<int:shiftId>/<AssignmentsIds>/', ShiftView.removeAssignments, name='removeAssignments'),
]
