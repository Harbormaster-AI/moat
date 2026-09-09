from django.urls import path
from hrOnDjango.views import BackgroundCheckView

urlpatterns = [
    path('', BackgroundCheckView.index, name='index'),
	path('create', BackgroundCheckView.get, name='create'),
	path('get/<int:backgroundCheckId>/', BackgroundCheckView.get, name='get'),
	path('save', BackgroundCheckView.save, name='save'),
	path('getAll', BackgroundCheckView.getAll, name='getAll'),
	path('delete/<int:backgroundCheckId>/', BackgroundCheckView.delete, name='delete'),
	path('assignCandidate/<int:backgroundCheckId>/<int:CandidateId>/', BackgroundCheckView.assignCandidate, name='assignCandidate'),
	path('unassignCandidate/<int:backgroundCheckId>/', BackgroundCheckView.unassignCandidate, name='unassignCandidate'),
	path('assignRequisition/<int:backgroundCheckId>/<int:RequisitionId>/', BackgroundCheckView.assignRequisition, name='assignRequisition'),
	path('unassignRequisition/<int:backgroundCheckId>/', BackgroundCheckView.unassignRequisition, name='unassignRequisition'),
	path('assignReport/<int:backgroundCheckId>/<int:ReportId>/', BackgroundCheckView.assignReport, name='assignReport'),
	path('unassignReport/<int:backgroundCheckId>/', BackgroundCheckView.unassignReport, name='unassignReport'),
]
