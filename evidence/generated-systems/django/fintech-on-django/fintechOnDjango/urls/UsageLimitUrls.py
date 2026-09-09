from django.urls import path
from fintechOnDjango.views import UsageLimitView

urlpatterns = [
    path('', UsageLimitView.index, name='index'),
	path('create', UsageLimitView.get, name='create'),
	path('get/<int:usageLimitId>/', UsageLimitView.get, name='get'),
	path('save', UsageLimitView.save, name='save'),
	path('getAll', UsageLimitView.getAll, name='getAll'),
	path('delete/<int:usageLimitId>/', UsageLimitView.delete, name='delete'),
	path('assignPricingPlan/<int:usageLimitId>/<int:PricingPlanId>/', UsageLimitView.assignPricingPlan, name='assignPricingPlan'),
	path('unassignPricingPlan/<int:usageLimitId>/', UsageLimitView.unassignPricingPlan, name='unassignPricingPlan'),
]
