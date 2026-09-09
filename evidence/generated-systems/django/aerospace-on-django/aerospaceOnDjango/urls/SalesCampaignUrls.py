from django.urls import path
from aerospaceOnDjango.views import SalesCampaignView

urlpatterns = [
    path('', SalesCampaignView.index, name='index'),
	path('create', SalesCampaignView.get, name='create'),
	path('get/<int:salesCampaignId>/', SalesCampaignView.get, name='get'),
	path('save', SalesCampaignView.save, name='save'),
	path('getAll', SalesCampaignView.getAll, name='getAll'),
	path('delete/<int:salesCampaignId>/', SalesCampaignView.delete, name='delete'),
	path('assignRegion/<int:salesCampaignId>/<int:RegionId>/', SalesCampaignView.assignRegion, name='assignRegion'),
	path('unassignRegion/<int:salesCampaignId>/', SalesCampaignView.unassignRegion, name='unassignRegion'),
	path('assignOperator/<int:salesCampaignId>/<int:OperatorId>/', SalesCampaignView.assignOperator, name='assignOperator'),
	path('unassignOperator/<int:salesCampaignId>/', SalesCampaignView.unassignOperator, name='unassignOperator'),
	path('addQuotes/<int:salesCampaignId>/<QuotesIds>/', SalesCampaignView.addQuotes, name='addQuotes'),
	path('removeQuotes/<int:salesCampaignId>/<QuotesIds>/', SalesCampaignView.removeQuotes, name='removeQuotes'),
]
