from django.urls import path
from crmOnDjango.views import OpportunityLineItemView

urlpatterns = [
    path('', OpportunityLineItemView.index, name='index'),
	path('create', OpportunityLineItemView.get, name='create'),
	path('get/<int:opportunityLineItemId>/', OpportunityLineItemView.get, name='get'),
	path('save', OpportunityLineItemView.save, name='save'),
	path('getAll', OpportunityLineItemView.getAll, name='getAll'),
	path('delete/<int:opportunityLineItemId>/', OpportunityLineItemView.delete, name='delete'),
	path('assignOpportunity/<int:opportunityLineItemId>/<int:OpportunityId>/', OpportunityLineItemView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:opportunityLineItemId>/', OpportunityLineItemView.unassignOpportunity, name='unassignOpportunity'),
	path('assignProduct/<int:opportunityLineItemId>/<int:ProductId>/', OpportunityLineItemView.assignProduct, name='assignProduct'),
	path('unassignProduct/<int:opportunityLineItemId>/', OpportunityLineItemView.unassignProduct, name='unassignProduct'),
	path('assignPriceBookEntry/<int:opportunityLineItemId>/<int:PriceBookEntryId>/', OpportunityLineItemView.assignPriceBookEntry, name='assignPriceBookEntry'),
	path('unassignPriceBookEntry/<int:opportunityLineItemId>/', OpportunityLineItemView.unassignPriceBookEntry, name='unassignPriceBookEntry'),
]
