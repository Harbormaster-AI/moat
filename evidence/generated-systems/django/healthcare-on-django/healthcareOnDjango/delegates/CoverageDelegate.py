from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.models.Patient import Patient
from healthcareOnDjango.models.InsurancePlan import InsurancePlan
from healthcareOnDjango.models.Claim import Claim
from healthcareOnDjango.models.Authorization import Authorization
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Coverage
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class CoverageDelegate Declaration
#======================================================================
class CoverageDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, coverageId ):
		try:	
			coverage = Coverage.objects.filter(id=coverageId)
			return coverage.first();
		except Coverage.DoesNotExist:
			raise ProcessingError("Coverage with id " + str(coverageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, coverage):
		for model in serializers.deserialize("json", coverage):
			model.save()
			return model;

	def create(self, coverage):
		coverage.save()
		return coverage;

	def saveFromJson(self, coverage):
		for model in serializers.deserialize("json", coverage):
			model.save()
			return coverage;
	
	def save(self, coverage):
		coverage.save()
		return coverage;
	
	def delete(self, coverageId ):
		errMsg = "Failed to delete Coverage from db using id " + str(coverageId)
		
		try:
			coverage = Coverage.objects.get(id=coverageId)
			coverage.delete()
			return True
		except Coverage.DoesNotExist:
			raise ProcessingError("Coverage with id " + str(coverageId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Coverage.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Coverage from db")
		except Exception:
			return None;
		
	def assignPatient( self, coverageId, patientId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.PatientDelegate import PatientDelegate

		errMsg = "Failed to assign element " + str(patientId) + " for Patient on Coverage"

		try:
			# get the Coverage from db
			coverage = self.get( coverageId ).first()	
			
			# get the Patient from db
			patient = PatientDelegate().get(patientId).first();
			
			# assign the Patient		
			coverage.patient = patient
			
			#save it
			coverage.save()

			# reload and return the appropriate version					
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Patient.DoesNotExist:
			raise ProcessingError(errMsg + " : Patient with id " + str(patientId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPatient( self, coverageId ):
		errMsg = "Failed to unassign element " + str(patientId) + " for Patient on Coverage"

		try:
			# get the Coverage from db
			coverage = self.get( coverageId ).first()	
			
			# assign to None for unassignment
			coverage.patient = None			

			#save it
			coverage.save()

			# reload and return the appropriate version					
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Exception:
			return None;
		
	def assignPlan( self, coverageId, planId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.InsurancePlanDelegate import InsurancePlanDelegate

		errMsg = "Failed to assign element " + str(planId) + " for Plan on Coverage"

		try:
			# get the Coverage from db
			coverage = self.get( coverageId ).first()	
			
			# get the InsurancePlan from db
			insurancePlan = InsurancePlanDelegate().get(planId).first();
			
			# assign the Plan		
			coverage.plan = insurancePlan
			
			#save it
			coverage.save()

			# reload and return the appropriate version					
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except InsurancePlan.DoesNotExist:
			raise ProcessingError(errMsg + " : InsurancePlan with id " + str(planId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPlan( self, coverageId ):
		errMsg = "Failed to unassign element " + str(planId) + " for Plan on Coverage"

		try:
			# get the Coverage from db
			coverage = self.get( coverageId ).first()	
			
			# assign to None for unassignment
			coverage.insurancePlan = None			

			#save it
			coverage.save()

			# reload and return the appropriate version					
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Exception:
			return None;
		
	def addClaims( self, coverageId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to add elements " + str(claimsIds) + " for Claims on Coverage"

		try:
			# get the Coverage
			coverage = self.get( coverageId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				coverage.claims.add(claim)
				
			# save it		
			coverage.save()
			
			# reload and return the appropriate version
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeClaims( self, coverageId, claimsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClaimDelegate import ClaimDelegate

		errMsg = "Failed to remove elements " + str(claimsIds) + " for Claims on Coverage"

		try:
			# get the Coverage
			coverage = self.get( coverageId ).first()
				
			# split on a comma with no spaces
			idList = claimsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Claim		
				claim = ClaimDelegate().get(id).first();	
				# add the Claim
				coverage.claims.remove(claim)
				
			# save it		
			coverage.save()
			
			# reload and return the appropriate version
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Claim.DoesNotExist:
			raise ProcessingError(errMsg + " : Claim does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addAuthorizations( self, coverageId, authorizationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

		errMsg = "Failed to add elements " + str(authorizationsIds) + " for Authorizations on Coverage"

		try:
			# get the Coverage
			coverage = self.get( coverageId ).first()
				
			# split on a comma with no spaces
			idList = authorizationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Authorization		
				authorization = AuthorizationDelegate().get(id).first();	
				# add the Authorization
				coverage.authorizations.add(authorization)
				
			# save it		
			coverage.save()
			
			# reload and return the appropriate version
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeAuthorizations( self, coverageId, authorizationsIds ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.AuthorizationDelegate import AuthorizationDelegate

		errMsg = "Failed to remove elements " + str(authorizationsIds) + " for Authorizations on Coverage"

		try:
			# get the Coverage
			coverage = self.get( coverageId ).first()
				
			# split on a comma with no spaces
			idList = authorizationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Authorization		
				authorization = AuthorizationDelegate().get(id).first();	
				# add the Authorization
				coverage.authorizations.remove(authorization)
				
			# save it		
			coverage.save()
			
			# reload and return the appropriate version
			return self.get( coverageId );
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
