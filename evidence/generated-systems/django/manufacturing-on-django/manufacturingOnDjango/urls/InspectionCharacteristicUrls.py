from django.urls import path
from manufacturingOnDjango.views import InspectionCharacteristicView

urlpatterns = [
    path('', InspectionCharacteristicView.index, name='index'),
	path('create', InspectionCharacteristicView.get, name='create'),
	path('get/<int:inspectionCharacteristicId>/', InspectionCharacteristicView.get, name='get'),
	path('save', InspectionCharacteristicView.save, name='save'),
	path('getAll', InspectionCharacteristicView.getAll, name='getAll'),
	path('delete/<int:inspectionCharacteristicId>/', InspectionCharacteristicView.delete, name='delete'),
	path('assignInspectionPlan/<int:inspectionCharacteristicId>/<int:InspectionPlanId>/', InspectionCharacteristicView.assignInspectionPlan, name='assignInspectionPlan'),
	path('unassignInspectionPlan/<int:inspectionCharacteristicId>/', InspectionCharacteristicView.unassignInspectionPlan, name='unassignInspectionPlan'),
]
