from django.urls import path
from ecommerceOnDjango.views import OrderLineView

urlpatterns = [
    path('', OrderLineView.index, name='index'),
	path('create', OrderLineView.get, name='create'),
	path('get/<int:orderLineId>/', OrderLineView.get, name='get'),
	path('save', OrderLineView.save, name='save'),
	path('getAll', OrderLineView.getAll, name='getAll'),
	path('delete/<int:orderLineId>/', OrderLineView.delete, name='delete'),
	path('assignOrder/<int:orderLineId>/<int:OrderId>/', OrderLineView.assignOrder, name='assignOrder'),
	path('unassignOrder/<int:orderLineId>/', OrderLineView.unassignOrder, name='unassignOrder'),
	path('assignVariant/<int:orderLineId>/<int:VariantId>/', OrderLineView.assignVariant, name='assignVariant'),
	path('unassignVariant/<int:orderLineId>/', OrderLineView.unassignVariant, name='unassignVariant'),
	path('addAppliedPromotions/<int:orderLineId>/<AppliedPromotionsIds>/', OrderLineView.addAppliedPromotions, name='addAppliedPromotions'),
	path('removeAppliedPromotions/<int:orderLineId>/<AppliedPromotionsIds>/', OrderLineView.removeAppliedPromotions, name='removeAppliedPromotions'),
]
