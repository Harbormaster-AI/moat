from django.urls import path
from fintechOnDjango.views import FeeScheduleView

urlpatterns = [
    path('', FeeScheduleView.index, name='index'),
	path('create', FeeScheduleView.get, name='create'),
	path('get/<int:feeScheduleId>/', FeeScheduleView.get, name='get'),
	path('save', FeeScheduleView.save, name='save'),
	path('getAll', FeeScheduleView.getAll, name='getAll'),
	path('delete/<int:feeScheduleId>/', FeeScheduleView.delete, name='delete'),
	path('assignPricingPlan/<int:feeScheduleId>/<int:PricingPlanId>/', FeeScheduleView.assignPricingPlan, name='assignPricingPlan'),
	path('unassignPricingPlan/<int:feeScheduleId>/', FeeScheduleView.unassignPricingPlan, name='unassignPricingPlan'),
]
