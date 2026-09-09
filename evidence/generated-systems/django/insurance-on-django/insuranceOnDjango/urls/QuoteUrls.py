from django.urls import path
from insuranceOnDjango.views import QuoteView

urlpatterns = [
    path('', QuoteView.index, name='index'),
	path('create', QuoteView.get, name='create'),
	path('get/<int:quoteId>/', QuoteView.get, name='get'),
	path('save', QuoteView.save, name='save'),
	path('getAll', QuoteView.getAll, name='getAll'),
	path('delete/<int:quoteId>/', QuoteView.delete, name='delete'),
	path('assignApplication/<int:quoteId>/<int:ApplicationId>/', QuoteView.assignApplication, name='assignApplication'),
	path('unassignApplication/<int:quoteId>/', QuoteView.unassignApplication, name='unassignApplication'),
	path('assignPolicy/<int:quoteId>/<int:PolicyId>/', QuoteView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:quoteId>/', QuoteView.unassignPolicy, name='unassignPolicy'),
	path('addUnderwritingDecisions/<int:quoteId>/<UnderwritingDecisionsIds>/', QuoteView.addUnderwritingDecisions, name='addUnderwritingDecisions'),
	path('removeUnderwritingDecisions/<int:quoteId>/<UnderwritingDecisionsIds>/', QuoteView.removeUnderwritingDecisions, name='removeUnderwritingDecisions'),
]
