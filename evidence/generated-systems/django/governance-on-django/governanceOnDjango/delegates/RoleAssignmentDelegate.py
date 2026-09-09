from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.RoleAssignment import RoleAssignment
from governanceOnDjango.models.Person import Person
from governanceOnDjango.models.Role import Role
from governanceOnDjango.models.GovernanceBody import GovernanceBody
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model RoleAssignment
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class RoleAssignmentDelegate Declaration
#======================================================================
class RoleAssignmentDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, roleAssignmentId ):
		try:	
			roleAssignment = RoleAssignment.objects.filter(id=roleAssignmentId)
			return roleAssignment.first();
		except RoleAssignment.DoesNotExist:
			raise ProcessingError("RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, roleAssignment):
		for model in serializers.deserialize("json", roleAssignment):
			model.save()
			return model;

	def create(self, roleAssignment):
		roleAssignment.save()
		return roleAssignment;

	def saveFromJson(self, roleAssignment):
		for model in serializers.deserialize("json", roleAssignment):
			model.save()
			return roleAssignment;
	
	def save(self, roleAssignment):
		roleAssignment.save()
		return roleAssignment;
	
	def delete(self, roleAssignmentId ):
		errMsg = "Failed to delete RoleAssignment from db using id " + str(roleAssignmentId)
		
		try:
			roleAssignment = RoleAssignment.objects.get(id=roleAssignmentId)
			roleAssignment.delete()
			return True
		except RoleAssignment.DoesNotExist:
			raise ProcessingError("RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = RoleAssignment.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all RoleAssignment from db")
		except Exception:
			return None;
		
	def assignPerson( self, roleAssignmentId, personId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PersonDelegate import PersonDelegate

		errMsg = "Failed to assign element " + str(personId) + " for Person on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# get the Person from db
			person = PersonDelegate().get(personId).first();
			
			# assign the Person		
			roleAssignment.person = person
			
			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person with id " + str(personId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPerson( self, roleAssignmentId ):
		errMsg = "Failed to unassign element " + str(personId) + " for Person on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# assign to None for unassignment
			roleAssignment.person = None			

			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignRole( self, roleAssignmentId, roleId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RoleDelegate import RoleDelegate

		errMsg = "Failed to assign element " + str(roleId) + " for Role on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# get the Role from db
			role = RoleDelegate().get(roleId).first();
			
			# assign the Role		
			roleAssignment.role = role
			
			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Role.DoesNotExist:
			raise ProcessingError(errMsg + " : Role with id " + str(roleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRole( self, roleAssignmentId ):
		errMsg = "Failed to unassign element " + str(roleId) + " for Role on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# assign to None for unassignment
			roleAssignment.role = None			

			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignGovernanceBody( self, roleAssignmentId, governanceBodyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.GovernanceBodyDelegate import GovernanceBodyDelegate

		errMsg = "Failed to assign element " + str(governanceBodyId) + " for GovernanceBody on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# get the GovernanceBody from db
			governanceBody = GovernanceBodyDelegate().get(governanceBodyId).first();
			
			# assign the GovernanceBody		
			roleAssignment.governanceBody = governanceBody
			
			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except GovernanceBody.DoesNotExist:
			raise ProcessingError(errMsg + " : GovernanceBody with id " + str(governanceBodyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignGovernanceBody( self, roleAssignmentId ):
		errMsg = "Failed to unassign element " + str(governanceBodyId) + " for GovernanceBody on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# assign to None for unassignment
			roleAssignment.governanceBody = None			

			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrganization( self, roleAssignmentId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			roleAssignment.organization = organization
			
			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, roleAssignmentId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on RoleAssignment"

		try:
			# get the RoleAssignment from db
			roleAssignment = self.get( roleAssignmentId ).first()	
			
			# assign to None for unassignment
			roleAssignment.organization = None			

			#save it
			roleAssignment.save()

			# reload and return the appropriate version					
			return self.get( roleAssignmentId );
		except RoleAssignment.DoesNotExist:
			raise ProcessingError(errMsg + " : RoleAssignment with id " + str(roleAssignmentId) + " does not exist.")
		except Exception:
			return None;
		
