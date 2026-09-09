from django.urls import path
from inventoryOnDjango.views import CycleCountView

urlpatterns = [
    path('', CycleCountView.index, name='index'),
	path('create', CycleCountView.get, name='create'),
	path('get/<int:cycleCountId>/', CycleCountView.get, name='get'),
	path('save', CycleCountView.save, name='save'),
	path('getAll', CycleCountView.getAll, name='getAll'),
	path('delete/<int:cycleCountId>/', CycleCountView.delete, name='delete'),
	path('assignWarehouse/<int:cycleCountId>/<int:WarehouseId>/', CycleCountView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:cycleCountId>/', CycleCountView.unassignWarehouse, name='unassignWarehouse'),
	path('addLocations/<int:cycleCountId>/<LocationsIds>/', CycleCountView.addLocations, name='addLocations'),
	path('removeLocations/<int:cycleCountId>/<LocationsIds>/', CycleCountView.removeLocations, name='removeLocations'),
	path('addEntries/<int:cycleCountId>/<EntriesIds>/', CycleCountView.addEntries, name='addEntries'),
	path('removeEntries/<int:cycleCountId>/<EntriesIds>/', CycleCountView.removeEntries, name='removeEntries'),
	path('addTransactions/<int:cycleCountId>/<TransactionsIds>/', CycleCountView.addTransactions, name='addTransactions'),
	path('removeTransactions/<int:cycleCountId>/<TransactionsIds>/', CycleCountView.removeTransactions, name='removeTransactions'),
]
