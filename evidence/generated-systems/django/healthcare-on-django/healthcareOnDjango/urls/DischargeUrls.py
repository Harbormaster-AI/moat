from django.urls import path
from healthcareOnDjango.views import DischargeView

urlpatterns = [
    path('', DischargeView.index, name='index'),
	path('create', DischargeView.get, name='create'),
	path('get/<int:dischargeId>/', DischargeView.get, name='get'),
	path('save', DischargeView.save, name='save'),
	path('getAll', DischargeView.getAll, name='getAll'),
	path('delete/<int:dischargeId>/', DischargeView.delete, name='delete'),
	path('assignEncounter/<int:dischargeId>/<int:EncounterId>/', DischargeView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:dischargeId>/', DischargeView.unassignEncounter, name='unassignEncounter'),
]
