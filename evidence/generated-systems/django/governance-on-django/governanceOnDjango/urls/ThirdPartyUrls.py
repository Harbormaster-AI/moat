from django.urls import path
from governanceOnDjango.views import ThirdPartyView

urlpatterns = [
    path('', ThirdPartyView.index, name='index'),
	path('create', ThirdPartyView.get, name='create'),
	path('get/<int:thirdPartyId>/', ThirdPartyView.get, name='get'),
	path('save', ThirdPartyView.save, name='save'),
	path('getAll', ThirdPartyView.getAll, name='getAll'),
	path('delete/<int:thirdPartyId>/', ThirdPartyView.delete, name='delete'),
	path('assignOrganization/<int:thirdPartyId>/<int:OrganizationId>/', ThirdPartyView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:thirdPartyId>/', ThirdPartyView.unassignOrganization, name='unassignOrganization'),
	path('addProcessingActivities/<int:thirdPartyId>/<ProcessingActivitiesIds>/', ThirdPartyView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:thirdPartyId>/<ProcessingActivitiesIds>/', ThirdPartyView.removeProcessingActivities, name='removeProcessingActivities'),
	path('addAssessments/<int:thirdPartyId>/<AssessmentsIds>/', ThirdPartyView.addAssessments, name='addAssessments'),
	path('removeAssessments/<int:thirdPartyId>/<AssessmentsIds>/', ThirdPartyView.removeAssessments, name='removeAssessments'),
	path('addContracts/<int:thirdPartyId>/<ContractsIds>/', ThirdPartyView.addContracts, name='addContracts'),
	path('removeContracts/<int:thirdPartyId>/<ContractsIds>/', ThirdPartyView.removeContracts, name='removeContracts'),
	path('addObligations/<int:thirdPartyId>/<ObligationsIds>/', ThirdPartyView.addObligations, name='addObligations'),
	path('removeObligations/<int:thirdPartyId>/<ObligationsIds>/', ThirdPartyView.removeObligations, name='removeObligations'),
	path('addDataBreaches/<int:thirdPartyId>/<DataBreachesIds>/', ThirdPartyView.addDataBreaches, name='addDataBreaches'),
	path('removeDataBreaches/<int:thirdPartyId>/<DataBreachesIds>/', ThirdPartyView.removeDataBreaches, name='removeDataBreaches'),
]
