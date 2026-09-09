from django.urls import path
from healthcareOnDjango.views import ImagingOrderView

urlpatterns = [
    path('', ImagingOrderView.index, name='index'),
	path('create', ImagingOrderView.get, name='create'),
	path('get/<int:imagingOrderId>/', ImagingOrderView.get, name='get'),
	path('save', ImagingOrderView.save, name='save'),
	path('getAll', ImagingOrderView.getAll, name='getAll'),
	path('delete/<int:imagingOrderId>/', ImagingOrderView.delete, name='delete'),
	path('assignOrder/<int:imagingOrderId>/<int:OrderId>/', ImagingOrderView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:imagingOrderId>/', ImagingOrderView.unassignOrder, name='unassignOrder'),
	path('assignImagingCenter/<int:imagingOrderId>/<int:ImagingCenterId>/', ImagingOrderView.assignImagingCenter, name='assignImagingCenter'),
	path('unassignImagingCenter/<int:imagingOrderId>/', ImagingOrderView.unassignImagingCenter, name='unassignImagingCenter'),
	path('addReports/<int:imagingOrderId>/<ReportsIds>/', ImagingOrderView.addReports, name='addReports'),
	path('removeReports/<int:imagingOrderId>/<ReportsIds>/', ImagingOrderView.removeReports, name='removeReports'),
]
