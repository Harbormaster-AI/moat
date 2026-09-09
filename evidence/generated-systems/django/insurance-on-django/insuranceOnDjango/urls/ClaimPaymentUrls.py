from django.urls import path
from insuranceOnDjango.views import ClaimPaymentView

urlpatterns = [
    path('', ClaimPaymentView.index, name='index'),
	path('create', ClaimPaymentView.get, name='create'),
	path('get/<int:claimPaymentId>/', ClaimPaymentView.get, name='get'),
	path('save', ClaimPaymentView.save, name='save'),
	path('getAll', ClaimPaymentView.getAll, name='getAll'),
	path('delete/<int:claimPaymentId>/', ClaimPaymentView.delete, name='delete'),
	path('assignClaim/<int:claimPaymentId>/<int:ClaimId>/', ClaimPaymentView.assignClaim, name='assignClaim'),
	path('unassignClaim/<int:claimPaymentId>/', ClaimPaymentView.unassignClaim, name='unassignClaim'),
	path('assignExposure/<int:claimPaymentId>/<int:ExposureId>/', ClaimPaymentView.assignExposure, name='assignExposure'),
	path('unassignExposure/<int:claimPaymentId>/', ClaimPaymentView.unassignExposure, name='unassignExposure'),
	path('assignBeneficiary/<int:claimPaymentId>/<int:BeneficiaryId>/', ClaimPaymentView.assignBeneficiary, name='assignBeneficiary'),
	path('unassignBeneficiary/<int:claimPaymentId>/', ClaimPaymentView.unassignBeneficiary, name='unassignBeneficiary'),
	path('assignServiceProvider/<int:claimPaymentId>/<int:ServiceProviderId>/', ClaimPaymentView.assignServiceProvider, name='assignServiceProvider'),
	path('unassignServiceProvider/<int:claimPaymentId>/', ClaimPaymentView.unassignServiceProvider, name='unassignServiceProvider'),
	path('assignCustomer/<int:claimPaymentId>/<int:CustomerId>/', ClaimPaymentView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:claimPaymentId>/', ClaimPaymentView.unassignCustomer, name='unassignCustomer'),
]
