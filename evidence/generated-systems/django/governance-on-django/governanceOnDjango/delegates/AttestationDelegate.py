from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Attestation import Attestation
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.ComplianceProgram import ComplianceProgram
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Attestation
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AttestationDelegate Declaration
#======================================================================
class AttestationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, attestationId ):
		try:	
			attestation = Attestation.objects.filter(id=attestationId)
			return attestation.first();
		except Attestation.DoesNotExist:
			raise ProcessingError("Attestation with id " + str(attestationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, attestation):
		for model in serializers.deserialize("json", attestation):
			model.save()
			return model;

	def create(self, attestation):
		attestation.save()
		return attestation;

	def saveFromJson(self, attestation):
		for model in serializers.deserialize("json", attestation):
			model.save()
			return attestation;
	
	def save(self, attestation):
		attestation.save()
		return attestation;
	
	def delete(self, attestationId ):
		errMsg = "Failed to delete Attestation from db using id " + str(attestationId)
		
		try:
			attestation = Attestation.objects.get(id=attestationId)
			attestation.delete()
			return True
		except Attestation.DoesNotExist:
			raise ProcessingError("Attestation with id " + str(attestationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Attestation.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Attestation from db")
		except Exception:
			return None;
		
	def assignControl( self, attestationId, controlId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to assign element " + str(controlId) + " for Control on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# get the Control from db
			control = ControlDelegate().get(controlId).first();
			
			# assign the Control		
			attestation.control = control
			
			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControl( self, attestationId ):
		errMsg = "Failed to unassign element " + str(controlId) + " for Control on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# assign to None for unassignment
			attestation.control = None			

			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicy( self, attestationId, policyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			attestation.policy = policy
			
			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, attestationId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# assign to None for unassignment
			attestation.policy = None			

			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignComplianceProgram( self, attestationId, complianceProgramId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ComplianceProgramDelegate import ComplianceProgramDelegate

		errMsg = "Failed to assign element " + str(complianceProgramId) + " for ComplianceProgram on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# get the ComplianceProgram from db
			complianceProgram = ComplianceProgramDelegate().get(complianceProgramId).first();
			
			# assign the ComplianceProgram		
			attestation.complianceProgram = complianceProgram
			
			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except ComplianceProgram.DoesNotExist:
			raise ProcessingError(errMsg + " : ComplianceProgram with id " + str(complianceProgramId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignComplianceProgram( self, attestationId ):
		errMsg = "Failed to unassign element " + str(complianceProgramId) + " for ComplianceProgram on Attestation"

		try:
			# get the Attestation from db
			attestation = self.get( attestationId ).first()	
			
			# assign to None for unassignment
			attestation.complianceProgram = None			

			#save it
			attestation.save()

			# reload and return the appropriate version					
			return self.get( attestationId );
		except Attestation.DoesNotExist:
			raise ProcessingError(errMsg + " : Attestation with id " + str(attestationId) + " does not exist.")
		except Exception:
			return None;
		
