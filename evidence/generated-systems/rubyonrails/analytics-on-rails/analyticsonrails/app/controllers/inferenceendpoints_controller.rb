class InferenceEndpointsController < ApplicationController
  def index
    @inferenceEndpoints = InferenceEndpoint.all
  end
 
  def show
    @inferenceEndpoint = InferenceEndpoint.find(params[:id])
  end
 
  def new
    @inferenceEndpoint = InferenceEndpoint.new
  end
 
  def edit
    @inferenceEndpoint = InferenceEndpoint.find(params[:id])
  end
 
  def create
    @inferenceEndpoint = InferenceEndpoint.new(inferenceEndpoint_params)
 
    if @inferenceEndpoint.save
      redirect_to inferenceEndpoints_path
    else
      render 'new'
    end
  end
 
  def update
    @inferenceEndpoint = InferenceEndpoint.find(params[:id])
 
    if @inferenceEndpoint.update(inferenceEndpoint_params)
      redirect_to inferenceEndpoints_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @inferenceEndpoint = InferenceEndpoint.find(params[:id])
    @inferenceEndpoint.destroy
    redirect_to inferenceEndpoints_path
  end

 
  private
    def inferenceEndpoint_params
      params.require(:inferenceEndpoint).permit(:name, :endpointUrl, :trafficShare, :Mode)
    end
end