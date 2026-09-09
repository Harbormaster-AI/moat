from django.urls import path
from aerospaceOnDjango.views import ServiceBulletinView

urlpatterns = [
    path('', ServiceBulletinView.index, name='index'),
	path('create', ServiceBulletinView.get, name='create'),
	path('get/<int:serviceBulletinId>/', ServiceBulletinView.get, name='get'),
	path('save', ServiceBulletinView.save, name='save'),
	path('getAll', ServiceBulletinView.getAll, name='getAll'),
	path('delete/<int:serviceBulletinId>/', ServiceBulletinView.delete, name='delete'),
	path('addWorkOrders/<int:serviceBulletinId>/<WorkOrdersIds>/', ServiceBulletinView.addWorkOrders, name='addWorkOrders'),
	path('removeWorkOrders/<int:serviceBulletinId>/<WorkOrdersIds>/', ServiceBulletinView.removeWorkOrders, name='removeWorkOrders'),
	path('addVariants/<int:serviceBulletinId>/<VariantsIds>/', ServiceBulletinView.addVariants, name='addVariants'),
	path('removeVariants/<int:serviceBulletinId>/<VariantsIds>/', ServiceBulletinView.removeVariants, name='removeVariants'),
]
