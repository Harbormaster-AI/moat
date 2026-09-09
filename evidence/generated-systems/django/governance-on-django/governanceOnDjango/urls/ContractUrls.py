from django.urls import path
from governanceOnDjango.views import ContractView

urlpatterns = [
    path('', ContractView.index, name='index'),
	path('create', ContractView.get, name='create'),
	path('get/<int:contractId>/', ContractView.get, name='get'),
	path('save', ContractView.save, name='save'),
	path('getAll', ContractView.getAll, name='getAll'),
	path('delete/<int:contractId>/', ContractView.delete, name='delete'),
	path('assignThirdParty/<int:contractId>/<int:ThirdPartyId>/', ContractView.assignThirdParty, name='assignThirdParty'),
	path('unassignThirdParty/<int:contractId>/', ContractView.unassignThirdParty, name='unassignThirdParty'),
	path('assignMatter/<int:contractId>/<int:MatterId>/', ContractView.assignMatter, name='assignMatter'),
	path('unassignMatter/<int:contractId>/', ContractView.unassignMatter, name='unassignMatter'),
	path('addObligations/<int:contractId>/<ObligationsIds>/', ContractView.addObligations, name='addObligations'),
	path('removeObligations/<int:contractId>/<ObligationsIds>/', ContractView.removeObligations, name='removeObligations'),
	path('addDataProcessingActivities/<int:contractId>/<DataProcessingActivitiesIds>/', ContractView.addDataProcessingActivities, name='addDataProcessingActivities'),
	path('removeDataProcessingActivities/<int:contractId>/<DataProcessingActivitiesIds>/', ContractView.removeDataProcessingActivities, name='removeDataProcessingActivities'),
]
