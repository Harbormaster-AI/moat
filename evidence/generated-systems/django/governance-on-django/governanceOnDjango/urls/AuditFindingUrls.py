from django.urls import path
from governanceOnDjango.views import AuditFindingView

urlpatterns = [
    path('', AuditFindingView.index, name='index'),
	path('create', AuditFindingView.get, name='create'),
	path('get/<int:auditFindingId>/', AuditFindingView.get, name='get'),
	path('save', AuditFindingView.save, name='save'),
	path('getAll', AuditFindingView.getAll, name='getAll'),
	path('delete/<int:auditFindingId>/', AuditFindingView.delete, name='delete'),
	path('assignEngagement/<int:auditFindingId>/<int:EngagementId>/', AuditFindingView.assignEngagement, name='assignEngagement'),
	path('unassignEngagement/<int:auditFindingId>/', AuditFindingView.unassignEngagement, name='unassignEngagement'),
	path('assignWorkpaper/<int:auditFindingId>/<int:WorkpaperId>/', AuditFindingView.assignWorkpaper, name='assignWorkpaper'),
	path('unassignWorkpaper/<int:auditFindingId>/', AuditFindingView.unassignWorkpaper, name='unassignWorkpaper'),
	path('addCorrectiveActions/<int:auditFindingId>/<CorrectiveActionsIds>/', AuditFindingView.addCorrectiveActions, name='addCorrectiveActions'),
	path('removeCorrectiveActions/<int:auditFindingId>/<CorrectiveActionsIds>/', AuditFindingView.removeCorrectiveActions, name='removeCorrectiveActions'),
	path('addRelatedRisks/<int:auditFindingId>/<RelatedRisksIds>/', AuditFindingView.addRelatedRisks, name='addRelatedRisks'),
	path('removeRelatedRisks/<int:auditFindingId>/<RelatedRisksIds>/', AuditFindingView.removeRelatedRisks, name='removeRelatedRisks'),
	path('addRelatedControls/<int:auditFindingId>/<RelatedControlsIds>/', AuditFindingView.addRelatedControls, name='addRelatedControls'),
	path('removeRelatedControls/<int:auditFindingId>/<RelatedControlsIds>/', AuditFindingView.removeRelatedControls, name='removeRelatedControls'),
	path('addIssues/<int:auditFindingId>/<IssuesIds>/', AuditFindingView.addIssues, name='addIssues'),
	path('removeIssues/<int:auditFindingId>/<IssuesIds>/', AuditFindingView.removeIssues, name='removeIssues'),
]
