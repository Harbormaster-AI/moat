from django.urls import path
from healthcareOnDjango.views import ImagingReportView

urlpatterns = [
    path('', ImagingReportView.index, name='index'),
	path('create', ImagingReportView.get, name='create'),
	path('get/<int:imagingReportId>/', ImagingReportView.get, name='get'),
	path('save', ImagingReportView.save, name='save'),
	path('getAll', ImagingReportView.getAll, name='getAll'),
	path('delete/<int:imagingReportId>/', ImagingReportView.delete, name='delete'),
	path('assignImagingOrder/<int:imagingReportId>/<int:ImagingOrderId>/', ImagingReportView.assignImagingOrder, name='assignImagingOrder'),
	path('unassignImagingOrder/<int:imagingReportId>/', ImagingReportView.unassignImagingOrder, name='unassignImagingOrder'),
	path('assignClinician/<int:imagingReportId>/<int:ClinicianId>/', ImagingReportView.assignClinician, name='assignClinician'),
	path('unassignClinician/<int:imagingReportId>/', ImagingReportView.unassignClinician, name='unassignClinician'),
	path('assignEncounter/<int:imagingReportId>/<int:EncounterId>/', ImagingReportView.assignEncounter, name='assignEncounter'),
	path('unassignEncounter/<int:imagingReportId>/', ImagingReportView.unassignEncounter, name='unassignEncounter'),
	path('assignImagingCenter/<int:imagingReportId>/<int:ImagingCenterId>/', ImagingReportView.assignImagingCenter, name='assignImagingCenter'),
	path('unassignImagingCenter/<int:imagingReportId>/', ImagingReportView.unassignImagingCenter, name='unassignImagingCenter'),
]
