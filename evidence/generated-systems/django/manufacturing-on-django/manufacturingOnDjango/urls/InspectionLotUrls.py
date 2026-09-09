from django.urls import path
from manufacturingOnDjango.views import InspectionLotView

urlpatterns = [
    path('', InspectionLotView.index, name='index'),
	path('create', InspectionLotView.get, name='create'),
	path('get/<int:inspectionLotId>/', InspectionLotView.get, name='get'),
	path('save', InspectionLotView.save, name='save'),
	path('getAll', InspectionLotView.getAll, name='getAll'),
	path('delete/<int:inspectionLotId>/', InspectionLotView.delete, name='delete'),
	path('assignItem/<int:inspectionLotId>/<int:ItemId>/', InspectionLotView.assignItem, name='assignItem'),
	path('unassignItem/<int:inspectionLotId>/', InspectionLotView.unassignItem, name='unassignItem'),
	path('assignWorkOrder/<int:inspectionLotId>/<int:WorkOrderId>/', InspectionLotView.assignWorkOrder, name='assignWorkOrder'),
	path('unassignWorkOrder/<int:inspectionLotId>/', InspectionLotView.unassignWorkOrder, name='unassignWorkOrder'),
	path('assignGoodsReceipt/<int:inspectionLotId>/<int:GoodsReceiptId>/', InspectionLotView.assignGoodsReceipt, name='assignGoodsReceipt'),
	path('unassignGoodsReceipt/<int:inspectionLotId>/', InspectionLotView.unassignGoodsReceipt, name='unassignGoodsReceipt'),
	path('addResults/<int:inspectionLotId>/<ResultsIds>/', InspectionLotView.addResults, name='addResults'),
	path('removeResults/<int:inspectionLotId>/<ResultsIds>/', InspectionLotView.removeResults, name='removeResults'),
]
