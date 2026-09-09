from django.urls import path
from governanceOnDjango.views import AuditEngagementView

urlpatterns = [
    path('', AuditEngagementView.index, name='index'),
	path('create', AuditEngagementView.get, name='create'),
	path('get/<int:auditEngagementId>/', AuditEngagementView.get, name='get'),
	path('save', AuditEngagementView.save, name='save'),
	path('getAll', AuditEngagementView.getAll, name='getAll'),
	path('delete/<int:auditEngagementId>/', AuditEngagementView.delete, name='delete'),
	path('assignAuditProgram/<int:auditEngagementId>/<int:AuditProgramId>/', AuditEngagementView.assignAuditProgram, name='assignAuditProgram'),
	path('unassignAuditProgram/<int:auditEngagementId>/', AuditEngagementView.unassignAuditProgram, name='unassignAuditProgram'),
	path('addBusinessUnits/<int:auditEngagementId>/<BusinessUnitsIds>/', AuditEngagementView.addBusinessUnits, name='addBusinessUnits'),
	path('removeBusinessUnits/<int:auditEngagementId>/<BusinessUnitsIds>/', AuditEngagementView.removeBusinessUnits, name='removeBusinessUnits'),
	path('addControlTests/<int:auditEngagementId>/<ControlTestsIds>/', AuditEngagementView.addControlTests, name='addControlTests'),
	path('removeControlTests/<int:auditEngagementId>/<ControlTestsIds>/', AuditEngagementView.removeControlTests, name='removeControlTests'),
	path('addWorkpapers/<int:auditEngagementId>/<WorkpapersIds>/', AuditEngagementView.addWorkpapers, name='addWorkpapers'),
	path('removeWorkpapers/<int:auditEngagementId>/<WorkpapersIds>/', AuditEngagementView.removeWorkpapers, name='removeWorkpapers'),
	path('addFindings/<int:auditEngagementId>/<FindingsIds>/', AuditEngagementView.addFindings, name='addFindings'),
	path('removeFindings/<int:auditEngagementId>/<FindingsIds>/', AuditEngagementView.removeFindings, name='removeFindings'),
]
