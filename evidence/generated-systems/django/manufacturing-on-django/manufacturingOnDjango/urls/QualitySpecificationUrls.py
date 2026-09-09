from django.urls import path
from manufacturingOnDjango.views import QualitySpecificationView

urlpatterns = [
    path('', QualitySpecificationView.index, name='index'),
	path('create', QualitySpecificationView.get, name='create'),
	path('get/<int:qualitySpecificationId>/', QualitySpecificationView.get, name='get'),
	path('save', QualitySpecificationView.save, name='save'),
	path('getAll', QualitySpecificationView.getAll, name='getAll'),
	path('delete/<int:qualitySpecificationId>/', QualitySpecificationView.delete, name='delete'),
	path('assignItem/<int:qualitySpecificationId>/<int:ItemId>/', QualitySpecificationView.assignItem, name='assignItem'),
	path('unassignItem/<int:qualitySpecificationId>/', QualitySpecificationView.unassignItem, name='unassignItem'),
]
