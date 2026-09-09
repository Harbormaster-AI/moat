from django.urls import path
from manufacturingOnDjango.views import BOMItemView

urlpatterns = [
    path('', BOMItemView.index, name='index'),
	path('create', BOMItemView.get, name='create'),
	path('get/<int:bOMItemId>/', BOMItemView.get, name='get'),
	path('save', BOMItemView.save, name='save'),
	path('getAll', BOMItemView.getAll, name='getAll'),
	path('delete/<int:bOMItemId>/', BOMItemView.delete, name='delete'),
	path('assignBom/<int:bOMItemId>/<int:BomId>/', BOMItemView.assignBom, name='assignBom'),
	path('unassignBom/<int:bOMItemId>/', BOMItemView.unassignBom, name='unassignBom'),
	path('assignComponent/<int:bOMItemId>/<int:ComponentId>/', BOMItemView.assignComponent, name='assignComponent'),
	path('unassignComponent/<int:bOMItemId>/', BOMItemView.unassignComponent, name='unassignComponent'),
]
