from django.urls import path
from healthcareOnDjango.views import CarePlanView

urlpatterns = [
    path('', CarePlanView.index, name='index'),
	path('create', CarePlanView.get, name='create'),
	path('get/<int:carePlanId>/', CarePlanView.get, name='get'),
	path('save', CarePlanView.save, name='save'),
	path('getAll', CarePlanView.getAll, name='getAll'),
	path('delete/<int:carePlanId>/', CarePlanView.delete, name='delete'),
	path('assignPatient/<int:carePlanId>/<int:PatientId>/', CarePlanView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:carePlanId>/', CarePlanView.unassignPatient, name='unassignPatient'),
	path('assignCareTeam/<int:carePlanId>/<int:CareTeamId>/', CarePlanView.assignCareTeam, name='assignCareTeam'),
	path('unassignCareTeam/<int:carePlanId>/', CarePlanView.unassignCareTeam, name='unassignCareTeam'),
	path('addEncounters/<int:carePlanId>/<EncountersIds>/', CarePlanView.addEncounters, name='addEncounters'),
	path('removeEncounters/<int:carePlanId>/<EncountersIds>/', CarePlanView.removeEncounters, name='removeEncounters'),
	path('addTasks/<int:carePlanId>/<TasksIds>/', CarePlanView.addTasks, name='addTasks'),
	path('removeTasks/<int:carePlanId>/<TasksIds>/', CarePlanView.removeTasks, name='removeTasks'),
]
