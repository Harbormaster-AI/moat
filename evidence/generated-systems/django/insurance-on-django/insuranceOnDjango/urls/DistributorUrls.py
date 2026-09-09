from django.urls import path
from insuranceOnDjango.views import DistributorView

urlpatterns = [
    path('', DistributorView.index, name='index'),
	path('create', DistributorView.get, name='create'),
	path('get/<int:distributorId>/', DistributorView.get, name='get'),
	path('save', DistributorView.save, name='save'),
	path('getAll', DistributorView.getAll, name='getAll'),
	path('delete/<int:distributorId>/', DistributorView.delete, name='delete'),
	path('addInsurers/<int:distributorId>/<InsurersIds>/', DistributorView.addInsurers, name='addInsurers'),
	path('removeInsurers/<int:distributorId>/<InsurersIds>/', DistributorView.removeInsurers, name='removeInsurers'),
	path('addAgents/<int:distributorId>/<AgentsIds>/', DistributorView.addAgents, name='addAgents'),
	path('removeAgents/<int:distributorId>/<AgentsIds>/', DistributorView.removeAgents, name='removeAgents'),
	path('addPolicies/<int:distributorId>/<PoliciesIds>/', DistributorView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:distributorId>/<PoliciesIds>/', DistributorView.removePolicies, name='removePolicies'),
]
