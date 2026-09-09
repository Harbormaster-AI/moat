from django.urls import path
from aerospaceOnDjango.views import QuoteView

urlpatterns = [
    path('', QuoteView.index, name='index'),
	path('create', QuoteView.get, name='create'),
	path('get/<int:quoteId>/', QuoteView.get, name='get'),
	path('save', QuoteView.save, name='save'),
	path('getAll', QuoteView.getAll, name='getAll'),
	path('delete/<int:quoteId>/', QuoteView.delete, name='delete'),
	path('assignAircraftOrder/<int:quoteId>/<int:AircraftOrderId>/', QuoteView.assignAircraftOrder, name='assignAircraftOrder'),
	path('unassignAircraftOrder/<int:quoteId>/', QuoteView.unassignAircraftOrder, name='unassignAircraftOrder'),
]
