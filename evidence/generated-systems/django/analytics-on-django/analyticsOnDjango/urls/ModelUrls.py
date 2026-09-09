from django.urls import path
from analyticsOnDjango.views import ModelView

urlpatterns = [
    path('', ModelView.index, name='index'),
	path('create', ModelView.get, name='create'),
	path('get/<int:modelId>/', ModelView.get, name='get'),
	path('save', ModelView.save, name='save'),
	path('getAll', ModelView.getAll, name='getAll'),
	path('delete/<int:modelId>/', ModelView.delete, name='delete'),
	path('assignWorkspace/<int:modelId>/<int:WorkspaceId>/', ModelView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:modelId>/', ModelView.unassignWorkspace, name='unassignWorkspace'),
	path('addVersions/<int:modelId>/<VersionsIds>/', ModelView.addVersions, name='addVersions'),
	path('removeVersions/<int:modelId>/<VersionsIds>/', ModelView.removeVersions, name='removeVersions'),
	path('addFeatureSets/<int:modelId>/<FeatureSetsIds>/', ModelView.addFeatureSets, name='addFeatureSets'),
	path('removeFeatureSets/<int:modelId>/<FeatureSetsIds>/', ModelView.removeFeatureSets, name='removeFeatureSets'),
	path('addExperiments/<int:modelId>/<ExperimentsIds>/', ModelView.addExperiments, name='addExperiments'),
	path('removeExperiments/<int:modelId>/<ExperimentsIds>/', ModelView.removeExperiments, name='removeExperiments'),
	path('addTags/<int:modelId>/<TagsIds>/', ModelView.addTags, name='addTags'),
	path('removeTags/<int:modelId>/<TagsIds>/', ModelView.removeTags, name='removeTags'),
]
