from django.urls import path
from insuranceOnDjango.views import CustomerView

urlpatterns = [
    path('', CustomerView.index, name='index'),
	path('create', CustomerView.get, name='create'),
	path('get/<int:customerId>/', CustomerView.get, name='get'),
	path('save', CustomerView.save, name='save'),
	path('getAll', CustomerView.getAll, name='getAll'),
	path('delete/<int:customerId>/', CustomerView.delete, name='delete'),
	path('addApplications/<int:customerId>/<ApplicationsIds>/', CustomerView.addApplications, name='addApplications'),
	path('removeApplications/<int:customerId>/<ApplicationsIds>/', CustomerView.removeApplications, name='removeApplications'),
	path('addPolicies/<int:customerId>/<PoliciesIds>/', CustomerView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:customerId>/<PoliciesIds>/', CustomerView.removePolicies, name='removePolicies'),
	path('addClaims/<int:customerId>/<ClaimsIds>/', CustomerView.addClaims, name='addClaims'),
	path('removeClaims/<int:customerId>/<ClaimsIds>/', CustomerView.removeClaims, name='removeClaims'),
	path('addAgents/<int:customerId>/<AgentsIds>/', CustomerView.addAgents, name='addAgents'),
	path('removeAgents/<int:customerId>/<AgentsIds>/', CustomerView.removeAgents, name='removeAgents'),
	path('addBeneficiaries/<int:customerId>/<BeneficiariesIds>/', CustomerView.addBeneficiaries, name='addBeneficiaries'),
	path('removeBeneficiaries/<int:customerId>/<BeneficiariesIds>/', CustomerView.removeBeneficiaries, name='removeBeneficiaries'),
]
