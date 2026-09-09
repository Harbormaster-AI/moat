from django.urls import path
from advertisingOnDjango.views import DealView

urlpatterns = [
    path('', DealView.index, name='index'),
	path('create', DealView.get, name='create'),
	path('get/<int:dealId>/', DealView.get, name='get'),
	path('save', DealView.save, name='save'),
	path('getAll', DealView.getAll, name='getAll'),
	path('delete/<int:dealId>/', DealView.delete, name='delete'),
	path('assignPublisher/<int:dealId>/<int:PublisherId>/', DealView.assignPublisher, name='assignPublisher'),
	path('unassignPublisher/<int:dealId>/', DealView.unassignPublisher, name='unassignPublisher'),
	path('addInventorySources/<int:dealId>/<InventorySourcesIds>/', DealView.addInventorySources, name='addInventorySources'),
	path('removeInventorySources/<int:dealId>/<InventorySourcesIds>/', DealView.removeInventorySources, name='removeInventorySources'),
	path('addPlacements/<int:dealId>/<PlacementsIds>/', DealView.addPlacements, name='addPlacements'),
	path('removePlacements/<int:dealId>/<PlacementsIds>/', DealView.removePlacements, name='removePlacements'),
]
