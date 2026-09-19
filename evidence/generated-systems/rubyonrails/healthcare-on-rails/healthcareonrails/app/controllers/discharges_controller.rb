class DischargesController < ApplicationController
  def index
    @discharges = Discharge.all
  end
 
  def show
    @discharge = Discharge.find(params[:id])
  end
 
  def new
    @discharge = Discharge.new
  end
 
  def edit
    @discharge = Discharge.find(params[:id])
  end
 
  def create
    @discharge = Discharge.new(discharge_params)
 
    if @discharge.save
      redirect_to discharges_path
    else
      render 'new'
    end
  end
 
  def update
    @discharge = Discharge.find(params[:id])
 
    if @discharge.update(discharge_params)
      redirect_to discharges_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @discharge = Discharge.find(params[:id])
    @discharge.destroy
    redirect_to discharges_path
  end

 
  private
    def discharge_params
      params.require(:discharge).permit(:dischargeDateTime, :Disposition)
    end
end