from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.Position import Position
from hrOnDjango.models.Department import Department
from hrOnDjango.models.JobProfile import JobProfile
from hrOnDjango.models.CostCenter import CostCenter
from hrOnDjango.models.Location import Location
from hrOnDjango.models.EmploymentAssignment import EmploymentAssignment
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Position
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PositionDelegate Declaration
#======================================================================
class PositionDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, positionId ):
		try:	
			position = Position.objects.filter(id=positionId)
			return position.first();
		except Position.DoesNotExist:
			raise ProcessingError("Position with id " + str(positionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, position):
		for model in serializers.deserialize("json", position):
			model.save()
			return model;

	def create(self, position):
		position.save()
		return position;

	def saveFromJson(self, position):
		for model in serializers.deserialize("json", position):
			model.save()
			return position;
	
	def save(self, position):
		position.save()
		return position;
	
	def delete(self, positionId ):
		errMsg = "Failed to delete Position from db using id " + str(positionId)
		
		try:
			position = Position.objects.get(id=positionId)
			position.delete()
			return True
		except Position.DoesNotExist:
			raise ProcessingError("Position with id " + str(positionId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Position.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Position from db")
		except Exception:
			return None;
		
	def assignDepartment( self, positionId, departmentId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.DepartmentDelegate import DepartmentDelegate

		errMsg = "Failed to assign element " + str(departmentId) + " for Department on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the Department from db
			department = DepartmentDelegate().get(departmentId).first();
			
			# assign the Department		
			position.department = department
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignDepartment( self, positionId ):
		errMsg = "Failed to unassign element " + str(departmentId) + " for Department on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.department = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignJobProfile( self, positionId, jobProfileId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.JobProfileDelegate import JobProfileDelegate

		errMsg = "Failed to assign element " + str(jobProfileId) + " for JobProfile on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the JobProfile from db
			jobProfile = JobProfileDelegate().get(jobProfileId).first();
			
			# assign the JobProfile		
			position.jobProfile = jobProfile
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except JobProfile.DoesNotExist:
			raise ProcessingError(errMsg + " : JobProfile with id " + str(jobProfileId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignJobProfile( self, positionId ):
		errMsg = "Failed to unassign element " + str(jobProfileId) + " for JobProfile on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.jobProfile = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignCostCenter( self, positionId, costCenterId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.CostCenterDelegate import CostCenterDelegate

		errMsg = "Failed to assign element " + str(costCenterId) + " for CostCenter on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the CostCenter from db
			costCenter = CostCenterDelegate().get(costCenterId).first();
			
			# assign the CostCenter		
			position.costCenter = costCenter
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except CostCenter.DoesNotExist:
			raise ProcessingError(errMsg + " : CostCenter with id " + str(costCenterId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCostCenter( self, positionId ):
		errMsg = "Failed to unassign element " + str(costCenterId) + " for CostCenter on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.costCenter = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignLocation( self, positionId, locationId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.LocationDelegate import LocationDelegate

		errMsg = "Failed to assign element " + str(locationId) + " for Location on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the Location from db
			location = LocationDelegate().get(locationId).first();
			
			# assign the Location		
			position.location = location
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Location.DoesNotExist:
			raise ProcessingError(errMsg + " : Location with id " + str(locationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignLocation( self, positionId ):
		errMsg = "Failed to unassign element " + str(locationId) + " for Location on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.location = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def assignManagerPosition( self, positionId, managerPositionId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to assign element " + str(managerPositionId) + " for ManagerPosition on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# get the Position from db
			position = PositionDelegate().get(managerPositionId).first();
			
			# assign the ManagerPosition		
			position.managerPosition = position
			
			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(managerPositionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignManagerPosition( self, positionId ):
		errMsg = "Failed to unassign element " + str(managerPositionId) + " for ManagerPosition on Position"

		try:
			# get the Position from db
			position = self.get( positionId ).first()	
			
			# assign to None for unassignment
			position.position = None			

			#save it
			position.save()

			# reload and return the appropriate version					
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Exception:
			return None;
		
	def addDirectReports( self, positionId, directReportsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to add elements " + str(directReportsIds) + " for DirectReports on Position"

		try:
			# get the Position
			position = self.get( positionId ).first()
				
			# split on a comma with no spaces
			idList = directReportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				position.directReports.add(position)
				
			# save it		
			position.save()
			
			# reload and return the appropriate version
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDirectReports( self, positionId, directReportsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PositionDelegate import PositionDelegate

		errMsg = "Failed to remove elements " + str(directReportsIds) + " for DirectReports on Position"

		try:
			# get the Position
			position = self.get( positionId ).first()
				
			# split on a comma with no spaces
			idList = directReportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Position		
				position = PositionDelegate().get(id).first();	
				# add the Position
				position.directReports.remove(position)
				
			# save it		
			position.save()
			
			# reload and return the appropriate version
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAssignments( self, positionId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

		errMsg = "Failed to add elements " + str(assignmentsIds) + " for Assignments on Position"

		try:
			# get the Position
			position = self.get( positionId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the EmploymentAssignment		
				employmentAssignment = EmploymentAssignmentDelegate().get(id).first();	
				# add the EmploymentAssignment
				position.assignments.add(employmentAssignment)
				
			# save it		
			position.save()
			
			# reload and return the appropriate version
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssignments( self, positionId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmploymentAssignmentDelegate import EmploymentAssignmentDelegate

		errMsg = "Failed to remove elements " + str(assignmentsIds) + " for Assignments on Position"

		try:
			# get the Position
			position = self.get( positionId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the EmploymentAssignment		
				employmentAssignment = EmploymentAssignmentDelegate().get(id).first();	
				# add the EmploymentAssignment
				position.assignments.remove(employmentAssignment)
				
			# save it		
			position.save()
			
			# reload and return the appropriate version
			return self.get( positionId );
		except Position.DoesNotExist:
			raise ProcessingError(errMsg + " : Position with id " + str(positionId) + " does not exist.")
		except EmploymentAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : EmploymentAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
