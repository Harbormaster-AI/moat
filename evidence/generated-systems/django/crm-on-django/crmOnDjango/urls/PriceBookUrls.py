from django.urls import path
from crmOnDjango.views import PriceBookView

urlpatterns = [
    path('', PriceBookView.index, name='index'),
	path('create', PriceBookView.get, name='create'),
	path('get/<int:priceBookId>/', PriceBookView.get, name='get'),
	path('save', PriceBookView.save, name='save'),
	path('getAll', PriceBookView.getAll, name='getAll'),
	path('delete/<int:priceBookId>/', PriceBookView.delete, name='delete'),
	path('assignOrganization/<int:priceBookId>/<int:OrganizationId>/', PriceBookView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:priceBookId>/', PriceBookView.unassignOrganization, name='unassignOrganization'),
	path('addEntries/<int:priceBookId>/<EntriesIds>/', PriceBookView.addEntries, name='addEntries'),
	path('removeEntries/<int:priceBookId>/<EntriesIds>/', PriceBookView.removeEntries, name='removeEntries'),
	path('addQuotes/<int:priceBookId>/<QuotesIds>/', PriceBookView.addQuotes, name='addQuotes'),
	path('removeQuotes/<int:priceBookId>/<QuotesIds>/', PriceBookView.removeQuotes, name='removeQuotes'),
	path('addOrders/<int:priceBookId>/<OrdersIds>/', PriceBookView.addOrders, name='addOrders'),
	path('removeOrders/<int:priceBookId>/<OrdersIds>/', PriceBookView.removeOrders, name='removeOrders'),
]
