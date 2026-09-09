from django.urls import path
from inventoryOnDjango.views import DemandSignalView

urlpatterns = [
    path('', DemandSignalView.index, name='index'),
	path('create', DemandSignalView.get, name='create'),
	path('get/<int:demandSignalId>/', DemandSignalView.get, name='get'),
	path('save', DemandSignalView.save, name='save'),
	path('getAll', DemandSignalView.getAll, name='getAll'),
	path('delete/<int:demandSignalId>/', DemandSignalView.delete, name='delete'),
	path('assignSku/<int:demandSignalId>/<int:SkuId>/', DemandSignalView.assignSku, name='assignSku'),
	path('unassignSku/<int:demandSignalId>/', DemandSignalView.unassignSku, name='unassignSku'),
	path('addReservations/<int:demandSignalId>/<ReservationsIds>/', DemandSignalView.addReservations, name='addReservations'),
	path('removeReservations/<int:demandSignalId>/<ReservationsIds>/', DemandSignalView.removeReservations, name='removeReservations'),
]
