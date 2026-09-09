from django.urls import path
from advertisingOnDjango.views import GeoRegionView

urlpatterns = [
    path('', GeoRegionView.index, name='index'),
	path('create', GeoRegionView.get, name='create'),
	path('get/<int:geoRegionId>/', GeoRegionView.get, name='get'),
	path('save', GeoRegionView.save, name='save'),
	path('getAll', GeoRegionView.getAll, name='getAll'),
	path('delete/<int:geoRegionId>/', GeoRegionView.delete, name='delete'),
	path('assignParent/<int:geoRegionId>/<int:ParentId>/', GeoRegionView.assignParent, name='assignParent'),
	path('unassignParent/<int:geoRegionId>/', GeoRegionView.unassignParent, name='unassignParent'),
	path('addChildren/<int:geoRegionId>/<ChildrenIds>/', GeoRegionView.addChildren, name='addChildren'),
	path('removeChildren/<int:geoRegionId>/<ChildrenIds>/', GeoRegionView.removeChildren, name='removeChildren'),
]
