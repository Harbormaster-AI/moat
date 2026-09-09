from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Exception_ import Exception_
from governanceOnDjango.models.RetentionSchedule import RetentionSchedule
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Exception_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class Exception_Delegate Declaration
#======================================================================
class Exception_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, exception_Id ):
		try:	
			exception_ = Exception_.objects.filter(id=exception_Id)
			return exception_.first();
		except Exception_.DoesNotExist:
			raise ProcessingError("Exception_ with id " + str(exception_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, exception_):
		for model in serializers.deserialize("json", exception_):
			model.save()
			return model;

	def create(self, exception_):
		exception_.save()
		return exception_;

	def saveFromJson(self, exception_):
		for model in serializers.deserialize("json", exception_):
			model.save()
			return exception_;
	
	def save(self, exception_):
		exception_.save()
		return exception_;
	
	def delete(self, exception_Id ):
		errMsg = "Failed to delete Exception_ from db using id " + str(exception_Id)
		
		try:
			exception_ = Exception_.objects.get(id=exception_Id)
			exception_.delete()
			return True
		except Exception_.DoesNotExist:
			raise ProcessingError("Exception_ with id " + str(exception_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Exception_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Exception_ from db")
		except Exception:
			return None;
		
	def assignRetentionSchedule( self, exception_Id, retentionScheduleId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RetentionScheduleDelegate import RetentionScheduleDelegate

		errMsg = "Failed to assign element " + str(retentionScheduleId) + " for RetentionSchedule on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# get the RetentionSchedule from db
			retentionSchedule = RetentionScheduleDelegate().get(retentionScheduleId).first();
			
			# assign the RetentionSchedule		
			exception_.retentionSchedule = retentionSchedule
			
			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except RetentionSchedule.DoesNotExist:
			raise ProcessingError(errMsg + " : RetentionSchedule with id " + str(retentionScheduleId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRetentionSchedule( self, exception_Id ):
		errMsg = "Failed to unassign element " + str(retentionScheduleId) + " for RetentionSchedule on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# assign to None for unassignment
			exception_.retentionSchedule = None			

			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignPolicy( self, exception_Id, policyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			exception_.policy = policy
			
			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, exception_Id ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# assign to None for unassignment
			exception_.policy = None			

			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignControl( self, exception_Id, controlId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to assign element " + str(controlId) + " for Control on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# get the Control from db
			control = ControlDelegate().get(controlId).first();
			
			# assign the Control		
			exception_.control = control
			
			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControl( self, exception_Id ):
		errMsg = "Failed to unassign element " + str(controlId) + " for Control on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# assign to None for unassignment
			exception_.control = None			

			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignRisk( self, exception_Id, riskId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to assign element " + str(riskId) + " for Risk on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# get the Risk from db
			risk = RiskDelegate().get(riskId).first();
			
			# assign the Risk		
			exception_.risk = risk
			
			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk with id " + str(riskId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignRisk( self, exception_Id ):
		errMsg = "Failed to unassign element " + str(riskId) + " for Risk on Exception_"

		try:
			# get the Exception_ from db
			exception_ = self.get( exception_Id ).first()	
			
			# assign to None for unassignment
			exception_.risk = None			

			#save it
			exception_.save()

			# reload and return the appropriate version					
			return self.get( exception_Id );
		except Exception_.DoesNotExist:
			raise ProcessingError(errMsg + " : Exception_ with id " + str(exception_Id) + " does not exist.")
		except Exception:
			return None;
		
