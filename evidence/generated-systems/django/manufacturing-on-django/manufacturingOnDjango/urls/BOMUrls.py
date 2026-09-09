from django.urls import path
from manufacturingOnDjango.views import BOMView

urlpatterns = [
    path('', BOMView.index, name='index'),
	path('create', BOMView.get, name='create'),
	path('get/<int:bOMId>/', BOMView.get, name='get'),
	path('save', BOMView.save, name='save'),
	path('getAll', BOMView.getAll, name='getAll'),
	path('delete/<int:bOMId>/', BOMView.delete, name='delete'),
	path('assignParentItem/<int:bOMId>/<int:ParentItemId>/', BOMView.assignParentItem, name='assignParentItem'),
	path('unassignParentItem/<int:bOMId>/', BOMView.unassignParentItem, name='unassignParentItem'),
	path('addBomItems/<int:bOMId>/<BomItemsIds>/', BOMView.addBomItems, name='addBomItems'),
	path('removeBomItems/<int:bOMId>/<BomItemsIds>/', BOMView.removeBomItems, name='removeBomItems'),
]
