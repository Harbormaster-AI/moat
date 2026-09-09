from django.urls import path
from insuranceOnDjango.views import ApplicationView

urlpatterns = [
    path('', ApplicationView.index, name='index'),
	path('create', ApplicationView.get, name='create'),
	path('get/<int:applicationId>/', ApplicationView.get, name='get'),
	path('save', ApplicationView.save, name='save'),
	path('getAll', ApplicationView.getAll, name='getAll'),
	path('delete/<int:applicationId>/', ApplicationView.delete, name='delete'),
	path('assignCustomer/<int:applicationId>/<int:CustomerId>/', ApplicationView.assignCustomer, name='assignCustomer'),
	path('unassignCustomer/<int:applicationId>/', ApplicationView.unassignCustomer, name='unassignCustomer'),
	path('assignProduct/<int:applicationId>/<int:ProductId>/', ApplicationView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:applicationId>/', ApplicationView.unassignProduct, name='unassignProduct'),
	path('assignDistributor/<int:applicationId>/<int:DistributorId>/', ApplicationView.assignDistributor, name='assignDistributor'),
	path('unassignDistributor/<int:applicationId>/', ApplicationView.unassignDistributor, name='unassignDistributor'),
	path('assignSelectedQuote/<int:applicationId>/<int:SelectedQuoteId>/', ApplicationView.assignSelectedQuote, name='assignSelectedQuote'),
	path('unassignSelectedQuote/<int:applicationId>/', ApplicationView.unassignSelectedQuote, name='unassignSelectedQuote'),
	path('addQuotes/<int:applicationId>/<QuotesIds>/', ApplicationView.addQuotes, name='addQuotes'),
	path('removeQuotes/<int:applicationId>/<QuotesIds>/', ApplicationView.removeQuotes, name='removeQuotes'),
]
