from django.urls import path
from fintechOnDjango.views import ProductOfferingView

urlpatterns = [
    path('', ProductOfferingView.index, name='index'),
	path('create', ProductOfferingView.get, name='create'),
	path('get/<int:productOfferingId>/', ProductOfferingView.get, name='get'),
	path('save', ProductOfferingView.save, name='save'),
	path('getAll', ProductOfferingView.getAll, name='getAll'),
	path('delete/<int:productOfferingId>/', ProductOfferingView.delete, name='delete'),
	path('assignInstitution/<int:productOfferingId>/<int:InstitutionId>/', ProductOfferingView.assignInstitution, name='assignInstitution'),
	path('unassignInstitution/<int:productOfferingId>/', ProductOfferingView.unassignInstitution, name='unassignInstitution'),
	path('addPricingPlans/<int:productOfferingId>/<PricingPlansIds>/', ProductOfferingView.addPricingPlans, name='addPricingPlans'),
	path('removePricingPlans/<int:productOfferingId>/<PricingPlansIds>/', ProductOfferingView.removePricingPlans, name='removePricingPlans'),
]
