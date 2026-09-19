class CarePlansController < ApplicationController
  def index
    @carePlans = CarePlan.all
  end
 
  def show
    @carePlan = CarePlan.find(params[:id])
  end
 
  def new
    @carePlan = CarePlan.new
  end
 
  def edit
    @carePlan = CarePlan.find(params[:id])
  end
 
  def create
    @carePlan = CarePlan.new(carePlan_params)
 
    if @carePlan.save
      redirect_to carePlans_path
    else
      render 'new'
    end
  end
 
  def update
    @carePlan = CarePlan.find(params[:id])
 
    if @carePlan.update(carePlan_params)
      redirect_to carePlans_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @carePlan = CarePlan.find(params[:id])
    @carePlan.destroy
    redirect_to carePlans_path
  end

 
  private
    def carePlan_params
      params.require(:carePlan).permit(:planNumber, :goalSummary, :Status)
    end
end