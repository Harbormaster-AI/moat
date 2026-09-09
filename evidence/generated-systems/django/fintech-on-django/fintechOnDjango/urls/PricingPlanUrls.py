from django.urls import path
from fintechOnDjango.views import PricingPlanView

urlpatterns = [
    path('', PricingPlanView.index, name='index'),
	path('create', PricingPlanView.get, name='create'),
	path('get/<int:pricingPlanId>/', PricingPlanView.get, name='get'),
	path('save', PricingPlanView.save, name='save'),
	path('getAll', PricingPlanView.getAll, name='getAll'),
	path('delete/<int:pricingPlanId>/', PricingPlanView.delete, name='delete'),
	path('assignProductOffering/<int:pricingPlanId>/<int:ProductOfferingId>/', PricingPlanView.assignProductOffering, name='assignProductOffering'),
	path('unassignProductOffering/<int:pricingPlanId>/', PricingPlanView.unassignProductOffering, name='unassignProductOffering'),
	path('addFeeSchedules/<int:pricingPlanId>/<FeeSchedulesIds>/', PricingPlanView.addFeeSchedules, name='addFeeSchedules'),
	path('removeFeeSchedules/<int:pricingPlanId>/<FeeSchedulesIds>/', PricingPlanView.removeFeeSchedules, name='removeFeeSchedules'),
	path('addLimits/<int:pricingPlanId>/<LimitsIds>/', PricingPlanView.addLimits, name='addLimits'),
	path('removeLimits/<int:pricingPlanId>/<LimitsIds>/', PricingPlanView.removeLimits, name='removeLimits'),
]
