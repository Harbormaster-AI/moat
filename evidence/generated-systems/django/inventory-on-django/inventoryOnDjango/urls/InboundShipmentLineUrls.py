from django.urls import path
from inventoryOnDjango.views import InboundShipmentLineView

urlpatterns = [
    path('', InboundShipmentLineView.index, name='index'),
	path('create', InboundShipmentLineView.get, name='create'),
	path('get/<int:inboundShipmentLineId>/', InboundShipmentLineView.get, name='get'),
	path('save', InboundShipmentLineView.save, name='save'),
	path('getAll', InboundShipmentLineView.getAll, name='getAll'),
	path('delete/<int:inboundShipmentLineId>/', InboundShipmentLineView.delete, name='delete'),
	path('assignInboundShipment/<int:inboundShipmentLineId>/<int:InboundShipmentId>/', InboundShipmentLineView.assignInboundShipment, name='assignInboundShipment'),
	path('unassignInboundShipment/<int:inboundShipmentLineId>/', InboundShipmentLineView.unassignInboundShipment, name='unassignInboundShipment'),
	path('assignSku/<int:inboundShipmentLineId>/<int:SkuId>/', InboundShipmentLineView.assignSku, name='assignSku'),
	path('unassignSku/<int:inboundShipmentLineId>/', InboundShipmentLineView.unassignSku, name='unassignSku'),
	path('assignLot/<int:inboundShipmentLineId>/<int:LotId>/', InboundShipmentLineView.assignLot, name='assignLot'),
	path('unassignLot/<int:inboundShipmentLineId>/', InboundShipmentLineView.unassignLot, name='unassignLot'),
	path('assignDestinationLocation/<int:inboundShipmentLineId>/<int:DestinationLocationId>/', InboundShipmentLineView.assignDestinationLocation, name='assignDestinationLocation'),
	path('unassignDestinationLocation/<int:inboundShipmentLineId>/', InboundShipmentLineView.unassignDestinationLocation, name='unassignDestinationLocation'),
	path('addSerialNumbers/<int:inboundShipmentLineId>/<SerialNumbersIds>/', InboundShipmentLineView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:inboundShipmentLineId>/<SerialNumbersIds>/', InboundShipmentLineView.removeSerialNumbers, name='removeSerialNumbers'),
]
