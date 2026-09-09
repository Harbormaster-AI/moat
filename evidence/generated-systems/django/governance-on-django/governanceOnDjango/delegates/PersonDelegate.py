from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Person import Person
from governanceOnDjango.models.RoleAssignment import RoleAssignment
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.CorrectiveAction import CorrectiveAction
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Person
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PersonDelegate Declaration
#======================================================================
class PersonDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, personId ):
		try:	
			person = Person.objects.filter(id=personId)
			return person.first();
		except Person.DoesNotExist:
			raise ProcessingError("Person with id " + str(personId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, person):
		for model in serializers.deserialize("json", person):
			model.save()
			return model;

	def create(self, person):
		person.save()
		return person;

	def saveFromJson(self, person):
		for model in serializers.deserialize("json", person):
			model.save()
			return person;
	
	def save(self, person):
		person.save()
		return person;
	
	def delete(self, personId ):
		errMsg = "Failed to delete Person from db using id " + str(personId)
		
		try:
			person = Person.objects.get(id=personId)
			person.delete()
			return True
		except Person.DoesNotExist:
			raise ProcessingError("Person with id " + str(personId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Person.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Person from db")
		except Exception:
			return None;
		
	def addRoleAssignments( self, personId, roleAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to add elements " + str(roleAssignmentsIds) + " for RoleAssignments on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = roleAssignmentsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				person.roleAssignments.add(roleAssignment)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRoleAssignments( self, personId, roleAssignmentsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleAssignmentDelegate import RoleAssignmentDelegate

		errMsg = "Failed to remove elements " + str(roleAssignmentsIds) + " for RoleAssignments on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = roleAssignmentsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the RoleAssignment		
				roleAssignment = RoleAssignmentDelegate().get(id).first();	
				# add the RoleAssignment
				person.roleAssignments.remove(roleAssignment)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addOwnedPolicies( self, personId, ownedPoliciesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(ownedPoliciesIds) + " for OwnedPolicies on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = ownedPoliciesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				person.ownedPolicies.add(policy)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwnedPolicies( self, personId, ownedPoliciesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(ownedPoliciesIds) + " for OwnedPolicies on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = ownedPoliciesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				person.ownedPolicies.remove(policy)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addCorrectiveActions( self, personId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to add elements " + str(correctiveActionsIds) + " for CorrectiveActions on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				person.correctiveActions.add(correctiveAction)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeCorrectiveActions( self, personId, correctiveActionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.CorrectiveActionDelegate import CorrectiveActionDelegate

		errMsg = "Failed to remove elements " + str(correctiveActionsIds) + " for CorrectiveActions on Person"

		try:
			# get the Person
			person = self.get( personId ).first()
				
			# split on a comma with no spaces
			idList = correctiveActionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the CorrectiveAction		
				correctiveAction = CorrectiveActionDelegate().get(id).first();	
				# add the CorrectiveAction
				person.correctiveActions.remove(correctiveAction)
				
			# save it		
			person.save()
			
			# reload and return the appropriate version
			return self.get( personId );
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except CorrectiveAction.DoesNotExist:
			raise ProcessingError(errMsg + " : CorrectiveAction does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
