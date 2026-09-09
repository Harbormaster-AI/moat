from django.urls import path
from hrOnDjango.views import JobProfileView

urlpatterns = [
    path('', JobProfileView.index, name='index'),
	path('create', JobProfileView.get, name='create'),
	path('get/<int:jobProfileId>/', JobProfileView.get, name='get'),
	path('save', JobProfileView.save, name='save'),
	path('getAll', JobProfileView.getAll, name='getAll'),
	path('delete/<int:jobProfileId>/', JobProfileView.delete, name='delete'),
	path('assignJobFamily/<int:jobProfileId>/<int:JobFamilyId>/', JobProfileView.assignJobFamily, name='assignJobFamily'),
	path('unassignJobFamily/<int:jobProfileId>/', JobProfileView.unassignJobFamily, name='unassignJobFamily'),
	path('addCompetencies/<int:jobProfileId>/<CompetenciesIds>/', JobProfileView.addCompetencies, name='addCompetencies'),
	path('removeCompetencies/<int:jobProfileId>/<CompetenciesIds>/', JobProfileView.removeCompetencies, name='removeCompetencies'),
	path('addTrainingRecommendations/<int:jobProfileId>/<TrainingRecommendationsIds>/', JobProfileView.addTrainingRecommendations, name='addTrainingRecommendations'),
	path('removeTrainingRecommendations/<int:jobProfileId>/<TrainingRecommendationsIds>/', JobProfileView.removeTrainingRecommendations, name='removeTrainingRecommendations'),
	path('addPositions/<int:jobProfileId>/<PositionsIds>/', JobProfileView.addPositions, name='addPositions'),
	path('removePositions/<int:jobProfileId>/<PositionsIds>/', JobProfileView.removePositions, name='removePositions'),
]
