from django.urls import path
from insuranceOnDjango.views import IncidentView

urlpatterns = [
    path('', IncidentView.index, name='index'),
	path('create', IncidentView.get, name='create'),
	path('get/<int:incidentId>/', IncidentView.get, name='get'),
	path('save', IncidentView.save, name='save'),
	path('getAll', IncidentView.getAll, name='getAll'),
	path('delete/<int:incidentId>/', IncidentView.delete, name='delete'),
	path('assignClaim/<int:incidentId>/<int:ClaimId>/', IncidentView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:incidentId>/', IncidentView.unassignClaim, name='unassignClaim'),
	path('addInsuredObjects/<int:incidentId>/<InsuredObjectsIds>/', IncidentView.addInsuredObjects, name='addInsuredObjects'),
	path('removeInsuredObjects/<int:incidentId>/<InsuredObjectsIds>/', IncidentView.removeInsuredObjects, name='removeInsuredObjects'),
]
