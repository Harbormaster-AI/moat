from django.urls import path
from fintechOnDjango.views import APIClientView

urlpatterns = [
    path('', APIClientView.index, name='index'),
	path('create', APIClientView.get, name='create'),
	path('get/<int:aPIClientId>/', APIClientView.get, name='get'),
	path('save', APIClientView.save, name='save'),
	path('getAll', APIClientView.getAll, name='getAll'),
	path('delete/<int:aPIClientId>/', APIClientView.delete, name='delete'),
	path('addConsents/<int:aPIClientId>/<ConsentsIds>/', APIClientView.addConsents, name='addConsents'),
	path('removeConsents/<int:aPIClientId>/<ConsentsIds>/', APIClientView.removeConsents, name='removeConsents'),
]
