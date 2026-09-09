from django.urls import path
from insuranceOnDjango.views import UnderwritingDecisionView

urlpatterns = [
    path('', UnderwritingDecisionView.index, name='index'),
	path('create', UnderwritingDecisionView.get, name='create'),
	path('get/<int:underwritingDecisionId>/', UnderwritingDecisionView.get, name='get'),
	path('save', UnderwritingDecisionView.save, name='save'),
	path('getAll', UnderwritingDecisionView.getAll, name='getAll'),
	path('delete/<int:underwritingDecisionId>/', UnderwritingDecisionView.delete, name='delete'),
	path('assignQuote/<int:underwritingDecisionId>/<int:QuoteId>/', UnderwritingDecisionView.assignQuote, name='assignQuote'),
	path('unassignQuote/<int:underwritingDecisionId>/', UnderwritingDecisionView.unassignQuote, name='unassignQuote'),
	path('assignUnderwriter/<int:underwritingDecisionId>/<int:UnderwriterId>/', UnderwritingDecisionView.assignUnderwriter, name='assignUnderwriter'),
	path('unassignUnderwriter/<int:underwritingDecisionId>/', UnderwritingDecisionView.unassignUnderwriter, name='unassignUnderwriter'),
]
