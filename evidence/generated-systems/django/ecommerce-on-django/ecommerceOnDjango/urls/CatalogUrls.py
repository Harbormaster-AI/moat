from django.urls import path
from ecommerceOnDjango.views import CatalogView

urlpatterns = [
    path('', CatalogView.index, name='index'),
	path('create', CatalogView.get, name='create'),
	path('get/<int:catalogId>/', CatalogView.get, name='get'),
	path('save', CatalogView.save, name='save'),
	path('getAll', CatalogView.getAll, name='getAll'),
	path('delete/<int:catalogId>/', CatalogView.delete, name='delete'),
	path('assignChannel/<int:catalogId>/<int:ChannelId>/', CatalogView.assignChannel, name='assignChannel'),
	path('unassignChannel/<int:catalogId>/', CatalogView.unassignChannel, name='unassignChannel'),
	path('addCategories/<int:catalogId>/<CategoriesIds>/', CatalogView.addCategories, name='addCategories'),
	path('removeCategories/<int:catalogId>/<CategoriesIds>/', CatalogView.removeCategories, name='removeCategories'),
]
