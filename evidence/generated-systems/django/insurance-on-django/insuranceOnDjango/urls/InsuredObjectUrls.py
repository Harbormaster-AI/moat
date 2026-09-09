from django.urls import path
from insuranceOnDjango.views import InsuredObjectView

urlpatterns = [
    path('', InsuredObjectView.index, name='index'),
	path('create', InsuredObjectView.get, name='create'),
	path('get/<int:insuredObjectId>/', InsuredObjectView.get, name='get'),
	path('save', InsuredObjectView.save, name='save'),
	path('getAll', InsuredObjectView.getAll, name='getAll'),
	path('delete/<int:insuredObjectId>/', InsuredObjectView.delete, name='delete'),
	path('assignPolicy/<int:insuredObjectId>/<int:PolicyId>/', InsuredObjectView.assignPolicy, name='assignPolicy'),
	path('unassignPolicy/<int:insuredObjectId>/', InsuredObjectView.unassignPolicy, name='unassignPolicy'),
	path('addCoverages/<int:insuredObjectId>/<CoveragesIds>/', InsuredObjectView.addCoverages, name='addCoverages'),
	path('removeCoverages/<int:insuredObjectId>/<CoveragesIds>/', InsuredObjectView.removeCoverages, name='removeCoverages'),
]
