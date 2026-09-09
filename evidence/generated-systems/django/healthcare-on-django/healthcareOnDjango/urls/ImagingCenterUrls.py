from django.urls import path
from healthcareOnDjango.views import ImagingCenterView

urlpatterns = [
    path('', ImagingCenterView.index, name='index'),
	path('create', ImagingCenterView.get, name='create'),
	path('get/<int:imagingCenterId>/', ImagingCenterView.get, name='get'),
	path('save', ImagingCenterView.save, name='save'),
	path('getAll', ImagingCenterView.getAll, name='getAll'),
	path('delete/<int:imagingCenterId>/', ImagingCenterView.delete, name='delete'),
	path('assignFacility/<int:imagingCenterId>/<int:FacilityId>/', ImagingCenterView.assignFacility, name='assignFacility'),
	path('unassignFacility/<int:imagingCenterId>/', ImagingCenterView.unassignFacility, name='unassignFacility'),
	path('addImagingOrders/<int:imagingCenterId>/<ImagingOrdersIds>/', ImagingCenterView.addImagingOrders, name='addImagingOrders'),
	path('removeImagingOrders/<int:imagingCenterId>/<ImagingOrdersIds>/', ImagingCenterView.removeImagingOrders, name='removeImagingOrders'),
	path('addImagingReports/<int:imagingCenterId>/<ImagingReportsIds>/', ImagingCenterView.addImagingReports, name='addImagingReports'),
	path('removeImagingReports/<int:imagingCenterId>/<ImagingReportsIds>/', ImagingCenterView.removeImagingReports, name='removeImagingReports'),
]
