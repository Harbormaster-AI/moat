from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.AccessPolicy import AccessPolicy
from analyticsOnDjango.models.AnalyticsWorkspace import AnalyticsWorkspace
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model AccessPolicy
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class AccessPolicyDelegate Declaration
#======================================================================
class AccessPolicyDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, accessPolicyId ):
		try:	
			accessPolicy = AccessPolicy.objects.filter(id=accessPolicyId)
			return accessPolicy.first();
		except AccessPolicy.DoesNotExist:
			raise ProcessingError("AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, accessPolicy):
		for model in serializers.deserialize("json", accessPolicy):
			model.save()
			return model;

	def create(self, accessPolicy):
		accessPolicy.save()
		return accessPolicy;

	def saveFromJson(self, accessPolicy):
		for model in serializers.deserialize("json", accessPolicy):
			model.save()
			return accessPolicy;
	
	def save(self, accessPolicy):
		accessPolicy.save()
		return accessPolicy;
	
	def delete(self, accessPolicyId ):
		errMsg = "Failed to delete AccessPolicy from db using id " + str(accessPolicyId)
		
		try:
			accessPolicy = AccessPolicy.objects.get(id=accessPolicyId)
			accessPolicy.delete()
			return True
		except AccessPolicy.DoesNotExist:
			raise ProcessingError("AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = AccessPolicy.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all AccessPolicy from db")
		except Exception:
			return None;
		
	def assignWorkspace( self, accessPolicyId, workspaceId ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.AnalyticsWorkspaceDelegate import AnalyticsWorkspaceDelegate

		errMsg = "Failed to assign element " + str(workspaceId) + " for Workspace on AccessPolicy"

		try:
			# get the AccessPolicy from db
			accessPolicy = self.get( accessPolicyId ).first()	
			
			# get the AnalyticsWorkspace from db
			analyticsWorkspace = AnalyticsWorkspaceDelegate().get(workspaceId).first();
			
			# assign the Workspace		
			accessPolicy.workspace = analyticsWorkspace
			
			#save it
			accessPolicy.save()

			# reload and return the appropriate version					
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except AnalyticsWorkspace.DoesNotExist:
			raise ProcessingError(errMsg + " : AnalyticsWorkspace with id " + str(workspaceId) + " does not exist.")
		except Exception:
			return None;
				
	def unassignWorkspace( self, accessPolicyId ):
		errMsg = "Failed to unassign element " + str(workspaceId) + " for Workspace on AccessPolicy"

		try:
			# get the AccessPolicy from db
			accessPolicy = self.get( accessPolicyId ).first()	
			
			# assign to None for unassignment
			accessPolicy.analyticsWorkspace = None			

			#save it
			accessPolicy.save()

			# reload and return the appropriate version					
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Exception:
			return None;
		
	def addDatasets( self, accessPolicyId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				accessPolicy.datasets.add(dataSet)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, accessPolicyId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				accessPolicy.datasets.remove(dataSet)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDashboards( self, accessPolicyId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to add elements " + str(dashboardsIds) + " for Dashboards on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				accessPolicy.dashboards.add(dashboard)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDashboards( self, accessPolicyId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to remove elements " + str(dashboardsIds) + " for Dashboards on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				accessPolicy.dashboards.remove(dashboard)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, accessPolicyId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				accessPolicy.reports.add(report)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, accessPolicyId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				accessPolicy.reports.remove(report)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, accessPolicyId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				accessPolicy.models.add(model)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, accessPolicyId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				accessPolicy.models.remove(model)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatureSets( self, accessPolicyId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to add elements " + str(featureSetsIds) + " for FeatureSets on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				accessPolicy.featureSets.add(featureSet)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatureSets( self, accessPolicyId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to remove elements " + str(featureSetsIds) + " for FeatureSets on AccessPolicy"

		try:
			# get the AccessPolicy
			accessPolicy = self.get( accessPolicyId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				accessPolicy.featureSets.remove(featureSet)
				
			# save it		
			accessPolicy.save()
			
			# reload and return the appropriate version
			return self.get( accessPolicyId );
		except AccessPolicy.DoesNotExist:
			raise ProcessingError(errMsg + " : AccessPolicy with id " + str(accessPolicyId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
