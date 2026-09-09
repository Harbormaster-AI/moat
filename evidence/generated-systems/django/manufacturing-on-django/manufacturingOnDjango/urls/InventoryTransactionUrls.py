from django.urls import path
from manufacturingOnDjango.views import InventoryTransactionView

urlpatterns = [
    path('', InventoryTransactionView.index, name='index'),
	path('create', InventoryTransactionView.get, name='create'),
	path('get/<int:inventoryTransactionId>/', InventoryTransactionView.get, name='get'),
	path('save', InventoryTransactionView.save, name='save'),
	path('getAll', InventoryTransactionView.getAll, name='getAll'),
	path('delete/<int:inventoryTransactionId>/', InventoryTransactionView.delete, name='delete'),
	path('assignItem/<int:inventoryTransactionId>/<int:ItemId>/', InventoryTransactionView.assignItem, name='assignItem'),
	path('unassignItem/<int:inventoryTransactionId>/', InventoryTransactionView.unassignItem, name='unassignItem'),
	path('assignLocation/<int:inventoryTransactionId>/<int:LocationId>/', InventoryTransactionView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:inventoryTransactionId>/', InventoryTransactionView.unassignLocation, name='unassignLocation'),
	path('assignWorkOrder/<int:inventoryTransactionId>/<int:WorkOrderId>/', InventoryTransactionView.assignWorkOrder, name='assignWorkOrder'),
	path('unassignWorkOrder/<int:inventoryTransactionId>/', InventoryTransactionView.unassignWorkOrder, name='unassignWorkOrder'),
	path('assignPurchaseOrder/<int:inventoryTransactionId>/<int:PurchaseOrderId>/', InventoryTransactionView.assignPurchaseOrder, name='assignPurchaseOrder'),
	path('unassignPurchaseOrder/<int:inventoryTransactionId>/', InventoryTransactionView.unassignPurchaseOrder, name='unassignPurchaseOrder'),
	path('assignSalesOrder/<int:inventoryTransactionId>/<int:SalesOrderId>/', InventoryTransactionView.assignSalesOrder, name='assignSalesOrder'),
	path('unassignSalesOrder/<int:inventoryTransactionId>/', InventoryTransactionView.unassignSalesOrder, name='unassignSalesOrder'),
]
