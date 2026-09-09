from django.urls import path
from fintechOnDjango.views import ConsentView

urlpatterns = [
    path('', ConsentView.index, name='index'),
	path('create', ConsentView.get, name='create'),
	path('get/<int:consentId>/', ConsentView.get, name='get'),
	path('save', ConsentView.save, name='save'),
	path('getAll', ConsentView.getAll, name='getAll'),
	path('delete/<int:consentId>/', ConsentView.delete, name='delete'),
	path('assignCustomer/<int:consentId>/<int:CustomerId>/', ConsentView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:consentId>/', ConsentView.unassignCustomer, name='unassignCustomer'),
	path('assignApiClient/<int:consentId>/<int:ApiClientId>/', ConsentView.assignApiClient, name='assignApiClient'),
	path('unassignApiClient/<int:consentId>/', ConsentView.unassignApiClient, name='unassignApiClient'),
]
