from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Evidence import Evidence
from governanceOnDjango.models.ControlTest_ import ControlTest_
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.AuditWorkpaper import AuditWorkpaper
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Evidence
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class EvidenceDelegate Declaration
#======================================================================
class EvidenceDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, evidenceId ):
		try:	
			evidence = Evidence.objects.filter(id=evidenceId)
			return evidence.first();
		except Evidence.DoesNotExist:
			raise ProcessingError("Evidence with id " + str(evidenceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, evidence):
		for model in serializers.deserialize("json", evidence):
			model.save()
			return model;

	def create(self, evidence):
		evidence.save()
		return evidence;

	def saveFromJson(self, evidence):
		for model in serializers.deserialize("json", evidence):
			model.save()
			return evidence;
	
	def save(self, evidence):
		evidence.save()
		return evidence;
	
	def delete(self, evidenceId ):
		errMsg = "Failed to delete Evidence from db using id " + str(evidenceId)
		
		try:
			evidence = Evidence.objects.get(id=evidenceId)
			evidence.delete()
			return True
		except Evidence.DoesNotExist:
			raise ProcessingError("Evidence with id " + str(evidenceId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Evidence.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Evidence from db")
		except Exception:
			return None;
		
	def assignControlTest( self, evidenceId, controlTestId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

		errMsg = "Failed to assign element " + str(controlTestId) + " for ControlTest on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# get the ControlTest_ from db
			controlTest_ = ControlTest_Delegate().get(controlTestId).first();
			
			# assign the ControlTest		
			evidence.controlTest = controlTest_
			
			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTestId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControlTest( self, evidenceId ):
		errMsg = "Failed to unassign element " + str(controlTestId) + " for ControlTest on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# assign to None for unassignment
			evidence.controlTest_ = None			

			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignControl( self, evidenceId, controlId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to assign element " + str(controlId) + " for Control on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# get the Control from db
			control = ControlDelegate().get(controlId).first();
			
			# assign the Control		
			evidence.control = control
			
			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControl( self, evidenceId ):
		errMsg = "Failed to unassign element " + str(controlId) + " for Control on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# assign to None for unassignment
			evidence.control = None			

			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignObligation( self, evidenceId, obligationId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to assign element " + str(obligationId) + " for Obligation on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# get the Obligation from db
			obligation = ObligationDelegate().get(obligationId).first();
			
			# assign the Obligation		
			evidence.obligation = obligation
			
			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation with id " + str(obligationId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignObligation( self, evidenceId ):
		errMsg = "Failed to unassign element " + str(obligationId) + " for Obligation on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# assign to None for unassignment
			evidence.obligation = None			

			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Exception:
			return None;
		
	def assignWorkpaper( self, evidenceId, workpaperId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditWorkpaperDelegate import AuditWorkpaperDelegate

		errMsg = "Failed to assign element " + str(workpaperId) + " for Workpaper on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# get the AuditWorkpaper from db
			auditWorkpaper = AuditWorkpaperDelegate().get(workpaperId).first();
			
			# assign the Workpaper		
			evidence.workpaper = auditWorkpaper
			
			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except AuditWorkpaper.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditWorkpaper with id " + str(workpaperId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkpaper( self, evidenceId ):
		errMsg = "Failed to unassign element " + str(workpaperId) + " for Workpaper on Evidence"

		try:
			# get the Evidence from db
			evidence = self.get( evidenceId ).first()	
			
			# assign to None for unassignment
			evidence.auditWorkpaper = None			

			#save it
			evidence.save()

			# reload and return the appropriate version					
			return self.get( evidenceId );
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence with id " + str(evidenceId) + " does not exist.")
		except Exception:
			return None;
		
