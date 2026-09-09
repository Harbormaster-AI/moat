from django.urls import path
from inventoryOnDjango.views import CycleCountEntryView

urlpatterns = [
    path('', CycleCountEntryView.index, name='index'),
	path('create', CycleCountEntryView.get, name='create'),
	path('get/<int:cycleCountEntryId>/', CycleCountEntryView.get, name='get'),
	path('save', CycleCountEntryView.save, name='save'),
	path('getAll', CycleCountEntryView.getAll, name='getAll'),
	path('delete/<int:cycleCountEntryId>/', CycleCountEntryView.delete, name='delete'),
	path('assignCycleCount/<int:cycleCountEntryId>/<int:CycleCountId>/', CycleCountEntryView.assignCycleCount, name='assignCycleCount'),
	path('unassignCycleCount/<int:cycleCountEntryId>/', CycleCountEntryView.unassignCycleCount, name='unassignCycleCount'),
	path('assignSku/<int:cycleCountEntryId>/<int:SkuId>/', CycleCountEntryView.assignSku, name='assignSku'),
	path('unassignSku/<int:cycleCountEntryId>/', CycleCountEntryView.unassignSku, name='unassignSku'),
	path('assignLot/<int:cycleCountEntryId>/<int:LotId>/', CycleCountEntryView.assignLot, name='assignLot'),
	path('unassignLot/<int:cycleCountEntryId>/', CycleCountEntryView.unassignLot, name='unassignLot'),
	path('assignLocation/<int:cycleCountEntryId>/<int:LocationId>/', CycleCountEntryView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:cycleCountEntryId>/', CycleCountEntryView.unassignLocation, name='unassignLocation'),
	path('addSerialNumbers/<int:cycleCountEntryId>/<SerialNumbersIds>/', CycleCountEntryView.addSerialNumbers, name='addSerialNumbers'),
	path('removeSerialNumbers/<int:cycleCountEntryId>/<SerialNumbersIds>/', CycleCountEntryView.removeSerialNumbers, name='removeSerialNumbers'),
]
