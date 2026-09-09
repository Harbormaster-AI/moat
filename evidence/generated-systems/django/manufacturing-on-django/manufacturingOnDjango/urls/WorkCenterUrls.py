from django.urls import path
from manufacturingOnDjango.views import WorkCenterView

urlpatterns = [
    path('', WorkCenterView.index, name='index'),
	path('create', WorkCenterView.get, name='create'),
	path('get/<int:workCenterId>/', WorkCenterView.get, name='get'),
	path('save', WorkCenterView.save, name='save'),
	path('getAll', WorkCenterView.getAll, name='getAll'),
	path('delete/<int:workCenterId>/', WorkCenterView.delete, name='delete'),
	path('assignProductionLine/<int:workCenterId>/<int:ProductionLineId>/', WorkCenterView.assignProductionLine, name='assignProductionLine'),
	path('unassignProductionLine/<int:workCenterId>/', WorkCenterView.unassignProductionLine, name='unassignProductionLine'),
	path('addAssets/<int:workCenterId>/<AssetsIds>/', WorkCenterView.addAssets, name='addAssets'),
	path('removeAssets/<int:workCenterId>/<AssetsIds>/', WorkCenterView.removeAssets, name='removeAssets'),
	path('addMaintenanceOrders/<int:workCenterId>/<MaintenanceOrdersIds>/', WorkCenterView.addMaintenanceOrders, name='addMaintenanceOrders'),
	path('removeMaintenanceOrders/<int:workCenterId>/<MaintenanceOrdersIds>/', WorkCenterView.removeMaintenanceOrders, name='removeMaintenanceOrders'),
]
