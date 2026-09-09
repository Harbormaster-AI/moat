from django.urls import path
from advertisingOnDjango.views import DeviceCriterionView

urlpatterns = [
    path('', DeviceCriterionView.index, name='index'),
	path('create', DeviceCriterionView.get, name='create'),
	path('get/<int:deviceCriterionId>/', DeviceCriterionView.get, name='get'),
	path('save', DeviceCriterionView.save, name='save'),
	path('getAll', DeviceCriterionView.getAll, name='getAll'),
	path('delete/<int:deviceCriterionId>/', DeviceCriterionView.delete, name='delete'),
	path('assignTargetingProfile/<int:deviceCriterionId>/<int:TargetingProfileId>/', DeviceCriterionView.assignTargetingProfile, name='assignTargetingProfile'),
	path('unassignTargetingProfile/<int:deviceCriterionId>/', DeviceCriterionView.unassignTargetingProfile, name='unassignTargetingProfile'),
]
