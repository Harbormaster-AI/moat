from django.urls import path
from governanceOnDjango.views import DataBreachView

urlpatterns = [
    path('', DataBreachView.index, name='index'),
	path('create', DataBreachView.get, name='create'),
	path('get/<int:dataBreachId>/', DataBreachView.get, name='get'),
	path('save', DataBreachView.save, name='save'),
	path('getAll', DataBreachView.getAll, name='getAll'),
	path('delete/<int:dataBreachId>/', DataBreachView.delete, name='delete'),
	path('assignOrganization/<int:dataBreachId>/<int:OrganizationId>/', DataBreachView.assignOrganization, name='assignOrganization'),
	path('unassignOrganization/<int:dataBreachId>/', DataBreachView.unassignOrganization, name='unassignOrganization'),
	path('assignMatter/<int:dataBreachId>/<int:MatterId>/', DataBreachView.assignMatter, name='assignMatter'),
	path('unassignMatter/<int:dataBreachId>/', DataBreachView.unassignMatter, name='unassignMatter'),
	path('addProcessingActivities/<int:dataBreachId>/<ProcessingActivitiesIds>/', DataBreachView.addProcessingActivities, name='addProcessingActivities'),
	path('removeProcessingActivities/<int:dataBreachId>/<ProcessingActivitiesIds>/', DataBreachView.removeProcessingActivities, name='removeProcessingActivities'),
	path('addDataCategories/<int:dataBreachId>/<DataCategoriesIds>/', DataBreachView.addDataCategories, name='addDataCategories'),
	path('removeDataCategories/<int:dataBreachId>/<DataCategoriesIds>/', DataBreachView.removeDataCategories, name='removeDataCategories'),
	path('addThirdParties/<int:dataBreachId>/<ThirdPartiesIds>/', DataBreachView.addThirdParties, name='addThirdParties'),
	path('removeThirdParties/<int:dataBreachId>/<ThirdPartiesIds>/', DataBreachView.removeThirdParties, name='removeThirdParties'),
]
