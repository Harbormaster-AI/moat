from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from governanceOnDjango.models.Control import Control
from governanceOnDjango.models.Policy import Policy
from governanceOnDjango.models.ControlTest_ import ControlTest_
from governanceOnDjango.models.Evidence import Evidence
from governanceOnDjango.models.Risk import Risk
from governanceOnDjango.models.Obligation import Obligation
from governanceOnDjango.models.Procedure import Procedure
from governanceOnDjango.models.Issue import Issue
from governanceOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Control
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class ControlDelegate Declaration
#======================================================================
class ControlDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, controlId ):
		try:	
			control = Control.objects.filter(id=controlId)
			return control.first();
		except Control.DoesNotExist:
			raise ProcessingError("Control with id " + str(controlId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, control):
		for model in serializers.deserialize("json", control):
			model.save()
			return model;

	def create(self, control):
		control.save()
		return control;

	def saveFromJson(self, control):
		for model in serializers.deserialize("json", control):
			model.save()
			return control;
	
	def save(self, control):
		control.save()
		return control;
	
	def delete(self, controlId ):
		errMsg = "Failed to delete Control from db using id " + str(controlId)
		
		try:
			control = Control.objects.get(id=controlId)
			control.delete()
			return True
		except Control.DoesNotExist:
			raise ProcessingError("Control with id " + str(controlId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Control.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Control from db")
		except Exception:
			return None;
		
	def assignPolicy( self, controlId, policyId ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.PolicyDelegate import PolicyDelegate

		errMsg = "Failed to assign element " + str(policyId) + " for Policy on Control"

		try:
			# get the Control from db
			control = self.get( controlId ).first()	
			
			# get the Policy from db
			policy = PolicyDelegate().get(policyId).first();
			
			# assign the Policy		
			control.policy = policy
			
			#save it
			control.save()

			# reload and return the appropriate version					
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Policy.DoesNotExist:
			raise ProcessingError(errMsg + " : Policy with id " + str(policyId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignPolicy( self, controlId ):
		errMsg = "Failed to unassign element " + str(policyId) + " for Policy on Control"

		try:
			# get the Control from db
			control = self.get( controlId ).first()	
			
			# assign to None for unassignment
			control.policy = None			

			#save it
			control.save()

			# reload and return the appropriate version					
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Exception:
			return None;
		
	def addControlTests( self, controlId, controlTestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

		errMsg = "Failed to add elements " + str(controlTestsIds) + " for ControlTests on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = controlTestsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ControlTest_		
				controlTest_ = ControlTest_Delegate().get(id).first();	
				# add the ControlTest_
				control.controlTests.add(controlTest_)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeControlTests( self, controlId, controlTestsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ControlTest_Delegate import ControlTest_Delegate

		errMsg = "Failed to remove elements " + str(controlTestsIds) + " for ControlTests on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = controlTestsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ControlTest_		
				controlTest_ = ControlTest_Delegate().get(id).first();	
				# add the ControlTest_
				control.controlTests.remove(controlTest_)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except ControlTest_.DoesNotExist:
			raise ProcessingError(errMsg + " : ControlTest_ does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addEvidence( self, controlId, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to add elements " + str(evidenceIds) + " for Evidence on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				control.evidence.add(evidence)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeEvidence( self, controlId, evidenceIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.EvidenceDelegate import EvidenceDelegate

		errMsg = "Failed to remove elements " + str(evidenceIds) + " for Evidence on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = evidenceIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Evidence		
				evidence = EvidenceDelegate().get(id).first();	
				# add the Evidence
				control.evidence.remove(evidence)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Evidence.DoesNotExist:
			raise ProcessingError(errMsg + " : Evidence does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addRisks( self, controlId, risksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to add elements " + str(risksIds) + " for Risks on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = risksIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				control.risks.add(risk)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeRisks( self, controlId, risksIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.RiskDelegate import RiskDelegate

		errMsg = "Failed to remove elements " + str(risksIds) + " for Risks on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = risksIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Risk		
				risk = RiskDelegate().get(id).first();	
				# add the Risk
				control.risks.remove(risk)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Risk.DoesNotExist:
			raise ProcessingError(errMsg + " : Risk does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addObligations( self, controlId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to add elements " + str(obligationsIds) + " for Obligations on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				control.obligations.add(obligation)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeObligations( self, controlId, obligationsIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ObligationDelegate import ObligationDelegate

		errMsg = "Failed to remove elements " + str(obligationsIds) + " for Obligations on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = obligationsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Obligation		
				obligation = ObligationDelegate().get(id).first();	
				# add the Obligation
				control.obligations.remove(obligation)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Obligation.DoesNotExist:
			raise ProcessingError(errMsg + " : Obligation does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addProcedures( self, controlId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to add elements " + str(proceduresIds) + " for Procedures on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				control.procedures.add(procedure)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeProcedures( self, controlId, proceduresIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.ProcedureDelegate import ProcedureDelegate

		errMsg = "Failed to remove elements " + str(proceduresIds) + " for Procedures on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = proceduresIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Procedure		
				procedure = ProcedureDelegate().get(id).first();	
				# add the Procedure
				control.procedures.remove(procedure)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Procedure.DoesNotExist:
			raise ProcessingError(errMsg + " : Procedure does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addIssues( self, controlId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to add elements " + str(issuesIds) + " for Issues on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				control.issues.add(issue)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeIssues( self, controlId, issuesIds ):
		# lazy importing avoids circular dependencies
		from governanceOnDjango.delegates.IssueDelegate import IssueDelegate

		errMsg = "Failed to remove elements " + str(issuesIds) + " for Issues on Control"

		try:
			# get the Control
			control = self.get( controlId ).first()
				
			# split on a comma with no spaces
			idList = issuesIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Issue		
				issue = IssueDelegate().get(id).first();	
				# add the Issue
				control.issues.remove(issue)
				
			# save it		
			control.save()
			
			# reload and return the appropriate version
			return self.get( controlId );
		except Control.DoesNotExist:
			raise ProcessingError(errMsg + " : Control with id " + str(controlId) + " does not exist.")
		except Issue.DoesNotExist:
			raise ProcessingError(errMsg + " : Issue does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
