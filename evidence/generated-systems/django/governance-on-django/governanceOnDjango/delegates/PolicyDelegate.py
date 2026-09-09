from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.Person import Person
from governanceOnDjango.models.ComplianceRequirement import ComplianceRequirement
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Procedure import Procedure
from governanceOnDjango.models.Exception_ import Exception_
from governanceOnDjango.models.Attestation import Attestation
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Policy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyDelegate Declaration
#======================================================================
class PolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, policyId ):
		try:	
			policy = Policy.objects.filter(id=policyId)
			return policy.first();
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return model;

	def create(self, policy):
		policy.save()
		return policy;

	def saveFromJson(self, policy):
		for model in serializers.deserialize("json", policy):
			model.save()
			return policy;
	
	def save(self, policy):
		policy.save()
		return policy;
	
	def delete(self, policyId ):
		errMsg = "Failed to delete Policy from db using id " + str(policyId)
		
		try:
			policy = Policy.objects.get(id=policyId)
			policy.delete()
			return True
		except Policy.DoesNotExist:
			raise ProcessingError("Policy with id " + str(policyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Policy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Policy from db")
		except Exception:
			return None;
		
	def assignOrganization( self, policyId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			policy.organization = organization
			
			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, policyId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on Policy"

		try:
			# get the Policy from db
			policy = self.get( policyId ).first()	
			
			# assign to None for unassignment
			policy.organization = None			

			#save it
			policy.save()

			# reload and return the appropriate version					
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
		
	def addOwners( self, policyId, ownersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PersonDelegate import PersonDelegate

		errMsg = "Failed to add elements " + str(ownersIds) + " for Owners on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = ownersIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Person		
				person = PersonDelegate().get(id).first();	
				# add the Person
				policy.owners.add(person)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeOwners( self, policyId, ownersIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PersonDelegate import PersonDelegate

		errMsg = "Failed to remove elements " + str(ownersIds) + " for Owners on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = ownersIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Person		
				person = PersonDelegate().get(id).first();	
				# add the Person
				policy.owners.remove(person)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Person.DoesNotExist:
			raise ProcessingError(errMsg + " : Person does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRelatedRequirements( self, policyId, relatedRequirementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

		errMsg = "Failed to add elements " + str(relatedRequirementsIds) + " for RelatedRequirements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = relatedRequirementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceRequirement		
				complianceRequirement = ComplianceRequirementDelegate().get(id).first();	
				# add the ComplianceRequirement
				policy.relatedRequirements.add(complianceRequirement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRelatedRequirements( self, policyId, relatedRequirementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

		errMsg = "Failed to remove elements " + str(relatedRequirementsIds) + " for RelatedRequirements on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = relatedRequirementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceRequirement		
				complianceRequirement = ComplianceRequirementDelegate().get(id).first();	
				# add the ComplianceRequirement
				policy.relatedRequirements.remove(complianceRequirement)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addControls( self, policyId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				policy.controls.add(control)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, policyId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				policy.controls.remove(control)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProcedures( self, policyId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to add elements " + str(proceduresIds) + " for Procedures on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				policy.procedures.add(procedure)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcedures( self, policyId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to remove elements " + str(proceduresIds) + " for Procedures on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				policy.procedures.remove(procedure)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addExceptions( self, policyId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

		errMsg = "Failed to add elements " + str(exceptionsIds) + " for Exceptions on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Exception_		
				exception_ = Exception_Delegate().get(id).first();	
				# add the Exception_
				policy.exceptions.add(exception_)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeExceptions( self, policyId, exceptionsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.Exception_Delegate import Exception_Delegate

		errMsg = "Failed to remove elements " + str(exceptionsIds) + " for Exceptions on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = exceptionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Exception_		
				exception_ = Exception_Delegate().get(id).first();	
				# add the Exception_
				policy.exceptions.remove(exception_)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAttestations( self, policyId, attestationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

		errMsg = "Failed to add elements " + str(attestationsIds) + " for Attestations on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = attestationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Attestation		
				attestation = AttestationDelegate().get(id).first();	
				# add the Attestation
				policy.attestations.add(attestation)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAttestations( self, policyId, attestationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

		errMsg = "Failed to remove elements " + str(attestationsIds) + " for Attestations on Policy"

		try:
			# get the Policy
			policy = self.get( policyId ).first()
				
			# split on a comma with no spaces
			idList = attestationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Attestation		
				attestation = AttestationDelegate().get(id).first();	
				# add the Attestation
				policy.attestations.remove(attestation)
				
			# save it		
			policy.save()
			
			# reload and return the appropriate version
			return self.get( policyId );
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
