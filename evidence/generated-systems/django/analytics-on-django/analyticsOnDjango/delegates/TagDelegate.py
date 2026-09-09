from django.core import exceptions
from django.core import serializers
from django.db import models
from django.db import utils

from analyticsOnDjango.models.Tag import Tag
from analyticsOnDjango.models.DataSet import DataSet
from analyticsOnDjango.models.Model import Model
from analyticsOnDjango.models.ModelVersion import ModelVersion
from analyticsOnDjango.models.Dashboard import Dashboard
from analyticsOnDjango.models.Report import Report
from analyticsOnDjango.models.FeatureSet import FeatureSet
from analyticsOnDjango.models.Metric import Metric
from analyticsOnDjango.exceptions import Exceptions

 #======================================================================
# 
# Encapsulates data for model Tag
#
# @author Harbormaster Dev Team
#
#======================================================================

#======================================================================
# Class TagDelegate Declaration
#======================================================================
class TagDelegate :

#======================================================================
# Function Declarations
#======================================================================

	def get(self, tagId ):
		try:	
			tag = Tag.objects.filter(id=tagId)
			return tag.first();
		except Tag.DoesNotExist:
			raise ProcessingError("Tag with id " + str(tagId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 

	def createFromJson(self, tag):
		for model in serializers.deserialize("json", tag):
			model.save()
			return model;

	def create(self, tag):
		tag.save()
		return tag;

	def saveFromJson(self, tag):
		for model in serializers.deserialize("json", tag):
			model.save()
			return tag;
	
	def save(self, tag):
		tag.save()
		return tag;
	
	def delete(self, tagId ):
		errMsg = "Failed to delete Tag from db using id " + str(tagId)
		
		try:
			tag = Tag.objects.get(id=tagId)
			tag.delete()
			return True
		except Tag.DoesNotExist:
			raise ProcessingError("Tag with id " + str(tagId) + " does not exist.")
		except utils.DatabaseError:
			raise StorageReadError()
		except Exception:
			raise GeneralError(errMsg) 
	
	def getAll(self):
		try:
			all = Tag.objects.all()
			return all;
		except utils.DatabaseError:
			raise StorageReadError("Failed to get all Tag from db")
		except Exception:
			return None;
		
	def addDatasets( self, tagId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to add elements " + str(datasetsIds) + " for Datasets on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				tag.datasets.add(dataSet)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDatasets( self, tagId, datasetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DataSetDelegate import DataSetDelegate

		errMsg = "Failed to remove elements " + str(datasetsIds) + " for Datasets on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = datasetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the DataSet		
				dataSet = DataSetDelegate().get(id).first();	
				# add the DataSet
				tag.datasets.remove(dataSet)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except DataSet.DoesNotExist:
			raise ProcessingError(errMsg + " : DataSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModels( self, tagId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to add elements " + str(modelsIds) + " for Models on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				tag.models.add(model)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModels( self, tagId, modelsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelDelegate import ModelDelegate

		errMsg = "Failed to remove elements " + str(modelsIds) + " for Models on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = modelsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Model		
				model = ModelDelegate().get(id).first();	
				# add the Model
				tag.models.remove(model)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Model.DoesNotExist:
			raise ProcessingError(errMsg + " : Model does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addModelVersions( self, tagId, modelVersionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to add elements " + str(modelVersionsIds) + " for ModelVersions on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = modelVersionsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				tag.modelVersions.add(modelVersion)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeModelVersions( self, tagId, modelVersionsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ModelVersionDelegate import ModelVersionDelegate

		errMsg = "Failed to remove elements " + str(modelVersionsIds) + " for ModelVersions on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = modelVersionsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the ModelVersion		
				modelVersion = ModelVersionDelegate().get(id).first();	
				# add the ModelVersion
				tag.modelVersions.remove(modelVersion)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except ModelVersion.DoesNotExist:
			raise ProcessingError(errMsg + " : ModelVersion does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addDashboards( self, tagId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to add elements " + str(dashboardsIds) + " for Dashboards on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				tag.dashboards.add(dashboard)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeDashboards( self, tagId, dashboardsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.DashboardDelegate import DashboardDelegate

		errMsg = "Failed to remove elements " + str(dashboardsIds) + " for Dashboards on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = dashboardsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Dashboard		
				dashboard = DashboardDelegate().get(id).first();	
				# add the Dashboard
				tag.dashboards.remove(dashboard)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Dashboard.DoesNotExist:
			raise ProcessingError(errMsg + " : Dashboard does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addReports( self, tagId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to add elements " + str(reportsIds) + " for Reports on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				tag.reports.add(report)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeReports( self, tagId, reportsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.ReportDelegate import ReportDelegate

		errMsg = "Failed to remove elements " + str(reportsIds) + " for Reports on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = reportsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Report		
				report = ReportDelegate().get(id).first();	
				# add the Report
				tag.reports.remove(report)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Report.DoesNotExist:
			raise ProcessingError(errMsg + " : Report does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addFeatureSets( self, tagId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to add elements " + str(featureSetsIds) + " for FeatureSets on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				tag.featureSets.add(featureSet)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeFeatureSets( self, tagId, featureSetsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.FeatureSetDelegate import FeatureSetDelegate

		errMsg = "Failed to remove elements " + str(featureSetsIds) + " for FeatureSets on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = featureSetsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the FeatureSet		
				featureSet = FeatureSetDelegate().get(id).first();	
				# add the FeatureSet
				tag.featureSets.remove(featureSet)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except FeatureSet.DoesNotExist:
			raise ProcessingError(errMsg + " : FeatureSet does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
	def addMetrics( self, tagId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to add elements " + str(metricsIds) + " for Metrics on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')

			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				tag.metrics.add(metric)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except Exception:
			raise ProcessingError(errMsg) 
		
	def removeMetrics( self, tagId, metricsIds ):
		# lazy importing avoids circular dependencies
		from analyticsOnDjango.delegates.MetricDelegate import MetricDelegate

		errMsg = "Failed to remove elements " + str(metricsIds) + " for Metrics on Tag"

		try:
			# get the Tag
			tag = self.get( tagId ).first()
				
			# split on a comma with no spaces
			idList = metricsIds.split(',')
			
			# iterate over ids
			for id in idList:
				# read the Metric		
				metric = MetricDelegate().get(id).first();	
				# add the Metric
				tag.metrics.remove(metric)
				
			# save it		
			tag.save()
			
			# reload and return the appropriate version
			return self.get( tagId );
		except Tag.DoesNotExist:
			raise ProcessingError(errMsg + " : Tag with id " + str(tagId) + " does not exist.")
		except Metric.DoesNotExist:
			raise ProcessingError(errMsg + " : Metric does not exist.")
		except utils.DatabaseError:
			raise StorageWriteError()
		except Exception:
			raise GeneralError(errMsg) 
		
