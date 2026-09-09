from django.urls import path
from healthcareOnDjango.views import InsurancePayerView

urlpatterns = [
    path('', InsurancePayerView.index, name='index'),
	path('create', InsurancePayerView.get, name='create'),
	path('get/<int:insurancePayerId>/', InsurancePayerView.get, name='get'),
	path('save', InsurancePayerView.save, name='save'),
	path('getAll', InsurancePayerView.getAll, name='getAll'),
	path('delete/<int:insurancePayerId>/', InsurancePayerView.delete, name='delete'),
	path('addPlans/<int:insurancePayerId>/<PlansIds>/', InsurancePayerView.addPlans, name='addPlans'),
	path('removePlans/<int:insurancePayerId>/<PlansIds>/', InsurancePayerView.removePlans, name='removePlans'),
	path('addClaims/<int:insurancePayerId>/<ClaimsIds>/', InsurancePayerView.addClaims, name='addClaims'),
	path('removeClaims/<int:insurancePayerId>/<ClaimsIds>/', InsurancePayerView.removeClaims, name='removeClaims'),
]
