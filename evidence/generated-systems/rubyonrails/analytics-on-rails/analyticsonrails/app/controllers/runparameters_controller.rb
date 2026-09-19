class RunParametersController < ApplicationController
  def index
    @runParameters = RunParameter.all
  end
 
  def show
    @runParameter = RunParameter.find(params[:id])
  end
 
  def new
    @runParameter = RunParameter.new
  end
 
  def edit
    @runParameter = RunParameter.find(params[:id])
  end
 
  def create
    @runParameter = RunParameter.new(runParameter_params)
 
    if @runParameter.save
      redirect_to runParameters_path
    else
      render 'new'
    end
  end
 
  def update
    @runParameter = RunParameter.find(params[:id])
 
    if @runParameter.update(runParameter_params)
      redirect_to runParameters_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @runParameter = RunParameter.find(params[:id])
    @runParameter.destroy
    redirect_to runParameters_path
  end

 
  private
    def runParameter_params
      params.require(:runParameter).permit(:name, :value)
    end
end