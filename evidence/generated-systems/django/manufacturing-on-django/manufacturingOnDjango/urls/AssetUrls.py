from django.urls import path
from manufacturingOnDjango.views import AssetView

urlpatterns = [
    path('', AssetView.index, name='index'),
	path('create', AssetView.get, name='create'),
	path('get/<int:assetId>/', AssetView.get, name='get'),
	path('save', AssetView.save, name='save'),
	path('getAll', AssetView.getAll, name='getAll'),
	path('delete/<int:assetId>/', AssetView.delete, name='delete'),
	path('assignPlant/<int:assetId>/<int:PlantId>/', AssetView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:assetId>/', AssetView.unassignPlant, name='unassignPlant'),
	path('assignWorkCenter/<int:assetId>/<int:WorkCenterId>/', AssetView.assignWorkCenter, name='assignWorkCenter'),
	path('unassignWorkCenter/<int:assetId>/', AssetView.unassignWorkCenter, name='unassignWorkCenter'),
	path('addMaintenanceOrders/<int:assetId>/<MaintenanceOrdersIds>/', AssetView.addMaintenanceOrders, name='addMaintenanceOrders'),
	path('removeMaintenanceOrders/<int:assetId>/<MaintenanceOrdersIds>/', AssetView.removeMaintenanceOrders, name='removeMaintenanceOrders'),
	path('addMaintenancePlans/<int:assetId>/<MaintenancePlansIds>/', AssetView.addMaintenancePlans, name='addMaintenancePlans'),
	path('removeMaintenancePlans/<int:assetId>/<MaintenancePlansIds>/', AssetView.removeMaintenancePlans, name='removeMaintenancePlans'),
]
