from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.Department import Department
from hrOnDjango.models.Location import Location
from hrOnDjango.models.JobFamily import JobFamily
from hrOnDjango.models.BenefitPlan import BenefitPlan
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.models.PayrollCalendar import PayrollCalendar
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Organization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class OrganizationDelegate Declaration
#======================================================================
class OrganizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, organizationId ):
		try:	
			organization = Organization.objects.filter(id=organizationId)
			return organization.first();
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return model;

	def create(self, organization):
		organization.save()
		return organization;

	def saveFromJson(self, organization):
		for model in serializers.deserialize("json", organization):
			model.save()
			return organization;
	
	def save(self, organization):
		organization.save()
		return organization;
	
	def delete(self, organizationId ):
		errMsg = "Failed to delete Organization from db using id " + str(organizationId)
		
		try:
			organization = Organization.objects.get(id=organizationId)
			organization.delete()
			return True
		except Organization.DoesNotExist:
			raise ProcessingError("Organization with id " + str(organizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Organization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Organization from db")
		except Exception:
			return None;
		
	def addDepartments( self, organizationId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to add elements " + str(departmentsIds) + " for Departments on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				organization.departments.add(department)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDepartments( self, organizationId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to remove elements " + str(departmentsIds) + " for Departments on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				organization.departments.remove(department)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addLocations( self, organizationId, locationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to add elements " + str(locationsIds) + " for Locations on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Location		
				location = LocationDelegate().get(id).first();	
				# add the Location
				organization.locations.add(location)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeLocations( self, organizationId, locationsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to remove elements " + str(locationsIds) + " for Locations on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = locationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Location		
				location = LocationDelegate().get(id).first();	
				# add the Location
				organization.locations.remove(location)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addJobFamilies( self, organizationId, jobFamiliesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobFamilyDelegate import JobFamilyDelegate

		errMsg = "Failed to add elements " + str(jobFamiliesIds) + " for JobFamilies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = jobFamiliesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the JobFamily		
				jobFamily = JobFamilyDelegate().get(id).first();	
				# add the JobFamily
				organization.jobFamilies.add(jobFamily)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeJobFamilies( self, organizationId, jobFamiliesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobFamilyDelegate import JobFamilyDelegate

		errMsg = "Failed to remove elements " + str(jobFamiliesIds) + " for JobFamilies on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = jobFamiliesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the JobFamily		
				jobFamily = JobFamilyDelegate().get(id).first();	
				# add the JobFamily
				organization.jobFamilies.remove(jobFamily)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except JobFamily.DoesNotExist:
			raise ProcessingError(errMsg + " : JobFamily does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addBenefitPlans( self, organizationId, benefitPlansIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitPlanDelegate import BenefitPlanDelegate

		errMsg = "Failed to add elements " + str(benefitPlansIds) + " for BenefitPlans on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = benefitPlansIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the BenefitPlan		
				benefitPlan = BenefitPlanDelegate().get(id).first();	
				# add the BenefitPlan
				organization.benefitPlans.add(benefitPlan)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeBenefitPlans( self, organizationId, benefitPlansIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.BenefitPlanDelegate import BenefitPlanDelegate

		errMsg = "Failed to remove elements " + str(benefitPlansIds) + " for BenefitPlans on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = benefitPlansIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the BenefitPlan		
				benefitPlan = BenefitPlanDelegate().get(id).first();	
				# add the BenefitPlan
				organization.benefitPlans.remove(benefitPlan)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except BenefitPlan.DoesNotExist:
			raise ProcessingError(errMsg + " : BenefitPlan does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCostCenters( self, organizationId, costCentersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to add elements " + str(costCentersIds) + " for CostCenters on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = costCentersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CostCenter		
				costCenter = CostCenterDelegate().get(id).first();	
				# add the CostCenter
				organization.costCenters.add(costCenter)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCostCenters( self, organizationId, costCentersIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to remove elements " + str(costCentersIds) + " for CostCenters on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = costCentersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CostCenter		
				costCenter = CostCenterDelegate().get(id).first();	
				# add the CostCenter
				organization.costCenters.remove(costCenter)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPayrollCalendars( self, organizationId, payrollCalendarsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

		errMsg = "Failed to add elements " + str(payrollCalendarsIds) + " for PayrollCalendars on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = payrollCalendarsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the PayrollCalendar		
				payrollCalendar = PayrollCalendarDelegate().get(id).first();	
				# add the PayrollCalendar
				organization.payrollCalendars.add(payrollCalendar)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePayrollCalendars( self, organizationId, payrollCalendarsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PayrollCalendarDelegate import PayrollCalendarDelegate

		errMsg = "Failed to remove elements " + str(payrollCalendarsIds) + " for PayrollCalendars on Organization"

		try:
			# get the Organization
			organization = self.get( organizationId ).first()
				
			# split on a comma with no spaces
			idList = payrollCalendarsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the PayrollCalendar		
				payrollCalendar = PayrollCalendarDelegate().get(id).first();	
				# add the PayrollCalendar
				organization.payrollCalendars.remove(payrollCalendar)
				
			# save it		
			organization.save()
			
			# reload and return the appropriate version
			return self.get( organizationId );
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except PayrollCalendar.DoesNotExist:
			raise ProcessingError(errMsg + " : PayrollCalendar does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
