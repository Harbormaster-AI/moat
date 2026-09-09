from django.urls import path
from healthcareOnDjango.views import InsurancePlanView

urlpatterns = [
    path('', InsurancePlanView.index, name='index'),
	path('create', InsurancePlanView.get, name='create'),
	path('get/<int:insurancePlanId>/', InsurancePlanView.get, name='get'),
	path('save', InsurancePlanView.save, name='save'),
	path('getAll', InsurancePlanView.getAll, name='getAll'),
	path('delete/<int:insurancePlanId>/', InsurancePlanView.delete, name='delete'),
	path('assignPayer/<int:insurancePlanId>/<int:PayerId>/', InsurancePlanView.assignPayer, name='assignPayer'),
	path('unassignPayer/<int:insurancePlanId>/', InsurancePlanView.unassignPayer, name='unassignPayer'),
	path('addCoverages/<int:insurancePlanId>/<CoveragesIds>/', InsurancePlanView.addCoverages, name='addCoverages'),
	path('removeCoverages/<int:insurancePlanId>/<CoveragesIds>/', InsurancePlanView.removeCoverages, name='removeCoverages'),
]
