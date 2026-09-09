from django.urls import path
from advertisingOnDjango.views import CreativeFileView

urlpatterns = [
    path('', CreativeFileView.index, name='index'),
	path('create', CreativeFileView.get, name='create'),
	path('get/<int:creativeFileId>/', CreativeFileView.get, name='get'),
	path('save', CreativeFileView.save, name='save'),
	path('getAll', CreativeFileView.getAll, name='getAll'),
	path('delete/<int:creativeFileId>/', CreativeFileView.delete, name='delete'),
	path('assignCreativeAsset/<int:creativeFileId>/<int:CreativeAssetId>/', CreativeFileView.assignCreativeAsset, name='assignCreativeAsset'),
	path('unassignCreativeAsset/<int:creativeFileId>/', CreativeFileView.unassignCreativeAsset, name='unassignCreativeAsset'),
]
