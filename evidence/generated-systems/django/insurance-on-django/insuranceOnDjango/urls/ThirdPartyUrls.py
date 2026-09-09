from django.urls import path
from insuranceOnDjango.views import ThirdPartyView

urlpatterns = [
    path('', ThirdPartyView.index, name='index'),
	path('create', ThirdPartyView.get, name='create'),
	path('get/<int:thirdPartyId>/', ThirdPartyView.get, name='get'),
	path('save', ThirdPartyView.save, name='save'),
	path('getAll', ThirdPartyView.getAll, name='getAll'),
	path('delete/<int:thirdPartyId>/', ThirdPartyView.delete, name='delete'),
	path('addSubrogations/<int:thirdPartyId>/<SubrogationsIds>/', ThirdPartyView.addSubrogations, name='addSubrogations'),
	path('removeSubrogations/<int:thirdPartyId>/<SubrogationsIds>/', ThirdPartyView.removeSubrogations, name='removeSubrogations'),
]
