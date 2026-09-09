from django.urls import path
from crmOnDjango.views import ProductView

urlpatterns = [
    path('', ProductView.index, name='index'),
	path('create', ProductView.get, name='create'),
	path('get/<int:productId>/', ProductView.get, name='get'),
	path('save', ProductView.save, name='save'),
	path('getAll', ProductView.getAll, name='getAll'),
	path('delete/<int:productId>/', ProductView.delete, name='delete'),
	path('assignOrganization/<int:productId>/<int:OrganizationId>/', ProductView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:productId>/', ProductView.unassignOrganization, name='unassignOrganization'),
	path('addPriceBookEntries/<int:productId>/<PriceBookEntriesIds>/', ProductView.addPriceBookEntries, name='addPriceBookEntries'),
	path('removePriceBookEntries/<int:productId>/<PriceBookEntriesIds>/', ProductView.removePriceBookEntries, name='removePriceBookEntries'),
	path('addOpportunityLineItems/<int:productId>/<OpportunityLineItemsIds>/', ProductView.addOpportunityLineItems, name='addOpportunityLineItems'),
	path('removeOpportunityLineItems/<int:productId>/<OpportunityLineItemsIds>/', ProductView.removeOpportunityLineItems, name='removeOpportunityLineItems'),
	path('addQuoteLineItems/<int:productId>/<QuoteLineItemsIds>/', ProductView.addQuoteLineItems, name='addQuoteLineItems'),
	path('removeQuoteLineItems/<int:productId>/<QuoteLineItemsIds>/', ProductView.removeQuoteLineItems, name='removeQuoteLineItems'),
	path('addOrderItems/<int:productId>/<OrderItemsIds>/', ProductView.addOrderItems, name='addOrderItems'),
	path('removeOrderItems/<int:productId>/<OrderItemsIds>/', ProductView.removeOrderItems, name='removeOrderItems'),
]
