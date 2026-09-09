from django.urls import path
from crmOnDjango.views import OpportunityStageHistoryView

urlpatterns = [
    path('', OpportunityStageHistoryView.index, name='index'),
	path('create', OpportunityStageHistoryView.get, name='create'),
	path('get/<int:opportunityStageHistoryId>/', OpportunityStageHistoryView.get, name='get'),
	path('save', OpportunityStageHistoryView.save, name='save'),
	path('getAll', OpportunityStageHistoryView.getAll, name='getAll'),
	path('delete/<int:opportunityStageHistoryId>/', OpportunityStageHistoryView.delete, name='delete'),
	path('assignOpportunity/<int:opportunityStageHistoryId>/<int:OpportunityId>/', OpportunityStageHistoryView.assignOpportunity, name='assignOpportunity'),
	path('unassignOpportunity/<int:opportunityStageHistoryId>/', OpportunityStageHistoryView.unassignOpportunity, name='unassignOpportunity'),
	path('assignChangedBy/<int:opportunityStageHistoryId>/<int:ChangedById>/', OpportunityStageHistoryView.assignChangedBy, name='assignChangedBy'),
	path('unassignChangedBy/<int:opportunityStageHistoryId>/', OpportunityStageHistoryView.unassignChangedBy, name='unassignChangedBy'),
]
