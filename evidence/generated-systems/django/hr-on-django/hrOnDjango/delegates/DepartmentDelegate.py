from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Department import Department
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.Employee import Employee
from hrOnDjango.models.Position import Position
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Department
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class DepartmentDelegate Declaration
#======================================================================
class DepartmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, departmentId ):
		try:	
			department = Department.objects.filter(id=departmentId)
			return department.first();
		except Department.DoesNotExist:
			raise ProcessingError("Department with id " + str(departmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, department):
		for model in serializers.deserialize("json", department):
			model.save()
			return model;

	def create(self, department):
		department.save()
		return department;

	def saveFromJson(self, department):
		for model in serializers.deserialize("json", department):
			model.save()
			return department;
	
	def save(self, department):
		department.save()
		return department;
	
	def delete(self, departmentId ):
		errMsg = "Failed to delete Department from db using id " + str(departmentId)
		
		try:
			department = Department.objects.get(id=departmentId)
			department.delete()
			return True
		except Department.DoesNotExist:
			raise ProcessingError("Department with id " + str(departmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Department.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Department from db")
		except Exception:
			return None;
		
	def assignOrganization( self, departmentId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			department.organization = organization
			
			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, departmentId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# assign to None for unassignment
			department.organization = None			

			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignManager( self, departmentId, managerId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(managerId) + " for Manager on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(managerId).first();
			
			# assign the Manager		
			department.manager = employee
			
			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(managerId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManager( self, departmentId ):
		errMsg = "Failed to unassign element " + str(managerId) + " for Manager on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# assign to None for unassignment
			department.employee = None			

			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCostCenter( self, departmentId, costCenterId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to assign element " + str(costCenterId) + " for CostCenter on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# get the CostCenter from db
			costCenter = CostCenterDelegate().get(costCenterId).first();
			
			# assign the CostCenter		
			department.costCenter = costCenter
			
			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCostCenter( self, departmentId ):
		errMsg = "Failed to unassign element " + str(costCenterId) + " for CostCenter on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# assign to None for unassignment
			department.costCenter = None			

			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addPositions( self, departmentId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(positionsIds) + " for Positions on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				department.positions.add(position)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePositions( self, departmentId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(positionsIds) + " for Positions on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				department.positions.remove(position)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmployees( self, departmentId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(employeesIds) + " for Employees on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				department.employees.add(employee)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmployees( self, departmentId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(employeesIds) + " for Employees on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				department.employees.remove(employee)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
