from django.urls import path
from insuranceOnDjango.views import ServiceProviderView

urlpatterns = [
    path('', ServiceProviderView.index, name='index'),
	path('create', ServiceProviderView.get, name='create'),
	path('get/<int:serviceProviderId>/', ServiceProviderView.get, name='get'),
	path('save', ServiceProviderView.save, name='save'),
	path('getAll', ServiceProviderView.getAll, name='getAll'),
	path('delete/<int:serviceProviderId>/', ServiceProviderView.delete, name='delete'),
	path('addClaims/<int:serviceProviderId>/<ClaimsIds>/', ServiceProviderView.addClaims, name='addClaims'),
	path('removeClaims/<int:serviceProviderId>/<ClaimsIds>/', ServiceProviderView.removeClaims, name='removeClaims'),
]
