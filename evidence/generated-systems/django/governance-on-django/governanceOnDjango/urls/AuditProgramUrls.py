from django.urls import path
from governanceOnDjango.views import AuditProgramView

urlpatterns = [
    path('', AuditProgramView.index, name='index'),
	path('create', AuditProgramView.get, name='create'),
	path('get/<int:auditProgramId>/', AuditProgramView.get, name='get'),
	path('save', AuditProgramView.save, name='save'),
	path('getAll', AuditProgramView.getAll, name='getAll'),
	path('delete/<int:auditProgramId>/', AuditProgramView.delete, name='delete'),
	path('assignOrganization/<int:auditProgramId>/<int:OrganizationId>/', AuditProgramView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:auditProgramId>/', AuditProgramView.unassignOrganization, name='unassignOrganization'),
	path('addEngagements/<int:auditProgramId>/<EngagementsIds>/', AuditProgramView.addEngagements, name='addEngagements'),
	path('removeEngagements/<int:auditProgramId>/<EngagementsIds>/', AuditProgramView.removeEngagements, name='removeEngagements'),
]
