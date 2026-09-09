from django.urls import path
from governanceOnDjango.views import RiskView

urlpatterns = [
    path('', RiskView.index, name='index'),
	path('create', RiskView.get, name='create'),
	path('get/<int:riskId>/', RiskView.get, name='get'),
	path('save', RiskView.save, name='save'),
	path('getAll', RiskView.getAll, name='getAll'),
	path('delete/<int:riskId>/', RiskView.delete, name='delete'),
	path('assignOrganization/<int:riskId>/<int:OrganizationId>/', RiskView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:riskId>/', RiskView.unassignOrganization, name='unassignOrganization'),
	path('addControls/<int:riskId>/<ControlsIds>/', RiskView.addControls, name='addControls'),
	path('removeControls/<int:riskId>/<ControlsIds>/', RiskView.removeControls, name='removeControls'),
	path('addAssessments/<int:riskId>/<AssessmentsIds>/', RiskView.addAssessments, name='addAssessments'),
	path('removeAssessments/<int:riskId>/<AssessmentsIds>/', RiskView.removeAssessments, name='removeAssessments'),
	path('addIssues/<int:riskId>/<IssuesIds>/', RiskView.addIssues, name='addIssues'),
	path('removeIssues/<int:riskId>/<IssuesIds>/', RiskView.removeIssues, name='removeIssues'),
	path('addFindings/<int:riskId>/<FindingsIds>/', RiskView.addFindings, name='addFindings'),
	path('removeFindings/<int:riskId>/<FindingsIds>/', RiskView.removeFindings, name='removeFindings'),
]
