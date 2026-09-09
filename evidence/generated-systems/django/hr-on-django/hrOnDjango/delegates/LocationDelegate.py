from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Location import Location
from hrOnDjango.models.Organization import Organization
from hrOnDjango.models.Department import Department
from hrOnDjango.models.Position import Position
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Location
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class LocationDelegate Declaration
#======================================================================
class LocationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, locationId ):
		try:	
			location = Location.objects.filter(id=locationId)
			return location.first();
		except Location.DoesNotExist:
			raise ProcessingError("Location with id " + str(locationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, location):
		for model in serializers.deserialize("json", location):
			model.save()
			return model;

	def create(self, location):
		location.save()
		return location;

	def saveFromJson(self, location):
		for model in serializers.deserialize("json", location):
			model.save()
			return location;
	
	def save(self, location):
		location.save()
		return location;
	
	def delete(self, locationId ):
		errMsg = "Failed to delete Location from db using id " + str(locationId)
		
		try:
			location = Location.objects.get(id=locationId)
			location.delete()
			return True
		except Location.DoesNotExist:
			raise ProcessingError("Location with id " + str(locationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Location.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Location from db")
		except Exception:
			return None;
		
	def assignOrganization( self, locationId, organizationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Location"

		try:
			# get the Location from db
			location = self.get( locationId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			location.organization = organization
			
			#save it
			location.save()

			# reload and return the appropriate version					
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, locationId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Location"

		try:
			# get the Location from db
			location = self.get( locationId ).first()	
			
			# assign to None for unassignment
			location.organization = None			

			#save it
			location.save()

			# reload and return the appropriate version					
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
		
	def addDepartments( self, locationId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to add elements " + str(departmentsIds) + " for Departments on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				location.departments.add(department)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDepartments( self, locationId, departmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to remove elements " + str(departmentsIds) + " for Departments on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = departmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Department		
				department = DepartmentDelegate().get(id).first();	
				# add the Department
				location.departments.remove(department)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addPositions( self, locationId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(positionsIds) + " for Positions on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				location.positions.add(position)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePositions( self, locationId, positionsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(positionsIds) + " for Positions on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = positionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				location.positions.remove(position)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEmployees( self, locationId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to add elements " + str(employeesIds) + " for Employees on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				location.employees.add(employee)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEmployees( self, locationId, employeesIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to remove elements " + str(employeesIds) + " for Employees on Location"

		try:
			# get the Location
			location = self.get( locationId ).first()
				
			# split on a comma with no spaces
			idList = employeesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Employee		
				employee = EmployeeDelegate().get(id).first();	
				# add the Employee
				location.employees.remove(employee)
				
			# save it		
			location.save()
			
			# reload and return the appropriate version
			return self.get( locationId );
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
