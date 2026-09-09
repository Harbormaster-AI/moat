from django.urls import path
from manufacturingOnDjango.views import MRPRunView

urlpatterns = [
    path('', MRPRunView.index, name='index'),
	path('create', MRPRunView.get, name='create'),
	path('get/<int:mRPRunId>/', MRPRunView.get, name='get'),
	path('save', MRPRunView.save, name='save'),
	path('getAll', MRPRunView.getAll, name='getAll'),
	path('delete/<int:mRPRunId>/', MRPRunView.delete, name='delete'),
	path('assignPlant/<int:mRPRunId>/<int:PlantId>/', MRPRunView.assignPlant, name='assignPlant'),
	path('unassignPlant/<int:mRPRunId>/', MRPRunView.unassignPlant, name='unassignPlant'),
	path('addPlannedOrders/<int:mRPRunId>/<PlannedOrdersIds>/', MRPRunView.addPlannedOrders, name='addPlannedOrders'),
	path('removePlannedOrders/<int:mRPRunId>/<PlannedOrdersIds>/', MRPRunView.removePlannedOrders, name='removePlannedOrders'),
]
