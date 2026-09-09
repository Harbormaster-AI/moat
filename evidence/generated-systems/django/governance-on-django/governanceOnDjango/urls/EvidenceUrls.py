from django.urls import path
from governanceOnDjango.views import EvidenceView

urlpatterns = [
    path('', EvidenceView.index, name='index'),
	path('create', EvidenceView.get, name='create'),
	path('get/<int:evidenceId>/', EvidenceView.get, name='get'),
	path('save', EvidenceView.save, name='save'),
	path('getAll', EvidenceView.getAll, name='getAll'),
	path('delete/<int:evidenceId>/', EvidenceView.delete, name='delete'),
	path('assignControlTest/<int:evidenceId>/<int:ControlTestId>/', EvidenceView.assignControlTest, name='assignControlTest'),
	path('unassignControlTest/<int:evidenceId>/', EvidenceView.unassignControlTest, name='unassignControlTest'),
	path('assignControl/<int:evidenceId>/<int:ControlId>/', EvidenceView.assignControl, name='assignControl'),
	path('unassignControl/<int:evidenceId>/', EvidenceView.unassignControl, name='unassignControl'),
	path('assignObligation/<int:evidenceId>/<int:ObligationId>/', EvidenceView.assignObligation, name='assignObligation'),
	path('unassignObligation/<int:evidenceId>/', EvidenceView.unassignObligation, name='unassignObligation'),
	path('assignWorkpaper/<int:evidenceId>/<int:WorkpaperId>/', EvidenceView.assignWorkpaper, name='assignWorkpaper'),
	path('unassignWorkpaper/<int:evidenceId>/', EvidenceView.unassignWorkpaper, name='unassignWorkpaper'),
]
