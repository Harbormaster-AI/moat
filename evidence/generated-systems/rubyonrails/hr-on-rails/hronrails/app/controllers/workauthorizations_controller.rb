class WorkAuthorizationsController < ApplicationController
  def index
    @workAuthorizations = WorkAuthorization.all
  end
 
  def show
    @workAuthorization = WorkAuthorization.find(params[:id])
  end
 
  def new
    @workAuthorization = WorkAuthorization.new
  end
 
  def edit
    @workAuthorization = WorkAuthorization.find(params[:id])
  end
 
  def create
    @workAuthorization = WorkAuthorization.new(workAuthorization_params)
 
    if @workAuthorization.save
      redirect_to workAuthorizations_path
    else
      render 'new'
    end
  end
 
  def update
    @workAuthorization = WorkAuthorization.find(params[:id])
 
    if @workAuthorization.update(workAuthorization_params)
      redirect_to workAuthorizations_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @workAuthorization = WorkAuthorization.find(params[:id])
    @workAuthorization.destroy
    redirect_to workAuthorizations_path
  end

 
  private
    def workAuthorization_params
      params.require(:workAuthorization).permit(:country, :expirationDate, :Status)
    end
end