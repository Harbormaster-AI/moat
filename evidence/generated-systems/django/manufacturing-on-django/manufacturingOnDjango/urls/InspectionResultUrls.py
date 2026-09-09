from django.urls import path
from manufacturingOnDjango.views import InspectionResultView

urlpatterns = [
    path('', InspectionResultView.index, name='index'),
	path('create', InspectionResultView.get, name='create'),
	path('get/<int:inspectionResultId>/', InspectionResultView.get, name='get'),
	path('save', InspectionResultView.save, name='save'),
	path('getAll', InspectionResultView.getAll, name='getAll'),
	path('delete/<int:inspectionResultId>/', InspectionResultView.delete, name='delete'),
	path('assignInspectionLot/<int:inspectionResultId>/<int:InspectionLotId>/', InspectionResultView.assignInspectionLot, name='assignInspectionLot'),
	path('unassignInspectionLot/<int:inspectionResultId>/', InspectionResultView.unassignInspectionLot, name='unassignInspectionLot'),
	path('assignCharacteristic/<int:inspectionResultId>/<int:CharacteristicId>/', InspectionResultView.assignCharacteristic, name='assignCharacteristic'),
	path('unassignCharacteristic/<int:inspectionResultId>/', InspectionResultView.unassignCharacteristic, name='unassignCharacteristic'),
]
