from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.Department import Department
from hrOnDjango.models.Position import Position
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CostCenter
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CostCenterDelegate Declaration
#======================================================================
class CostCenterDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, costCenterId ):
		try:	
			costCenter = CostCenter.objects.filter(id=costCenterId)
			return costCenter.first();
		except CostCenter.DoesNotExist:
			raise ProcessingError("CostCenter with id " + str(costCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, costCenter):
		for model in serializers.deserialize("json", costCenter):
			model.save()
			return model;

	def create(self, costCenter):
		costCenter.save()
		return costCenter;

	def saveFromJson(self, costCenter):
		for model in serializers.deserialize("json", costCenter):
			model.save()
			return costCenter;
	
	def save(self, costCenter):
		costCenter.save()
		return costCenter;
	
	def delete(self, costCenterId ):
		errMsg = "Failed to delete CostCenter from db using id " + str(costCenterId)
		
		try:
			costCenter = CostCenter.objects.get(id=costCenterId)
			costCenter.delete()
			return True
		except CostCenter.DoesNotExist:
			raise ProcessingError("CostCenter with id " + str(costCenterId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CostCenter.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CostCenter from db")
		except Exception:
			return None;
		
	def assignOrganization( self, costCenterId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on CostCenter"

		try:
			# get the CostCenter from db
			costCenter = self.get( costCenterId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			costCenter.organization = organization
			
			#save it
			costCenter.save()

			# reload and return the appropriate version					
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, costCenterId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on CostCenter"

		try:
			# get the CostCenter from db
			costCenter = self.get( costCenterId ).first()	
			
			# assign to None for unassignment
			costCenter.organization = None			

			#save it
			costCenter.save()

			# reload and return the appropriate version					
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Exception:
			return None;
		
	def addDepartments( self, costCenterId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to add elements " + str(departmentsIds) + " for Departments on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				costCenter.departments.add(department)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDepartments( self, costCenterId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to remove elements " + str(departmentsIds) + " for Departments on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				costCenter.departments.remove(department)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPositions( self, costCenterId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(positionsIds) + " for Positions on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				costCenter.positions.add(position)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePositions( self, costCenterId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(positionsIds) + " for Positions on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				costCenter.positions.remove(position)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmployees( self, costCenterId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(employeesIds) + " for Employees on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				costCenter.employees.add(employee)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmployees( self, costCenterId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(employeesIds) + " for Employees on CostCenter"

		try:
			# get the CostCenter
			costCenter = self.get( costCenterId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				costCenter.employees.remove(employee)
				
			# save it		
			costCenter.save()
			
			# reload and return the appropriate version
			return self.get( costCenterId );
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
