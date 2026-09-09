from django.urls import path
from governanceOnDjango.views import ObligationView

urlpatterns = [
    path('', ObligationView.index, name='index'),
	path('create', ObligationView.get, name='create'),
	path('get/<int:obligationId>/', ObligationView.get, name='get'),
	path('save', ObligationView.save, name='save'),
	path('getAll', ObligationView.getAll, name='getAll'),
	path('delete/<int:obligationId>/', ObligationView.delete, name='delete'),
	path('assignRegulation/<int:obligationId>/<int:RegulationId>/', ObligationView.assignRegulation, name='assignRegulation'),
	path('unassignRegulation/<int:obligationId>/', ObligationView.unassignRegulation, name='unassignRegulation'),
	path('addControls/<int:obligationId>/<ControlsIds>/', ObligationView.addControls, name='addControls'),
	path('removeControls/<int:obligationId>/<ControlsIds>/', ObligationView.removeControls, name='removeControls'),
	path('addPolicies/<int:obligationId>/<PoliciesIds>/', ObligationView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:obligationId>/<PoliciesIds>/', ObligationView.removePolicies, name='removePolicies'),
	path('addContracts/<int:obligationId>/<ContractsIds>/', ObligationView.addContracts, name='addContracts'),
	path('removeContracts/<int:obligationId>/<ContractsIds>/', ObligationView.removeContracts, name='removeContracts'),
]
