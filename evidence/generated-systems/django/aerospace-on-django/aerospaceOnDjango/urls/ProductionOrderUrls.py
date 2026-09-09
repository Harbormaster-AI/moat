from django.urls import path
from aerospaceOnDjango.views import ProductionOrderView

urlpatterns = [
    path('', ProductionOrderView.index, name='index'),
	path('create', ProductionOrderView.get, name='create'),
	path('get/<int:productionOrderId>/', ProductionOrderView.get, name='get'),
	path('save', ProductionOrderView.save, name='save'),
	path('getAll', ProductionOrderView.getAll, name='getAll'),
	path('delete/<int:productionOrderId>/', ProductionOrderView.delete, name='delete'),
	path('assignVariant/<int:productionOrderId>/<int:VariantId>/', ProductionOrderView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:productionOrderId>/', ProductionOrderView.unassignVariant, name='unassignVariant'),
	path('assignPlant/<int:productionOrderId>/<int:PlantId>/', ProductionOrderView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:productionOrderId>/', ProductionOrderView.unassignPlant, name='unassignPlant'),
	path('assignAircraftOrder/<int:productionOrderId>/<int:AircraftOrderId>/', ProductionOrderView.assignAircraftOrder, name='assignAircraftOrder'),
	path('unassignAircraftOrder/<int:productionOrderId>/', ProductionOrderView.unassignAircraftOrder, name='unassignAircraftOrder'),
]
