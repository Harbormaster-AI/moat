from django.urls import path
from advertisingOnDjango.views import CreativeApprovalView

urlpatterns = [
    path('', CreativeApprovalView.index, name='index'),
	path('create', CreativeApprovalView.get, name='create'),
	path('get/<int:creativeApprovalId>/', CreativeApprovalView.get, name='get'),
	path('save', CreativeApprovalView.save, name='save'),
	path('getAll', CreativeApprovalView.getAll, name='getAll'),
	path('delete/<int:creativeApprovalId>/', CreativeApprovalView.delete, name='delete'),
	path('assignCreativeAsset/<int:creativeApprovalId>/<int:CreativeAssetId>/', CreativeApprovalView.assignCreativeAsset, name='assignCreativeAsset'),
	path('unassignCreativeAsset/<int:creativeApprovalId>/', CreativeApprovalView.unassignCreativeAsset, name='unassignCreativeAsset'),
	path('assignPublisher/<int:creativeApprovalId>/<int:PublisherId>/', CreativeApprovalView.assignPublisher, name='assignPublisher'),
	path('unassignPublisher/<int:creativeApprovalId>/', CreativeApprovalView.unassignPublisher, name='unassignPublisher'),
]
