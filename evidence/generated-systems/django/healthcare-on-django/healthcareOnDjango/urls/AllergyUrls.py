from django.urls import path
from healthcareOnDjango.views import AllergyView

urlpatterns = [
    path('', AllergyView.index, name='index'),
	path('create', AllergyView.get, name='create'),
	path('get/<int:allergyId>/', AllergyView.get, name='get'),
	path('save', AllergyView.save, name='save'),
	path('getAll', AllergyView.getAll, name='getAll'),
	path('delete/<int:allergyId>/', AllergyView.delete, name='delete'),
	path('assignPatient/<int:allergyId>/<int:PatientId>/', AllergyView.assignPatient, name='assignPatient'),
	path('unassignPatient/<int:allergyId>/', AllergyView.unassignPatient, name='unassignPatient'),
]
