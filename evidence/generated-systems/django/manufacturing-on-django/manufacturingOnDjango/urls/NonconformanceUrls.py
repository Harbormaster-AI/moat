from django.urls import path
from manufacturingOnDjango.views import NonconformanceView

urlpatterns = [
    path('', NonconformanceView.index, name='index'),
	path('create', NonconformanceView.get, name='create'),
	path('get/<int:nonconformanceId>/', NonconformanceView.get, name='get'),
	path('save', NonconformanceView.save, name='save'),
	path('getAll', NonconformanceView.getAll, name='getAll'),
	path('delete/<int:nonconformanceId>/', NonconformanceView.delete, name='delete'),
	path('assignItem/<int:nonconformanceId>/<int:ItemId>/', NonconformanceView.assignItem, name='assignItem'),
	path('unassignItem/<int:nonconformanceId>/', NonconformanceView.unassignItem, name='unassignItem'),
	path('assignWorkOrder/<int:nonconformanceId>/<int:WorkOrderId>/', NonconformanceView.assignWorkOrder, name='assignWorkOrder'),
	path('unassignWorkOrder/<int:nonconformanceId>/', NonconformanceView.unassignWorkOrder, name='unassignWorkOrder'),
	path('assignInspectionLot/<int:nonconformanceId>/<int:InspectionLotId>/', NonconformanceView.assignInspectionLot, name='assignInspectionLot'),
	path('unassignInspectionLot/<int:nonconformanceId>/', NonconformanceView.unassignInspectionLot, name='unassignInspectionLot'),
	path('assignCorrectiveAction/<int:nonconformanceId>/<int:CorrectiveActionId>/', NonconformanceView.assignCorrectiveAction, name='assignCorrectiveAction'),
	path('unassignCorrectiveAction/<int:nonconformanceId>/', NonconformanceView.unassignCorrectiveAction, name='unassignCorrectiveAction'),
]
