from django.urls import path
from insuranceOnDjango.views import AdjusterView

urlpatterns = [
    path('', AdjusterView.index, name='index'),
	path('create', AdjusterView.get, name='create'),
	path('get/<int:adjusterId>/', AdjusterView.get, name='get'),
	path('save', AdjusterView.save, name='save'),
	path('getAll', AdjusterView.getAll, name='getAll'),
	path('delete/<int:adjusterId>/', AdjusterView.delete, name='delete'),
	path('addClaims/<int:adjusterId>/<ClaimsIds>/', AdjusterView.addClaims, name='addClaims'),
	path('removeClaims/<int:adjusterId>/<ClaimsIds>/', AdjusterView.removeClaims, name='removeClaims'),
	path('addServiceProviders/<int:adjusterId>/<ServiceProvidersIds>/', AdjusterView.addServiceProviders, name='addServiceProviders'),
	path('removeServiceProviders/<int:adjusterId>/<ServiceProvidersIds>/', AdjusterView.removeServiceProviders, name='removeServiceProviders'),
]
