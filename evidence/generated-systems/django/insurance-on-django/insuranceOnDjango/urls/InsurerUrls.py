from django.urls import path
from insuranceOnDjango.views import InsurerView

urlpatterns = [
    path('', InsurerView.index, name='index'),
	path('create', InsurerView.get, name='create'),
	path('get/<int:insurerId>/', InsurerView.get, name='get'),
	path('save', InsurerView.save, name='save'),
	path('getAll', InsurerView.getAll, name='getAll'),
	path('delete/<int:insurerId>/', InsurerView.delete, name='delete'),
	path('addProducts/<int:insurerId>/<ProductsIds>/', InsurerView.addProducts, name='addProducts'),
	path('removeProducts/<int:insurerId>/<ProductsIds>/', InsurerView.removeProducts, name='removeProducts'),
	path('addDistributionPartners/<int:insurerId>/<DistributionPartnersIds>/', InsurerView.addDistributionPartners, name='addDistributionPartners'),
	path('removeDistributionPartners/<int:insurerId>/<DistributionPartnersIds>/', InsurerView.removeDistributionPartners, name='removeDistributionPartners'),
	path('addPolicies/<int:insurerId>/<PoliciesIds>/', InsurerView.addPolicies, name='addPolicies'),
	path('removePolicies/<int:insurerId>/<PoliciesIds>/', InsurerView.removePolicies, name='removePolicies'),
	path('addClaims/<int:insurerId>/<ClaimsIds>/', InsurerView.addClaims, name='addClaims'),
	path('removeClaims/<int:insurerId>/<ClaimsIds>/', InsurerView.removeClaims, name='removeClaims'),
	path('addReinsuranceAgreements/<int:insurerId>/<ReinsuranceAgreementsIds>/', InsurerView.addReinsuranceAgreements, name='addReinsuranceAgreements'),
	path('removeReinsuranceAgreements/<int:insurerId>/<ReinsuranceAgreementsIds>/', InsurerView.removeReinsuranceAgreements, name='removeReinsuranceAgreements'),
]
