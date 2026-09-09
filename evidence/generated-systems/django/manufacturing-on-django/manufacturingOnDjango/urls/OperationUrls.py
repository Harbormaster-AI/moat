from django.urls import path
from manufacturingOnDjango.views import OperationView

urlpatterns = [
    path('', OperationView.index, name='index'),
	path('create', OperationView.get, name='create'),
	path('get/<int:operationId>/', OperationView.get, name='get'),
	path('save', OperationView.save, name='save'),
	path('getAll', OperationView.getAll, name='getAll'),
	path('delete/<int:operationId>/', OperationView.delete, name='delete'),
	path('assignRouting/<int:operationId>/<int:RoutingId>/', OperationView.assignRouting, name='assignRouting'),
	path('unassignRouting/<int:operationId>/', OperationView.unassignRouting, name='unassignRouting'),
	path('assignWorkCenter/<int:operationId>/<int:WorkCenterId>/', OperationView.assignWorkCenter, name='assignWorkCenter'),
	path('unassignWorkCenter/<int:operationId>/', OperationView.unassignWorkCenter, name='unassignWorkCenter'),
	path('assignInspectionPlan/<int:operationId>/<int:InspectionPlanId>/', OperationView.assignInspectionPlan, name='assignInspectionPlan'),
	path('unassignInspectionPlan/<int:operationId>/', OperationView.unassignInspectionPlan, name='unassignInspectionPlan'),
]
