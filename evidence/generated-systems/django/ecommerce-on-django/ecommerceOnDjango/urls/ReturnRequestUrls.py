from django.urls import path
from ecommerceOnDjango.views import ReturnRequestView

urlpatterns = [
    path('', ReturnRequestView.index, name='index'),
	path('create', ReturnRequestView.get, name='create'),
	path('get/<int:returnRequestId>/', ReturnRequestView.get, name='get'),
	path('save', ReturnRequestView.save, name='save'),
	path('getAll', ReturnRequestView.getAll, name='getAll'),
	path('delete/<int:returnRequestId>/', ReturnRequestView.delete, name='delete'),
	path('assignOrder/<int:returnRequestId>/<int:OrderId>/', ReturnRequestView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:returnRequestId>/', ReturnRequestView.unassignOrder, name='unassignOrder'),
	path('assignRefund/<int:returnRequestId>/<int:RefundId>/', ReturnRequestView.assignRefund, name='assignRefund'),
	path('unassignRefund/<int:returnRequestId>/', ReturnRequestView.unassignRefund, name='unassignRefund'),
	path('assignShipment/<int:returnRequestId>/<int:ShipmentId>/', ReturnRequestView.assignShipment, name='assignShipment'),
	path('unassignShipment/<int:returnRequestId>/', ReturnRequestView.unassignShipment, name='unassignShipment'),
	path('addItems/<int:returnRequestId>/<ItemsIds>/', ReturnRequestView.addItems, name='addItems'),
	path('removeItems/<int:returnRequestId>/<ItemsIds>/', ReturnRequestView.removeItems, name='removeItems'),
]
