from django.urls import path
from aerospaceOnDjango.views import SalesRegionView

urlpatterns = [
    path('', SalesRegionView.index, name='index'),
	path('create', SalesRegionView.get, name='create'),
	path('get/<int:salesRegionId>/', SalesRegionView.get, name='get'),
	path('save', SalesRegionView.save, name='save'),
	path('getAll', SalesRegionView.getAll, name='getAll'),
	path('delete/<int:salesRegionId>/', SalesRegionView.delete, name='delete'),
	path('addOperators/<int:salesRegionId>/<OperatorsIds>/', SalesRegionView.addOperators, name='addOperators'),
	path('removeOperators/<int:salesRegionId>/<OperatorsIds>/', SalesRegionView.removeOperators, name='removeOperators'),
	path('addSalesCampaigns/<int:salesRegionId>/<SalesCampaignsIds>/', SalesRegionView.addSalesCampaigns, name='addSalesCampaigns'),
	path('removeSalesCampaigns/<int:salesRegionId>/<SalesCampaignsIds>/', SalesRegionView.removeSalesCampaigns, name='removeSalesCampaigns'),
]
