from django.urls import path
from inventoryOnDjango.views import ReplenishmentPolicyView

urlpatterns = [
    path('', ReplenishmentPolicyView.index, name='index'),
	path('create', ReplenishmentPolicyView.get, name='create'),
	path('get/<int:replenishmentPolicyId>/', ReplenishmentPolicyView.get, name='get'),
	path('save', ReplenishmentPolicyView.save, name='save'),
	path('getAll', ReplenishmentPolicyView.getAll, name='getAll'),
	path('delete/<int:replenishmentPolicyId>/', ReplenishmentPolicyView.delete, name='delete'),
	path('assignSku/<int:replenishmentPolicyId>/<int:SkuId>/', ReplenishmentPolicyView.assignSku, name='assignSku'),
	path('unassignSku/<int:replenishmentPolicyId>/', ReplenishmentPolicyView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:replenishmentPolicyId>/<int:WarehouseId>/', ReplenishmentPolicyView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:replenishmentPolicyId>/', ReplenishmentPolicyView.unassignWarehouse, name='unassignWarehouse'),
	path('assignLocation/<int:replenishmentPolicyId>/<int:LocationId>/', ReplenishmentPolicyView.assignLocation, name='assignLocation'),
	path('unassignLocation/<int:replenishmentPolicyId>/', ReplenishmentPolicyView.unassignLocation, name='unassignLocation'),
]
