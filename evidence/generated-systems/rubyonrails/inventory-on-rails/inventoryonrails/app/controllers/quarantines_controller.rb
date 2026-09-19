class QuarantinesController < ApplicationController
  def index
    @quarantines = Quarantine.all
  end
 
  def show
    @quarantine = Quarantine.find(params[:id])
  end
 
  def new
    @quarantine = Quarantine.new
  end
 
  def edit
    @quarantine = Quarantine.find(params[:id])
  end
 
  def create
    @quarantine = Quarantine.new(quarantine_params)
 
    if @quarantine.save
      redirect_to quarantines_path
    else
      render 'new'
    end
  end
 
  def update
    @quarantine = Quarantine.find(params[:id])
 
    if @quarantine.update(quarantine_params)
      redirect_to quarantines_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @quarantine = Quarantine.find(params[:id])
    @quarantine.destroy
    redirect_to quarantines_path
  end

 
  private
    def quarantine_params
      params.require(:quarantine).permit(:reason, :startedAt, :releasedAt, :Disposition)
    end
end