from django.urls import path
from advertisingOnDjango.views import AdSlotView

urlpatterns = [
    path('', AdSlotView.index, name='index'),
	path('create', AdSlotView.get, name='create'),
	path('get/<int:adSlotId>/', AdSlotView.get, name='get'),
	path('save', AdSlotView.save, name='save'),
	path('getAll', AdSlotView.getAll, name='getAll'),
	path('delete/<int:adSlotId>/', AdSlotView.delete, name='delete'),
	path('assignInventorySource/<int:adSlotId>/<int:InventorySourceId>/', AdSlotView.assignInventorySource, name='assignInventorySource'),
	path('unassignInventorySource/<int:adSlotId>/', AdSlotView.unassignInventorySource, name='unassignInventorySource'),
	path('addPlacements/<int:adSlotId>/<PlacementsIds>/', AdSlotView.addPlacements, name='addPlacements'),
	path('removePlacements/<int:adSlotId>/<PlacementsIds>/', AdSlotView.removePlacements, name='removePlacements'),
	path('addRates/<int:adSlotId>/<RatesIds>/', AdSlotView.addRates, name='addRates'),
	path('removeRates/<int:adSlotId>/<RatesIds>/', AdSlotView.removeRates, name='removeRates'),
]
