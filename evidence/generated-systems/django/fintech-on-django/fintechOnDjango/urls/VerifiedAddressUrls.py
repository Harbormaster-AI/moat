from django.urls import path
from fintechOnDjango.views import VerifiedAddressView

urlpatterns = [
    path('', VerifiedAddressView.index, name='index'),
	path('create', VerifiedAddressView.get, name='create'),
	path('get/<int:verifiedAddressId>/', VerifiedAddressView.get, name='get'),
	path('save', VerifiedAddressView.save, name='save'),
	path('getAll', VerifiedAddressView.getAll, name='getAll'),
	path('delete/<int:verifiedAddressId>/', VerifiedAddressView.delete, name='delete'),
	path('assignKycProfile/<int:verifiedAddressId>/<int:KycProfileId>/', VerifiedAddressView.assignKycProfile, name='assignKycProfile'),
	path('unassignKycProfile/<int:verifiedAddressId>/', VerifiedAddressView.unassignKycProfile, name='unassignKycProfile'),
]
