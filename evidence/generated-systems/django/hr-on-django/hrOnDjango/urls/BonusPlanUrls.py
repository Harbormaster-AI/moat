from django.urls import path
from hrOnDjango.views import BonusPlanView

urlpatterns = [
    path('', BonusPlanView.index, name='index'),
	path('create', BonusPlanView.get, name='create'),
	path('get/<int:bonusPlanId>/', BonusPlanView.get, name='get'),
	path('save', BonusPlanView.save, name='save'),
	path('getAll', BonusPlanView.getAll, name='getAll'),
	path('delete/<int:bonusPlanId>/', BonusPlanView.delete, name='delete'),
	path('addCompensationPackages/<int:bonusPlanId>/<CompensationPackagesIds>/', BonusPlanView.addCompensationPackages, name='addCompensationPackages'),
	path('removeCompensationPackages/<int:bonusPlanId>/<CompensationPackagesIds>/', BonusPlanView.removeCompensationPackages, name='removeCompensationPackages'),
]
