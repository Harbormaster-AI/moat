from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.models.Organization import Organization
from governanceOnDjango.models.ComplianceRequirement import ComplianceRequirement
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Attestation import Attestation
from governanceOnDjango.models.Regulation import Regulation
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ComplianceProgram
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ComplianceProgramDelegate Declaration
#======================================================================
class ComplianceProgramDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, complianceProgramId ):
		try:	
			complianceProgram = ComplianceProgram.objects.filter(id=complianceProgramId)
			return complianceProgram.first();
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError("ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, complianceProgram):
		for model in serializers.deserialize("json", complianceProgram):
			model.save()
			return model;

	def create(self, complianceProgram):
		complianceProgram.save()
		return complianceProgram;

	def saveFromJson(self, complianceProgram):
		for model in serializers.deserialize("json", complianceProgram):
			model.save()
			return complianceProgram;
	
	def save(self, complianceProgram):
		complianceProgram.save()
		return complianceProgram;
	
	def delete(self, complianceProgramId ):
		errMsg = "Failed to delete ComplianceProgram from db using id " + str(complianceProgramId)
		
		try:
			complianceProgram = ComplianceProgram.objects.get(id=complianceProgramId)
			complianceProgram.delete()
			return True
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError("ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ComplianceProgram.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ComplianceProgram from db")
		except Exception:
			return None;
		
	def assignOrganization( self, complianceProgramId, organizationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.OrganizationDelegate import OrganizationDelegate

		errMsg = "Failed to assign element " + str(organizationId) + " for Organization on ComplianceProgram"

		try:
			# get the ComplianceProgram from db
			complianceProgram = self.get( complianceProgramId ).first()	
			
			# get the Organization from db
			organization = OrganizationDelegate().get(organizationId).first();
			
			# assign the Organization		
			complianceProgram.organization = organization
			
			#save it
			complianceProgram.save()

			# reload and return the appropriate version					
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Organization.DoesNotExist:
			raise ProcessingError(errMsg + " : Organization with id " + str(organizationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrganization( self, complianceProgramId ):
		errMsg = "Failed to unassign element " + str(organizationId) + " for Organization on ComplianceProgram"

		try:
			# get the ComplianceProgram from db
			complianceProgram = self.get( complianceProgramId ).first()	
			
			# assign to None for unassignment
			complianceProgram.organization = None			

			#save it
			complianceProgram.save()

			# reload and return the appropriate version					
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Exception:
			return None;
		
	def addRequirements( self, complianceProgramId, requirementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

		errMsg = "Failed to add elements " + str(requirementsIds) + " for Requirements on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = requirementsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ComplianceRequirement		
				complianceRequirement = ComplianceRequirementDelegate().get(id).first();	
				# add the ComplianceRequirement
				complianceProgram.requirements.add(complianceRequirement)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRequirements( self, complianceProgramId, requirementsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceRequirementDelegate import ComplianceRequirementDelegate

		errMsg = "Failed to remove elements " + str(requirementsIds) + " for Requirements on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = requirementsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ComplianceRequirement		
				complianceRequirement = ComplianceRequirementDelegate().get(id).first();	
				# add the ComplianceRequirement
				complianceProgram.requirements.remove(complianceRequirement)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except ComplianceRequirement.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceRequirement does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addControls( self, complianceProgramId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to add elements " + str(controlsIds) + " for Controls on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				complianceProgram.controls.add(control)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControls( self, complianceProgramId, controlsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to remove elements " + str(controlsIds) + " for Controls on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = controlsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Control		
				control = ControlDelegate().get(id).first();	
				# add the Control
				complianceProgram.controls.remove(control)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAttestations( self, complianceProgramId, attestationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

		errMsg = "Failed to add elements " + str(attestationsIds) + " for Attestations on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = attestationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Attestation		
				attestation = AttestationDelegate().get(id).first();	
				# add the Attestation
				complianceProgram.attestations.add(attestation)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAttestations( self, complianceProgramId, attestationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AttestationDelegate import AttestationDelegate

		errMsg = "Failed to remove elements " + str(attestationsIds) + " for Attestations on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = attestationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Attestation		
				attestation = AttestationDelegate().get(id).first();	
				# add the Attestation
				complianceProgram.attestations.remove(attestation)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRegulations( self, complianceProgramId, regulationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RegulationDelegate import RegulationDelegate

		errMsg = "Failed to add elements " + str(regulationsIds) + " for Regulations on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = regulationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Regulation		
				regulation = RegulationDelegate().get(id).first();	
				# add the Regulation
				complianceProgram.regulations.add(regulation)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRegulations( self, complianceProgramId, regulationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RegulationDelegate import RegulationDelegate

		errMsg = "Failed to remove elements " + str(regulationsIds) + " for Regulations on ComplianceProgram"

		try:
			# get the ComplianceProgram
			complianceProgram = self.get( complianceProgramId ).first()
				
			# split on a comma with no spaces
			idList = regulationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Regulation		
				regulation = RegulationDelegate().get(id).first();	
				# add the Regulation
				complianceProgram.regulations.remove(regulation)
				
			# save it		
			complianceProgram.save()
			
			# reload and return the appropriate version
			return self.get( complianceProgramId );
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Regulation.DoesNotExist:
			raise ProcessingError(errMsg + " : Regulation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
