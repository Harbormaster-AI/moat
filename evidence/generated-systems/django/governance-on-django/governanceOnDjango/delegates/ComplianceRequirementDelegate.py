from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.ComplianceRequirement import ComplianceRequirement
from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ComplianceRequirement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceRequirementDelegate Declaration
#======================================================================
class ComplianceRequirementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, complianceRequirementId ):
		try:	
			complianceRequirement = ComplianceRequirement.objects.filter(id=complianceRequirementId)
			return complianceRequirement.first();
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError("ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, complianceRequirement):
		for model in serializers.deserialize("json", complianceRequirement):
			model.save()
			return model;

	def create(self, complianceRequirement):
		complianceRequirement.save()
		return complianceRequirement;

	def saveFromJson(self, complianceRequirement):
		for model in serializers.deserialize("json", complianceRequirement):
			model.save()
			return complianceRequirement;
	
	def save(self, complianceRequirement):
		complianceRequirement.save()
		return complianceRequirement;
	
	def delete(self, complianceRequirementId ):
		errMsg = "Failed to delete ComplianceRequirement from db using id " + str(complianceRequirementId)
		
		try:
			complianceRequirement = ComplianceRequirement.objects.get(id=complianceRequirementId)
			complianceRequirement.delete()
			return True
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError("ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ComplianceRequirement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ComplianceRequirement from db")
		except Exception:
			return None;
		
	def assignComplianceProgram( self, complianceRequirementId, complianceProgramId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to assign element " + str(complianceProgramId) + " for ComplianceProgram on ComplianceRequirement"

		try:
			# get the ComplianceRequirement from db
			complianceRequirement = self.get( complianceRequirementId ).first()	
			
			# get the ComplianceProgram from db
			complianceProgram = ComplianceProgramDelegate().get(complianceProgramId).first();
			
			# assign the ComplianceProgram		
			complianceRequirement.complianceProgram = complianceProgram
			
			#save it
			complianceRequirement.save()

			# reload and return the appropriate version					
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignComplianceProgram( self, complianceRequirementId ):
		errMsg = "Failed to unassign element " + str(complianceProgramId) + " for ComplianceProgram on ComplianceRequirement"

		try:
			# get the ComplianceRequirement from db
			complianceRequirement = self.get( complianceRequirementId ).first()	
			
			# assign to None for unassignment
			complianceRequirement.complianceProgram = None			

			#save it
			complianceRequirement.save()

			# reload and return the appropriate version					
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Exception:
			return None;
		
	def addPolicies( self, complianceRequirementId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to add elements " + str(policiesIds) + " for Policies on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				complianceRequirement.policies.add(policy)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removePolicies( self, complianceRequirementId, policiesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to remove elements " + str(policiesIds) + " for Policies on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = policiesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Policy		
				policy = PolicyDelegate().get(id).first();	
				# add the Policy
				complianceRequirement.policies.remove(policy)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addControls( self, complianceRequirementId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				complianceRequirement.controls.add(control)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, complianceRequirementId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				complianceRequirement.controls.remove(control)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addObligations( self, complianceRequirementId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to add elements " + str(obligationsIds) + " for Obligations on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				complianceRequirement.obligations.add(obligation)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObligations( self, complianceRequirementId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to remove elements " + str(obligationsIds) + " for Obligations on ComplianceRequirement"

		try:
			# get the ComplianceRequirement
			complianceRequirement = self.get( complianceRequirementId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				complianceRequirement.obligations.remove(obligation)
				
			# save it		
			complianceRequirement.save()
			
			# reload and return the appropriate version
			return self.get( complianceRequirementId );
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement with id " + str(complianceRequirementId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
