from django.urls import path
from governanceOnDjango.views import CorrectiveActionView

urlpatterns = [
    path('', CorrectiveActionView.index, name='index'),
	path('create', CorrectiveActionView.get, name='create'),
	path('get/<int:correctiveActionId>/', CorrectiveActionView.get, name='get'),
	path('save', CorrectiveActionView.save, name='save'),
	path('getAll', CorrectiveActionView.getAll, name='getAll'),
	path('delete/<int:correctiveActionId>/', CorrectiveActionView.delete, name='delete'),
	path('assignFinding/<int:correctiveActionId>/<int:FindingId>/', CorrectiveActionView.assignFinding, name='assignFinding'),
	path('unassignFinding/<int:correctiveActionId>/', CorrectiveActionView.unassignFinding, name='unassignFinding'),
	path('assignIssue/<int:correctiveActionId>/<int:IssueId>/', CorrectiveActionView.assignIssue, name='assignIssue'),
	path('unassignIssue/<int:correctiveActionId>/', CorrectiveActionView.unassignIssue, name='unassignIssue'),
]
