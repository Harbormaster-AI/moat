from django.urls import path
from insuranceOnDjango.views import ReinsuranceAgreementView

urlpatterns = [
    path('', ReinsuranceAgreementView.index, name='index'),
	path('create', ReinsuranceAgreementView.get, name='create'),
	path('get/<int:reinsuranceAgreementId>/', ReinsuranceAgreementView.get, name='get'),
	path('save', ReinsuranceAgreementView.save, name='save'),
	path('getAll', ReinsuranceAgreementView.getAll, name='getAll'),
	path('delete/<int:reinsuranceAgreementId>/', ReinsuranceAgreementView.delete, name='delete'),
	path('assignInsurer/<int:reinsuranceAgreementId>/<int:InsurerId>/', ReinsuranceAgreementView.assignInsurer, name='assignInsurer'),
	path('unassignInsurer/<int:reinsuranceAgreementId>/', ReinsuranceAgreementView.unassignInsurer, name='unassignInsurer'),
	path('addPolicies/<int:reinsuranceAgreementId>/<PoliciesIds>/', ReinsuranceAgreementView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:reinsuranceAgreementId>/<PoliciesIds>/', ReinsuranceAgreementView.removePolicies, name='removePolicies'),
]
