from django.urls import path
from fintechOnDjango.views import PositionView

urlpatterns = [
    path('', PositionView.index, name='index'),
	path('create', PositionView.get, name='create'),
	path('get/<int:positionId>/', PositionView.get, name='get'),
	path('save', PositionView.save, name='save'),
	path('getAll', PositionView.getAll, name='getAll'),
	path('delete/<int:positionId>/', PositionView.delete, name='delete'),
	path('assignPortfolio/<int:positionId>/<int:PortfolioId>/', PositionView.assignPortfolio, name='assignPortfolio'),
	path('unassignPortfolio/<int:positionId>/', PositionView.unassignPortfolio, name='unassignPortfolio'),
	path('assignSecurity/<int:positionId>/<int:SecurityId>/', PositionView.assignSecurity, name='assignSecurity'),
	path('unassignSecurity/<int:positionId>/', PositionView.unassignSecurity, name='unassignSecurity'),
]
