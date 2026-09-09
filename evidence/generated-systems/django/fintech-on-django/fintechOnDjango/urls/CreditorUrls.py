from django.urls import path
from fintechOnDjango.views import CreditorView

urlpatterns = [
    path('', CreditorView.index, name='index'),
	path('create', CreditorView.get, name='create'),
	path('get/<int:creditorId>/', CreditorView.get, name='get'),
	path('save', CreditorView.save, name='save'),
	path('getAll', CreditorView.getAll, name='getAll'),
	path('delete/<int:creditorId>/', CreditorView.delete, name='delete'),
	path('addMandates/<int:creditorId>/<MandatesIds>/', CreditorView.addMandates, name='addMandates'),
	path('removeMandates/<int:creditorId>/<MandatesIds>/', CreditorView.removeMandates, name='removeMandates'),
]
