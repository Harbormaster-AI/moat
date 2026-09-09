from django.urls import path
from fintechOnDjango.views import CardTokenizationView

urlpatterns = [
    path('', CardTokenizationView.index, name='index'),
	path('create', CardTokenizationView.get, name='create'),
	path('get/<int:cardTokenizationId>/', CardTokenizationView.get, name='get'),
	path('save', CardTokenizationView.save, name='save'),
	path('getAll', CardTokenizationView.getAll, name='getAll'),
	path('delete/<int:cardTokenizationId>/', CardTokenizationView.delete, name='delete'),
	path('assignCard/<int:cardTokenizationId>/<int:CardId>/', CardTokenizationView.assignCard, name='assignCard'),
	path('unassignCard/<int:cardTokenizationId>/', CardTokenizationView.unassignCard, name='unassignCard'),
]
