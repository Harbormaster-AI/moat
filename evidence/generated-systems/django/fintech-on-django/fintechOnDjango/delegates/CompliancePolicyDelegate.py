from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from fintechOnDjango.models.CompliancePolicy import CompliancePolicy
from fintechOnDjango.models.FinancialInstitution import FinancialInstitution
from fintechOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model CompliancePolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CompliancePolicyDelegate Declaration
#======================================================================
class CompliancePolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, compliancePolicyId ):
		try:	
			compliancePolicy = CompliancePolicy.objects.filter(id=compliancePolicyId)
			return compliancePolicy.first();
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError("CompliancePolicy with id " + str(compliancePolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, compliancePolicy):
		for model in serializers.deserialize("json", compliancePolicy):
			model.save()
			return model;

	def create(self, compliancePolicy):
		compliancePolicy.save()
		return compliancePolicy;

	def saveFromJson(self, compliancePolicy):
		for model in serializers.deserialize("json", compliancePolicy):
			model.save()
			return compliancePolicy;
	
	def save(self, compliancePolicy):
		compliancePolicy.save()
		return compliancePolicy;
	
	def delete(self, compliancePolicyId ):
		errMsg = "Failed to delete CompliancePolicy from db using id " + str(compliancePolicyId)
		
		try:
			compliancePolicy = CompliancePolicy.objects.get(id=compliancePolicyId)
			compliancePolicy.delete()
			return True
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError("CompliancePolicy with id " + str(compliancePolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = CompliancePolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all CompliancePolicy from db")
		except Exception:
			return None;
		
	def assignInstitution( self, compliancePolicyId, institutionId ):
		# lazy importing avoids circular dependencies
		from fintechOnDjango.delegates.FinancialInstitutionDelegate import FinancialInstitutionDelegate

		errMsg = "Failed to assign element " + str(institutionId) + " for Institution on CompliancePolicy"

		try:
			# get the CompliancePolicy from db
			compliancePolicy = self.get( compliancePolicyId ).first()	
			
			# get the FinancialInstitution from db
			financialInstitution = FinancialInstitutionDelegate().get(institutionId).first();
			
			# assign the Institution		
			compliancePolicy.institution = financialInstitution
			
			#save it
			compliancePolicy.save()

			# reload and return the appropriate version					
			return self.get( compliancePolicyId );
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : CompliancePolicy with id " + str(compliancePolicyId) + " does not exist.")
		except FinancialInstitution.DoesNotExist:
			raise ProcessingError(errMsg + " : FinancialInstitution with id " + str(institutionId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignInstitution( self, compliancePolicyId ):
		errMsg = "Failed to unassign element " + str(institutionId) + " for Institution on CompliancePolicy"

		try:
			# get the CompliancePolicy from db
			compliancePolicy = self.get( compliancePolicyId ).first()	
			
			# assign to None for unassignment
			compliancePolicy.financialInstitution = None			

			#save it
			compliancePolicy.save()

			# reload and return the appropriate version					
			return self.get( compliancePolicyId );
		except CompliancePolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : CompliancePolicy with id " + str(compliancePolicyId) + " does not exist.")
		except Exception:
			return None;
		
