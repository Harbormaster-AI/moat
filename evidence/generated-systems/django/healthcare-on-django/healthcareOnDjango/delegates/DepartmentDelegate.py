from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Department import Department
from healthcareOnDjango.models.Facility import Facility
from healthcareOnDjango.models.CareTeam import CareTeam
from healthcareOnDjango.exceptions import Exceptions

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
		
	def assignFacility( self, departmentId, facilityId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.FacilityDelegate import FacilityDelegate

		errMsg = "Failed to assign element " + str(facilityId) + " for Facility on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# get the Facility from db
			facility = FacilityDelegate().get(facilityId).first();
			
			# assign the Facility		
			department.facility = facility
			
			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Facility.DoesNotExist:
			raise ProcessingError(errMsg + " : Facility with id " + str(facilityId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignFacility( self, departmentId ):
		errMsg = "Failed to unassign element " + str(facilityId) + " for Facility on Department"

		try:
			# get the Department from db
			department = self.get( departmentId ).first()	
			
			# assign to None for unassignment
			department.facility = None			

			#save it
			department.save()

			# reload and return the appropriate version					
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except Exception:
			return None;
		
	def addCareTeams( self, departmentId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to add elements " + str(careTeamsIds) + " for CareTeams on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				department.careTeams.add(careTeam)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCareTeams( self, departmentId, careTeamsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CareTeamDelegate import CareTeamDelegate

		errMsg = "Failed to remove elements " + str(careTeamsIds) + " for CareTeams on Department"

		try:
			# get the Department
			department = self.get( departmentId ).first()
				
			# split on a comma with no spaces
			idList = careTeamsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CareTeam		
				careTeam = CareTeamDelegate().get(id).first();	
				# add the CareTeam
				department.careTeams.remove(careTeam)
				
			# save it		
			department.save()
			
			# reload and return the appropriate version
			return self.get( departmentId );
		except Department.DoesNotExist:
			raise ProcessingError(errMsg + " : Department with id " + str(departmentId) + " does not exist.")
		except CareTeam.DoesNotExist:
			raise ProcessingError(errMsg + " : CareTeam does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
