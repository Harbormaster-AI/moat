from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.ControlTest_ import ControlTest_
from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Evidence import Evidence
from governanceOnDjango.models.AuditEngagement import AuditEngagement
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model ControlTest_
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlTest_Delegate Declaration
#======================================================================
class ControlTest_Delegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, controlTest_Id ):
		try:	
			controlTest_ = ControlTest_.objects.filter(id=controlTest_Id)
			return controlTest_.first();
		except ControlTest_.DoesNotExist:
			raise ProcessingError("ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, controlTest_):
		for model in serializers.deserialize("json", controlTest_):
			model.save()
			return model;

	def create(self, controlTest_):
		controlTest_.save()
		return controlTest_;

	def saveFromJson(self, controlTest_):
		for model in serializers.deserialize("json", controlTest_):
			model.save()
			return controlTest_;
	
	def save(self, controlTest_):
		controlTest_.save()
		return controlTest_;
	
	def delete(self, controlTest_Id ):
		errMsg = "Failed to delete ControlTest_ from db using id " + str(controlTest_Id)
		
		try:
			controlTest_ = ControlTest_.objects.get(id=controlTest_Id)
			controlTest_.delete()
			return True
		except ControlTest_.DoesNotExist:
			raise ProcessingError("ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = ControlTest_.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all ControlTest_ from db")
		except Exception:
			return None;
		
	def assignControl( self, controlTest_Id, controlId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlDelegate import ControlDelegate

		errMsg = "Failed to assign element " + str(controlId) + " for Control on ControlTest_"

		try:
			# get the ControlTest_ from db
			controlTest_ = self.get( controlTest_Id ).first()	
			
			# get the Control from db
			control = ControlDelegate().get(controlId).first();
			
			# assign the Control		
			controlTest_.control = control
			
			#save it
			controlTest_.save()

			# reload and return the appropriate version					
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignControl( self, controlTest_Id ):
		errMsg = "Failed to unassign element " + str(controlId) + " for Control on ControlTest_"

		try:
			# get the ControlTest_ from db
			controlTest_ = self.get( controlTest_Id ).first()	
			
			# assign to None for unassignment
			controlTest_.control = None			

			#save it
			controlTest_.save()

			# reload and return the appropriate version					
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except Exception:
			return None;
		
	def assignEngagement( self, controlTest_Id, engagementId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.AuditEngagementDelegate import AuditEngagementDelegate

		errMsg = "Failed to assign element " + str(engagementId) + " for Engagement on ControlTest_"

		try:
			# get the ControlTest_ from db
			controlTest_ = self.get( controlTest_Id ).first()	
			
			# get the AuditEngagement from db
			auditEngagement = AuditEngagementDelegate().get(engagementId).first();
			
			# assign the Engagement		
			controlTest_.engagement = auditEngagement
			
			#save it
			controlTest_.save()

			# reload and return the appropriate version					
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except AuditEngagement.DoesNotExist:
			raise ProcessingError(errMsg + " : AuditEngagement with id " + str(engagementId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignEngagement( self, controlTest_Id ):
		errMsg = "Failed to unassign element " + str(engagementId) + " for Engagement on ControlTest_"

		try:
			# get the ControlTest_ from db
			controlTest_ = self.get( controlTest_Id ).first()	
			
			# assign to None for unassignment
			controlTest_.auditEngagement = None			

			#save it
			controlTest_.save()

			# reload and return the appropriate version					
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except Exception:
			return None;
		
	def addEvidence( self, controlTest_Id, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to add elements " + str(evidenceIds) + " for Evidence on ControlTest_"

		try:
			# get the ControlTest_
			controlTest_ = self.get( controlTest_Id ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				controlTest_.evidence.add(evidence)
				
			# save it		
			controlTest_.save()
			
			# reload and return the appropriate version
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEvidence( self, controlTest_Id, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to remove elements " + str(evidenceIds) + " for Evidence on ControlTest_"

		try:
			# get the ControlTest_
			controlTest_ = self.get( controlTest_Id ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				controlTest_.evidence.remove(evidence)
				
			# save it		
			controlTest_.save()
			
			# reload and return the appropriate version
			return self.get( controlTest_Id );
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ with id " + str(controlTest_Id) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
