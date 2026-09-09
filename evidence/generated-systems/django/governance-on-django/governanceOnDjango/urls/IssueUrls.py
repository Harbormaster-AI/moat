from django.urls import path
from governanceOnDjango.views import IssueView

urlpatterns = [
    path('', IssueView.index, name='index'),
	path('create', IssueView.get, name='create'),
	path('get/<int:issueId>/', IssueView.get, name='get'),
	path('save', IssueView.save, name='save'),
	path('getAll', IssueView.getAll, name='getAll'),
	path('delete/<int:issueId>/', IssueView.delete, name='delete'),
	path('assignRisk/<int:issueId>/<int:RiskId>/', IssueView.assignRisk, name='assignRisk'),
	path('unassignRisk/<int:issueId>/', IssueView.unassignRisk, name='unassignRisk'),
	path('assignFinding/<int:issueId>/<int:FindingId>/', IssueView.assignFinding, name='assignFinding'),
	path('unassignFinding/<int:issueId>/', IssueView.unassignFinding, name='unassignFinding'),
	path('assignControl/<int:issueId>/<int:ControlId>/', IssueView.assignControl, name='assignControl'),
	path('unassignControl/<int:issueId>/', IssueView.unassignControl, name='unassignControl'),
	path('addCorrectiveActions/<int:issueId>/<CorrectiveActionsIds>/', IssueView.addCorrectiveActions, name='addCorrectiveActions'),
	path('removeCorrectiveActions/<int:issueId>/<CorrectiveActionsIds>/', IssueView.removeCorrectiveActions, name='removeCorrectiveActions'),
]
