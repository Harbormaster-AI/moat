from django.urls import path
from crmOnDjango.views import OrderView

urlpatterns = [
    path('', OrderView.index, name='index'),
	path('create', OrderView.get, name='create'),
	path('get/<int:orderId>/', OrderView.get, name='get'),
	path('save', OrderView.save, name='save'),
	path('getAll', OrderView.getAll, name='getAll'),
	path('delete/<int:orderId>/', OrderView.delete, name='delete'),
	path('assignOrganization/<int:orderId>/<int:OrganizationId>/', OrderView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:orderId>/', OrderView.unassignOrganization, name='unassignOrganization'),
	path('assignAccount/<int:orderId>/<int:AccountId>/', OrderView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:orderId>/', OrderView.unassignAccount, name='unassignAccount'),
	path('assignOpportunity/<int:orderId>/<int:OpportunityId>/', OrderView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:orderId>/', OrderView.unassignOpportunity, name='unassignOpportunity'),
	path('assignQuote/<int:orderId>/<int:QuoteId>/', OrderView.assignQuote, name='assignQuote'),
	path('unassignQuote/<int:orderId>/', OrderView.unassignQuote, name='unassignQuote'),
	path('assignOwner/<int:orderId>/<int:OwnerId>/', OrderView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:orderId>/', OrderView.unassignOwner, name='unassignOwner'),
	path('assignContract/<int:orderId>/<int:ContractId>/', OrderView.assignContract, name='assignContract'),
	path('unassignContract/<int:orderId>/', OrderView.unassignContract, name='unassignContract'),
	path('assignPriceBook/<int:orderId>/<int:PriceBookId>/', OrderView.assignPriceBook, name='assignPriceBook'),
	path('unassignPriceBook/<int:orderId>/', OrderView.unassignPriceBook, name='unassignPriceBook'),
	path('addItems/<int:orderId>/<ItemsIds>/', OrderView.addItems, name='addItems'),
	path('removeItems/<int:orderId>/<ItemsIds>/', OrderView.removeItems, name='removeItems'),
]
