from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from healthcareOnDjango.models.Authorization import Authorization
from healthcareOnDjango.models.Coverage import Coverage
from healthcareOnDjango.models.ClinicalOrder import ClinicalOrder
from healthcareOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Authorization
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AuthorizationDelegate Declaration
#======================================================================
class AuthorizationDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, authorizationId ):
		try:	
			authorization = Authorization.objects.filter(id=authorizationId)
			return authorization.first();
		except Authorization.DoesNotExist:
			raise ProcessingError("Authorization with id " + str(authorizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, authorization):
		for model in serializers.deserialize("json", authorization):
			model.save()
			return model;

	def create(self, authorization):
		authorization.save()
		return authorization;

	def saveFromJson(self, authorization):
		for model in serializers.deserialize("json", authorization):
			model.save()
			return authorization;
	
	def save(self, authorization):
		authorization.save()
		return authorization;
	
	def delete(self, authorizationId ):
		errMsg = "Failed to delete Authorization from db using id " + str(authorizationId)
		
		try:
			authorization = Authorization.objects.get(id=authorizationId)
			authorization.delete()
			return True
		except Authorization.DoesNotExist:
			raise ProcessingError("Authorization with id " + str(authorizationId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Authorization.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Authorization from db")
		except Exception:
			return None;
		
	def assignCoverage( self, authorizationId, coverageId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.CoverageDelegate import CoverageDelegate

		errMsg = "Failed to assign element " + str(coverageId) + " for Coverage on Authorization"

		try:
			# get the Authorization from db
			authorization = self.get( authorizationId ).first()	
			
			# get the Coverage from db
			coverage = CoverageDelegate().get(coverageId).first();
			
			# assign the Coverage		
			authorization.coverage = coverage
			
			#save it
			authorization.save()

			# reload and return the appropriate version					
			return self.get( authorizationId );
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization with id " + str(authorizationId) + " does not exist.")
		except Coverage.DoesNotExist:
			raise ProcessingError(errMsg + " : Coverage with id " + str(coverageId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignCoverage( self, authorizationId ):
		errMsg = "Failed to unassign element " + str(coverageId) + " for Coverage on Authorization"

		try:
			# get the Authorization from db
			authorization = self.get( authorizationId ).first()	
			
			# assign to None for unassignment
			authorization.coverage = None			

			#save it
			authorization.save()

			# reload and return the appropriate version					
			return self.get( authorizationId );
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization with id " + str(authorizationId) + " does not exist.")
		except Exception:
			return None;
		
	def assignOrder( self, authorizationId, orderId ):
		# lazy importing avoids circular dependencies
		from healthcareOnDjango.delegates.ClinicalOrderDelegate import ClinicalOrderDelegate

		errMsg = "Failed to assign element " + str(orderId) + " for Order on Authorization"

		try:
			# get the Authorization from db
			authorization = self.get( authorizationId ).first()	
			
			# get the ClinicalOrder from db
			clinicalOrder = ClinicalOrderDelegate().get(orderId).first();
			
			# assign the Order		
			authorization.order = clinicalOrder
			
			#save it
			authorization.save()

			# reload and return the appropriate version					
			return self.get( authorizationId );
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization with id " + str(authorizationId) + " does not exist.")
		except ClinicalOrder.DoesNotExist:
			raise ProcessingError(errMsg + " : ClinicalOrder with id " + str(orderId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignOrder( self, authorizationId ):
		errMsg = "Failed to unassign element " + str(orderId) + " for Order on Authorization"

		try:
			# get the Authorization from db
			authorization = self.get( authorizationId ).first()	
			
			# assign to None for unassignment
			authorization.clinicalOrder = None			

			#save it
			authorization.save()

			# reload and return the appropriate version					
			return self.get( authorizationId );
		except Authorization.DoesNotExist:
			raise ProcessingError(errMsg + " : Authorization with id " + str(authorizationId) + " does not exist.")
		except Exception:
			return None;
		
