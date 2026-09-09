from django.urls import path
from advertisingOnDjango.views import ContentCategoryView

urlpatterns = [
    path('', ContentCategoryView.index, name='index'),
	path('create', ContentCategoryView.get, name='create'),
	path('get/<int:contentCategoryId>/', ContentCategoryView.get, name='get'),
	path('save', ContentCategoryView.save, name='save'),
	path('getAll', ContentCategoryView.getAll, name='getAll'),
	path('delete/<int:contentCategoryId>/', ContentCategoryView.delete, name='delete'),
]
