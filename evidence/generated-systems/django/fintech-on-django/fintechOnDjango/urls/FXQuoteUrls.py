from django.urls import path
from fintechOnDjango.views import FXQuoteView

urlpatterns = [
    path('', FXQuoteView.index, name='index'),
	path('create', FXQuoteView.get, name='create'),
	path('get/<int:fXQuoteId>/', FXQuoteView.get, name='get'),
	path('save', FXQuoteView.save, name='save'),
	path('getAll', FXQuoteView.getAll, name='getAll'),
	path('delete/<int:fXQuoteId>/', FXQuoteView.delete, name='delete'),
	path('assignRequestedBy/<int:fXQuoteId>/<int:RequestedById>/', FXQuoteView.assignRequestedBy, name='assignRequestedBy'),
	path('unassignRequestedBy/<int:fXQuoteId>/', FXQuoteView.unassignRequestedBy, name='unassignRequestedBy'),
]
