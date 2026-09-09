from django.urls import path
from healthcareOnDjango.views import ConditionView

urlpatterns = [
    path('', ConditionView.index, name='index'),
	path('create', ConditionView.get, name='create'),
	path('get/<int:conditionId>/', ConditionView.get, name='get'),
	path('save', ConditionView.save, name='save'),
	path('getAll', ConditionView.getAll, name='getAll'),
	path('delete/<int:conditionId>/', ConditionView.delete, name='delete'),
	path('assignPatient/<int:conditionId>/<int:PatientId>/', ConditionView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:conditionId>/', ConditionView.unassignPatient, name='unassignPatient'),
]
