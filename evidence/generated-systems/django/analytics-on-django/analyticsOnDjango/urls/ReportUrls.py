from django.urls import path
from analyticsOnDjango.views import ReportView

urlpatterns = [
    path('', ReportView.index, name='index'),
	path('create', ReportView.get, name='create'),
	path('get/<int:reportId>/', ReportView.get, name='get'),
	path('save', ReportView.save, name='save'),
	path('getAll', ReportView.getAll, name='getAll'),
	path('delete/<int:reportId>/', ReportView.delete, name='delete'),
	path('assignWorkspace/<int:reportId>/<int:WorkspaceId>/', ReportView.assignWorkspace, name='assignWorkspace'),
	path('unassignWorkspace/<int:reportId>/', ReportView.unassignWorkspace, name='unassignWorkspace'),
	path('addVisualizations/<int:reportId>/<VisualizationsIds>/', ReportView.addVisualizations, name='addVisualizations'),
	path('removeVisualizations/<int:reportId>/<VisualizationsIds>/', ReportView.removeVisualizations, name='removeVisualizations'),
	path('addDatasets/<int:reportId>/<DatasetsIds>/', ReportView.addDatasets, name='addDatasets'),
	path('removeDatasets/<int:reportId>/<DatasetsIds>/', ReportView.removeDatasets, name='removeDatasets'),
	path('addSemanticModels/<int:reportId>/<SemanticModelsIds>/', ReportView.addSemanticModels, name='addSemanticModels'),
	path('removeSemanticModels/<int:reportId>/<SemanticModelsIds>/', ReportView.removeSemanticModels, name='removeSemanticModels'),
	path('addQueries/<int:reportId>/<QueriesIds>/', ReportView.addQueries, name='addQueries'),
	path('removeQueries/<int:reportId>/<QueriesIds>/', ReportView.removeQueries, name='removeQueries'),
	path('addTags/<int:reportId>/<TagsIds>/', ReportView.addTags, name='addTags'),
	path('removeTags/<int:reportId>/<TagsIds>/', ReportView.removeTags, name='removeTags'),
]
