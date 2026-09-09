from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from hrOnDjango.models.PolicyAcknowledgement import PolicyAcknowledgement
from hrOnDjango.models.Policy import Policy
from hrOnDjango.models.Employee import Employee
from hrOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model PolicyAcknowledgement
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class PolicyAcknowledgementDelegate Declaration
#======================================================================
class PolicyAcknowledgementDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, policyAcknowledgementId ):
		try:	
			policyAcknowledgement = PolicyAcknowledgement.objects.filter(id=policyAcknowledgementId)
			return policyAcknowledgement.first();
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError("PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, policyAcknowledgement):
		for model in serializers.deserialize("json", policyAcknowledgement):
			model.save()
			return model;

	def create(self, policyAcknowledgement):
		policyAcknowledgement.save()
		return policyAcknowledgement;

	def saveFromJson(self, policyAcknowledgement):
		for model in serializers.deserialize("json", policyAcknowledgement):
			model.save()
			return policyAcknowledgement;
	
	def save(self, policyAcknowledgement):
		policyAcknowledgement.save()
		return policyAcknowledgement;
	
	def delete(self, policyAcknowledgementId ):
		errMsg = "Failed to delete PolicyAcknowledgement from db using id " + str(policyAcknowledgementId)
		
		try:
			policyAcknowledgement = PolicyAcknowledgement.objects.get(id=policyAcknowledgementId)
			policyAcknowledgement.delete()
			return True
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError("PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = PolicyAcknowledgement.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all PolicyAcknowledgement from db")
		except Exception:
			return None;
		
	def assignPolicy( self, policyAcknowledgementId, policyId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on PolicyAcknowledgement"

		try:
			# get the PolicyAcknowledgement from db
			policyAcknowledgement = self.get( policyAcknowledgementId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			policyAcknowledgement.policy = policy
			
			#save it
			policyAcknowledgement.save()

			# reload and return the appropriate version					
			return self.get( policyAcknowledgementId );
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, policyAcknowledgementId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on PolicyAcknowledgement"

		try:
			# get the PolicyAcknowledgement from db
			policyAcknowledgement = self.get( policyAcknowledgementId ).first()	
			
			# assign to None for unassignment
			policyAcknowledgement.policy = None			

			#save it
			policyAcknowledgement.save()

			# reload and return the appropriate version					
			return self.get( policyAcknowledgementId );
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except Exception:
			return None;
		
	def assignEmployee( self, policyAcknowledgementId, employeeId ):
		# lazy importing avoids circular dependencies
		from hrOnDjango.delegates.EmployeeDelegate import EmployeeDelegate

		errMsg = "Failed to assign element " + str(employeeId) + " for Employee on PolicyAcknowledgement"

		try:
			# get the PolicyAcknowledgement from db
			policyAcknowledgement = self.get( policyAcknowledgementId ).first()	
			
			# get the Employee from db
			employee = EmployeeDelegate().get(employeeId).first();
			
			# assign the Employee		
			policyAcknowledgement.employee = employee
			
			#save it
			policyAcknowledgement.save()

			# reload and return the appropriate version					
			return self.get( policyAcknowledgementId );
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except Employee.DoesNotExist:
			raise ProcessingError(errMsg + " : Employee with id " + str(employeeId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEmployee( self, policyAcknowledgementId ):
		errMsg = "Failed to unassign element " + str(employeeId) + " for Employee on PolicyAcknowledgement"

		try:
			# get the PolicyAcknowledgement from db
			policyAcknowledgement = self.get( policyAcknowledgementId ).first()	
			
			# assign to None for unassignment
			policyAcknowledgement.employee = None			

			#save it
			policyAcknowledgement.save()

			# reload and return the appropriate version					
			return self.get( policyAcknowledgementId );
		except PolicyAcknowledgement.DoesNotExist:
			raise ProcessingError(errMsg + " : PolicyAcknowledgement with id " + str(policyAcknowledgementId) + " does not exist.")
		except Exception:
			return None;
		
