from django.urls import path
from advertisingOnDjango.views import PublisherView

urlpatterns = [
    path('', PublisherView.index, name='index'),
	path('create', PublisherView.get, name='create'),
	path('get/<int:publisherId>/', PublisherView.get, name='get'),
	path('save', PublisherView.save, name='save'),
	path('getAll', PublisherView.getAll, name='getAll'),
	path('delete/<int:publisherId>/', PublisherView.delete, name='delete'),
	path('addInventorySources/<int:publisherId>/<InventorySourcesIds>/', PublisherView.addInventorySources, name='addInventorySources'),
	path('removeInventorySources/<int:publisherId>/<InventorySourcesIds>/', PublisherView.removeInventorySources, name='removeInventorySources'),
	path('addDeals/<int:publisherId>/<DealsIds>/', PublisherView.addDeals, name='addDeals'),
	path('removeDeals/<int:publisherId>/<DealsIds>/', PublisherView.removeDeals, name='removeDeals'),
	path('addCreativeApprovals/<int:publisherId>/<CreativeApprovalsIds>/', PublisherView.addCreativeApprovals, name='addCreativeApprovals'),
	path('removeCreativeApprovals/<int:publisherId>/<CreativeApprovalsIds>/', PublisherView.removeCreativeApprovals, name='removeCreativeApprovals'),
	path('addInsertionOrders/<int:publisherId>/<InsertionOrdersIds>/', PublisherView.addInsertionOrders, name='addInsertionOrders'),
	path('removeInsertionOrders/<int:publisherId>/<InsertionOrdersIds>/', PublisherView.removeInsertionOrders, name='removeInsertionOrders'),
	path('addRateCards/<int:publisherId>/<RateCardsIds>/', PublisherView.addRateCards, name='addRateCards'),
	path('removeRateCards/<int:publisherId>/<RateCardsIds>/', PublisherView.removeRateCards, name='removeRateCards'),
]
