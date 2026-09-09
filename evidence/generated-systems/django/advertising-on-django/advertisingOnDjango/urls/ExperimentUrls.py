from django.urls import path
from advertisingOnDjango.views import ExperimentView

urlpatterns = [
    path('', ExperimentView.index, name='index'),
	path('create', ExperimentView.get, name='create'),
	path('get/<int:experimentId>/', ExperimentView.get, name='get'),
	path('save', ExperimentView.save, name='save'),
	path('getAll', ExperimentView.getAll, name='getAll'),
	path('delete/<int:experimentId>/', ExperimentView.delete, name='delete'),
	path('assignCampaign/<int:experimentId>/<int:CampaignId>/', ExperimentView.assignCampaign, name='assignCampaign'),
	path('unassignCampaign/<int:experimentId>/', ExperimentView.unassignCampaign, name='unassignCampaign'),
	path('addVariants/<int:experimentId>/<VariantsIds>/', ExperimentView.addVariants, name='addVariants'),
	path('removeVariants/<int:experimentId>/<VariantsIds>/', ExperimentView.removeVariants, name='removeVariants'),
]
