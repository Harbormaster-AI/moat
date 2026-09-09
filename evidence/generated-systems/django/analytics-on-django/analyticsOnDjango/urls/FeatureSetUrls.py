from django.urls import path
from analyticsOnDjango.views import FeatureSetView

urlpatterns = [
    path('', FeatureSetView.index, name='index'),
	path('create', FeatureSetView.get, name='create'),
	path('get/<int:featureSetId>/', FeatureSetView.get, name='get'),
	path('save', FeatureSetView.save, name='save'),
	path('getAll', FeatureSetView.getAll, name='getAll'),
	path('delete/<int:featureSetId>/', FeatureSetView.delete, name='delete'),
	path('assignWorkspace/<int:featureSetId>/<int:WorkspaceId>/', FeatureSetView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:featureSetId>/', FeatureSetView.unassignWorkspace, name='unassignWorkspace'),
	path('addFeatures/<int:featureSetId>/<FeaturesIds>/', FeatureSetView.addFeatures, name='addFeatures'),
	path('removeFeatures/<int:featureSetId>/<FeaturesIds>/', FeatureSetView.removeFeatures, name='removeFeatures'),
	path('addDatasets/<int:featureSetId>/<DatasetsIds>/', FeatureSetView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:featureSetId>/<DatasetsIds>/', FeatureSetView.removeDatasets, name='removeDatasets'),
	path('addModels/<int:featureSetId>/<ModelsIds>/', FeatureSetView.addModels, name='addModels'),
	path('removeModels/<int:featureSetId>/<ModelsIds>/', FeatureSetView.removeModels, name='removeModels'),
	path('addModelVersions/<int:featureSetId>/<ModelVersionsIds>/', FeatureSetView.addModelVersions, name='addModelVersions'),
	path('removeModelVersions/<int:featureSetId>/<ModelVersionsIds>/', FeatureSetView.removeModelVersions, name='removeModelVersions'),
	path('addTags/<int:featureSetId>/<TagsIds>/', FeatureSetView.addTags, name='addTags'),
	path('removeTags/<int:featureSetId>/<TagsIds>/', FeatureSetView.removeTags, name='removeTags'),
]
