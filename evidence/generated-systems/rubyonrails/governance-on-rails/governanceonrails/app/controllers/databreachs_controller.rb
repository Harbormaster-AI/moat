class DataBreachsController < ApplicationController
  def index
    @dataBreachs = DataBreach.all
  end
 
  def show
    @dataBreach = DataBreach.find(params[:id])
  end
 
  def new
    @dataBreach = DataBreach.new
  end
 
  def edit
    @dataBreach = DataBreach.find(params[:id])
  end
 
  def create
    @dataBreach = DataBreach.new(dataBreach_params)
 
    if @dataBreach.save
      redirect_to dataBreachs_path
    else
      render 'new'
    end
  end
 
  def update
    @dataBreach = DataBreach.find(params[:id])
 
    if @dataBreach.update(dataBreach_params)
      redirect_to dataBreachs_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataBreach = DataBreach.find(params[:id])
    @dataBreach.destroy
    redirect_to dataBreachs_path
  end

 
  private
    def dataBreach_params
      params.require(:dataBreach).permit(:incidentDate, :description, :recordsAffected, :notificationRequired, :Severity, :Status)
    end
end