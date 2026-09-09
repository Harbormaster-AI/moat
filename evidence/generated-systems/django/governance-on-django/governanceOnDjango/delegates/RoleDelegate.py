from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Role import Role
from governanceOnDjango.models.RoleAssignment import RoleAssignment
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Role
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleDelegate Declaration
#======================================================================
class RoleDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, roleId ):
		try:	
			role = Role.objects.filter(id=roleId)
			return role.first();
		except Role.DoesNotExist:
			raise ProcessingError("Role with id " + str(roleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, role):
		for model in serializers.deserialize("json", role):
			model.save()
			return model;

	def create(self, role):
		role.save()
		return role;

	def saveFromJson(self, role):
		for model in serializers.deserialize("json", role):
			model.save()
			return role;
	
	def save(self, role):
		role.save()
		return role;
	
	def delete(self, roleId ):
		errMsg = "Failed to delete Role from db using id " + str(roleId)
		
		try:
			role = Role.objects.get(id=roleId)
			role.delete()
			return True
		except Role.DoesNotExist:
			raise ProcessingError("Role with id " + str(roleId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Role.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Role from db")
		except Exception:
			return None;
		
	def addAssignments( self, roleId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to add elements " + str(assignmentsIds) + " for Assignments on Role"

		try:
			# get the Role
			role = self.get( roleId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				role.assignments.add(roleAssignment)
				
			# save it		
			role.save()
			
			# reload and return the appropriate version
			return self.get( roleId );
		except Role.DoesNotExist:
			raise ProcessingError(errMsg + " : Role with id " + str(roleId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAssignments( self, roleId, assignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to remove elements " + str(assignmentsIds) + " for Assignments on Role"

		try:
			# get the Role
			role = self.get( roleId ).first()
				
			# split on a comma with no spaces
			idList = assignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				role.assignments.remove(roleAssignment)
				
			# save it		
			role.save()
			
			# reload and return the appropriate version
			return self.get( roleId );
		except Role.DoesNotExist:
			raise ProcessingError(errMsg + " : Role with id " + str(roleId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
