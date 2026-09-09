from django.urls import path
from governanceOnDjango.views import MatterView

urlpatterns = [
    path('', MatterView.index, name='index'),
	path('create', MatterView.get, name='create'),
	path('get/<int:matterId>/', MatterView.get, name='get'),
	path('save', MatterView.save, name='save'),
	path('getAll', MatterView.getAll, name='getAll'),
	path('delete/<int:matterId>/', MatterView.delete, name='delete'),
	path('assignOrganization/<int:matterId>/<int:OrganizationId>/', MatterView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:matterId>/', MatterView.unassignOrganization, name='unassignOrganization'),
	path('addLegalHolds/<int:matterId>/<LegalHoldsIds>/', MatterView.addLegalHolds, name='addLegalHolds'),
	path('removeLegalHolds/<int:matterId>/<LegalHoldsIds>/', MatterView.removeLegalHolds, name='removeLegalHolds'),
	path('addDataBreaches/<int:matterId>/<DataBreachesIds>/', MatterView.addDataBreaches, name='addDataBreaches'),
	path('removeDataBreaches/<int:matterId>/<DataBreachesIds>/', MatterView.removeDataBreaches, name='removeDataBreaches'),
	path('addContracts/<int:matterId>/<ContractsIds>/', MatterView.addContracts, name='addContracts'),
	path('removeContracts/<int:matterId>/<ContractsIds>/', MatterView.removeContracts, name='removeContracts'),
]
