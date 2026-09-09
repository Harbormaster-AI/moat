from django.urls import path
from inventoryOnDjango.views import InventoryThresholdAlertView

urlpatterns = [
    path('', InventoryThresholdAlertView.index, name='index'),
	path('create', InventoryThresholdAlertView.get, name='create'),
	path('get/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.get, name='get'),
	path('save', InventoryThresholdAlertView.save, name='save'),
	path('getAll', InventoryThresholdAlertView.getAll, name='getAll'),
	path('delete/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.delete, name='delete'),
	path('assignSku/<int:inventoryThresholdAlertId>/<int:SkuId>/', InventoryThresholdAlertView.assignSku, name='assignSku'),
	path('unassignSku/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:inventoryThresholdAlertId>/<int:WarehouseId>/', InventoryThresholdAlertView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLocation/<int:inventoryThresholdAlertId>/<int:LocationId>/', InventoryThresholdAlertView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.unassignLocation, name='unassignLocation'),
	path('assignRelatedPolicy/<int:inventoryThresholdAlertId>/<int:RelatedPolicyId>/', InventoryThresholdAlertView.assignRelatedPolicy, name='assignRelatedPolicy'),
	path('unassignRelatedPolicy/<int:inventoryThresholdAlertId>/', InventoryThresholdAlertView.unassignRelatedPolicy, name='unassignRelatedPolicy'),
]
