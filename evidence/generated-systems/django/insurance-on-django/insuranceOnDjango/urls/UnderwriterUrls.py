from django.urls import path
from insuranceOnDjango.views import UnderwriterView

urlpatterns = [
    path('', UnderwriterView.index, name='index'),
	path('create', UnderwriterView.get, name='create'),
	path('get/<int:underwriterId>/', UnderwriterView.get, name='get'),
	path('save', UnderwriterView.save, name='save'),
	path('getAll', UnderwriterView.getAll, name='getAll'),
	path('delete/<int:underwriterId>/', UnderwriterView.delete, name='delete'),
	path('assignInsurer/<int:underwriterId>/<int:InsurerId>/', UnderwriterView.assignInsurer, name='assignInsurer'),
	path('unassignInsurer/<int:underwriterId>/', UnderwriterView.unassignInsurer, name='unassignInsurer'),
	path('addDecisions/<int:underwriterId>/<DecisionsIds>/', UnderwriterView.addDecisions, name='addDecisions'),
	path('removeDecisions/<int:underwriterId>/<DecisionsIds>/', UnderwriterView.removeDecisions, name='removeDecisions'),
]
