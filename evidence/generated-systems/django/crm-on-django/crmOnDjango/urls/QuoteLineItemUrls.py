from django.urls import path
from crmOnDjango.views import QuoteLineItemView

urlpatterns = [
    path('', QuoteLineItemView.index, name='index'),
	path('create', QuoteLineItemView.get, name='create'),
	path('get/<int:quoteLineItemId>/', QuoteLineItemView.get, name='get'),
	path('save', QuoteLineItemView.save, name='save'),
	path('getAll', QuoteLineItemView.getAll, name='getAll'),
	path('delete/<int:quoteLineItemId>/', QuoteLineItemView.delete, name='delete'),
	path('assignQuote/<int:quoteLineItemId>/<int:QuoteId>/', QuoteLineItemView.assignQuote, name='assignQuote'),
	path('unassignQuote/<int:quoteLineItemId>/', QuoteLineItemView.unassignQuote, name='unassignQuote'),
	path('assignProduct/<int:quoteLineItemId>/<int:ProductId>/', QuoteLineItemView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:quoteLineItemId>/', QuoteLineItemView.unassignProduct, name='unassignProduct'),
	path('assignPriceBookEntry/<int:quoteLineItemId>/<int:PriceBookEntryId>/', QuoteLineItemView.assignPriceBookEntry, name='assignPriceBookEntry'),
	path('unassignPriceBookEntry/<int:quoteLineItemId>/', QuoteLineItemView.unassignPriceBookEntry, name='unassignPriceBookEntry'),
	path('assignOpportunityLineItem/<int:quoteLineItemId>/<int:OpportunityLineItemId>/', QuoteLineItemView.assignOpportunityLineItem, name='assignOpportunityLineItem'),
	path('unassignOpportunityLineItem/<int:quoteLineItemId>/', QuoteLineItemView.unassignOpportunityLineItem, name='unassignOpportunityLineItem'),
]
