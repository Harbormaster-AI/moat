from django.urls import path
from crmOnDjango.views import QuoteView

urlpatterns = [
    path('', QuoteView.index, name='index'),
	path('create', QuoteView.get, name='create'),
	path('get/<int:quoteId>/', QuoteView.get, name='get'),
	path('save', QuoteView.save, name='save'),
	path('getAll', QuoteView.getAll, name='getAll'),
	path('delete/<int:quoteId>/', QuoteView.delete, name='delete'),
	path('assignOrganization/<int:quoteId>/<int:OrganizationId>/', QuoteView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:quoteId>/', QuoteView.unassignOrganization, name='unassignOrganization'),
	path('assignAccount/<int:quoteId>/<int:AccountId>/', QuoteView.assignAccount, name='assignAccount'),
	path('unassignAccount/<int:quoteId>/', QuoteView.unassignAccount, name='unassignAccount'),
	path('assignOpportunity/<int:quoteId>/<int:OpportunityId>/', QuoteView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:quoteId>/', QuoteView.unassignOpportunity, name='unassignOpportunity'),
	path('assignOwner/<int:quoteId>/<int:OwnerId>/', QuoteView.assignOwner, name='assignOwner'),
	path('unassignOwner/<int:quoteId>/', QuoteView.unassignOwner, name='unassignOwner'),
	path('assignPriceBook/<int:quoteId>/<int:PriceBookId>/', QuoteView.assignPriceBook, name='assignPriceBook'),
	path('unassignPriceBook/<int:quoteId>/', QuoteView.unassignPriceBook, name='unassignPriceBook'),
	path('assignOrder/<int:quoteId>/<int:OrderId>/', QuoteView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:quoteId>/', QuoteView.unassignOrder, name='unassignOrder'),
	path('addLineItems/<int:quoteId>/<LineItemsIds>/', QuoteView.addLineItems, name='addLineItems'),
	path('removeLineItems/<int:quoteId>/<LineItemsIds>/', QuoteView.removeLineItems, name='removeLineItems'),
]
