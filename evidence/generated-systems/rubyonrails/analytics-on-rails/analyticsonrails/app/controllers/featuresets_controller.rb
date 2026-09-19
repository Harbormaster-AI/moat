class FeatureSetsController < ApplicationController
  def index
    @featureSets = FeatureSet.all
  end
 
  def show
    @featureSet = FeatureSet.find(params[:id])
  end
 
  def new
    @featureSet = FeatureSet.new
  end
 
  def edit
    @featureSet = FeatureSet.find(params[:id])
  end
 
  def create
    @featureSet = FeatureSet.new(featureSet_params)
 
    if @featureSet.save
      redirect_to featureSets_path
    else
      render 'new'
    end
  end
 
  def update
    @featureSet = FeatureSet.find(params[:id])
 
    if @featureSet.update(featureSet_params)
      redirect_to featureSets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @featureSet = FeatureSet.find(params[:id])
    @featureSet.destroy
    redirect_to featureSets_path
  end

 
  private
    def featureSet_params
      params.require(:featureSet).permit(:name, :refreshSchedule, :StoreType)
    end
end