from django.urls import path
from healthcareOnDjango.views import CareTaskView

urlpatterns = [
    path('', CareTaskView.index, name='index'),
	path('create', CareTaskView.get, name='create'),
	path('get/<int:careTaskId>/', CareTaskView.get, name='get'),
	path('save', CareTaskView.save, name='save'),
	path('getAll', CareTaskView.getAll, name='getAll'),
	path('delete/<int:careTaskId>/', CareTaskView.delete, name='delete'),
	path('assignCarePlan/<int:careTaskId>/<int:CarePlanId>/', CareTaskView.assignCarePlan, name='assignCarePlan'),
	path('unassignCarePlan/<int:careTaskId>/', CareTaskView.unassignCarePlan, name='unassignCarePlan'),
	path('assignAssignedTo/<int:careTaskId>/<int:AssignedToId>/', CareTaskView.assignAssignedTo, name='assignAssignedTo'),
	path('unassignAssignedTo/<int:careTaskId>/', CareTaskView.unassignAssignedTo, name='unassignAssignedTo'),
	path('assignEncounter/<int:careTaskId>/<int:EncounterId>/', CareTaskView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:careTaskId>/', CareTaskView.unassignEncounter, name='unassignEncounter'),
]
