from django.urls import path
from insuranceOnDjango.views import AgentView

urlpatterns = [
    path('', AgentView.index, name='index'),
	path('create', AgentView.get, name='create'),
	path('get/<int:agentId>/', AgentView.get, name='get'),
	path('save', AgentView.save, name='save'),
	path('getAll', AgentView.getAll, name='getAll'),
	path('delete/<int:agentId>/', AgentView.delete, name='delete'),
	path('assignDistributor/<int:agentId>/<int:DistributorId>/', AgentView.assignDistributor, name='assignDistributor'),
	path('unassignDistributor/<int:agentId>/', AgentView.unassignDistributor, name='unassignDistributor'),
	path('addPolicies/<int:agentId>/<PoliciesIds>/', AgentView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:agentId>/<PoliciesIds>/', AgentView.removePolicies, name='removePolicies'),
	path('addCustomers/<int:agentId>/<CustomersIds>/', AgentView.addCustomers, name='addCustomers'),
	path('removeCustomers/<int:agentId>/<CustomersIds>/', AgentView.removeCustomers, name='removeCustomers'),
]
