from django.urls import path
from inventoryOnDjango.views import ExpirationPolicyView

urlpatterns = [
    path('', ExpirationPolicyView.index, name='index'),
	path('create', ExpirationPolicyView.get, name='create'),
	path('get/<int:expirationPolicyId>/', ExpirationPolicyView.get, name='get'),
	path('save', ExpirationPolicyView.save, name='save'),
	path('getAll', ExpirationPolicyView.getAll, name='getAll'),
	path('delete/<int:expirationPolicyId>/', ExpirationPolicyView.delete, name='delete'),
	path('assignSku/<int:expirationPolicyId>/<int:SkuId>/', ExpirationPolicyView.assignSku, name='assignSku'),
	path('unassignSku/<int:expirationPolicyId>/', ExpirationPolicyView.unassignSku, name='unassignSku'),
	path('assignWarehouse/<int:expirationPolicyId>/<int:WarehouseId>/', ExpirationPolicyView.assignWarehouse, name='assignWarehouse'),
	path('unassignWarehouse/<int:expirationPolicyId>/', ExpirationPolicyView.unassignWarehouse, name='unassignWarehouse'),
]
