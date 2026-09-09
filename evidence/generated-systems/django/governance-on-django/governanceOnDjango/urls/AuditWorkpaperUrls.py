from django.urls import path
from governanceOnDjango.views import AuditWorkpaperView

urlpatterns = [
    path('', AuditWorkpaperView.index, name='index'),
	path('create', AuditWorkpaperView.get, name='create'),
	path('get/<int:auditWorkpaperId>/', AuditWorkpaperView.get, name='get'),
	path('save', AuditWorkpaperView.save, name='save'),
	path('getAll', AuditWorkpaperView.getAll, name='getAll'),
	path('delete/<int:auditWorkpaperId>/', AuditWorkpaperView.delete, name='delete'),
	path('assignEngagement/<int:auditWorkpaperId>/<int:EngagementId>/', AuditWorkpaperView.assignEngagement, name='assignEngagement'),
	path('unassignEngagement/<int:auditWorkpaperId>/', AuditWorkpaperView.unassignEngagement, name='unassignEngagement'),
	path('addEvidence/<int:auditWorkpaperId>/<EvidenceIds>/', AuditWorkpaperView.addEvidence, name='addEvidence'),
	path('removeEvidence/<int:auditWorkpaperId>/<EvidenceIds>/', AuditWorkpaperView.removeEvidence, name='removeEvidence'),
	path('addFindings/<int:auditWorkpaperId>/<FindingsIds>/', AuditWorkpaperView.addFindings, name='addFindings'),
	path('removeFindings/<int:auditWorkpaperId>/<FindingsIds>/', AuditWorkpaperView.removeFindings, name='removeFindings'),
]
