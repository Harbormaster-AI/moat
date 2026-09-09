from django.urls import path
from hrOnDjango.views import CompetencyView

urlpatterns = [
    path('', CompetencyView.index, name='index'),
	path('create', CompetencyView.get, name='create'),
	path('get/<int:competencyId>/', CompetencyView.get, name='get'),
	path('save', CompetencyView.save, name='save'),
	path('getAll', CompetencyView.getAll, name='getAll'),
	path('delete/<int:competencyId>/', CompetencyView.delete, name='delete'),
	path('addJobProfiles/<int:competencyId>/<JobProfilesIds>/', CompetencyView.addJobProfiles, name='addJobProfiles'),
	path('removeJobProfiles/<int:competencyId>/<JobProfilesIds>/', CompetencyView.removeJobProfiles, name='removeJobProfiles'),
	path('addCompetencyRatings/<int:competencyId>/<CompetencyRatingsIds>/', CompetencyView.addCompetencyRatings, name='addCompetencyRatings'),
	path('removeCompetencyRatings/<int:competencyId>/<CompetencyRatingsIds>/', CompetencyView.removeCompetencyRatings, name='removeCompetencyRatings'),
]
