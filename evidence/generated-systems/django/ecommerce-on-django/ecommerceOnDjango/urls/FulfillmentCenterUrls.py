from django.urls import path
from ecommerceOnDjango.views import FulfillmentCenterView

urlpatterns = [
    path('', FulfillmentCenterView.index, name='index'),
	path('create', FulfillmentCenterView.get, name='create'),
	path('get/<int:fulfillmentCenterId>/', FulfillmentCenterView.get, name='get'),
	path('save', FulfillmentCenterView.save, name='save'),
	path('getAll', FulfillmentCenterView.getAll, name='getAll'),
	path('delete/<int:fulfillmentCenterId>/', FulfillmentCenterView.delete, name='delete'),
	path('assignMerchant/<int:fulfillmentCenterId>/<int:MerchantId>/', FulfillmentCenterView.assignMerchant, name='assignMerchant'),
	path('unassignMerchant/<int:fulfillmentCenterId>/', FulfillmentCenterView.unassignMerchant, name='unassignMerchant'),
	path('addInventoryItems/<int:fulfillmentCenterId>/<InventoryItemsIds>/', FulfillmentCenterView.addInventoryItems, name='addInventoryItems'),
	path('removeInventoryItems/<int:fulfillmentCenterId>/<InventoryItemsIds>/', FulfillmentCenterView.removeInventoryItems, name='removeInventoryItems'),
	path('addShipments/<int:fulfillmentCenterId>/<ShipmentsIds>/', FulfillmentCenterView.addShipments, name='addShipments'),
	path('removeShipments/<int:fulfillmentCenterId>/<ShipmentsIds>/', FulfillmentCenterView.removeShipments, name='removeShipments'),
]
