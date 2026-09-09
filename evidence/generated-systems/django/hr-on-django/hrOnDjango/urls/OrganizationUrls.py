from django.urls import path
from hrOnDjango.views import OrganizationView

urlpatterns = [
    path('', OrganizationView.index, name='index'),
	path('create', OrganizationView.get, name='create'),
	path('get/<int:organizationId>/', OrganizationView.get, name='get'),
	path('save', OrganizationView.save, name='save'),
	path('getAll', OrganizationView.getAll, name='getAll'),
	path('delete/<int:organizationId>/', OrganizationView.delete, name='delete'),
	path('addDepartments/<int:organizationId>/<DepartmentsIds>/', OrganizationView.addDepartments, name='addDepartments'),
	path('removeDepartments/<int:organizationId>/<DepartmentsIds>/', OrganizationView.removeDepartments, name='removeDepartments'),
	path('addLocations/<int:organizationId>/<LocationsIds>/', OrganizationView.addLocations, name='addLocations'),
	path('removeLocations/<int:organizationId>/<LocationsIds>/', OrganizationView.removeLocations, name='removeLocations'),
	path('addJobFamilies/<int:organizationId>/<JobFamiliesIds>/', OrganizationView.addJobFamilies, name='addJobFamilies'),
	path('removeJobFamilies/<int:organizationId>/<JobFamiliesIds>/', OrganizationView.removeJobFamilies, name='removeJobFamilies'),
	path('addBenefitPlans/<int:organizationId>/<BenefitPlansIds>/', OrganizationView.addBenefitPlans, name='addBenefitPlans'),
	path('removeBenefitPlans/<int:organizationId>/<BenefitPlansIds>/', OrganizationView.removeBenefitPlans, name='removeBenefitPlans'),
	path('addCostCenters/<int:organizationId>/<CostCentersIds>/', OrganizationView.addCostCenters, name='addCostCenters'),
	path('removeCostCenters/<int:organizationId>/<CostCentersIds>/', OrganizationView.removeCostCenters, name='removeCostCenters'),
	path('addPayrollCalendars/<int:organizationId>/<PayrollCalendarsIds>/', OrganizationView.addPayrollCalendars, name='addPayrollCalendars'),
	path('removePayrollCalendars/<int:organizationId>/<PayrollCalendarsIds>/', OrganizationView.removePayrollCalendars, name='removePayrollCalendars'),
]
