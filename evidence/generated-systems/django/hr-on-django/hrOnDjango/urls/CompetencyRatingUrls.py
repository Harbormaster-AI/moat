from django.urls import path
from hrOnDjango.views import CompetencyRatingView

urlpatterns = [
    path('', CompetencyRatingView.index, name='index'),
	path('create', CompetencyRatingView.get, name='create'),
	path('get/<int:competencyRatingId>/', CompetencyRatingView.get, name='get'),
	path('save', CompetencyRatingView.save, name='save'),
	path('getAll', CompetencyRatingView.getAll, name='getAll'),
	path('delete/<int:competencyRatingId>/', CompetencyRatingView.delete, name='delete'),
	path('assignReview/<int:competencyRatingId>/<int:ReviewId>/', CompetencyRatingView.assignReview, name='assignReview'),
	path('unassignReview/<int:competencyRatingId>/', CompetencyRatingView.unassignReview, name='unassignReview'),
	path('assignCompetency/<int:competencyRatingId>/<int:CompetencyId>/', CompetencyRatingView.assignCompetency, name='assignCompetency'),
	path('unassignCompetency/<int:competencyRatingId>/', CompetencyRatingView.unassignCompetency, name='unassignCompetency'),
]
