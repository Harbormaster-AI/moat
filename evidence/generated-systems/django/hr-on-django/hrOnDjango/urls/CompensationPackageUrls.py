from django.urls import path
from hrOnDjango.views import CompensationPackageView

urlpatterns = [
    path('', CompensationPackageView.index, name='index'),
	path('create', CompensationPackageView.get, name='create'),
	path('get/<int:compensationPackageId>/', CompensationPackageView.get, name='get'),
	path('save', CompensationPackageView.save, name='save'),
	path('getAll', CompensationPackageView.getAll, name='getAll'),
	path('delete/<int:compensationPackageId>/', CompensationPackageView.delete, name='delete'),
	path('assignContract/<int:compensationPackageId>/<int:ContractId>/', CompensationPackageView.assignContract, name='assignContract'),
	path('unassignContract/<int:compensationPackageId>/', CompensationPackageView.unassignContract, name='unassignContract'),
	path('addSalaryComponents/<int:compensationPackageId>/<SalaryComponentsIds>/', CompensationPackageView.addSalaryComponents, name='addSalaryComponents'),
	path('removeSalaryComponents/<int:compensationPackageId>/<SalaryComponentsIds>/', CompensationPackageView.removeSalaryComponents, name='removeSalaryComponents'),
	path('addBonusPlans/<int:compensationPackageId>/<BonusPlansIds>/', CompensationPackageView.addBonusPlans, name='addBonusPlans'),
	path('removeBonusPlans/<int:compensationPackageId>/<BonusPlansIds>/', CompensationPackageView.removeBonusPlans, name='removeBonusPlans'),
	path('addEquityGrants/<int:compensationPackageId>/<EquityGrantsIds>/', CompensationPackageView.addEquityGrants, name='addEquityGrants'),
	path('removeEquityGrants/<int:compensationPackageId>/<EquityGrantsIds>/', CompensationPackageView.removeEquityGrants, name='removeEquityGrants'),
]
