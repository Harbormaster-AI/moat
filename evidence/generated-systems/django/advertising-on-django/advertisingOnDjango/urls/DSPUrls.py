from django.urls import path
from advertisingOnDjango.views import DSPView

urlpatterns = [
    path('', DSPView.index, name='index'),
	path('create', DSPView.get, name='create'),
	path('get/<int:dSPId>/', DSPView.get, name='get'),
	path('save', DSPView.save, name='save'),
	path('getAll', DSPView.getAll, name='getAll'),
	path('delete/<int:dSPId>/', DSPView.delete, name='delete'),
	path('addAdAccounts/<int:dSPId>/<AdAccountsIds>/', DSPView.addAdAccounts, name='addAdAccounts'),
	path('removeAdAccounts/<int:dSPId>/<AdAccountsIds>/', DSPView.removeAdAccounts, name='removeAdAccounts'),
]
